package main

/*
Go에서 빌트인 패키지만 사용하여
rest API로 영화 세부 정보를 검색하는 예시
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// omdbapi.com API key
const APIKEY = "193ef3a"

// omdbapi.com에서 반환하는 JSON 구조체
// 예시의 간략화를 위해 일부 값은
// 구조체에 매핑하지 않음

type MovieInfo struct {
	Title string `json:"Title"`
	Year  string `json:Year`
}

func main() {
	body, _ := SearchById("tt3896198")
	fmt.Printf("Title = %v\n", body.Title)
	body, _ = SearchByName("Game of")
	fmt.Printf("Title = %v\n", body.Title)
}

func SearchByName(name string) (*MovieInfo, error) {
	parms := url.Values{}
	parms.Set("apikey", APIKEY)
	parms.Set("t", name)
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	fmt.Printf("siteURL = %v\n", siteURL)
	body, err := sendGetRequest(siteURL)
	if err != nil {

		return nil, errors.New(err.Error() + "\nBody:" + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func SearchById(id string) (*MovieInfo, error) {
	parms := url.Values{}
	parms.Set("apikey", APIKEY)
	parms.Set("i", id)
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	fmt.Printf("siteURL = %v\n", siteURL)
	body, err := sendGetRequest(siteURL)
	if err != nil {

		return nil, errors.New(err.Error() + "\nBody:" + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func sendGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status)
	}
	return string(body), nil
}
