package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var config Config

func downloadUrl(url string) {
	fmt.Println(url)
	var buf bytes.Buffer
	x := strings.Split(config.Command, " ")
	x = append(x, url)
	c := exec.Command(x[0], x[1:]...)
	c.Stderr = &buf
	c.Dir = config.Directory
	if err := c.Run(); err != nil {
		log.Printf("Command failed for url %s: %v", url, err)
		log.Printf("Stderr: %s", buf.String())
	}
}

func main() {
	viper.SetConfigName("downloaderConfig")
	viper.AddConfigPath("/etc/downloader")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Viper threw err: %v", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Viper unmarshall err: %v", err)
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request to: %v", req.URL)
		tmp := strings.Split(req.URL.Path, "/")
		b64data := tmp[len(tmp)-1]
		url, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			log.Printf("Decoder for %v returned error %v", b64data, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		go downloadUrl(string(url))
		w.WriteHeader(http.StatusOK)
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
