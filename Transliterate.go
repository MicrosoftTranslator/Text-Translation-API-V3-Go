package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"log"
    "net/http"
	"os"
    "strconv"
    "strings"
    "time"
)

func main() {
	if "" == os.Getenv("TRANSLATOR_TEXT_SUBSCRIPTION_KEY") {
		log.Fatal("Please set/export the environment variable TRANSLATOR_TEXT_SUBSCRIPTION_KEY.")
	}
	var subscriptionKey string = os.Getenv("TRANSLATOR_TEXT_SUBSCRIPTION_KEY")
	if "" == os.Getenv("TRANSLATOR_TEXT_ENDPOINT") {
		log.Fatal("Please set/export the environment variable TRANSLATOR_TEXT_ENDPOINT.")
	}
	var uriBase string = os.Getenv("TRANSLATOR_TEXT_ENDPOINT")

	const uriPath = "/transliterate?api-version=3.0"
	// Transliterate text in Japanese from Japanese script (i.e. Hiragana/Katakana/Kanji) to Latin script
	const params = "&language=ja&fromScript=jpan&toScript=latn"
    var uri string = uriBase + uriPath + params

    // Transliterate "good afternoon".
	const text = "こんにちは"

    r := strings.NewReader("[{\"Text\" : \"" + text + "\"}]")

    client := &http.Client{
        Timeout: time.Second * 2,
    }

    req, err := http.NewRequest("POST", uri, r)
    if err != nil {
        fmt.Printf("Error creating request: %v\n", err)
        return
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
    req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("Error on request: %v\n", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response body: %v\n", err)
        return
    }

    var f interface{}
    json.Unmarshal(body, &f)

    jsonFormatted, err := json.MarshalIndent(f, "", "  ")
    if err != nil {
        fmt.Printf("Error producing JSON: %v\n", err)
        return
    }
    fmt.Println(string(jsonFormatted))
}
