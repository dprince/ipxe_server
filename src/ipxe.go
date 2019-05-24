package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    addr_file := strings.Split(r.RemoteAddr, ":")[0]
    file, err := os.Open(addr_file)
    if err != nil {
        fmt.Fprintln(w, "#!ipxe")
        fmt.Fprintln(w, "exit")
    } else {
        fmt.Printf("%s Serving: %s\n", time.Now().Format("2006-01-02 15:04:05"), addr_file);
        io.Copy(w, file)
        os.Remove(addr_file)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":10080", nil)
}
