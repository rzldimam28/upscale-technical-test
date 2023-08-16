package fetch

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	URL = "https://jsonplaceholder.typicode.com/posts"
	ErrNotFoundURL = errors.New("can not find url")
)

type ResponseBody struct {
	ID    int64 `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func FetchFromAPI(url string) ([]*ResponseBody, error) {

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrNotFoundURL
	}
	defer resp.Body.Close()

	var respBody []*ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}