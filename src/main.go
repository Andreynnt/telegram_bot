package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:]{
		if !(strings.HasPrefix(url, "http://")){
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return
		}
		st := resp.Status
		fmt.Printf("%s\n", st)
		resp.Body.Close()
		st = resp.Status
		fmt.Printf("After closing: %s\n", st)
	}

}

