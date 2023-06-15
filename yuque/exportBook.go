package yuque

import (
	"ChangePicBed/model"
	"ChangePicBed/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RequestData struct {
	Type    string `json:"type"`
	Force   int    `json:"force"`
	Options string `json:"options"`
}

type ResponseData struct {
	Data data `mapstructure:"data"`
}

type data struct {
	State       string `mapstructure:"state"`
	BookName    string `mapstructure:"book_name"`
	SummaryName string `mapstructure:"summary_name"`
	Url         string `mapstructure:"url"`
}

func ExportBook(bookStacks model.YuqueBookStacks, bookId int, config model.Config, options int) {
	var err error
	var req *http.Request
	var resp *http.Response
	var client *http.Client
	var summarys []ResponseData
	var flag int = 0

	for _, _data := range bookStacks.Data {
		for _, book := range _data.Books {
			if book.Id == bookId {
				flag = 1
				for _, summary := range book.Summary {
					apiUrl := fmt.Sprintf("https://www.yuque.com/api/docs/%d/export", summary.Id)
					method := "POST"
					reqData := RequestData{
						Type:    "markdown",
						Force:   0,
						Options: `{"latexType":2}`,
					}
					jsonData, _ := json.Marshal(reqData)

					req, err = http.NewRequest(method, apiUrl, bytes.NewBuffer(jsonData))
					if err != nil {

					}
					req.AddCookie(&http.Cookie{Name: "_yuque_session", Value: config.YuqueConfig.YuqueSession})
					req.AddCookie(&http.Cookie{Name: "yuque_ctoken", Value: config.YuqueConfig.YuqueCToken})
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("Origin", "https://www.yuque.com")
					req.Header.Set("Referer", fmt.Sprintf("https://www.yuque.com/%s/%s/%s", book.User.Login, book.Slug, summary.Slug))
					req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

					client = &http.Client{}
					resp, err = client.Do(req)
					respBytes, _ := io.ReadAll(resp.Body)

					var respData ResponseData
					err = json.Unmarshal(respBytes, &respData)
					respData.Data.BookName = book.Name
					respData.Data.SummaryName = summary.Title

					summarys = append(summarys, respData)
				}
			}
		}
	}

	for _, summary := range summarys {
		exportPath := fmt.Sprintf("%s/%s", config.YuqueConfig.ExportPath, summary.Data.BookName)
		err = utils.CheckDir(exportPath)
		filePath := fmt.Sprintf("%s/%s/%s.md", config.YuqueConfig.ExportPath, summary.Data.BookName, summary.Data.SummaryName)
		file, _ := os.Create(filePath)

		client = &http.Client{}
		req, _ = http.NewRequest("GET", summary.Data.Url, nil)
		req.AddCookie(&http.Cookie{Name: "_yuque_session", Value: config.YuqueConfig.YuqueSession})
		req.AddCookie(&http.Cookie{Name: "yuque_ctoken", Value: config.YuqueConfig.YuqueCToken})

		resp, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		_, err = io.Copy(file, resp.Body)
		fmt.Println("Exported:", filePath)
		_ = file.Close()

		if options != 0 {
			tmpConfig := config
			tmpConfig.OutputDir = exportPath
			err = utils.ChangePicBed(filePath, tmpConfig)
		}
	}

	if flag != 0 {
		fmt.Println("Export complete!")
	} else {
		fmt.Println("BookID error, Knowledge base not found.")
	}
}
