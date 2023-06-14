package helper

import (
	"euler_scraper/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// WriteToFile function writes content to a file
func WriteToFile(filepath string, content string) error {
	// Convert <ul> tags to markdown list
	content = strings.ReplaceAll(content, "<ul>", "")
	content = strings.ReplaceAll(content, "</ul>", "")
	content = strings.ReplaceAll(content, "<li>", "- ")
	content = strings.ReplaceAll(content, "</li>", "\n")

	// Convert <br> tags to new lines
	content = strings.ReplaceAll(content, "<br>", "\n")

	// Convert <b> tags to bold (surround with '**')
	content = strings.ReplaceAll(content, "<b>", "**")
	content = strings.ReplaceAll(content, "</b>", "**")

	// Convert <i> tags to italic (surround with '_')
	content = strings.ReplaceAll(content, "<i>", "_")
	content = strings.ReplaceAll(content, "</i>", "_")

	// Remove remaining HTML tags from content
	re := regexp.MustCompile(`<(.*?)>`)
	content = re.ReplaceAllString(content, "")

	data := []byte(content)
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// CreateFolders function creates necessary folders for problem files
func CreateFolders(problemNumber string, folderPath string, solutionFolder string) error {
	problemFolder := filepath.Join(folderPath, problemNumber)
	codeFolder := filepath.Join(problemFolder, solutionFolder)
	err := os.MkdirAll(codeFolder, 0755)
	if err != nil {
		return err
	}
	return nil
}

// DashifyTitle function replaces non-alphanumeric characters in a string with dashes
func DashifyTitle(title string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	dashified := reg.ReplaceAllString(title, "-")
	dashified = strings.ToLower(strings.Trim(dashified, "-"))
	return dashified
}

// LoadConfig function loads the config file and returns a Config struct
func LoadConfig(filename string) (*model.Config, error) {
	if filename == "" {
		return &model.Config{ProgrammingLanguages: map[string]string{}}, nil
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config model.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
