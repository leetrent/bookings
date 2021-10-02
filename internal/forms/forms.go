package forms

import (
	"net/http"
	"net/url"
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
		f.Errors.Add(field, "This field cannot be blank.")
		return false
	}
	return true
}

// Valid return true if there are no form errors, otherwise true
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
