package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strconv"
)

var ReplacerCmd = &cobra.Command{
	Use:   "replacer",
	Short: "Replace the contents of the original files",
	Long:  `Replace the contents of the original files`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "ReplacerTest"

		create_folder(folder)
		create_files(folder, 500)
		files := get_files(folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			i, err := strconv.Atoi(file.Name())
			check(err)
			write_to_file(rand_string(i * 42), file_name)
		}
		return nil
	},
}
