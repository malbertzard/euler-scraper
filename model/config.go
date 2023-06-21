package model

// Config struct defines the structure of the config file
type Config struct {
	SolutionFileName     string            `yaml:"fileName"`
	SolutionFolderName   string            `yaml:"folderName"`
	ProgrammingLanguages map[string]string `yaml:"programmingLanguages"`
}
