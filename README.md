![heading](https://gitea.koodsisu.fi/petrkubec/art/raw/branch/main/static/Heading.jpg)

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Description
This program provides encoding and decoding functionality for text input. It encodes repeating substrings using a bracketed format `[count substring]` and decodes such formats back to their original form. The program supports single-line and multi-line processing and can read input from files or command-line arguments.

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Features
- Encode text by compressing repeating substrings.
- Decode encoded text back to its original form.
- Supports file input and direct string input.
- Multi-line mode for processing multiple lines of text.
- Web interface for encoding and decoding text.

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Usage
### Running the Program
The program can be executed using the `go run` command:
```sh
go run . -h                    # Show help message
go run . -d "encoded text"      # Decode the input
go run . -e "text to encode"    # Encode the input
go run . -m -d input.txt        # Decode multi-line input from a file
go run . -s                     # Run a web server
```

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

### Flags
- `-h` : Show help information
- `-d` : Decode input (default mode if no flag is provided)
- `-e` : Encode input
- `-m` : Enable multi-line mode for file inputs
- `-s` : Run a web server

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Web Interface
The program includes a web interface that allows users to encode and decode text through a browser. To use the web interface, run the program with the `-s` flag:
```sh
go run . -s
```
This will start a web server on `http://localhost:8080`. Open this URL in your browser to access the web interface. You can enter text to encode or decode, and the results will be displayed on the page.

Textbox adjusts height to your art, use buttons "+" and "-" to resize the font in both textbox and result

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Input Format
- Strings containing repeating sequences will be encoded into `[count substring]` format.
- Decoding reverses the process, restoring the original text.
- Input can be provided as a direct string or a file.
(Note: The encoder only handles files in multi-line mode)

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Examples
### Encoding
#### Input:
```
aaaabbbbcccc
```
#### Command:
```sh
go run . -e "aaaabbbbbcccccc"
```
#### Output:
```
aaaabbbbb[6 c]
```
(Note: The encoder only compresses substrings if doing so provides meaningful compression.)

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

### Decoding
#### Input:
```
[3 hello] world!
```
#### Command:
```sh
go run . -d "[3 hello] world!"
```
#### Output:
```
hellohellohello world!
```

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Output
The program outputs the results to the terminal and also writes them to a file named `latestOutput.txt`.

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Error Handling
- Invalid input formats (e.g., missing spaces in brackets) will result in an error message.
- Multi-line input must be explicitly enabled using `-m`.

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Dependencies
This program is written in Go and requires the Go compiler to run.

![separator](https://64.media.tumblr.com/885d08abf15908e6e49257d8653b1e4f/9cf3dc62c374aba1-67/s2048x3072/a5a084384f964aab9500719967cf090b53aeccd7.pnj)

## Author
- Developed by Petr Kubec
