package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Checks to see if form field is in POST request and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	//return r.Form.Get(field) != ""
	value := r.Form.Get(field)
	if value == "" {
		return false
	}
	return true
}

// Valid return true if there are no form errors, otherwise true
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

// MinLength check for minimum string length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	value := r.Form.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("%s must be at least %d characters long.", field, length))
		return false
	}
	return true
}
