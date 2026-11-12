package installer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"mpm/internal/aur"
)

func Run(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Install(name string) error {
	check := exec.Command("pacman", "-Si", name)

	if check.Run() == nil {
		fmt.Println("Installing from official repository...")
		return Run("", "sudo", "pacman", "-S", "--", name)
	}

	packages, err := aur.Info(name)
	if err != nil {
		return err
	}

	if len(packages) == 0 {
		return fmt.Errorf("package %s not found", name)
	}

	return InstallAUR(name, packages[0].PackageBase)
}

func InstallAUR(name, base string) error {
	if os.Geteuid() == 0 {
		return fmt.Errorf("do not run MPM as root")
	}

	if base == "" {
		base = name
	}

	temp, err := os.MkdirTemp("", "mpm-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(temp)

	repo := "https://aur.archlinux.org/" + base + ".git"

	fmt.Println("downloading ", name)
	err = Run(temp, "git", "clone", "--", repo, base)
	if err != nil {
		return err
	}
	dir := filepath.Join(temp, base)
	content, err := os.ReadFile(filepath.Join(dir, "PKGBUILD"))
	if err != nil {
		return err
	}

	fmt.Println("\nPKGBUILD")
	fmt.Println(string(content))
	fmt.Println("")

	fmt.Println("review other build files in:", dir)
	fmt.Print("build and install? [y/N]: ")

	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	if strings.TrimSpace(strings.ToLower(answer)) != "y" {
		return fmt.Errorf("installation cancelled")
	}
	return Run(dir, "makepkg", "-si")
}
