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

	if resp.StatusCode == 200 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("token"))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Wrong auth credentials"))
	}
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/token", token)

	http.ListenAndServe(":8888", nil)
}
