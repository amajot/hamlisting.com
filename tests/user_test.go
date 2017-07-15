package tests

import (
	beetest "github.com/astaxie/beego/testing"
	"testing"
	"strconv"
)


func TestLoginGet(t *testing.T) {
	request := beetest.Get("/login")
	response, _ := request.Response()    
	defer response.Body.Close()
	
	if 	response.StatusCode != 200 {
		t.Fatal("/login route is throwing HTTP code " + strconv.Itoa(response.StatusCode) + " instead of 200")
	}
}

func TestLogoutGet(t *testing.T) {
	request := beetest.Get("/logout")
	response, _ := request.Response()    
	defer response.Body.Close()
	
	if 	response.StatusCode != 200 {
		t.Fatal("/logout route is throwing HTTP code " + strconv.Itoa(response.StatusCode) + " instead of 200")
	}
}

func TestForgotGet(t *testing.T) {
	request := beetest.Get("/forgot")
	response, _ := request.Response()    
	defer response.Body.Close()
	
	if 	response.StatusCode != 200 {
		t.Fatal("/forgot route is throwing HTTP code " + strconv.Itoa(response.StatusCode) + " instead of 200")
	}
}

func TestRegistertGet(t *testing.T) {
	request := beetest.Get("/register")
	response, _ := request.Response()    
	defer response.Body.Close()
	
	if 	response.StatusCode != 200 {
		t.Fatal("/register route is throwing HTTP code " + strconv.Itoa(response.StatusCode) + " instead of 200")
	}
}