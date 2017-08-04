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

const locky_folder = "LockyTest"

var locky_key string

var LockyCmd = &cobra.Command{
	Use:   "locky",
	Short: "It simulates one of the countless variables of ransomware Locky",
	Long:  `It simulates one of the countless variables of ransomware Thor`,
	PreRun: func(cmd *cobra.Command, args []string) {
		create_folder(locky_folder)
		create_files(locky_folder, num_files)
		locky_key = generate_rsa_key()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		files := get_files(locky_folder)
		for _, file := range files {
			file_path := fmt.Sprintf(pwd+"/%s/%s", locky_folder, file.Name())
			text, err := read_from_file(file_path)
			check(err)
			ciphertext := encrypt_aes(string(text), locky_key)
			write_to_file(ciphertext, fmt.Sprintf(pwd+"/%s/%s.locky", locky_folder, generate_key()))
			remove(file_path)
		}
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		files := get_files(locky_folder)
		if len(files) == num_files*2 {
			fmt.Println("Vulnerable!!!")
		} else {
			fmt.Println("Passed :)")
		}
		for _, file := range files {
			file_path := fmt.Sprintf(pwd+"/%s/%s", locky_folder, file.Name())
			remove(file_path)
		}
		remove_all(fmt.Sprintf(pwd+"/%s", locky_folder))
	},
}
