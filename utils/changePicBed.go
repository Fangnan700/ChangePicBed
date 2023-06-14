package utils

import (
	"ChangePicBed/model"
	"fmt"
	"os"
)

func ChangePicBed(filePath string, config model.Config) error {
	var err error
	var markdown model.MarkdownInfo

	markdown, err = ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	markdown, err = DownloadImages(markdown, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	markdown, err = UploadCOS(markdown, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	err = WriteFile(markdown, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return nil
}
