package main

import "os/exec"

func main() {
	exec.Command("xdg-open", "http://mail.google.com/").Run()
}
