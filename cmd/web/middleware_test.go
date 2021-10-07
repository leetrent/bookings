package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var th testHandler
	handler := NoSurf(&th)

	switch v := handler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not an HTTP handler but is instead a %T", v))
	}
}
