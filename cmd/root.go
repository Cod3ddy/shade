/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cod3ddy/shade/pkg/lib"
	"github.com/spf13/cobra"
)

	var targetPath string
	var targetFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shade",
	Short: "A brief description of your application",
	Run: func (cmd *cobra.Command, args[]string)  {


		info, err := os.Stat(targetPath)
		if err != nil{
			log.Fatalf("error checking path: %v", err)
		}

		if info.IsDir(){

			err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if filepath.Ext(path) == ".go" {
					err := lib.CleanFile(path)
					if err != nil {
						log.Printf("Error cleaning file %s: %v", path, err)
					}
				}
				return nil
			})

			if err != nil {
				log.Fatalf("Error scanning directory: %v", err)
			}

		}else{
			err := lib.CleanFile(targetFile)
			if err != nil {
				log.Fatalf("Error cleaning file: %v", err)
			}
		}

	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&targetPath, "path", "w", ".", "File or directory to process")
	rootCmd.Flags().StringVarP(&targetFile, "file", "f", "", "Specific Go file to process")
	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


