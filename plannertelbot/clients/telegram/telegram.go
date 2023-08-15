package telegram

import (
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	host   string
	path   string
	Client http.Client
}

func New(host string, token string) Client {
	return Client{
		host:   host,
		path:   newBasePath(token),
		Client: http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))
}

func (c *Client) doRequest(query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   c.path + "methhod",
	}
}

func (c *Client) Message() {

}
