package utils

import (
	"ChangePicBed/model"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

func UploadCOS(markdown model.MarkdownInfo, config model.Config) (model.MarkdownInfo, error) {
	targetUrl, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", config.CosConfig.BucketName, config.CosConfig.BucketArea))
	bucketUrl := &cos.BaseURL{BucketURL: targetUrl}
	client := cos.NewClient(bucketUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.CosConfig.SecretID,
			SecretKey: config.CosConfig.SecretKey,
		},
	})
	for i := 0; i < len(markdown.ImagesInfo); i++ {
		result, _, err := client.Object.Upload(
			context.Background(),
			fmt.Sprintf("%s/%s.jpg", config.CosConfig.PicPath, markdown.ImagesInfo[i].ImageMd5),
			fmt.Sprintf("%s/%s.jpg", config.TempDir, markdown.ImagesInfo[i].ImageMd5),
			nil,
		)
		if err != nil {
			return markdown, err
		}
		markdown.ImagesInfo[i].ImageDes = result.Location
		fmt.Println(fmt.Sprintf("Uploaded image: %s.jpg", markdown.ImagesInfo[i].ImageMd5))
	}
	return markdown, nil
}
