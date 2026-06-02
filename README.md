<!--suppress HtmlDeprecatedAttribute -->

<p align="center">
    <img src="https://github.com/coldtanky1/cli-clicker/blob/master/ascii-art-text.png" alt="logo">
</p>

## Requirements
- [Golang](https://go.dev/dl/)
- [ydotool](https://github.com/ReimuNotMoe/ydotool)

## Making ydotoold automatically run on startup
ydotool is important for this program to work because it's what handles the autoclicking. I supplied a custom made ydotoold.service file that can be used with systemd.
Below is a way on how to use it:

```sh
# assuming you're in the repo folder
sudo cp ydotoold.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now ydotoold.service

# check if it's running
systemctl status ydotoold.service
```

Don't forget to add the following line to your `~/.bashrc`, `~/.zshrc` or fish config
```sh
export YDOTOOL_SOCKET="/run/ydotoold.socket"
```
Or, if you're using fish shell
```sh
set -Ux YDOTOOL_SOCKET "/run/ydotoold.socket"
```

### Why do this?
Because I wasn't able to find an autoclicker that worked on wayland

### Why no GUI?
I suck at making GUIs. But you're free to make one for this.

## Build
```sh
git clone https://github.com/coldtanky1/cli-clicker.git
cd cli-clicker
go mod download
go build -o cli-clicker
```
