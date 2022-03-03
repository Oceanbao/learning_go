package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const port string = "8085"

func main() {

	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/", helloWorldHandler)
	http.Handle("/welcome", inejctMsgID(helloWorldHandler))

	specialHandler := createSpecialHandler()
	http.Handle("/special", specialHandler)

	// CookieJar
	docHandler := http.HandlerFunc(docHandler)
	http.Handle("/doc", docHandler)

	docGetID := http.HandlerFunc(docGetID)
	http.Handle("/doc/id", docGetID)

	fmt.Println("Serving :8085...")
	http.ListenAndServe(":"+port, nil)
}

//HelloWorld hello world handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := ""
	if m := r.Context().Value(keyMsgID); m != nil {
		if value, ok := m.(string); ok {
			msgID = value
		}
	}
	w.Header().Add("msgId", msgID)
	w.Write([]byte("Hello, world"))
}

type ctxKey string

const (
	keyMsgID ctxKey = "msgID"
)

func inejctMsgID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgID := uuid.New().String()
		ctx := context.WithValue(r.Context(), keyMsgID, msgID)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func docHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Cookie in First API call")
	for _, c := range r.Cookies() {
		fmt.Println(c)
	}
	fmt.Println()
	w.WriteHeader(200)
	w.Write([]byte("Doc Get OK"))
	return
}

func docGetID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Cookie in Second API call")
	for _, c := range r.Cookies() {
		fmt.Println(c)
	}
	fmt.Println()
	w.WriteHeader(200)
	w.Write([]byte("Doc Get ID OK"))
	return
}

func createSpecialHandler() http.Handler {
	counter := 0
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cookies
		cookie := &http.Cookie{
			Name:  "counterCookie",
			Value: fmt.Sprint(counter),
		}
		r.AddCookie(cookie)
		fmt.Println("--- Cookies ---")
		for _, c := range r.Cookies() {
			fmt.Println(c)
		}
		// Header (verify at header client-side)
		w.Header().Set("counterHeader", fmt.Sprint(counter))
		// Context value

		counter++
	})
}
