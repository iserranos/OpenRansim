package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var InsideCryptoCmd = &cobra.Command{
	Use:   "inside-crypto",
	Short: "Encrypt the data and overwrite the original files",
	Long:  `Encrypt the data and overwrite the original files`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "InsideCryptorTest"
		var key = generate_rsa_key()

		create_folder(folder)
		create_files(folder, 500)
		files := get_files(folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt(string(text), key)
			write_to_file(ciphertext, file_name)
		}
		return nil
	},
}
