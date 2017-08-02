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

const thor_folder = "ThorTest"
const url_wallpaper = "https://github.com/iserranos/OpenRansim/raw/master/images/thor-ransomware.png"

var thor_key string

var ThorCmd = &cobra.Command{
	Use:   "thor",
	Short: "It simulates one of the countless variables of ransomware Thor",
	Long:  `It simulates one of the countless variables of ransomware Thor`,
	PreRun: func(cmd *cobra.Command, args []string) {
		create_folder(thor_folder)
		create_files(thor_folder, 500)
		thor_key = generate_rsa_key()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		files := get_files(thor_folder)
		for _, file := range files {
			file_path := fmt.Sprintf(pwd+"/%s/%s", thor_folder, file.Name())
			text, err := read_from_file(file_path)
			check(err)
			ciphertext := encrypt(string(text), thor_key)
			write_to_file(ciphertext, fmt.Sprintf(pwd+"/%s/%s.thor", thor_folder, file.Name()))
		}
		change_wallpaper(url_wallpaper)
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		files := get_files(thor_folder)
		if len(files) == 500*2{
			fmt.Println("Vulnerable!!!")
		}else{
			fmt.Println("Passed :)")
		}
		for _, file := range files {
			file_path := fmt.Sprintf(pwd+"/%s/%s", thor_folder, file.Name())
			remove(file_path)
		}
		remove(fmt.Sprintf(pwd+"/%s", thor_folder))
	},
}
