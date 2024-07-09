# ASCII Art Banner Generator

This Go program generates ASCII art banners from input text, optionally applying color to specified characters. It supports multiple font styles and allows for custom colorization.

## Features

- Generate ASCII art using different fonts.
- Apply colors to specified characters in the output.
- Handle special characters and newlines.

## Requirements

- Go 1.18 or higher

## Installation

1. Clone the repository:
    ```sh
    git clone https://learn.zone01kisumu.ke/git/hiombima/ascii-art-color.git
    ```
2. Navigate to the project directory:
    ```sh
    cd ascii-art/color
    ```

## Usage

```sh
go run . [OPTION] [STRING]

go run . --color=red/blue "He" "Hello World"

```

### Examples


1. Generate an ASCII art banner with color:
    ```sh
    go run . --color=red H "Hello, World!"
    ```

### Options

- `--color=<color>`: Specifies the color to be applied to the characters in the STRING. Replace `<color>` with the desired colorsS name.
- `standard`, `thinkertoy`, `shadow`: Specifies the font style to be used for the ASCII art.

## How It Works

The program processes command-line arguments to determine the font style, color, and input string. It reads the appropriate font file and generates the ASCII art line by line. If a color option is provided, it applies the specified color to the designated characters.


## License

This project is licensed under the MIT License.

## Contribution


Feel free to reach out if you have any questions or suggestions!