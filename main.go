package main

import (
	"euler_scraper/helper"
	"euler_scraper/model"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// main function is the entry point of the program
func main() {
    // Define command line flags
    problemNumberPtr := flag.Int("p", 0, "problem number")
    problemNumberRangePtr := flag.Int("r", 0, "problem number range end")
    folderPathPtr := flag.String("f", "", "folder path")
    configPathPtr := flag.String("c", "", "config file path")
    flag.Parse()

    // Validate required flag
    if *problemNumberPtr == 0 {
        fmt.Println("Usage: go run main.go -p <problem_number> [-r <problem_range_end>] [-f <folder_path>] [-c <config_file_path>]")
        os.Exit(1)
    }

    // Assign flag values to variables
    problemNumber := *problemNumberPtr
    problemNumberRange := *problemNumberRangePtr
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

    if (problemNumberRange != 0) {
        for i := problemNumber; i <= problemNumberRange; i++ {
            scrapeProblem(i, folderPath, config);
        }
    } else {
        scrapeProblem(problemNumber, folderPath, config);
    }

}

func scrapeProblem(problemNumber int, folderPath string, config *model.Config)  {
    // Extract problem content from Project Euler website

    problemString := strconv.Itoa(problemNumber)
    title, content, err := helper.ExtractContent(problemString)
    if err != nil {
        fmt.Printf("Failed to extract content: %v\n", err)
        os.Exit(1)
    }

    // Set default folder path if not provided
    if folderPath == "" {
        folderPath = "."
    }

    // Create necessary folders for problem files
    err = helper.CreateFolders(problemString, folderPath, config.SolutionFolderName)
    if err != nil {
        fmt.Printf("Failed to create folders: %v\n", err)
        os.Exit(1)
    }

    // Write problem file
    problemFile := filepath.Join(folderPath, problemString, fmt.Sprintf("%s.md", helper.DashifyTitle(title)))
    err = helper.WriteToFile(problemFile, fmt.Sprintf("# %s\n\n%s", title, content))
    if err != nil {
        fmt.Printf("Failed to write problem file: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Successfully written %s\n", problemFile)

    codeFolder := filepath.Join(folderPath, problemString, config.SolutionFolderName)
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
