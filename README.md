# MPM

a lightweight AUR helper for Arch Linux
MPM allows you to search, install, remove, and update packages from official repositories and the AUR

## features

- install packages from AUR and official repositories
- remove installed packages
- update system and AUR packages
- search and view package information

## installation

**requirements:** Go (+1.27.3), Git, pacman, base-devel

```bash
git clone https://github.com/lock3d9/mpm.git
cd mpm
go build -ldflags="-s -w" -o mpm .
sudo install -Dm755 mpm /usr/local/bin/mpm
```

## usage

```bash
mpm install <package>
mpm remove <package>
mpm search <package>
mpm info <package>
mpm update
```

## status

early development
bugs and missing features are expected

## License

MIT
