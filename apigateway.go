package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func token(w http.ResponseWriter, r *http.Request) {

	authProviderURL, exists := os.LookupEnv("AUTH_PROVIDER_URL")

	if !exists {
		log.Fatal("AUTH_PROVIDER_URL is not defined")
	}

	resp, err := http.Get(authProviderURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(body))

	w.Header().Set("Content-Type", "application/json")

	if resp.StatusCode == 200 {
		w.WriteHeader(http.StatusOK)
		// todo: generate token
		w.Write([]byte("{\"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c\"}"))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("{\"error\": \"Wrong auth credentials\""))
	}
}

func resourse1Handler(w http.ResponseWriter, r *http.Request) {

	log.Print("Handling request to resourse 1")

	resourse1Url, exists := os.LookupEnv("RESOURSE_1_URL")

	if !exists {
		log.Fatal("RESOURSE_1_URL is not defined")
	}

	resp, err := http.Get(resourse1Url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(body))
}

func resourse2Handler(w http.ResponseWriter, r *http.Request) {

	log.Print("Handling request to resourse 2")

	resourse2Url, exists := os.LookupEnv("RESOURSE_2_URL")

	if !exists {
		log.Fatal("RESOURSE_2_URL is not defined")
	}

	resp, err := http.Get(resourse2Url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(body))
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/token", token)
	http.HandleFunc("/resourse1", resourse1Handler)
	http.HandleFunc("/resourse2", resourse2Handler)

	log.Print("About to start ApiGateway  http://127.0.0.1:8888")

	err = http.ListenAndServe(":8888", nil)

	if err != nil {
		log.Fatal(err)
	}
}
