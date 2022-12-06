package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mockHandler myHandler
	handler := NoSurf(&mockHandler)

	switch handlerType := handler.(type) {
	case http.Handler:
		//success
	default:
		t.Errorf("type is not an http.Handler, but is: %T", handlerType)
	}
}

func TestSessionLoad(t *testing.T) {
	var mockHandler myHandler
	handler := SessionLoad(&mockHandler)

	switch handlerType := handler.(type) {
	case http.Handler:
		//success
	default:
		t.Errorf("type is not an http.Handler, but is: %T", handlerType)
	}
}
