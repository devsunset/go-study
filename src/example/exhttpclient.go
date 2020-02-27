package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const TEST_SERVER_GET_URL = "http://httpbin.org/get"
const TEST_SERVER_POST_URL = "http://httpbin.org/post"

type ReqObj struct {
	A string
	B string
}

// HTTP Requst Example
func main() {
	reqGet()
	reqGetCustomHeader()
	reqPostBody()
	reqPostForm()
	reqPostJson()
	reqPostXml()
	reqPostCustomHeader()
}

func reqGet() {
	res, err := http.Get(TEST_SERVER_GET_URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func reqGetCustomHeader() {
	req, err := http.NewRequest("GET", TEST_SERVER_GET_URL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "NewRequest")
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}

func reqPostBody() {
	reqBody := bytes.NewBufferString("Post req body")
	res, err := http.Post(TEST_SERVER_POST_URL, "text/plain", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err == nil {
		fmt.Println(string(resBody))
	}
}

func reqPostForm() {
	res, err := http.PostForm(TEST_SERVER_POST_URL, url.Values{"A": {"A_VAL"}, "B": {"B_VAL"}})
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err == nil {
		fmt.Println(string(resBody))
	}
}

func reqPostJson() {
	reqobj := ReqObj{"A_VAL", "B_VAL"}
	reqobjbytes, _ := json.Marshal(reqobj)
	fmt.Println(string(reqobjbytes))
	buff := bytes.NewBuffer(reqobjbytes)

	res, err := http.Post(TEST_SERVER_POST_URL, "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err == nil {
		fmt.Println(string(resBody))
	}
}

func reqPostXml() {
	reqobj := ReqObj{"A_VAL", "B_VAL"}
	reqobjbytes, _ := xml.Marshal(reqobj)
	fmt.Println(string(reqobjbytes))
	buff := bytes.NewBuffer(reqobjbytes)

	res, err := http.Post(TEST_SERVER_POST_URL, "application/xml", buff)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err == nil {
		fmt.Println(string(resBody))
	}
}

func reqPostCustomHeader() {
	reqobj := ReqObj{"A_VAL", "B_VAL"}
	reqobjbytes, _ := xml.Marshal(reqobj)
	buff := bytes.NewBuffer(reqobjbytes)
	fmt.Println(string(reqobjbytes))

	req, err := http.NewRequest("POST", TEST_SERVER_POST_URL, buff)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("User-Agent", "NewRequest")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err == nil {
		fmt.Println(string(resBody))
	}
}
