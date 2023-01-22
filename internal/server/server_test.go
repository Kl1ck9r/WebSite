package server

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPageLogin(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(PageLogin))
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("excepted 200 but got: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	req, _ := http.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	PageLogin(rec, req)

	file, err := os.Open("./templates/auth/login.html")
	if err!=nil{
		t.Error(err)
	}

	read, err := ioutil.ReadAll(file)
	if err!=nil{
		t.Error(err)
	}

	assert.Equal(t, rec.Body.Bytes(),read)
}

func TestPageRegistration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(PageRegistration))
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("excepted 200 but got: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	req, _ := http.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	PageLogin(rec, req)

	file, err := os.Open("./templates/auth/signup.html")
	if err!=nil{
		t.Error(err)
	}

	read, err := ioutil.ReadAll(file)
	if err!=nil{
		t.Error(err)
	}

	assert.Equal(t, rec.Body.Bytes(),read)
}

func TestPageResetPassword(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(PageResetPassword))
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("excepted 200 but got: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	req, _ := http.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	PageLogin(rec, req)

	file, err := os.Open("./templates/recovery.html")
	if err!=nil{
		t.Error(err)
	}

	read, err := ioutil.ReadAll(file)
	if err!=nil{
		t.Error(err)
	}

	assert.Equal(t, rec.Body.Bytes(),read)
}

func TestPageMain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(PageMain))
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("excepted 200 but got: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	req, _ := http.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	PageLogin(rec, req)

	file, err := os.Open("./templates/home.html")
	if err!=nil{
		t.Error(err)
	}

	read, err := ioutil.ReadAll(file)
	if err!=nil{
		t.Error(err)
	}

	assert.Equal(t, rec.Body.Bytes(),read)
}
