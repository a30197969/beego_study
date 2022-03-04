package models

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// CURL POST请求
func CurlPost(url string, contentType string, reqBody string) (string, error) {
	client := http.Client{}
	resp, err := client.Post(url, contentType, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func CurlPost2(url string, contentType string, reqBody string) (string, error) {
	client := http.Client{}
	resp, err := client.Post(url, contentType, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
