package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageTestData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "no error rendering non-exisiting go template when expected"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering non-exisiting jet template when expected"},
	{"invalid_renderer_engine", "foo", "home", true, "no error rendering with non-exisiting template engine"},
}

func TestRender_Page(t *testing.T) {
	for _, e := range pageTestData {
		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()
		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"
		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("get", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata/"

	err = testRenderer.GoPage(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering go page", err)
	}

}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("get", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata/"

	err = testRenderer.GoPage(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering jet page", err)
	}

}
