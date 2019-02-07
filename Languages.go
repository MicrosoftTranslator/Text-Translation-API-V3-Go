
package main

import (    
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"    
)

func main() {
   getLanguages()
}

func getLanguages() {
    u, _ := url.Parse("https://api.cognitive.microsofttranslator.com/languages")
    q := u.Query()
    q.Add("api-version", "3.0")
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
