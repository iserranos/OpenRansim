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
	"fmt"
	"github.com/spf13/cobra"
)

const weak_cryptor_folder = "WeakCryptorTest"

var weak_cryptor_key string

var WeakCryptorCmd = &cobra.Command{
	Use:   "weak-cryptor",
	Short: "Uses weak encryption to encrypt_aes data and removes original files",
	Long:  `Uses weak encryption to encrypt_aes data and removes original files`,
	PreRun: func(cmd *cobra.Command, args []string) {
		create_folder(weak_cryptor_folder)
		create_files(weak_cryptor_folder, num_files)
		weak_cryptor_key = generate_rsa_key()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		files := get_files(weak_cryptor_folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", weak_cryptor_folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt_des(string(text), weak_cryptor_key)
			write_to_file(ciphertext, file_name+".copy")
			remove(file_name)
			check(err)
		}
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		files := get_files(weak_cryptor_folder)
		if len(files) == num_files {
			fmt.Println("Vulnerable!!!")
		} else {
			fmt.Println("Passed :)")
		}
		remove_all(fmt.Sprintf(pwd+"/%s", weak_cryptor_folder))
	},
}
