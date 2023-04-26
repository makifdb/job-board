package sources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	httpClient = &http.Client{}

	bufferPool = sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
)

const HIMALAYAS_URL = "https://himalayas.app/jobs/api"

type HimalayasResponse struct {
	UpdatedAt  int `json:"updated_at"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	TotalCount int `json:"total_count"`
	Jobs       []struct {
		Title       string   `json:"title"`
		Excerpt     string   `json:"excerpt"`
		Image       string   `json:"image"`
		CompanyName string   `json:"companyName"`
		CompanyLogo string   `json:"companyLogo"`
		Categories  []string `json:"categories"`
		Description string   `json:"description"`
	} `json:"jobs"`
}

func GetHimalayaJobList(limit, offset int) (*HimalayasResponse, error) {

	// get buffer from pool
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)
	buf.Reset()

	url := fmt.Sprintf("%s?limit=%d&offset=%d", HIMALAYAS_URL, limit, offset)

	// create a new request
	req, err := http.NewRequest("GET", url, buf)
	if err != nil {
		return nil, err
	}

	// send request
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// close response body
	defer res.Body.Close()

	// decode response body to response struct
	var response HimalayasResponse

	// decode response body to json

	dec := json.NewDecoder(res.Body)

	if err := dec.Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
