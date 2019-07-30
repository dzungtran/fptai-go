package fptai_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultApiVersion = "v3"
	ApiBaseUrl        = "https://v3-api.fpt.ai/api"
)

// API docs: https://docs.fpt.ai/#general
type Client struct {
	BotToken   string
	Version    string
	ApiBaseUrl string

	// private access
	httpClient *http.Client
}

type CommonResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SetHttpClient - Allows to use custom http.Client
func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// SetBotToken - Changes Bot Token.\n You can get bot token follows instruction: https://docs.fpt.ai/#authentication
func (c *Client) SetBotToken(token string) {
	c.BotToken = token
}

// NewClient - returns fpt.ai client
func NewClient(token string) *Client {
	defaultHttpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &Client{
		BotToken:   token,
		Version:    DefaultApiVersion,
		httpClient: defaultHttpClient,
		ApiBaseUrl: ApiBaseUrl,
	}
}

func (c *Client) getAuthHeader() string {
	return fmt.Sprintf("Bearer %s", c.BotToken)
}

func (c *Client) getApiEndpoint(apiPath string) string {
	return fmt.Sprintf("%s/%s/%s", c.ApiBaseUrl, c.Version, apiPath)
}

func (c *Client) request(method, apiPath string, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.getApiEndpoint(apiPath), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.getAuthHeader())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		defer resp.Body.Close()

		var e interface{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&e)
		if err != nil {
			return nil, fmt.Errorf("unable to decode error message: %s", err.Error())
		}

		return nil, fmt.Errorf("error while make api request. resp: %v", e)
	}
	return resp.Body, nil
}
