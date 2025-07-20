package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Struct for POST /sum
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}
type SumResponse struct {
	Sum int `json:"sum"`
}

func main() {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/sum", sumHandler)
	mux.HandleFunc("/data", dataHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      logMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

// --------- HANDLERS -----------

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Fprintln(w, "Hello, World!")
	case <-ctx.Done():
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
	}
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req SumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// simulate processing delay
	select {
	case <-time.After(1 * time.Second):
		resp := SumResponse{Sum: req.A + req.B}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	case <-ctx.Done():
		http.Error(w, "Request timeout", http.StatusRequestTimeout)
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	resultCh := make(chan string)

	// simulate a slow data fetch (random delay)
	go func() {
		delay := time.Duration(rand.Intn(5)) * time.Second
		time.Sleep(delay)
		resultCh <- fmt.Sprintf("Data fetched after %v", delay)
	}()

	select {
	case res := <-resultCh:
		fmt.Fprintln(w, res)
	case <-ctx.Done():
		http.Error(w, "Data fetch timed out", http.StatusRequestTimeout)
	}
}

// ---------- Middleware -------------

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := fmt.Sprintf("%d", rand.Intn(1000000))
		ctx := context.WithValue(r.Context(), "reqID", reqID)
		log.Printf("→ %s %s [reqID=%s]", r.Method, r.URL.Path, reqID)

		next.ServeHTTP(w, r.WithContext(ctx))

		log.Printf("← %s completed in %v [reqID=%s]", r.URL.Path, time.Since(start), reqID)
	})
}
