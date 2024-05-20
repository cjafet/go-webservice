package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cjafet/go-webservice/models"
	"golang.org/x/exp/slices"
)

type mockResponseWriter struct {
	bytes.Buffer
}

func (mrw mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (mrw mockResponseWriter) WriteHeader(statusCode int) {}

func (mrw mockResponseWriter) Write(b []byte) (int, error) {
	return mrw.Write(b)
}

func (mrw mockResponseWriter) getData() []byte {
	return mrw.Bytes()
}

func TestUserController_Handler(t *testing.T) {
	users := []*models.User{
		{
			ID:        1,
			FirstName: "Carlos",
			LastName:  "Jafet",
		},
		{
			ID:        2,
			FirstName: "Carlos",
			LastName:  "Neto",
		},
	}

	req := httptest.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()

	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	uc := newUserController()
	uc.ServeHTTP(rec, req)

	res := rec.Result()

	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	got = got[:len(got)-1]

	if !slices.Equal(got, expect) {
		t.Errorf("got \n%s, want \n%s", got, expect)
	}

}

func TestUserController_server(t *testing.T) {
	users := []*models.User{
		{
			ID:        1,
			FirstName: "Carlos",
			LastName:  "Jafet",
		},
		{
			ID:        2,
			FirstName: "Carlos",
			LastName:  "Neto",
		},
	}

	// Marshal user slice into json
	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	uc := newUserController()
	// request sent to this server will be routed to the ServeHTTP handler
	s := httptest.NewServer(http.HandlerFunc(uc.ServeHTTP))

	// sending requests to the server with an http client
	c := s.Client()

	// the url will be interpreted by the handler
	res, err := c.Get(s.URL + "/users")
	if err != nil {
		t.Fail()
	}

	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fail()
	}

	got = got[:len(got)-1]

	if !slices.Equal(got, expect) {
		t.Errorf("got \n%s, want \n%s", got, expect)
	}

}

func BenchmarkHandler(b *testing.B) {
	uc := newUserController()
	// request sent to this server will be routed to the ServeHTTP handler
	s := httptest.NewServer(http.HandlerFunc(uc.ServeHTTP))
	// sending requests to the server with an http client
	c := s.Client()
	// I do not want to benchmark my setup code/logic
	// This will ignore our setup by resetting the timer to 0
	b.ResetTimer()
	// Shows how goes get into the number of iteration that it can run in one second, its default value
	// N will receive different values depending on the duration we pass for it. default is 1s.
	// If we pass 10s we gonna need a higher number of iterations to achieve test duration
	// so go test runner will increase n until we achieve the test duration we ask for 
	b.Log(b.N)
	for i := 0; i < b.N; i++ {
		// the url will be interpreted by the handler
		_, err := c.Get(s.URL + "/users")
		// b.Log(res.Body)
		if err != nil {
			b.Fail()
		}
	}

}
