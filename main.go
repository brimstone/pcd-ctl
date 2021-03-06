package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var cmds []*cobra.Command

var BASE_URL = "http://127.0.0.1:8080/"

var API_VERSION = "1"

func APIGet(path string) string {
	resp, err := http.Get(BASE_URL + API_VERSION + "/" + path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(body)
}

func APIPost(path string, payload string) {
	_, err := http.Post(BASE_URL+API_VERSION+"/"+path,
		"text/plain",
		bytes.NewBufferString(payload))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:  os.Args[0],
		Long: "Control application for PCD API",
	}

	for cmd := range cmds {
		rootCmd.AddCommand(cmds[cmd])
	}
	rootCmd.Execute()
}
