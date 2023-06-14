package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func Clear(tempPath string) error {
	var err error
	fmt.Println("Cleaning up temporary files...")
	err = filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = os.Remove(path)
			if err != nil {
				return err
			}
			fmt.Println("Deleted file:", path)
		}
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("Done!")
	return nil
}
