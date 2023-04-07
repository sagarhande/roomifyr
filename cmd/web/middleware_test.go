package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH MyHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//  Do nothing
	default:
		t.Error("type is not http.Handler, but is ", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH MyHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//  Do nothing
	default:
		t.Error("type is not http.Handler, but is ", v)
	}
}
