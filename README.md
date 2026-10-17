# iwantto

`iwantto` is a small command-line TODO list that stores items in `~/.iwantto`.

## Requirements

- Go 1.23+

## Build

```bash
go build -o iwantto .
```

## Usage

```bash
./iwantto [command] [value]
```

Commands:

- `l` or `list`: show current items
- `a` or `add <text>`: add a new item
- `r` or `remove <number>`: remove an item by its 1-based number
- `r` or `remove all`: remove all items
- `h`, `help`, `-h`, or `--help`: show help

If you pass text without a command (for example `./iwantto buy milk`), it is added as a new item.

Running `./iwantto` with no arguments prints the current list.
