package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	go mustRun(`C:\code\caddy\caddy.exe`, `C:\code\caddy`, "run", "--config", "Caddyfile")
	go mustRun("python", `C:\code\localFile\src`, "-m", "http.server", "3005")
	go mustRun("python", `C:\code\secretsHolder`, "-m", "http.server", "3004")
	go mustRun("npm", `C:\code\ACTMail-admin`, "run", "dev", "--", "-p", "3000")
	go mustRun("npm", `C:\code\ACTMail-client`, "run", "dev", "--", "-p", "3001")
	mustRun("npm", `C:\code\ethanc`, "run", "dev", "--", "-p", "3002")
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
