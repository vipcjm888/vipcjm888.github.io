/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "dir",
	Long:  `dir`,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := cmd.Flags().GetString("str")
		if str == "" {
			panic("no start str")
		}
		directory := "." // 替换为你的目录路径
		mdFiles := getMDFilesRecursively(directory)

		fmt.Println("Markdown Files:")
		for _, file := range mdFiles {
			s, l := getFileName(file)
			for i := 1; i <= l; i++ {
				fmt.Print("  ")
			}
			fmt.Printf(" * [%s](%s)\n", s, str+file)
		}

	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
	dirCmd.Flags().StringP("str", "S", "", "start Str")
}

func getMDFilesRecursively(directory string) []string {
	var mdFiles []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			mdFiles = append(mdFiles, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return mdFiles
}

func getFileName(str string) (string, int) {
	s := strings.Split(str, "/")
	return s[(len(s) - 1)], len(s) - 1
}
