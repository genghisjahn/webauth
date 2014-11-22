package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const TESTURL = "http://localhost:8080"

const EXPECTEDRESULT = "\nEXPECTED:%v\nRESULT:%v\n"

func TestValidLogin(t *testing.T) {
	r, err := http.NewRequest("POST", fmt.Sprintf("%v/auth", TESTURL), nil)
	v := url.Values{}
	v.Add("un", "test@test.com")
	v.Add("pw", "password")
	r.Form = v
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	authHandler(w, r)
	if w.Code != 200 {
		t.Errorf(EXPECTEDRESULT, 200, w.Code)
	}

}
