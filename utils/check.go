package utils

import (
	"fmt"
	"os"
)

func CheckDir(folderPath string) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		fmt.Println("Created folder:", folderPath)
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
