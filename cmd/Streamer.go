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

const streamer_folder = "StreamerTest"

var streamer_key string

var StreamerCmd = &cobra.Command{
	Use:   "streamer",
	Short: "Encrypts all data and groups it into a single file",
	Long:  `Encrypts all data and groups it into a single file`,
	PreRun: func(cmd *cobra.Command, args []string) {
		create_folder(streamer_folder)
		create_files(streamer_folder, num_files)
		streamer_key = generate_rsa_key()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		files := get_files(streamer_folder)
		var all_text string
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", streamer_folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt_aes(string(text), streamer_key)
			all_text += ciphertext
			remove(file_name)
		}
		write_to_file(all_text, fmt.Sprintf(pwd+"/%s/streamer", streamer_folder))
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		files := get_files(streamer_folder)
		if len(files) == num_files+1 {
			fmt.Println("Vulnerable!!!")
		} else {
			fmt.Println("Passed :)")
		}
		remove_all(fmt.Sprintf(pwd+"/%s", streamer_folder))
	},
}
