package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var StrongCryptorCmd = &cobra.Command{
	Use:   "strong-cryptor",
	Short: "Encrypt the data and delete the originals safely",
	Long:  `Encrypt the data and delete the originals safely`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "StrongCryptorTest"
		var key = generate_rsa_key()

		create_folder(folder)
		create_files(folder, 50)
		files := get_files(folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt(string(text), key)
			write_to_file(ciphertext, file_name+".copy")
			remove(file_name)
		}
		return nil
	},
}
