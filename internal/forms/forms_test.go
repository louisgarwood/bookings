package forms

import (
	"net/url"
	"testing"
)

func TestForms_Has(t *testing.T) {

	form := New(url.Values{})

	if form.Has("testField") {
		t.Error("The form should not have this field: testField")
	}

	postData := url.Values{}
	postData.Add("testField", "testValue")

	form = New(postData)
	if !form.Has("testField") {
		t.Error("The form should have this field: testField")
	}
}

func TestForms_Valid(t *testing.T) {
	form := New(url.Values{})

	if !form.Valid() {
		t.Error("this form should be valid")
	}

	form.Required("a")
	if form.Valid() {
		t.Error("this form should be invalid")
	}
}

func TestForms_Required(t *testing.T) {
	form := New(url.Values{})

	form.Required("testField")
	if len(form.Errors) == 0 {
		t.Error("this form should be invalid")
	}

	postData := url.Values{}
	postData.Add("testField", "testValue")
	form = New(postData)

	form.Required("testField")
	if len(form.Errors) > 0 {
		t.Error("The form should be valid")
	}
}

func TestForms_IsEmail(t *testing.T) {

	postData := url.Values{}
	postData.Add("email", "not-an-email")
	form := New(postData)

	form.IsEmail("email")
	if len(form.Errors) == 0 {
		t.Error("this email should have returned an error as it is not an email")
	}

	postData = url.Values{}
	postData.Add("email", "is@an.email")
	form = New(postData)

	form.IsEmail("email")
	if len(form.Errors) > 0 {
		t.Error("The email should have validated")
	}
}
