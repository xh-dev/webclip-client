# webclip-client

Command-line client for [webclip2-server](https://github.com/xh-dev/webclip2-server), a temporary message "clipboard" you can use to move text between machines.

Sending a message returns a numeric **id** and a 3-letter uppercase **secret**. Both are required to retrieve it. Messages expire on the server after a timeout (see the server's `/config`).

## Installation

### linux - arm
```sh
curl -L https://github.com/xh-dev/webclip-client/releases/latest/download/webclip-linux-arm --output webclip && chmod +x webclip
```

### linux - arm64
```sh
curl -L https://github.com/xh-dev/webclip-client/releases/latest/download/webclip-linux-arm64 --output webclip && chmod +x webclip
```

### linux - 386
```sh
curl -L https://github.com/xh-dev/webclip-client/releases/latest/download/webclip-linux-386 --output webclip && chmod +x webclip
```

### linux - amd64
```sh
curl -L https://github.com/xh-dev/webclip-client/releases/latest/download/webclip-linux-amd64 --output webclip && chmod +x webclip
```

### [win - 386](https://github.com/xh-dev/webclip-client/releases/latest/download/webclip-win-386.exe)
### [win - amd64](https://github.com/xh-dev/webclip-client/releases/latest/download/webclip-win-amd64.exe)

Each release also ships `checksums.txt` (SHA-256) and a `.tar.gz` archive with its `.sha256` file; verify with `sha256sum -c checksums.txt`.

## Usage

```
webclip -action <send|retrive> [flags] [args...]
```

### Send

```sh
$ webclip -action send hello world
567014 QXT
```

The output is `<id> <secret>`. Arguments are joined with spaces to form the message.

Read the message from the clipboard instead:
```sh
webclip -action send -ic
```

Print a shareable web URL (`<host>/#/get?id=<id>&secret=<secret>`) instead:
```sh
webclip -action send -asUrl hello world
```

### Retrieve

```sh
$ webclip -action retrive 567014 QXT
hello world
```

Copy the retrieved message to the clipboard as well:
```sh
webclip -action retrive -oc 567014 QXT
```

## Flags

| Flag      | Default                           | Description                                              |
|-----------|-----------------------------------|----------------------------------------------------------|
| `-action` | `none`                            | `send` or `retrive`. Prints help when omitted.           |
| `-host`   | `https://webclip.webutils.space`  | Server base URL. The client calls `<host>/api/msg/...`.  |
| `-asUrl`  | `false`                           | `send` only: output the retrieve URL.                    |
| `-ic`     | `false`                           | `send` only: read the message from the clipboard.        |
| `-oc`     | `false`                           | Also write the output to the clipboard.                  |

## Server API used

| Action   | Request                                                   | Response                    |
|----------|-----------------------------------------------------------|-----------------------------|
| send     | `POST /api/msg/create` `{"msg": "..."}`                   | `{"id": "...", "secret": "..."}` |
| retrieve | `POST /api/msg/retrieve` `{"code": "<id>", "secret": "..."}` | `{"msg": "..."}`         |

## Building

Requires Go 1.16+.

```sh
go build -o webclip cmd/webclip.go
./build.sh   # cross-compile all targets into build/
```
