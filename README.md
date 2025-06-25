# Password Generator

A command-line tool for generating passwords with support for different character types and column-based output.

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

To create an executable:

```bash
go build -o pgen main.go
```

Then run it as:

```bash
./pgen --special-chars --length 20
```
