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
	"os"
	"path/filepath"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Clean workspace",
	Long:  `Clean workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {

		pwd, _ := os.Getwd()
		d, err := os.Open(pwd + "/TestDirectory")
		check(err)

		defer d.Close()
		names, err := d.Readdirnames(-1)
		check(err)

		fmt.Println("Removing all files from 'TestDirectory' folder")
		for _, name := range names {
			err = os.RemoveAll(filepath.Join(pwd+"/TestDirectory", name))
			if err != nil {
				return err
			}
			fmt.Print(".")
		}
		fmt.Println("")
		return nil
	},
}
