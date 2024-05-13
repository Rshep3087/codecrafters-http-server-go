package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHTTPRequest(t *testing.T) {
	request := "GET /echo/hello HTTP/1.1\r\nHost: localhost\r\n\r\n"
	expectedStartLine := StartLine{
		Method:  "GET",
		Path:    "/echo/hello",
		Version: "HTTP/1.1",
	}
	expectedHeaders := map[string]string{
		"Host": "localhost",
	}

	r := strings.NewReader(request)
	httpRequest, err := ParseHTTPRequest(bufio.NewReader(r))

	assert.NoError(t, err)
	assert.Equal(t, expectedStartLine, httpRequest.StartLine)
	assert.Equal(t, expectedHeaders, httpRequest.Headers)
}

func TestReadStartLine(t *testing.T) {
	startLineStr := "GET /echo/hello HTTP/1.1\r\n"
	expectedStartLine := StartLine{
		Method:  "GET",
		Path:    "/echo/hello",
		Version: "HTTP/1.1",
	}

	r := strings.NewReader(startLineStr)
	startLine, err := readStartLine(bufio.NewReader(r))

	assert.NoError(t, err)
	assert.Equal(t, expectedStartLine, startLine)
}

func TestReadHeaders(t *testing.T) {
	headersStr := "Host: localhost\r\nContent-Type: application/json\r\n\r\n"
	expectedHeaders := map[string][]string{
		"Host":         {"localhost"},
		"Content-Type": {"application/json"},
	}

	r := strings.NewReader(headersStr)
	headers, err := readHeaders(bufio.NewReader(r))

	assert.NoError(t, err)
	assert.Equal(t, expectedHeaders, headers)

	headersStrWithMultipleValues := "Host: localhost\r\nContent-Type: application/json, text/plain\r\n\r\n"
	expectedHeadersWithMultipleValues := map[string][]string{
		"Host":         {"localhost"},
		"Content-Type": {"application/json", "text/plain"},
	}

	r = strings.NewReader(headersStrWithMultipleValues)
	headers, err = readHeaders(bufio.NewReader(r))

	assert.NoError(t, err)
	assert.Equal(t, expectedHeadersWithMultipleValues, headers)
}

func TestHTTPResponseBytes(t *testing.T) {
	response := HTTPResponse{
		StatusLine: "HTTP/1.1 200 OK",
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": "5",
		},
		Body: "Hello",
	}
	expectedBytes := []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 5\r\n\r\nHello")

	bytes := response.Bytes()

	assert.Equal(t, expectedBytes, bytes)
}
