package dial

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Spider struct {
	Client *http.Client
}

func NewSpider(timeout time.Duration, maxConn int) *Spider {
	nt := &http.Transport{
		MaxConnsPerHost:       maxConn,
		ResponseHeaderTimeout: timeout,
	}
	client := &http.Client{
		Transport: nt,
		Timeout:   timeout,
	}
	return &Spider{
		Client: client,
	}
}

// Get request
func (s *Spider) Get(url, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if token != "" {
		req.Header.Set("token", token)
	}
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n", resBytes)
	return resBytes, nil
}

// Post request
func (s *Spider) Post(url, token string, data []byte) ([]byte, error) {
	//data, err := json.Marshal(res)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", token)
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
