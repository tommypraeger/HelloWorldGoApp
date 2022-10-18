package main

import (
    "fmt"
    "net/http"
    "log"
    "os/exec"
)

func main() {
    version, err := exec.Command("go", "version").Output()
    if err != nil {
        log.Fatal(err)
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, string(version))
    })
    fmt.Println("Starting the server on :3000...")
    http.ListenAndServe(":3000", nil)
}
