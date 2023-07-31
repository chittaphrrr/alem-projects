# Text Editing and Auto-correction Tool

This project is a simple text completion, editing, and auto-correction tool written in Go. It receives as input a text file that needs modifications and outputs a new file with the modified text.

## Project Objectives

The objective of this project is to use your old functions made in your old repository to create a tool that can perform the following modifications to a text file:

- Replace every instance of (hex) with the decimal version of the preceding hexadecimal number.
- Replace every instance of (bin) with the decimal version of the preceding binary number.
- Convert the word before (up) to uppercase.
- Convert the word before (low) to lowercase.
- Convert the word before (cap) to capitalized.
- For (up), (low), and (cap), if a number appears next to it, like so: (up, 2), it turns the previously specified number of words in uppercase, lowercase, or capitalized, accordingly.
- Ensure that every punctuation mark (., ,, !, ?, :, and ;) is close to the previous word and with space apart from the next one. Except for groups of punctuation like: ... or !?, in which case the program should format the text as follows: "I was thinking ... You were right" -> "I was thinking... You were right".
- Place the punctuation mark ' to the right and left of the word in the middle of them, without any spaces. If there are more than one word between the two '' marks, the program should place the marks next to the corresponding words.
- Turn every instance of 'a' into 'an' if the next word begins with a vowel (a, e, i, o, u) or a 'h'.

## Usage

To use the tool, run the following command:

`go run . <input.txt> <output.txt>`


Where `input.txt` is the file to be modified and `output.txt` is the file to which the modified text should be saved.

## Authors

[@nsabyrov](https://01.alem.school/git/nsabyrov)

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

