# Project Euler Website Scraper

This is a command-line tool written in Go that scrapes problem content from the Project Euler website. It retrieves the content of the first `<h2>` element and the content of the `problem_content` class `<div>` for a given problem number and saves it as a Markdown file.
As well as creating the files for the diffrent programming languages.

## Prerequisites

- Go 1.16 or later
- Internet connection

## Usage

1. Clone the repository:

   ```shell
   git clone https://github.com/malbertzard/euler-scraper
   ```

2. Navigate to the project directory:

   ```shell
   cd euler_scraper
   ```

3. Build the Go program:

   ```shell
   go build euler_scraper.go
   ```

4. Run the program with the problem number as an argument:

   ```shell
   ./euler_scraper <problem_number>
   ```

   Replace `<problem_number>` with the desired problem number from Project Euler.

5. The program will scrape the problem content from the Project Euler website and save it as a Markdown file. The Markdown file will be located in the directory `<problem_number>/<title>.md`, where `<programming_language>` is the programming language and `<title>` is the dashified version of the language name.

   For example, if you run the command `./euler_scraper 1` for problem 1 and the Go language, the Markdown file will be saved as `1/x.md`.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

This tool utilizes the following open-source libraries:

- [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) - For HTML parsing and traversal.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
