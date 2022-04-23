package app

import (
	"encoding/json"
	"fmt"
	"gh/auth"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

func CreateClient() *http.Client {
	//Define a client for making the requests
	client := &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	return client
}

func (r *RequestHandler) createRequest() *http.Request {
	//Check validity of the URL
	endpoint := r.URL
	_, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal("The given URL cannot be parsed")
	}
	//Create a new HTTP request from the receiver data
	req, err := http.NewRequest(r.Method.String(), r.URL, r.Body)
	if err != nil {
		log.Fatal("Creating request", err)
	}
	return req
}

// SetAuthorization takes in a http request from the receiver RequestHandler and assigns appropriate headers to the http request
func (r *RequestHandler) setAuthorization(req *http.Request) *http.Request {
	loginObjectName := r.Login
	if loginObjectName == "" {
		return req
	}
	secretsString, err := auth.GetLoginSecrets(loginObjectName)
	if err != nil {
		log.Fatal("Getting secrets from AWS", err)
	}

	var jsonMap map[string]string
	json.Unmarshal([]byte(secretsString), &jsonMap)
	switch jsonMap["auth"] {
	case "Basic":
		baa := auth.BasicAuthParameter{
			Username: jsonMap["username"],
			Password: jsonMap["password"],
		}
		req = baa.SetAuthorizationHeader(req)
	case "Bearer":
		bea := auth.BearerAuthParameter{
			Token: jsonMap["token"],
		}
		req = bea.SetAuthorizationHeader(req)
		// TODO: add additional use cases handling here
	default:
		log.Fatal("The auth type cannot be determined or unimplemented")
	}
	return req
}

func (r *RequestHandler) setHeaders(req *http.Request) *http.Request {
	mapHeaders := r.Headers
	for key, value := range mapHeaders {
		req.Header.Add(key, value)
	}
	return req
}

func (r *RequestHandler) PerformRequest() {
	var resp *http.Response

	//Create http client and construct a request from the receiver data
	client := CreateClient()
	req := r.createRequest()
	req = r.setHeaders(req)
	req = r.setAuthorization(req)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
	fmt.Printf("Body : %s", body)
}
