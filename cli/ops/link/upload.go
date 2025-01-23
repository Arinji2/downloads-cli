package link

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Upload struct {
	uploadType LinkType
	filePath   string
}

type Endpoint struct {
	domain         string
	responsePrefix string
}

func getEndpoint(linkType LinkType) Endpoint {
	switch linkType {
	case LinkTemp:
		return Endpoint{
			domain:         "https://litterbox.catbox.moe/resources/internals/api.php",
			responsePrefix: "https://litter.catbox.moe",
		}
	case LinkPerm:
		return Endpoint{
			domain:         "https://catbox.moe/user/api.php",
			responsePrefix: "https://files.catbox.moe",
		}
	default:
		return Endpoint{}
	}
}

func (u *Upload) UploadData() (string, error) {
	file, err := os.ReadFile(u.filePath)
	if err != nil {
		return "", err
	}

	fileName := filepath.Base(u.filePath)

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		if u.uploadType == LinkTemp {
			m.WriteField("time", "1h")
		}
		part, err := m.CreateFormFile("fileToUpload", filepath.Base(fileName))
		if err != nil {
			return
		}
		if _, err = io.Copy(part, bytes.NewBuffer(file)); err != nil {
			return
		}
	}()
	endpoint := getEndpoint(u.uploadType)
	if endpoint.domain == "" {
		return "", errors.New("invalid upload type")
	}
	req, err := http.NewRequest(http.MethodPost, endpoint.domain, r)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", m.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(string(body), endpoint.responsePrefix) {
		return string(body), nil
	}
	return "", errors.New("invalid response from endpoint")
}
