package str

const BLANK = `package main

import "fmt"
	
func main() {
	fmt.Println("Hello, World!")
}
`

const WEB = `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

` // this will need to create a middleware folder, a models folder, and a helper folder

const GRPC = `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
`

const CLI = `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// Import <project>/cmd
	// Call the Execute function on the import
}
` // this will need to create a cmd folder, a helpers folder, and a models folder

const MW = `
package middleware

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

// Middleware for logging requests
// Feel free to adjust as neccesary
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		log.Println(r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Middleware for rate limiting
func RateLimiterMiddleware(next http.Handler, limiter *rate.Limiter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		}
	})
}
`

const MDL = `
package models

// Feel free to add your models here
`

const CLIROOT = `
package cmd
import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "",
	Version: version,
	Short:   "",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
`
