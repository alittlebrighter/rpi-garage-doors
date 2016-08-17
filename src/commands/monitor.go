package commands

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

// Monitor captures an image from an attached webcam and streams the jpg result
func Monitor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "HTTP Method "+r.Method+" not valid for this endpoint.")
		return
	}

	log.Println("Fetching image from webcam.")
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error parsing query parameters.")
		return
	}

	args := []string{"-q"}
	for option, value := range r.Form {
		if strings.HasPrefix(option, "o_") {
			args = append(args, "-s", fmt.Sprintf("%s=%s", option[2:], value[0]))
		}
	}

	cmd := exec.Command("fswebcam", append(args, "-")...)
	cmd.Stdout = w
	err = cmd.Run()

	if err != nil {
		fmt.Fprintf(w, "Error fetching image.")
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
}
