# Password Generator

A command-line tool for generating passwords with support for different character types and column-based output.

## Installation

### Using Homebrew

You can install the password generator using Homebrew:

```bash
brew tap mexoboy/tools
brew install pgen
```
### Download pre-built binaries

You can download pre-built binaries from the [GitHub Releases](https://github.com/mexoboy/pgen/releases) page. Simply:

1. Go to the [Releases section](https://github.com/mexoboy/pgen/releases)
2. Download the appropriate binary for your operating system
3. Make it executable (on Unix-like systems): `chmod +x pgen`
4. Run it: `./pgen --special-chars --length 20`

### Manual Build

Alternatively, you can build from source (see [Build](#build) section below).

## Usage

```bash
go run main.go [options]
```

## Options

- `--lowercase` – include only lowercase letters in the password (default: false)
- `--chars` – include Latin alphabet characters (default: true)
- `--numbers` – include digits (default: true)
- `--special-chars` – include special characters "!@#$%^&*()-_=+[]{}|;:,.<>?'`~" (default: false)
- `--length` – password length (default: 16)
- `--count` – number of passwords to generate (default: 20)

## Usage Examples

### Generate passwords with default settings
```bash
go run main.go
```

### Generate passwords with lowercase letters only
```bash
go run main.go --lowercase --length 12 --count 10
```

### Generate passwords including special characters
```bash
go run main.go --special-chars --length 20 --count 5
```

### Generate passwords using digits only
```bash
go run main.go --chars=false --length 8
```

### Generate long passwords with all character types
```bash
go run main.go --special-chars --length 32 --count 10
```

## Build

To create an executable from source:

```bash
go build -o pgen main.go
```

Then run it as:

```bash
./pgen --special-chars --length 20
```
