<p align="center">
	<img src="https://file.anotherhadi.com/logo-gml-ui.png"
</p>

# Deprecated
# Now here: [anotherhadi/wtui-components](https://github.com/anotherhadi/wtui-components)

# GML-UI
## Go Missing Library - UI

The 'GML-UI' package is a Go library that offers a set of components for developing Text-based User Interfaces (TUI).
It includes features like input fields, selections, and other essential components for building interactive text-based applications.
Easy to set up, use, customize and modify. Some components are inspired by [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles).


## Contents
- [Go Missing Library - UI](#gml-ui)
  - [Contents](#contents)
  - [Installation](#installation)
- [Components](#components)
- [Inputs](#inputs)
  - [Selection](#selection)
  - [List](#list)
  - [Checkbox](#checkbox)
  - [Confirm](#confirm)
  - [Confirm Inline](#confirm-inline)
  - [Input](#input)
  - [Number Picker](#number-picker)
- [Outputs](#outputs)
  - [Paragraph](#paragraph)
  - [Asciitext](#asciitext)
  - [ANSI](#ansi)
  - [Asciimoji](#asciimoji)
  - [Loading](#loading)
  - [Progress](#progress)
  - [Table](#table)
- [Utils](#utils)
  - [Get Char](#get-char)
  - [Get Size](#get-size)

## Installation

```bash
go get https://github.com/anotherhadi/gml-ui@latest
```

## Components

Almost all components accept the Settings struct (from gml-ui/settings) which allows you to modify various parameters: Colors, Default values, Maximum and Minimum, Max rows & columns, etc.
You can refer to the examples folder for each component to learn more.

## Inputs

### Selection

![Selection Example](https://file.anotherhadi.com/selection.gif)

The `selection` component allows you to quickly prompt the user to choose an option.
Move with the arrow keys or JK, select an option with CR.
You can add a Filter to allow the user to filter the options.

### List

![List Example](https://file.anotherhadi.com/list.gif)

The `list` component allows you to quickly prompt the user to choose an option (With Title/Description). Move with the arrow keys or HL, select an option with CR

### Checkbox

![Checkbox Example](https://file.anotherhadi.com/checkbox.gif)

The `checkbox` component allows you to quicky prompt the user to choose multiple options.
Move with the arrow keys or JK, select an option with SPACE

### Confirm

![Confirm Example](https://file.anotherhadi.com/confirm.gif)

The `confirm` component allows you to quickly prompt the user to choose between Yes/No, True/False, etc..
Move with the arrow keys or HL, select an option with CR

### Confirm Inline

![Confirm Inline Example](https://file.anotherhadi.com/confirm_inline.gif)

The `confirm_inline` component allows you to quickly prompt the user to choose between Yes/No, True/False, etc.. But inline.

### Input

![Input Example](https://file.anotherhadi.com/input.gif)

The `input` component allows you to prompt the user to type a string

### Number Picker

![Number Picker Example](https://file.anotherhadi.com/number_picker.gif)

The `number_picker` component allows you to quickly prompt the user to choose a number (int or float).
Increment/Decrement with the arrow keys or HJ/KL, type a number to change the input, and validate with CR.
You can change the Maximum and Minimum through Settings.

## Outputs

### Paragraph

![Paragraph Example](https://file.anotherhadi.com/paragraph.png)

The `paragraph` component is used to print strings with the same look and feel as other components.

### Asciitext

![Asciitext Example](https://file.anotherhadi.com/asciitext.png)

The `asciitext` component is used to print large ASCII art text.

### ANSI

![ANSI Example](https://file.anotherhadi.com/ansi.png)

The `ansi` package provides a simple and user-friendly way to add color and formatting to terminal output in Go applications. Enhance your command-line interfaces with vibrant text, background colors, and text styles, making your output more readable and visually appealing.

### Asciimoji

![Asciimoji Example](https://file.anotherhadi.com/asciimoji.png)

A collection of emojis made from ASCII characters. From [asciimoji](https://asciimoji.com)

### Loading

![Loading Example](https://file.anotherhadi.com/loading.gif)

The `loading` component is a loading spinner, useful for indicating that some kind of operation is in progress.

### Progress

![Progress Example](https://file.anotherhadi.com/progress.gif)

The `progress` component is a progress bar, helpful for visualizing the advancement of a task or process.

### Table

![Table Example](https://file.anotherhadi.com/table.png)

The `table` component is ideal for creating tables with columns and other features for organized data representation.

## Utils

### Get Char

The  `getchar` package enables you to capture a single character or key input without the need for the user to press enter.

Example:

```go
ascii, arrow, err := getchar.GetChar()
if err != nil {
	panic(err)
}
fmt.Print(ascii, arrow)
```

### Get Size

The `getsize` package allows you to obtain the size in columns and rows of the terminal.

Example:

```go
cols, rows, err := getsize.GetSize()
if err != nil {
	panic(err)
}
fmt.Print(cols, rows)
```
