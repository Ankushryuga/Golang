/***
Go HTTP client is part of Go's standard library and is used to send HTTP requests and receive responses from servers - Including APIs, websites and microservices.

It built around the net/http package, which provides:
- http.Client  - Customizable HTTP client.
- http.Get, http.Post  - Convenience functions.
- http.NewRequest  - for full control(headers, body etc)

| Task                     | Method                                     |
| ------------------------ | ------------------------------------------ |
| Simple GET               | `http.Get(url)`                            |
| Simple POST              | `http.Post(url, contentType, body)`        |
| Custom headers / timeout | `http.Client` + `http.NewRequest`          |
| JSON API requests        | Use `encoding/json` + custom `http.Client` |
| Field           | Purpose                              |
| --------------- | ------------------------------------ |
| `Timeout`       | Total request timeout (recommended!) |
| `Transport`     | Customize DNS, TLS, proxy, etc.      |
| `Jar`           | Manage cookies                       |
| `CheckRedirect` | Control redirect behavior            |

*/


package main
import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "time"
  )

//struct for parsing JSON response.
type HttpBinResponse struct{
  URL        string                    `json:"url"`
  Headers    map[string]interface{}    `json:"headers"`
  Json       map[string]interface{}    `json:"json"`
}

func main(){
  client := &http.Client{
    Timeout : 10*time.Second,
  }
  //--------Get Reequest..
  resp, err := client.Get("https://httpbin.org/get")
  if err!=nil{
    panic(err)
  }
  defer resp.Body.Close()

  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("Get response")
  fmt.Println(string(bodyBytes)


              // ----------- POST Request (with JSON) -----------
	fmt.Println("\n---- POST Request ----")

	payload := map[string]string{
		"name":  "Alice",
		"email": "alice@example.com",
	}
	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Custom-Header", "my-custom-value")

	resp2, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	var result HttpBinResponse
	if err := json.NewDecoder(resp2.Body).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Parsed POST Response:")
	fmt.Println("URL:     ", result.URL)
	fmt.Println("Headers: ", result.Headers)
	fmt.Println("JSON:    ", result.Json)
}



















