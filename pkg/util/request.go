package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HttpPOST[R interface{}](url string, payload string) (*R, error) {
	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	var token R
	if err := decoder.Decode(&token); err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return &token, err
}

func BuildQueryParams(params map[string]interface{}) string {
	var queries []string
	for key, val := range params {
		query := fmt.Sprintf("%v=%v", key, val)
		queries = append(queries, query)
	}

	if len(queries) > 0 {
		return "?" + strings.Join(queries, "&")
	}
	return ""
}
