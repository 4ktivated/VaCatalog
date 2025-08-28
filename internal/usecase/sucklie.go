package usecase

import (
	"fmt"
	"net/http"
)

func SuckLie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "suck lie")
}