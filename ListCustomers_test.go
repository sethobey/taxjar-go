package taxjar_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
	mock "github.com/taxjar/taxjar-go/mocks"
)

func TestListCustomersOnSuccess(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/customers", req.URL.String())
		return &http.Response{
			StatusCode: 201,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(mock.ListCustomersJSON))),
			Header:     make(http.Header),
		}
	})

	res, err := client.ListCustomers()

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, *mock.ListCustomers, *res)
}

func TestListCustomersOnError(t *testing.T) {
	client := mock.Client(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://api.taxjar.com/v2/customers", req.URL.String())
		return &http.Response{
			StatusCode: 401,
			Header:     make(http.Header),
		}
	})

	res, err := client.ListCustomers()

	assert.Nil(t, res)
	assert.NotNil(t, err)
}
