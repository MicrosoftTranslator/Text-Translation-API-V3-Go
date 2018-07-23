package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
    "time"
)

func main() {
    // Replace the subscriptionKey string value with your valid subscription key
    const subscriptionKey = "<Subscription Key>"

	const uriBase = "https://api.cognitive.microsofttranslator.com"
	const uriPath = "/languages?api-version=3.0"

    const uri = uriBase + uriPath

    client := &http.Client{
        Timeout: time.Second * 2,
    }

    req, err := http.NewRequest("GET", uri, nil)
    if err != nil {
        fmt.Printf("Error creating request: %v\n", err)
        return
    }

    req.Header.Add("Content-Type", "application/json")
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