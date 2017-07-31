package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var StreamerCmd = &cobra.Command{
	Use:   "streamer",
	Short: "Encrypts all data and groups it into a single file",
	Long:  `Encrypts all data and groups it into a single file`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "StreamerTest"
		var key = generate_rsa_key()

		create_folder(folder)
		create_files(folder, 500)
		files := get_files(folder)
		var all_text string
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt(string(text), key)
			all_text += ciphertext
		}
		write_to_file(all_text, fmt.Sprintf(pwd+"/%s/streamer", folder))
		return nil
	},
}
