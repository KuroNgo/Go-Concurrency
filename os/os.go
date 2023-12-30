package _os

import (
	"os"
	"os/exec"
)

func OperatingSystem() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
