/**
In production go servers(e.g., using net/http), you must recover from panics to prevent crashing the whole app.
*/

package main
import (
  "fmt"
  "log"
  "net/http"
  )

func recoveryMiddleware(next http.Handler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    defer func(){
      if err:= recover(); err!=nil{
          log.Printf("[panic recovered] %", err )
        http.Error(w "Internal server error", http.StatusInternalServerError)
      }
    }()
    next.ServeHTTP(w,r)
  })
}

func handlerThatPanics(w http.ResponseWriter, r *http.Request){
  panic("oh no something went wrong")
}

func main(){
  mux := http.NewServeMux()
  mux.HandleFunc("/", handlerThatPanics)

  //wrap the mux with recovery middleware.
  http.ListenAndServe(":8080", recoveryMiddleware(mux))
}
//Visit http://localhost:8080 and see how the panic is recovered and logged.

