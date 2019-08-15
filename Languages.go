package main

import (    
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
	"os"
)

func main() {
   getLanguages()
}

func getLanguages() {
	if "" == os.Getenv("TRANSLATOR_TEXT_ENDPOINT") {
		log.Fatal("Please set/export the environment variable TRANSLATOR_TEXT_ENDPOINT.")
	}
	var uriBase string = os.Getenv("TRANSLATOR_TEXT_ENDPOINT")
	const uriPath = "/languages?api-version=3.0"
    var uri string = uriBase + uriPath

    u, _ := url.Parse(uri)
    q := u.Query()
    u.RawQuery = q.Encode()
    
    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        log.Fatal(err)
    }    
    req.Header.Add("Content-Type", "application/json")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    var result interface{}
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
      log.Fatal(err)
    }

    prettyJSON, _ := json.MarshalIndent(result, "", "  ")
    fmt.Printf("%s\n", prettyJSON)
}
