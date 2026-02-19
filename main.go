package main

import (
	"os"
	"os/exec"
)

func main() {
	go run("npm", "C:\\code\\localFile", "start")                          //example of running node server
	run("python3", "C:\\code\\secretsHolder", "-m", "http.server", "9000") //example of running html (using python http.server)
}

func run(cmd, dir string, args ...string) {
	c := exec.Command(cmd, args...) //builds the command with the provided arguments
	c.Dir = dir                     //sets the working directory for the command
	c.Stdout = os.Stdout            //reroutes command output
	c.Stderr = os.Stderr            //reroutes command errors
	c.Run()                         //runs the command
}
