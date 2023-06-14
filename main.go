package main

import (
	"euler_scraper/helper"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

	// Set default config path if not provided
	if configPath == "" {
		configPath = "config.yaml"
	}

	// Write solution files for each programming language specified in config file
	config, err := helper.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}


	// Extract problem content from Project Euler website
	title, content, err := helper.ExtractContent(problemNumber)
	if err != nil {
		fmt.Printf("Failed to extract content: %v\n", err)
		os.Exit(1)
	}

	// Set default folder path if not provided
	if folderPath == "" {
		folderPath = "."
	}

	// Create necessary folders for problem files
	err = helper.CreateFolders(problemNumber, folderPath, config.SolutionFolderName)
	if err != nil {
		fmt.Printf("Failed to create folders: %v\n", err)
		os.Exit(1)
	}

	// Write problem file
	problemFile := filepath.Join(folderPath, problemNumber, fmt.Sprintf("%s.md", helper.DashifyTitle(title)))
	err = helper.WriteToFile(problemFile, fmt.Sprintf("# %s\n\n%s", title, content))
	if err != nil {
		fmt.Printf("Failed to write problem file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully written %s\n", problemFile)

	codeFolder := filepath.Join(folderPath, problemNumber, config.SolutionFolderName)
	for _, ext := range config.ProgrammingLanguages {
		solutionFilename := fmt.Sprintf("%s.%s", config.SolutionFileName, ext)
		solutionPath := filepath.Join(codeFolder, solutionFilename)
		err = helper.WriteToFile(solutionPath, content)
		if err != nil {
			fmt.Printf("Failed to write solution file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully written %s\n", solutionPath)
	}
}
