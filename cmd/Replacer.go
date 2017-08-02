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
	"strconv"
)

const replacer_folder = "ReplacerTest"

var ReplacerCmd = &cobra.Command{
	Use:   "replacer",
	Short: "Replace the contents of the original files",
	Long:  `Replace the contents of the original files`,
	PreRun: func(cmd *cobra.Command, args []string) {
		create_folder(replacer_folder)
		create_files(replacer_folder, 500)
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		files := get_files(replacer_folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", replacer_folder, file.Name())
			i, err := strconv.Atoi(file.Name())
			check(err)
			write_to_file(rand_string(i*42), file_name)
		}
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		files := get_files(replacer_folder)
		if len(files) == 500{
			fmt.Println("Vulnerable!!!")
		}else{
			fmt.Println("Passed :)")
		}
		remove(fmt.Sprintf(pwd+"/%s", replacer_folder))
	},
}
