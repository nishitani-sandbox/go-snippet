package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	cases := []struct {
		Title  string
		Wanted int64
	}{
		{
			Title:  "Zero",
			Wanted: 0,
		},
		{
			Title:  "MinusOne",
			Wanted: -1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Title, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			req.ContentLength = tc.Wanted
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, r.ContentLength)
			}(w, req)
			res := w.Result()
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(b, []byte(strconv.FormatInt(tc.Wanted, 10))) {
				t.Fatalf("want: %d, got: %s", tc.Wanted, b)
			}
		})
	}
}
