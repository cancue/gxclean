# gxclean

**gxclean** is a File and Directory Deletion Tool

This tool is a simple CLI (Command Line Interface) tool written in Go that searches for and deletes files or directories with a specific name.

## Installation

```bash
go install github.com/cancue/gxclean@latest
```

## Usage

```bash
gxclean -n <name> [-d] [-f]
```

- `-n`: Specifies the name to delete. All files and directories with this name are targeted.
- `-d`: If this option is used, only directories are targeted.
- `-f`: If this option is used, only files are targeted.

The `-d` and `-f` options cannot be used at the same time. You must choose one of them.

## Example

```bash
gxclean -n test -d
```

The above command searches for and deletes all directories named 'test'.

## Caution

This tool actually deletes files and directories. Therefore, please make sure to back up before using it. The tool displays a confirmation message before proceeding with the deletion, and deletion only proceeds if 'y' is entered.

## License

This project is distributed under the MIT license. For more details, please refer to the `LICENSE` file.