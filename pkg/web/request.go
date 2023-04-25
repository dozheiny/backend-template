// Package web provides functionality for making HTTP requests and binding JSON output.
package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Base is a struct for using web requests.
type Base struct {
	Host      string
	Params    []Params
	Headers   []Params
	Method    string
	Body      interface{}
	Output    interface{}
	useOutput bool
}

type Params struct {
	Key   string
	Value string
}

// Init is an initializer to use all functions for sending requests.
func Init() Base {
	return Base{Method: http.MethodGet}
}

// SetHost will set Host on Base.Host and return Base.
func (b Base) SetHost(host string) Base {
	b.Host = host
	return b
}

// SetParam gets key and value and set into Base.Params,
// in Do function SetParam use for set query params to send HTTP request.
func (b Base) SetParam(key, value string) Base {
	p := Params{
		Key:   key,
		Value: value,
	}

	b.Params = append(b.Params, p)

	return b
}

// SetHeader given key and value as string and set into Base.Headers,
// in Do function headers will be set for HTTP request.
func (b Base) SetHeader(key, value string) Base {
	h := Params{
		Key:   key,
		Value: value,
	}

	b.Headers = append(b.Headers, h)

	return b
}

// SetMethod gets method and set method for HTTP requests,
// method will store in Base.Method;
// HINT: it's good to use http package for set methods;
// Default is GET method.
func (b Base) SetMethod(method string) Base {
	b.Method = method
	return b
}

// SetOutput gets interface called output and set it into Base.Output.
func (b Base) SetOutput(output interface{}) Base {
	b.Output = output
	b.useOutput = true
	return b
}

func (b Base) SetBody(input interface{}) Base {
	b.Body = input
	return b
}

// Do sends HTTP request by all variables sets on Base struct.
func (b Base) Do() (interface{}, error) {

	query := url.Values{}
	var client http.Client
	var req *http.Request
	var inputBody io.Reader

	if len(b.Params) != 0 {
		for number := range b.Params {
			query.Add(b.Params[number].Key, b.Params[number].Value)
		}
	}

	if b.Body != nil {
		reqBody, err := json.Marshal(b.Body)
		if err != nil {
			return nil, err
		}

		inputBody = bytes.NewBuffer(reqBody)
	}

	b.Host = fmt.Sprintf("%s?%s", b.Host, query.Encode())

	req, err := http.NewRequest(b.Method, b.Host, inputBody)
	if err != nil {
		return nil, err
	}

	if len(b.Headers) != 0 {
		for number := range b.Headers {
			req.Header.Set(b.Headers[number].Key, b.Headers[number].Value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if b.useOutput {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(body, b.Output); err != nil {
			return nil, err
		}
	}

	return b.Output, nil
}
