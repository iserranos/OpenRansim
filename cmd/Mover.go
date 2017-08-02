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

const mover_folder = "MoverTest"
const new_mover_folder = "MoverTest_2"

var mover_key string

var MoverCmd = &cobra.Command{
	Use:   "mover",
	Short: "Encrypts the data in a different folder than the original one and deletes the original",
	Long:  `Encrypts the data in a different folder than the original one and deletes the original`,
	PreRun: func(cmd *cobra.Command, args []string) {
		create_folder(mover_folder)
		create_folder(new_mover_folder)
		create_files(mover_folder, 500)
		mover_key = generate_rsa_key()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		files := get_files(mover_folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", mover_folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt(string(text), mover_key)
			new_file_name := fmt.Sprintf(pwd+"/%s/%s", new_mover_folder, file.Name())
			write_to_file(ciphertext, new_file_name)
			remove(file_name)
		}
		remove(fmt.Sprintf(pwd+"/%s", mover_folder))
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		files := get_files(new_mover_folder)
		if len(files) == 500 {
			fmt.Println("Vulnerable!!!")
		} else {
			fmt.Println("Passed :)")
		}
		remove(fmt.Sprintf(pwd+"/%s", new_mover_folder))
		remove(fmt.Sprintf(pwd+"/%s", mover_folder))
	},
}
