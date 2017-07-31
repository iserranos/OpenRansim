package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var ThorCmd = &cobra.Command{
	Use:   "thor",
	Short: "It simulates one of the countless variables of ransomware Thor",
	Long:  `It simulates one of the countless variables of ransomware Thor`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "ThorTest"
		var key = generate_rsa_key()

		create_folder(folder)
		create_files(folder, 500)
		files := get_files(folder)
		for _, file := range files {
			file_path := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			text, err := read_from_file(file_path)
			check(err)
			ciphertext := encrypt(string(text), key)
			write_to_file(ciphertext, fmt.Sprintf(pwd+"/%s/%s.thor", folder, file.Name()))
		}
		change_wallpaper("https://github.com/iserranos/OpenRansim/raw/master/images/thor-ransomware.png")
		return nil
	},
}
