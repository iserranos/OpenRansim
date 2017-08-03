//	This file is part of OpenRansim.
//
//	Foobar is free software: you can redistribute it and/or modify
//	it under the terms of the GNU General Public License as published by
//	the Free Software Foundation, either version 3 of the License, or
//	(at your option) any later version.
//
//	Foobar is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even the implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with Foobar.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

const strong_cryptor_folder = "StrongCryptorTest"

var strong_cryptor_key string

var StrongCryptorCmd = &cobra.Command{
	Use:   "strong-cryptor",
	Short: "Encrypt the data and delete the originals safely",
	Long:  `Encrypt the data and delete the originals safely`,
	PreRun: func(cmd *cobra.Command, args []string) {
		create_folder(strong_cryptor_folder)
		create_files(strong_cryptor_folder, 50)
		strong_cryptor_key = generate_rsa_key()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		files := get_files(strong_cryptor_folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", strong_cryptor_folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt_aes(string(text), strong_cryptor_key)
			write_to_file(ciphertext, file_name+".copy")
			for i := 0; i < 30; i++ {
				text, err := read_from_file(file_name)
				check(err)
				ciphertext := encrypt_aes(string(text), strong_cryptor_key)
				write_to_file(ciphertext, fmt.Sprintf("%s.%d", file_name, i))
			}
			remove(file_name)
		}
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		files := get_files(strong_cryptor_folder)
		if len(files) == 500 {
			fmt.Println("Vulnerable!!!")
		} else {
			fmt.Println("Passed :)")
		}
		for _, file := range files {
			file_path := fmt.Sprintf(pwd+"/%s/%s", strong_cryptor_folder, file.Name())
			remove(file_path)
		}
		remove(fmt.Sprintf(pwd+"/%s", strong_cryptor_folder))
	},
}
