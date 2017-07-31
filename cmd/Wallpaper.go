package cmd

import (
	"github.com/reujab/wallpaper"
	"github.com/spf13/cobra"
	"fmt"
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
			ciphertext := encrypt(string(text), key)
			all_text += ciphertext
		}
		write_to_file(all_text, fmt.Sprintf(pwd+"/%s/streamer", folder))
		return nil
	},
}

func change_wallpaper(url string)  {
	wallpaper.SetFromURL(url)
}

