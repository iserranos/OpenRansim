package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var MoverCmd = &cobra.Command{
	Use:   "mover",
	Short: "Encrypts the data in a different folder than the original one and deletes the original",
	Long:  `Encrypts the data in a different folder than the original one and deletes the original`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "MoverTest"
		var new_folder = "MoverTest 2"
		var key = generate_rsa_key()

		create_folder(folder)
		create_folder(new_folder)
		create_files(folder, 500)
		files := get_files(folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt(string(text), key)
			new_file_name := fmt.Sprintf(pwd+"/%s/%s", new_folder, file.Name())
			write_to_file(ciphertext, new_file_name)
			remove(file_name)
		}
		remove(fmt.Sprintf(pwd+"/%s", folder))
		return nil
	},
}
