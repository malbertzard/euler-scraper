package model

// Config struct defines the structure of the config file
type Config struct {
	SolutionFolderName string `yaml:"folderName"`
	ProgrammingLanguages map[string]string `yaml:"programmingLanguages"`
}
