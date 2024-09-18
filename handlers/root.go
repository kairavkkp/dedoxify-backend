// handlers/root.go
package handlers

import (
	"fmt"
	"net/http"
)

func PublicRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Public: Dedoxify Backend!")
}

func PrivateRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Private: Dedoxify Backend!")
}
