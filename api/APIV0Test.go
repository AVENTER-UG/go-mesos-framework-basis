package api

import (
	"fmt"
	"net/http"
)

// V0Test Handle
func V0Test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cmd := r.Form["cmd"][0]
	config.CommandChan <- cmd
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Scheduled: %s", cmd)
}
