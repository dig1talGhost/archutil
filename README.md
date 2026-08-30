## Archutil 󱡕
Simple TUI application to run linux utilities with style.

> [!NOTE]
> - This is meant to run on Arch Linux.

## Installation
I recommend you download the latest release's binary and it's `sha512sum.txt` from \
the [releases page](https://github.com/dig1talGhost/archutil/releases), so you can verify your download.

#### Using curl:
```bash
curl -L -O "https://github.com/dig1talGhost/archutil/releases/download/v0.0.1/archutil-v0.0.1-linux-amd64"
curl -L -O "https://github.com/dig1talGhost/archutil/releases/download/v0.0.1/sha512sum.txt"
```
Verify your download:
```bash
sha512sum -c sha512sum.txt
```

Here's what the output should look like:
```
archutil-v0.0.1-linux-amd64: OK
```

Then, make sure the binary is located anywhere on your `$PATH` and make it executable. \
In this example, we assume you have `$GOBIN` set to `~/go/bin`
```bash
mv archutil-v0.0.1-linux-amd64 "$GOBIN/archutil"
chmod +x "$GOBIN/archutil"
```

#### Using git:
Clone the repo in your `~/Downloads` directory:
```bash
cd ~/Downloads
git clone --depth=1 https://github.com/dig1talGhost/archutil.git
```

Generate the binary using `make` (will be stored in `$GOBIN` or `$GOPATH` by default):
```bash
cd archutil && make install
```

## Usage
Print the current version:
```bash
archutil --version
```

Initialize TUI:
```bash
archutil
```

Print `Makefile` help:
```bash
make help
```

### Local http server (ghosttp)
Deploy a simple HTTP server locally:

> [!TIP]
> - The application will prompt you to specify the full path to serve your website locally.
> - As for an example: `/home/user/path/to/website-folder`
> - In your browser, visit `http://localhost:8080/` to access the local server.
> - To close the server and go back to the main menu, press `CTRL + C`.

<br>

### Getarch
Quickly download latest archiso and it's signature. \
By default, download is done via https and the following mirror:
```go
mirror = "https://mirror.quantum5.ca/archlinux/iso/latest/"
```

Verify your download with `gpg`:
```bash
# The name of the iso and the signsture will be different on your system
gpg --verify archlinux.iso.sig archlinux.iso
```

 You should get something similar to the following:
```
Good signature from Pierre Schmitz!
```

### Glyph menu
Unicode symbols on the command line, built using [`huh?`](https://github.com/charmbracelet/huh):

- Add more glyphs easily to `glyphs/glyphs.json`.
- To leave the glyphs menu and go back to the main menu, press `CTRL + C`.

```json
[
  {
    "name": "Arch Linux",
    "icon": "󰣇"
  },
  {
    "name": "Secure",
    "icon": "󰦝 "
  },
  {
    "name": "Digital Key",
    "icon": "󰷖 "
  },
  {
    "name": "Lightning",
    "icon": "󰉁"
  }
]
```

### Password generator
Generate a random 16 characters password by default using the follwing characters:
```go
abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()
```

You can modify the number of characters in `password_generator/password_generator.go`:
```go
func Gen() {

  // Define the number of characters (passwordLength)
  passwordLength := 16
  password, err := generatePassword(passwordLength, defaultChars)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to generate password:", err)
		return
	}

	defer func() {

		for i := range password {
			password[i] = 0
		}
	}()

	lipgloss.Println(simpleStyle.Render(string(password)))
}
```