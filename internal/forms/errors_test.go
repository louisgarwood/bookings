package forms

import "testing"

func TestErrors_Add(t *testing.T) {
	var e = make(errors)
	e.Add("testField", "testField error")
	if e["testField"][0] != "testField error" {
		t.Error("Add failed to add the error")
	}
}

func TestErrors_Get(t *testing.T) {
	var e = make(errors)
	e.Add("testField", "testField error")
	if e.Get("testField") != "testField error" {
		t.Error("Get failed to get the error from the map")
	}
}

func TestErrors_GetEmpty(t *testing.T) {
	var e = make(errors)
	if e.Get("testField") != "" {
		t.Error("Get failed to return empty string when the error map is empty")
	}
}
