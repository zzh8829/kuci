package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func dockerImageExists(tag string) bool {
	splits := strings.Split(tag, ":")
	image, tag := splits[0], splits[1]
	resp, err := http.Get(fmt.Sprintf("https://index.docker.io/v1/repositories/%v/tags/%v", image, tag))
	if err == nil && resp.StatusCode == 200 {
		return true
	}
	return false
}

func main() {
	tag := "zihao/play:kubox"
	fmt.Printf("%v", dockerImageExists(tag))

	os.Setenv("GIT_SHA", "LOL")
	fmt.Printf("%v", os.ExpandEnv("${GIT_SHA}"))
}
