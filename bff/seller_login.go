package bff

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HealthCheck interface {
	Run() string
	Description() string
}

type LoginHealthCheck struct {
}

func (l LoginHealthCheck) Run() string {
	jwt, err := login()
	if err != nil {
		return "ERROR"
	}
	status, err := getSeller(jwt, "1")
	if err != nil {
		return "ERROR"
	}
	if status == http.StatusOK {
		return "SUCCESS"
	}
	return "FAILED"
}

func (l LoginHealthCheck) Description() string {
	return "login health check"
}

var client = http.Client{}

func getSeller(jwt string, id string) (int, error) {
	requestURL := fmt.Sprintf("http://localhost:%d/seller/%s", 8080, id)
	req, _ := http.NewRequest(http.MethodGet, requestURL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	res, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	return res.StatusCode, nil
}

func login() (string, error) {
	jsonBody := []byte(`{"user": "ralf", "password": "password"}`)
	requestURL := fmt.Sprintf("http://localhost:%d/login", 8080)
	req, _ := http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(body), nil
}
