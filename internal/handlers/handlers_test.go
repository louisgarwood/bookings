package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	httpMethod         string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"captains-quarters", "/captains", "GET", []postData{}, http.StatusOK},
	{"oceanview", "/oceanview", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"reservation-summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},

	{"search-availability-post", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-11-20"},
		{key: "start", value: "2020-11-22"},
	}, http.StatusOK},
	{"search-availability-json-post", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-11-20"},
		{key: "start", value: "2020-11-22"},
	}, http.StatusOK},
	{"make-reservation-post", "/make-reservation", "POST", []postData{
		{key: "first-name", value: "John"},
		{key: "last-name", value: "Doe"},
		{key: "email-address", value: "jdoe@msn.com"},
		{key: "phone-number", value: "444-444-4444"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, endpoint := range theTests {
		if endpoint.httpMethod == "GET" {
			response, error := testServer.Client().Get(testServer.URL + endpoint.url)
			if error != nil {
				t.Log(error)
				t.Fatal(error)
			}

			if response.StatusCode != endpoint.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", endpoint.name, endpoint.expectedStatusCode, response.StatusCode)
			}
		} else {

			values := url.Values{}

			for _, value := range endpoint.params {
				values.Add(value.key, value.value)
			}

			response, error := testServer.Client().PostForm(testServer.URL+endpoint.url, values)
			if error != nil {
				t.Log(error)
				t.Fatal(error)
			}

			if response.StatusCode != endpoint.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", endpoint.name, endpoint.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
