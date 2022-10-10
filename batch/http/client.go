package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

type (
	Client struct{}

	ClientInterface interface {
		Get(u string) ([]byte, error)
	}
)

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Get(u string) ([]byte, error) {
	res, err := http.Get(u)
	if err != nil {
		log.Printf("failed HTTP request err: %v\n", err.Error())
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed read response body err: %v\n", err.Error())
		return nil, err
	}
	return b, nil
}
