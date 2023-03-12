package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Error errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		make(errors),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := f.Error.Get(field)

	if x == "" {
		f.Error.Add(field, "This field cann't be blank")
		return false
	} else {
		return true
	}

}

func (f *Form) Valid() bool {
	return len(f.Error) == 0
}
