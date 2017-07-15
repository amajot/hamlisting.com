package tests

import (
	beetest "github.com/astaxie/beego/testing"
	"testing"
	"strconv"
)


func TestGetHome(t *testing.T) {
	request := beetest.Get("/")
	response, _ := request.Response()    
	defer response.Body.Close()
	
	if 	response.StatusCode != 200 {
		t.Fatal("/ route is throwing HTTP code " + strconv.Itoa(response.StatusCode) + " instead of 200")
	}

}

