package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	params, err := url.ParseQuery(r.URL.RawQuery)
        mac := params["mac"][0]
	fmt.Printf("%s Looking for MAC : %s\n", time.Now().Format("2006-01-02 15:04:05"), mac)
	file, err := os.Open(mac)
	if err != nil {
		fmt.Fprintln(w, "#!ipxe")
		fmt.Fprintln(w, "echo no custom ipxe script found. booting locally...")
		//fmt.Fprintln(w, "exit")
		fmt.Fprintln(w, "sanboot --no-describe --drive 0x80")
	} else {
		fmt.Printf("%s Serving: %s\n", time.Now().Format("2006-01-02 15:04:05"), mac)
		io.Copy(w, file)
		os.Remove(mac)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":10080", nil)
}
