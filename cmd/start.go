package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Prepare workspace",
	Long:  `Prepare workspace`,
	RunE: func(cmd *cobra.Command, args []string) error {

		pwd, _ := os.Getwd()
		os.MkdirAll(pwd+"/TestDirectory", os.ModePerm)
		fmt.Println("Folder created: TestDirectory")
		time.Sleep(1000)
		fmt.Print("Creating random files ")
		num := 1
		for ; num < 500; num++ {
			f, err := os.Create(fmt.Sprintf(pwd+"/TestDirectory/%d", num))
			check(err)
			defer f.Close()
			_, err = f.WriteString(rand_string(num * 42))
			check(err)
			f.Sync()
			fmt.Print(".")
		}
		fmt.Println("")
		return nil
	},
}
