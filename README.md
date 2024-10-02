# bravesalsa

A command-line tool to sort file paths based on modification time.

## Example usage

```bash
# Sort files based on modification time
echo -e "/path/to/file1\n/path/to/file2\n/path/to/file3" | bravesalsa sort

# Sort files in reverse order
echo -e "/path/to/file1\n/path/to/file2\n/path/to/file3" | bravesalsa sort --reverse

# Sort files from a text file
bravesalsa sort < file_list.txt

# Sort files and save the output
bravesalsa sort < file_list.txt > sorted_files.txt
```

## Installation

On macOS/Linux:
```bash
brew install gkwa/homebrew-tools/bravesalsa
```

On Windows:
```powershell
# Installation method for Windows to be determined
```

## Building from source

```bash
go build -o bravesalsa main.go
```

## Features

- Sorts file paths based on most recent modification time
- Handles both regular files and directories
- Supports reverse sorting with --reverse flag
- Gracefully handles errors for unreadable files
- Works with file paths containing spaces and special characters

## License

[MIT License](LICENSE)
```
