<!--suppress HtmlDeprecatedAttribute -->

<p align="center">
    <img src="https://github.com/coldtanky1/cli-clicker/blob/master/ascii-art-text.png" alt="logo">
</p>

## Requirements
- [Golang](https://go.dev/dl/)

### Why do this?
Because I wasn't able to find an autoclicker that worked on wayland

### Why no GUI?
I suck at making GUIs. But you're free to make one for this.

## Build
```sh
git clone https://github.com/coldtanky1/cli-autoclicker.git
cd cli-autoclicker
go mod download
go build -o cli-clicker
```

## Changelog
v2.1.1 - Removed ydotool as a dependency. Fixed a bug where changing CPS did not in fact change CPS.
v2.1.2 - Added a debug option
