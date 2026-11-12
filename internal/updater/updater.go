package updater

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"mpm/internal/aur"
	"mpm/internal/installer"
	"mpm/internal/pacman"
)
func Foreign() (map[string]string, error) {
	output, err := exec.Command("pacman", "-Qm").Output()
	if err != nil {
		return nil, err
	}

	packages := make(map[string]string)

	for _, line := range strings.Split(string(output), "\n") {
		fields := strings.Fields(line)

		if len(fields) >= 2 {
			packages[fields[0]] = fields[1]
		}
	}

	return packages, nil
}
func Newer(remote, local string) (bool, error) {
	output, err := exec.Command(
		"vercmp", remote, local,
	).Output()
	if err != nil {
		return false, err
	}

	result, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return false, err
	}

	return result > 0, nil
}

func Update() error {
	fmt.Println("Updating official packages...")
	if err := pacman.Update(); err != nil {
		return err
	}

	packages, err := Foreign()
	if err != nil {
		return err
	}

	fmt.Println("\nChecking AUR updates...")

	for name, localVersion := range packages {
		info, err := aur.Info(name)
		if err != nil {
			fmt.Println("Failed to check:", name, err)
			continue
		}
		if len(info) == 0 {
			continue
		}

		pkg := info[0]

		newer, err := Newer(pkg.Version, localVersion)
		if err != nil {
			return err
		}

		if !newer {
			continue
		}

		fmt.Printf(
			"\n%s: %s -> %s\n",
			name, localVersion, pkg.Version,
		)

		err = installer.InstallAUR(name, pkg.PackageBase)
		if err != nil {
			return fmt.Errorf("failed to update %s: %w", name, err)
		}
	}

	fmt.Println("\nUpdate finished.")
	return nil
}
