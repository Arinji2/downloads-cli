package link

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Upload struct {
	UploadType LinkType
	FilePath   string
	UserHash   string
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
	file, err := os.ReadFile(u.FilePath)
	if err != nil {
		return "", err
	}

	fileName := filepath.Base(u.FilePath)

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		if u.UploadType == LinkTemp {
			m.WriteField("time", "1h")
		}
		if u.UserHash != "" {
			m.WriteField("userhash", u.UserHash)
		}
		part, err := m.CreateFormFile("fileToUpload", filepath.Base(fileName))
		if err != nil {
			return
		}
		if _, err = io.Copy(part, bytes.NewBuffer(file)); err != nil {
			return
		}
	}()
	endpoint := getEndpoint(u.UploadType)
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

func (u *Upload) DeletedUploadedFile(uploadID string) error {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		if err := m.WriteField("reqtype", "deletefiles"); err != nil {
			w.CloseWithError(err)
			return
		}
		if err := m.WriteField("userhash", u.UserHash); err != nil {
			w.CloseWithError(err)
			return
		}
		if err := m.WriteField("files", uploadID); err != nil {
			w.CloseWithError(err)
			return
		}
	}()

	endpoint := getEndpoint(u.UploadType)
	if endpoint.domain == "" {
		return errors.New("invalid upload type")
	}

	req, err := http.NewRequest(http.MethodPost, endpoint.domain, r)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", m.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	return nil
}
