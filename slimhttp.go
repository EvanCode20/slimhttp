package slimhttp

import (
	"encoding/json"
	"net/http"
)

func Get[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func Post() {

}

func Put() {

}

func Delete() {

}
