package pacman

import (
	"os"
	"os/exec"
)

func Run(args ...string) error {
	cmd := exec.Command("sudo", append([]string{"pacman"}, args...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Search(name string) error {
	cmd := exec.Command("pacman", "-Ss", name)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Remove(name string) error {
	return Run("-Rns", "--", name)
}

func Update() error {
	return Run("-Syu")
}
