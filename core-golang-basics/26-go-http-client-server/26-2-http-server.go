/**
Creating an HTTP server in Go is simple and powerful thanks to the standard net/http package. You can handle routes, return JSON, serve static files, or even build full REST APIs — all without third-party libraries.
*/

package main
import (
  "fmt"
  "net/http"
  "encoding/json"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
  fmt.Println(w, "Hello, world!")
}


//Handle Get and Post differently:
func handler(w http.ResponseWriter, r *http.Request){
  switch r.Method{
    case "GET":
      fmt.Println(w, "This is a Get Request")
    case "POST":
      fmt.Println(w, "This is a Post Request")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
  }
}


//return json response:
func jsonHandler(w http.ResposneWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  response := map[string]string{"message":"Hello JSON"}
  json.NewEncoder(w).Encode(response)
}



func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("🚀 Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
///Route handling::
  http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "About Page")
  })

  

}
