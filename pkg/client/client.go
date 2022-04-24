package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
)

type MultipartFormClient struct {
	url    string
	writer *multipart.Writer
	body   *bytes.Buffer
	header http.Header
	err    error
}

func NewMultipartFormClient(url string) *MultipartFormClient {
	body := &bytes.Buffer{}
	return &MultipartFormClient{
		url:    url,
		writer: multipart.NewWriter(body),
		body:   body,
		header: http.Header{},
	}
}

func (c *MultipartFormClient) SetParam(key, value string) *MultipartFormClient {
	if c == nil || c.writer == nil {
		return nil
	}
	if c.err != nil {
		return c
	}

	c.err = c.writer.WriteField(key, value)
	return c
}

func (c *MultipartFormClient) SetFile(field, filename string, buff *bytes.Buffer) *MultipartFormClient {
	if c == nil || c.writer == nil {
		return nil
	}
	if c.err != nil {
		return c
	}

	part, err := c.writer.CreateFormFile(field, filename)
	if err != nil {
		c.err = err
		return c
	}

	_, err = io.Copy(part, buff)
	if err != nil {
		c.err = err
		return c
	}

	return c
}

func (c *MultipartFormClient) SetHeader(key, value string) *MultipartFormClient {
	if c == nil || c.header == nil {
		return nil
	}
	if c.err != nil {
		return c
	}

	c.header.Add(key, value)
	return c
}

func (c *MultipartFormClient) Do(result interface{}) error {
	if c == nil {
		return errors.New("empty client")
	}
	if c.err != nil {
		return c.err
	}

	if err := c.writer.Close(); err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, c.url, c.body)
	if err != nil {
		return err
	}

	request.Header = c.header
	request.Header.Add("Content-Type", c.writer.FormDataContentType())

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(result)
}
