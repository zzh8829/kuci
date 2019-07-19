package main

import (
	// "fmt"
	// "net/http"
	// "os"
	// "strings"
	"bufio"
	"io"
	"log"
	"os/exec"
)

func shellCmd(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	reader := io.MultiReader(stdout, stderr)
	in := bufio.NewScanner(reader)

	outputs := ""
	for in.Scan() {
		outputs += in.Text()
		log.Printf(in.Text())
	}

	if err := in.Err(); err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return outputs, nil
}

func main() {

}
