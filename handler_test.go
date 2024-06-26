package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	main "playsee.co/interview"
)

func Test_SuccessRequest(t *testing.T) {
	const expectResult = `head: aa
node1: bb
node2: cc
tail: dd
`
	const contentType = "application/json"
	APIKey := "qwerklj1230dsa350123l2k1j4kl1j24"
	data := []byte(`{"Array":["aa", "bb", "cc", "dd"]}`)
	req, err := http.NewRequest("POST", "/api-test", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("api-key", APIKey)
	rr := httptest.NewRecorder()
	handler := http.NewServeMux()
	handler.HandleFunc("/api-test", main.RequestMiddleware(main.Test1))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusCreated)
	}
	var result main.Node
	err = json.NewDecoder(rr.Body).Decode(&result)
	if err != nil {
		t.Error(err)
	}
	if strings.Compare(result.ToString(), expectResult) != 0 {
		t.Errorf("result %v, expect %v", result.ToString(), expectResult)
	}
}

func Test_FailedWithoutApiKeyRequest(t *testing.T) {
	const contentType = "application/json"
	req, err := http.NewRequest("POST", "/api-test", nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", contentType)
	rr := httptest.NewRecorder()
	handler := http.NewServeMux()
	handler.HandleFunc("/api-test", main.RequestMiddleware(main.Test1))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func Test_FailedWithInvalidPayload(t *testing.T) {
	const contentType = "application/json"
	APIKey := "qwerklj1230dsa350123l2k1j4kl1j24"
	data := []byte(`["aa", "bb", "cc", "dd"]`)
	req, err := http.NewRequest("POST", "/api-test", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("api-key", APIKey)
	rr := httptest.NewRecorder()
	handler := http.NewServeMux()
	handler.HandleFunc("/api-test", main.RequestMiddleware(main.Test1))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
