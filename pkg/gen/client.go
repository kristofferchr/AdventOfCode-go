package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	baseUrl string
	http.Client
}

func NewClient(url string) Client {
	return Client{
		url,
		http.Client{Timeout: time.Duration(3) * time.Second},
	}
}

type SubmitPayload struct {
	Level  string `json:"level"`
	Answer string `json:"answer"`
}

func (c *Client) SubmitAnswer(day int, year int, level string, answer string) (io.ReadCloser, error) {
	data := url.Values{}
	data.Set("answer", answer)
	data.Set("level", level)

	return c.Post(fmt.Sprintf("%d/day/%d/answer", year, day), data)
}

func (c *Client) Post(path string, data url.Values) (io.ReadCloser, error) {
	fullUrl := fmt.Sprintf("%s/%s", c.baseUrl, path)
	req, err := http.NewRequest("POST", fullUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Error with HTTP request: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return c.DoAndHandleRequest(req)
}

func (c *Client) Get(path string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.baseUrl, path), nil)
	if err != nil {
		log.Fatalf("Error with HTTP request: %s", err.Error())
	}

	return c.DoAndHandleRequest(req)
}

// DoAndHandleRequest invoke request
func (c Client) DoAndHandleRequest(req *http.Request) (io.ReadCloser, error) {
	session := &http.Cookie{
		Name:   "session",
		Value:  os.Getenv("AOC_SESSION"),
		MaxAge: 0,
	}
	req.AddCookie(session)
	req.Header.Set("User-Agent", "kristoffer.mob.chr@gmail.com")

	response, err := c.Do(req)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to execute request=%s", req.Method)
	}

	if response.StatusCode != 200 {
		err := errors.New(fmt.Sprintf("Warning: Received non 200 status code, code=%s", strconv.Itoa(response.StatusCode)))
		return nil, err
	}

	return response.Body, nil
}

func payloadToReader(payload interface{}) (io.Reader, error) {
	marshalledBody, err := json.Marshal(payload)
	return bytes.NewBuffer(marshalledBody), err
}
