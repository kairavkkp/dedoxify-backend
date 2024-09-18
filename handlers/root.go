// handlers/root.go
package handlers

import (
	"fmt"
	"net/http"
)

// RootHandler is a temporary handler that responds with "Hello, World!"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dedoxify Backend!")
}
