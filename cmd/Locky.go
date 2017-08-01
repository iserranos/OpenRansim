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

var LockyCmd = &cobra.Command{
	Use:   "locky",
	Short: "It simulates one of the countless variables of ransomware Locky",
	Long:  `It simulates one of the countless variables of ransomware Thor`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "LockyTest"
		var key = generate_rsa_key()

		create_folder(folder)
		create_files(folder, 500)
		files := get_files(folder)
		for _, file := range files {
			file_path := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			text, err := read_from_file(file_path)
			check(err)
			ciphertext := encrypt(string(text), key)
			write_to_file(ciphertext, fmt.Sprintf(pwd+"/%s/%s.locky", folder, generate_key()))
			remove(file_path)
		}
		return nil
	},
}
