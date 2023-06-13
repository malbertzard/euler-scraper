package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// main function is the entry point of the program
func main() {
	// Define command line flags
	problemNumberPtr := flag.String("p", "", "problem number")
	folderPathPtr := flag.String("f", "", "folder path")
	configPathPtr := flag.String("c", "", "config file path")
	flag.Parse()

	// Validate required flag
	if *problemNumberPtr == "" {
		fmt.Println("Usage: go run main.go -p <problem_number> [-f <folder_path>] [-c <config_file_path>]")
		os.Exit(1)
	}

	// Assign flag values to variables
	problemNumber := *problemNumberPtr
	folderPath := *folderPathPtr
	configPath := *configPathPtr

	// Extract problem content from Project Euler website
	url := fmt.Sprintf("https://projecteuler.net/problem=%s", problemNumber)
	title, content, err := extractContent(url)
	if err != nil {
		fmt.Printf("Failed to extract content: %v\n", err)
		os.Exit(1)
	}

	// Set default folder path if not provided
	if folderPath == "" {
		folderPath = "."
	}

	// Create necessary folders for problem files
	err = createFolders(problemNumber, folderPath)
	if err != nil {
		fmt.Printf("Failed to create folders: %v\n", err)
		os.Exit(1)
	}

	// Write problem file
	problemFile := filepath.Join(folderPath, problemNumber, fmt.Sprintf("%s.md", dashifyTitle(title)))
	err = writeToFile(problemFile, fmt.Sprintf("# %s\n\n%s", title, content))
	if err != nil {
		fmt.Printf("Failed to write problem file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully written %s\n", problemFile)

	// Set default config path if not provided
	if configPath == "" {
		configPath = "config.yaml"
	}

	// Write solution files for each programming language specified in config file
	codeFolder := filepath.Join(folderPath, problemNumber, "code")
	config, err := loadConfig(configPath)
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

// Config struct defines the structure of the config file
type Config struct {
	ProgrammingLanguages map[string]string `yaml:"programmingLanguages"`
}

// extractContent function extracts the problem title and content from the Project Euler website
func extractContent(url string) (string, string, error) {
	// Create HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}

	// Set user agent header to avoid 403 error
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.9999.999 Safari/537.36")

	// Send request and read response body
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Check if response status code is OK
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to fetch URL: %s", resp.Status)
	}

	// Extract problem content from response body
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

	// Extract problem title from response body
	reTitle := regexp.MustCompile(`(?s)<h2>(.*?)</h2>`)
	titleMatches := reTitle.FindStringSubmatch(body)
	if len(titleMatches) < 2 {
		return "", "", fmt.Errorf("failed to extract problem title")
	}
	title := titleMatches[1]

	return strings.TrimSpace(title), strings.TrimSpace(content), nil
}

// writeToFile function writes content to a file
func writeToFile(filepath string, content string) error {
	data := []byte(content)
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// createFolders function creates necessary folders for problem files
func createFolders(problemNumber string, folderPath string) error {
	problemFolder := filepath.Join(folderPath, problemNumber)
	codeFolder := filepath.Join(problemFolder, "code")
	err := os.MkdirAll(codeFolder, 0755)
	if err != nil {
		return err
	}
	return nil
}

// dashifyTitle function replaces non-alphanumeric characters in a string with dashes
func dashifyTitle(title string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	dashified := reg.ReplaceAllString(title, "-")
	dashified = strings.ToLower(strings.Trim(dashified, "-"))
	return dashified
}

// loadConfig function loads the config file and returns a Config struct
func loadConfig(filename string) (*Config, error) {
	if filename == "" {
		return &Config{ProgrammingLanguages: map[string]string{}}, nil
	}
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
