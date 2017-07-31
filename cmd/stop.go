package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"fmt"
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
