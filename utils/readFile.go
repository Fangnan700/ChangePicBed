package utils

import (
	"ChangePicBed/model"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func ReadFile(filePath string) (model.MarkdownInfo, error) {
	var err error
	var stat os.FileInfo
	var file *os.File
	var fileName string
	var fileSize int64
	var lines []string
	var images []model.Images
	var markdown model.MarkdownInfo

	file, err = os.Open(filePath)
	if err != nil {
		return markdown, err
	}
	stat, err = file.Stat()
	fileName = stat.Name()
	fileSize = stat.Size()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	_ = file.Close()

	for index, line := range lines {
		if len(line) > 0 && line[:2] == "![" {
			var image model.Images

			re := regexp.MustCompile(`\((.*?)\)`)
			match := re.FindStringSubmatch(line)
			image.LineIndex = index
			image.ImageSrc = match[1]

			images = append(images, image)
		}
	}

	markdown.FileName = fileName
	markdown.FileSize = fileSize
	markdown.ContentLines = lines
	markdown.ImagesInfo = images

	fmt.Println("Read file:", markdown.FileName)
	return markdown, nil
}
