package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ProgrammingLanguages map[string]string `yaml:"programmingLanguages"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <problem_number> [<folder_path>]")
		os.Exit(1)
	}

	var problemNumber string
	var folderPath string

	if len(os.Args) == 2 {
		problemNumber = os.Args[1]
	} else {
		problemNumber = os.Args[1]
		folderPath = os.Args[2]
	}

	url := fmt.Sprintf("https://projecteuler.net/problem=%s", problemNumber)

	title, content, err := extractContent(url)
	if err != nil {
		fmt.Printf("Failed to extract content: %v\n", err)
		os.Exit(1)
	}

	if folderPath == "" {
		folderPath = "."
	}

	err = createFolders(problemNumber, folderPath)
	if err != nil {
		fmt.Printf("Failed to create folders: %v\n", err)
		os.Exit(1)
	}

	problemFile := filepath.Join(folderPath, problemNumber, fmt.Sprintf("%s.md", dashifyTitle(title)))

	err = writeToFile(problemFile, fmt.Sprintf("# %s\n\n%s", title, content))
	if err != nil {
		fmt.Printf("Failed to write problem file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully written %s\n", problemFile)

	codeFolder := filepath.Join(folderPath, problemNumber, "code")

	config, err := loadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	for _, ext := range config.ProgrammingLanguages {
		solutionFilename := fmt.Sprintf("%s.%s", "solution", ext)
		solutionPath := filepath.Join(codeFolder, solutionFilename)

		err = writeToFile(solutionPath, content)
		if err != nil {
			fmt.Printf("Failed to write solution file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully written %s\n", solutionPath)
	}
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

	reTitle := regexp.MustCompile(`(?s)<h2>(.*?)</h2>`)
	titleMatches := reTitle.FindStringSubmatch(body)
	if len(titleMatches) < 2 {
		return "", "", fmt.Errorf("failed to extract problem title")
	}

	title := titleMatches[1]

	return strings.TrimSpace(title), strings.TrimSpace(content), nil
}

func writeToFile(filepath string, content string) error {
	data := []byte(content)
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func createFolders(problemNumber string, folderPath string) error {
	problemFolder := filepath.Join(folderPath, problemNumber)
	codeFolder := filepath.Join(problemFolder, "code")

	err := os.MkdirAll(codeFolder, 0755)
	if err != nil {
		return err
	}

	return nil
}

func dashifyTitle(title string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	dashified := reg.ReplaceAllString(title, "-")
	dashified = strings.ToLower(strings.Trim(dashified, "-"))
	return dashified
}

func loadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

