package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func pingMotor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	ip := r.FormValue("ip")
	out, err := exec.Command("ping", "-c", "1", "-W", "5", ip).Output()

	if err != nil || strings.Contains(string(out), "100% packet loss") {
		fmt.Fprint(w, "false")
	} else {
		fmt.Fprint(w, "true")
	}
}

func main() {
	http.HandleFunc("/pingMotor", pingMotor)
	http.ListenAndServe(":8080", nil)
}
