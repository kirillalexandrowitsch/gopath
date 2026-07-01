package challenges

import "strings"

const testPrelude = `package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type doResult struct {
	response *http.Response
	err      error
}

type fakeDoer struct {
	results []doResult
	calls   int
}

func (d *fakeDoer) Do(req *http.Request) (*http.Response, error) {
	d.calls++
	if len(d.results) == 0 {
		return nil, errors.New("unexpected extra request")
	}

	result := d.results[0]
	d.results = d.results[1:]

	return result.response, result.err
}

type temporaryError struct{}

func (e temporaryError) Error() string {
	return "temporary error"
}

func (e temporaryError) Temporary() bool {
	return true
}

func setClient(doer Doer) {
	client = doer
}

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
`

func buildTestFile(tests []challengeTest) string {
	var builder strings.Builder
	builder.WriteString(testPrelude)

	for _, test := range tests {
		builder.WriteString("\n\n")
		builder.WriteString(test.Code)
	}

	return builder.String()
}
