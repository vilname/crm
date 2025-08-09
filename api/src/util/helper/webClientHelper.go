package helper

import (
	"api/src/util/constant"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func GetWebClient(url string) []byte {
	client := http.Client{
		Timeout: constant.Timeout * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return body
}

//func PostWebClient(url string, body []byte) ([]byte, error) {
//	client := http.Client{
//		Timeout: enum.Timeout * time.Second,
//	}
//	bodyReader := bytes.NewReader(body)
//
//	resp, err := client.Post(url, "application/json", bodyReader)
//
//	if err != nil {
//		return nil, err
//	}
//
//	defer resp.Body.Close()
//
//	respBody, _ := io.ReadAll(resp.Body)
//
//	return respBody, nil
//}

func PostWebClient(url string, body []byte) ([]byte, error) {
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}

	apiKey := os.Getenv("API_KEY_DS")
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	return respBody, nil
}
