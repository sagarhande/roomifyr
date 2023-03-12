package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
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
		return false
	} else {
		return true
	}

}

func (f *Form) Valid() bool {
	return len(f.Error) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Error.Add(field, "This field is required")
		}

	}
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Error.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

func (f *Form) ValidateEmail(email string) bool {
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Match email against regular expression
	if !regex.MatchString(email) {
		f.Error.Add("email", "invalid email address")
		return false
	}
	return true
}
