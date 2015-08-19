package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	switch {
	case flag.Arg(0) == "hostname":
		if len(flag.Args()) == 2 {
			_, err := http.Post("http://127.0.0.1:8080/hostname", "text/plain", bytes.NewBufferString(flag.Arg(1)))
			if err != nil {
				log.Fatal(err.Error())
			}
		} else if len(flag.Args()) == 1 {
			// handle sets
			// handle gets
			resp, err := http.Get("http://127.0.0.1:8080/hostname")
			if err != nil {
				log.Fatal(err.Error())
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Printf("%s", string(body))
		} else {
			log.Fatal("Usage: " + os.Args[0] + " hostname <new hostname>")
		}
	default:
		log.Fatal("Usage: " + os.Args[0] + " <command>\nhostname: set hostname")
	}

}
