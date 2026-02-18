package main

import (
	"os"
	"os/exec"
)

func main() {
	go run("npm", "/home/ethan/localFile", "start")
	run("python3", "/home/ethan/secretsHolder", "-m", "http.server", "9000")
}

func run(cmd, dir string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}
