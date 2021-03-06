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
	"github.com/reujab/wallpaper"
	"github.com/spf13/cobra"
)

var WallpaperCmd = &cobra.Command{
	Use:   "wallpaper",
	Short: "Change the wallpaper of your computer",
	Long:  `Change the wallpaper of your computer`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var folder = "StreamerTest"
		var key = "streamer00000000"

		create_folder(folder)
		create_files(folder, 500)
		files := get_files(folder)
		var all_text string
		for _, file := range files {
			file_name := fmt.Sprintf(pwd+"/%s/%s", folder, file.Name())
			text, err := read_from_file(file_name)
			check(err)
			ciphertext := encrypt_aes(string(text), key)
			all_text += ciphertext
		}
		write_to_file(all_text, fmt.Sprintf(pwd+"/%s/streamer", folder))
		return nil
	},
}

func change_wallpaper(url string) {
	wallpaper.SetFromURL(url)
}
