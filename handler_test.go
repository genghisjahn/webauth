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
	v.Add("un", "test_user")
	v.Add("pw", "password")
	r.Form = v
	//r.Body = ioutil.NopCloser(strings.NewReader(fmt.Sprintf("{\"language\":\"%v\"}", invalidLang)))
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	log.Println("about to call auth")
	authHandler(w, r)
	if w.Code != 200 {
		t.Errorf(EXPECTEDRESULT, 200, w.Code)
	}

}
