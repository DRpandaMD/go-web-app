//main_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(test *testing.T) {

	// instatiate a router we definde over in main.go
	router := newRouter()

	//create a new server from the httptest library
	//pass in the router we made to test
	mockServer := httptest.NewServer(router)

	//Now the mocker server will need to trigger a GET request
	// at the "/hello" route we just made
	// also protip: it already exposes the url for us
	resp, err := http.Get(mockServer.URL + "/hello")

	//we need to account for any error response
	if err != nil {
		test.Fatal(err)
	}

	// we want our status code we get back to be 200(ok)
	// the assmption here is that this will work properly
	// so if we get any other reponse other than ok
	// we need to be made aware of it
	if resp.StatusCode != http.StatusOK {
		test.Errorf("Status should be OK got %d", resp.StatusCode)
	}

	// now we got to get the response of the body
	// read it  and convert it to a string
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		test.Fatal(err)
	}

	//convert the bytes we go out of the body into a string
	responseString := string(bytes)
	expected := "Hello World!"

	// now to pass the test we want to make sure they mach
	// duh lol! :D

	if responseString != expected {
		test.Errorf("Response should be %s, got %s instead", expected, responseString)
	}

} //end test

func TestForNonExistentRoute(test *testing.T) {

	router := newRouter()
	mockServer := httptest.NewServer(router)

	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		test.Fatal(err)
	}

	// here we ant are status code to be 405
	// method not allowd
	if resp.StatusCode != http.StatusMethodNotAllowed {
		test.Errorf("Statues should be 405, got %d", resp.StatusCode)
	}

	//since the method POST is not allowed
	// our message body will be EMPTY

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		test.Fatal(err)
	}
	responseString := string(bytes)
	expected := ""

	if responseString != expected {
		test.Errorf("Response should be %s, but got %s", expected, responseString)
	}
} // end test

func TestStaticFileServer(test *testing.T) {
	router := newRouter()
	mockServer := httptest.NewServer(router)

	resp, err := http.Get(mockServer.URL + "/assests/")

	if err != nil {
		test.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		test.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	//now we can see if we pull HTML
	contentType := resp.Header.Get("Content-Type")
	expextedContentType := "text/html; charset=utf-8"

	if expextedContentType != contentType {
		test.Errorf("Wrong content type: expected %s got %s", expextedContentType, contentType)
	}

}
