# Euler Scraper

Euler Scraper is a command-line tool for scraping problem content from Project Euler and saving it to local files.

## Usage

```
go run main.go -p <problem_number> [-r <problem_number_end>] [-f <folder_path>] [-c <config_file_path>]
```

The tool supports the following command-line flags:

- `-p <problem_number>`: Specifies the problem number to scrape from Project Euler. This flag is required.
- `-r <problem_number_end>`: Specifies the range of problem numbers to scrape from Project Euler.
- `-f <folder_path>`: Specifies the folder path where the scraped content will be saved. If not provided, the current directory will be used.
- `-c <config_file_path>`: Specifies the path to a YAML config file that defines the programming languages and their corresponding file extensions. If not provided, the tool will use a default configuration.

## Configuration File

The optional configuration file allows you to define the programming languages and their corresponding file extensions. The file should be in YAML format and have the following structure:

```yaml
fileName: name
folderName: name
programmingLanguages:
  language1: extension1
  language2: extension2
  ...
```

For example:

```yaml
fileName: solution
folderName: solutions
programmingLanguages:
  golang: go
  python: py
  ruby: rb
```

When the configuration file is provided, the tool will create solution files in the specified programming languages based on the defined extensions.

## Examples

Scrape problem 1 and save the content to the current directory:

```
go run main.go -p 1
```

Scrape problem 1 and save the content to a specific folder:

```
go run main.go -p 1 -f /path/to/folder
```

Scrape problem 1 and use a custom configuration file:

```
go run main.go -p 1 -c /path/to/config.yaml
```

## Dependencies

The tool relies on the following external package:

- `gopkg.in/yaml.v2`: A YAML parser for Go. You can install it using the command `go get gopkg.in/yaml.v2`.
