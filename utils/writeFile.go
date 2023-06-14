package utils

import (
	"ChangePicBed/model"
	"bufio"
	"fmt"
	"os"
)

func WriteFile(markdown model.MarkdownInfo, config model.Config) error {
	var err error

	var file *os.File

	for i := 0; i < len(markdown.ImagesInfo); i++ {
		lineIndex := markdown.ImagesInfo[i].LineIndex
		imageMd5 := markdown.ImagesInfo[i].ImageMd5
		imageDes := markdown.ImagesInfo[i].ImageDes

		markdown.ContentLines[lineIndex] = fmt.Sprintf("![%s.jpg](%s)", imageMd5, imageDes)
	}

	file, err = os.Create(fmt.Sprintf("%s/%s", config.OutputDir, markdown.FileName))
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	for _, line := range markdown.ContentLines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
