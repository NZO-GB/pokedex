package pokeapi

import (
	"net/http"
	"time"
	"fmt"
)

type Client struct {
	HTTPClient http.Client
}

func NewClient(timeout time.Duration) Client {
	fmt.Println("Created client")
	return Client{
		HTTPClient: http.Client {
			Timeout: timeout,
		},
	}
}