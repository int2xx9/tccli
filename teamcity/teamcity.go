package teamcity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrEmptyLocator = fmt.Errorf("locator is empty")
)

type Client struct {
	Server     string
	Token      string
	HttpClient *http.Client
}

func NewClient(server string, token string) *Client {
	return &Client{
		Server:     server,
		Token:      token,
		HttpClient: http.DefaultClient,
	}
}

func (c *Client) getUrl(path string, queries map[string]string) (*url.URL, error) {
	parseUrl, err := url.Parse(c.Server)
	if err != nil {
		return nil, err
	}

	splitPath := strings.SplitN(path, "?", 2)
	baseUrl := splitPath[0]
	queryString := ""
	if len(splitPath) > 1 {
		queryString = splitPath[1]
	}

	newUrl := parseUrl.JoinPath(baseUrl)
	newUrl.RawQuery = queryString

	query := newUrl.Query()
	for key, value := range queries {
		query.Add(key, value)
	}
	newUrl.RawQuery = query.Encode()

	return newUrl, nil
}

func (c *Client) get(path string, queries map[string]string, requestFunc func(*http.Request)) (*http.Response, error) {
	requestUrl, err := c.getUrl(path, queries)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", requestUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+c.Token)

	if requestFunc != nil {
		requestFunc(request)
	}

	resp, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, ResponseError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	return resp, nil
}

func (c *Client) getJson(path string, queries map[string]string, requestFunc func(*http.Request), v any) error {
	response, err := c.get(path, queries, func(r *http.Request) {
		r.Header.Add("Accept", "application/json")
		if requestFunc != nil {
			requestFunc(r)
		}
	})
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, v)
}

func (c *Client) getBinary(path string, queries map[string]string, requestFunc func(*http.Request)) (io.ReadCloser, error) {
	response, err := c.get(path, queries, func(r *http.Request) {
		if requestFunc != nil {
			requestFunc(r)
		}
	})
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
