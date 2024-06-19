package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/charmbracelet/glamour"
)

var matter struct {
	Title string   `yaml:"title"`
	Tags  []string `yaml:"tags"`
}

func PrintFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("No cheats found with %s\n", filePath)
			os.Exit(1)
		}
		return err
	}
	return printMarkdown(string(data))
}

func printMarkdown(markdownText string) error {
	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width (default is 80)
		glamour.WithWordWrap(40),
	)
	rest, err := frontmatter.Parse(strings.NewReader(markdownText), &matter)
	if err != nil {
		panic(err)
	}

	out, err := r.Render(string(rest))
	if err != nil {
		return err
	}
	fmt.Print(out)
	return nil
}

func OpenEditor(filePath string) error {
	cmd := exec.Command("nvim", filePath)

	// Connect the command's standard input, output, and error to the Go program's standard streams
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func ListFiles() error {
	// Specify the directory path
  path, err := ReadConfig()


	// Open the directory
	d, err := os.Open(path.Path)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return err
	}
	defer d.Close()

	// Read directory contents
	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return err
	}

	// Print names of all files in the directory
	for _, file := range files {
		if file.Mode().IsRegular() {
			fmt.Println(file.Name())
		}
	}
  return nil
}


