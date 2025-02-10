package options

import (
	"fmt"
	"os"
)

func verifyFolder(folderName, folderLocation string) (bool, error) {
	_, err := os.ReadDir(folderLocation)
	if err != nil {
		err = fmt.Errorf("failed to read %s folder: %v", folderName, err)
		return false, err
	}
	file, err := os.CreateTemp(folderLocation, "test")
	if err != nil {
		err = fmt.Errorf("failed to write to %s folder: %v", folderName, err)
		return false, err
	}
	file.Close()
	os.Remove(file.Name())

	return true, nil
}
