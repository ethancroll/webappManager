package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	go mustRun("python", `C:\code\localFile\src`, "-m", "http.server", "3500")
	go mustRun("npm", `C:\code\amiworking`, "run", "dev", "--", "-p", "3000")
	mustRun("python", `C:\code\secretsHolder`, "-m", "http.server", "9000")
}

func mustRun(cmd, dir string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed: %s %v (dir=%s): %v\n", cmd, args, dir, err)
	}
}
