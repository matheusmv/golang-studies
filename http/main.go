package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fmt.Println("***GET***")
	getRequest()

	fmt.Println("***POST***")
	postRequest()

	fmt.Println("***POST FORM***")
	formRequest()

	fmt.Println("***CUSTOM REQUEST***")
	customRequest()

	fmt.Println("***FILE UPLOAD***")
	fileUploadRequest()
}

func getRequest() {

	resp, err := http.Get("https://httpbin.org/get")

	checkError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	checkError(err)

	log.Println(string(body))
}

func postRequest() {

	requestBody, err := json.Marshal(map[string]string{
		"name":  "FirstName LastName",
		"email": "user@email.com",
	})

	checkError(err)

	resp, err := http.Post(
		"https://httpbin.org/post",
		"application/json",
		bytes.NewBuffer(requestBody),
	)

	checkError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	checkError(err)

	log.Println(string(body))
}

func formRequest() {

	formData := url.Values{
		"name": {"FirstName"},
	}

	resp, err := http.PostForm("http://httpbin.org/post", formData)

	checkError(err)

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}

func customRequest() {

	requestBody, err := json.Marshal(map[string]string{
		"name":  "FirstName LastName",
		"email": "user@email.com",
	})

	checkError(err)

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(
		"POST",
		"http://httpbin.org/post",
		bytes.NewBuffer(requestBody),
	)
	request.Header.Set("Content-type", "application/json")

	checkError(err)

	resp, err := client.Do(request)

	checkError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	checkError(err)

	log.Println(string(body))
}

func fileUploadRequest() {
	fileName := "file.txt"
	currentPath, err := os.Getwd()

	checkError(err)

	filePath := path.Join(currentPath, fileName)

	file, err := os.Open(filePath)

	checkError(err)

	defer file.Close()

	var requestBody bytes.Buffer

	multiPartWriter := multipart.NewWriter(&requestBody)

	fieldWriter, err := multiPartWriter.CreateFormFile("file_field", fileName)

	checkError(err)

	_, err = fieldWriter.Write([]byte("Value"))

	checkError(err)

	multiPartWriter.Close()

	req, err := http.NewRequest(
		"POST",
		"http://httpbin.org/post",
		&requestBody,
	)

	checkError(err)

	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)

	checkError(err)

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	log.Println(result)
}
