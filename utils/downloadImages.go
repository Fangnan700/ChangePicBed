package utils

import (
	"ChangePicBed/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadImages(markdown model.MarkdownInfo, config model.Config) (model.MarkdownInfo, error) {
	var err error
	var resp *http.Response
	var file *os.File

	fmt.Println("Downloading images of", markdown.FileName)
	for i := 0; i < len(markdown.ImagesInfo); i++ {
		resp, err = http.Get(markdown.ImagesInfo[i].ImageSrc)
		if err != nil {
			return markdown, err
		}

		hash := md5.New()
		hash.Write([]byte(markdown.ImagesInfo[i].ImageSrc))
		_md5 := hex.EncodeToString(hash.Sum(nil))
		markdown.ImagesInfo[i].ImageMd5 = _md5

		imagePath := fmt.Sprintf("%s", config.TempDir)
		file, err = os.Create(fmt.Sprintf("%s/%s.jpg", imagePath, _md5))
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return markdown, err
		}
		fmt.Println(fmt.Sprintf("Downloaded image: %s.jpg", _md5))
		_ = file.Close()
	}
	return markdown, nil
}
