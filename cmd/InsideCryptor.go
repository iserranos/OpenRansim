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

const inside_folder = "InsideCryptorTest"

var inside_key string

var InsideCryptoCmd = &cobra.Command{
	Use:   "inside-crypto",
	Short: "Encrypt the data and overwrite the original files",
	Long:  `Encrypt the data and overwrite the original files`,
	PreRun: func(cmd *cobra.Command, args []string) {
		create_folder(inside_folder)
		create_files(inside_folder, 500)
		inside_key = generate_rsa_key()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		files := get_files(inside_folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", inside_folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt(string(text), inside_key)
			write_to_file(ciphertext, file_name)
		}
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		files := get_files(inside_folder)
		if len(files) == 500 {
			fmt.Println("Vulnerable!!!")
		} else {
			fmt.Println("Passed :)")
		}
		remove(fmt.Sprintf(pwd+"/%s", inside_folder))
	},
}
