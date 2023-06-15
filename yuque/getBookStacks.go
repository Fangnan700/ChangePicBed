package yuque

import (
	"ChangePicBed/model"
	"encoding/json"
	"io"
	"net/http"
)

func GetBookStacks(config model.Config) (model.YuqueBookStacks, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var client *http.Client
	var bookStacks model.YuqueBookStacks

	apiUrl := "https://www.yuque.com/api/mine/book_stacks"
	method := "GET"

	req, err = http.NewRequest(method, apiUrl, nil)
	if err != nil {
		return bookStacks, err
	}
	req.AddCookie(&http.Cookie{Name: "_yuque_session", Value: config.YuqueConfig.YuqueSession})
	req.AddCookie(&http.Cookie{Name: "yuque_ctoken", Value: config.YuqueConfig.YuqueCToken})
	req.Header.Set("Referer", "https://www.yuque.com/dashboard/books")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return bookStacks, err
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &bookStacks)

	return bookStacks, nil
}
