# Project Euler Website Scraper

<!--toc:start-->
- [Project Euler Website Scraper](#project-euler-website-scraper)
  - [Prerequisites](#prerequisites)
  - [Usage](#usage)
  - [Example](#example)
  - [Project Structure](#project-structure)
  - [License](#license)
<!--toc:end-->

This is a simple command-line tool written in Go that scrapes problem content from Project Euler website and generates a project structure for each problem, including the problem description and code files for multiple programming languages.

## Prerequisites

To run this tool, you need to have Go installed on your system. You can download and install Go from the official website: [https://golang.org](https://golang.org)

## Usage

```
go run website_scraper.go <problem_number> [<folder_path>]
```

- `<problem_number>`: The specific problem number you want to scrape from Project Euler.
- `<folder_path>` (optional): The path to the folder where you want to create the project. If not provided, the current directory will be used.

## Example

To scrape problem 1 and create the project in the current directory, run the following command:

```
go run website_scraper.go 1
```

To scrape problem 2 and create the project in a specific folder, run the following command:

```
go run website_scraper.go 2 /path/to/projects
```

## Project Structure

After running the tool, it will create a project structure for the specified problem in the specified folder (or the current directory). The structure will be as follows:

```
<folder_path>/
└── <problem_number>/
    ├── <dashified_title>.md
    └── code/
        ├── solution.go
        ├── solution.nim
        ├── solution.c
        ├── solution.rb
        └── solution.py
```

- `<folder_path>`: The folder where the project is created. If not provided, the current directory is used.
- `<problem_number>`: The problem number specified during execution.
- `<dashified_title>.md`: The markdown file containing the problem description.
- `code/`: The folder containing code files for different programming languages.
- `solution.go`, `solution.nim`, `solution.c`, `solution.rb`, `solution.py`: Sample code files for each programming language.

You can modify the code files according to your solutions for the specific problem.

## License

This project is licensed under the [MIT License](LICENSE).
