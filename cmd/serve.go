package cmd

import (
	"bufio"
	"fmt"
	"log"
	"mhttp/defaults"
	"mhttp/server"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	port     string
	file     string
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts a HTTP server based on specification",
		Long:  `Starts a HTTP server based on specification from file or standard input`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Starting HTTP server on port", port)
			var definition *string
			if file == "-" {
				definition = scan(os.Stdin)
			} else {
				fmt.Println("Endpoints specification based on file [", file, "]")
				hFile, err := os.OpenFile(file, os.O_RDONLY, 0755)
				if err != nil {
					fmt.Println("Unable to open a specified file [", file, "]")
					os.Exit(1)
				}
				definition = scan(hFile)
			}

			var spec = server.SpecBoundle{}
			yamlErr := yaml.Unmarshal([]byte(*definition), &spec)
			if yamlErr != nil {
				log.Fatalf("error: %v", yamlErr)
			}

			fmt.Printf("Number of specifications %d\n", len(spec.Spec))
			server.Start(&port, &spec.Spec)
		},
	}
)

func scan(file *os.File) *string {
	var definition string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		definition += line + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading specification:", err)
	}

	return &definition
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringVarP(&port, "port", "p", defaults.Defaults.ServerPort, "port number")
	serveCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "API definition file or - in case pasing a configuration using standard input")

}
