package cli

import (
	"fmt"
	"os"

	"mpm/internal/aur"
	"mpm/internal/installer"
	"mpm/internal/pacman"
	"mpm/internal/updater"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mpm <command> [package]")
		fmt.Println("Commands: install, remove, update, search, info")
		return
	}
	command := os.Args[1]
	if command == "update" {
		if err := updater.Update(); err != nil {
			fmt.Println("Error:", err)
		}
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Package name required")
		return
	}

	name := os.Args[2]

	switch command {
	case "install":
		if err := installer.Install(name); err != nil {
			fmt.Println("Error:", err)
		}

	case "remove":
		if err := pacman.Remove(name); err != nil {
			fmt.Println("Error:", err)
		}

	case "search":
		packages, err := aur.Search(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, pkg := range packages {
			fmt.Printf("%s %s\n", pkg.Name, pkg.Version)
			fmt.Println(" ", pkg.Description)
		}

	case "info":
		packages, err := aur.Info(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(packages) == 0 {
			fmt.Println("Package not found")
			return
		}

		pkg := packages[0]
		fmt.Println("Name:", pkg.Name)
		fmt.Println("Version:", pkg.Version)
		fmt.Println("Description:", pkg.Description)
		fmt.Println("Maintainer:", pkg.Maintainer)

	default:
		fmt.Println("Unknown command:", command)
	}
}
