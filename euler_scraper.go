package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var programmingLanguages = map[string]string{
	"golang":  "go",
	"nim":     "nim",
	"c":       "c",
	"ruby":    "rb",
	"python3": "py",
}

func extractContent(url string) (string, string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.9999.999 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to fetch URL: %s", resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	body := string(bodyBytes)

	re := regexp.MustCompile(`(?s)<div class="problem_content" role="problem">(.*?)</div>`)
	matches := re.FindStringSubmatch(body)
	if len(matches) < 2 {
		return "", "", fmt.Errorf("failed to extract problem content")
	}

	content := matches[1]

	return "", strings.TrimSpace(content), nil
}

func writeToFile(filepath string, content string) error {
	data := []byte(content)
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func createFolders(problemNumber string) error {
	problemFolder := filepath.Join(".", problemNumber)
	codeFolder := filepath.Join(problemFolder, "code")

	err := os.MkdirAll(codeFolder, 0755)
	if err != nil {
		return err
	}

	return nil
}

func dashifyTitle(title string) string {
	// Remove non-alphanumeric characters and replace spaces with dashes
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	dashified := reg.ReplaceAllString(title, "-")
	dashified = strings.ToLower(strings.Trim(dashified, "-"))
	return dashified
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run website_scraper.go <problem_number>")
		os.Exit(1)
	}

	problemNumber := os.Args[1]
	url := fmt.Sprintf("https://projecteuler.net/problem=%s", problemNumber)

	_, content, err := extractContent(url)
	if err != nil {
		fmt.Printf("Failed to extract content: %v\n", err)
		os.Exit(1)
	}

	err = createFolders(problemNumber)
	if err != nil {
		fmt.Printf("Failed to create folders: %v\n", err)
		os.Exit(1)
	}

	codeFolder := filepath.Join(problemNumber, "code")

	for lang, ext := range programmingLanguages {
		dashifiedTitle := dashifyTitle(lang)
		solutionFilename := fmt.Sprintf("%s.%s", "solution", ext)
		solutionPath := filepath.Join(codeFolder, dashifiedTitle, solutionFilename)

		err = os.MkdirAll(filepath.Dir(solutionPath), 0755)
		if err != nil {
			fmt.Printf("Failed to create folder: %v\n", err)
			os.Exit(1)
		}

		err = writeToFile(solutionPath, content)
		if err != nil {
			fmt.Printf("Failed to write to file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully written %s\n", solutionPath)
	}
}

