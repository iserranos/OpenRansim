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

var ReplacerCmd = &cobra.Command{
	Use:   "replacer",
	Short: "Replace the contents of the original files",
	Long:  `Replace the contents of the original files`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "ReplacerTest"

		create_folder(folder)
		create_files(folder, 500)
		files := get_files(folder)
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			i, err := strconv.Atoi(file.Name())
			check(err)
			write_to_file(rand_string(i * 42), file_name)
		}
		return nil
	},
}
