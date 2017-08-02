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
	"os"
	"fmt"
	"io/ioutil"
)

var pwd, _ = os.Getwd()

var FilesCmd = &cobra.Command{
	Use:   "files",
	Short: "List files, write over their, update, remove...",
	Long:  `List files, write over their, update, remove...`,
	RunE: func(cmd *cobra.Command, args []string) error {

		go create_folder("TestDirectory")
		go create_files("TestDirectory", 50)
		return nil
	},
}

func create_folder(folderName string) {
	pwd, _ := os.Getwd()
	path := pwd + "/" + folderName
	if folder_exists(path) == false {
		os.MkdirAll(pwd+"/"+folderName, os.ModePerm)
		fmt.Println("Folder created: " + folderName)
	} else {
		fmt.Println("Folder was previous created: " + folderName)
	}
}

func folder_exists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func create_files(folderName string, num_files int) {
	fmt.Print("Creating random files")
	num := 1
	for ; num <= num_files; num++ {
		f, err := os.Create(fmt.Sprintf(pwd+"/%s/%d", folderName, num))
		check(err)
		defer f.Close()
		_, err = f.WriteString(rand_string(num * 142))
		check(err)
		f.Sync()
		fmt.Print(".")
	}
	fmt.Println("")
}

func rand_string(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func get_files(folder string) []os.FileInfo {
	files, err := ioutil.ReadDir(folder)
	check(err)
	return files
}

func write_to_file(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func read_from_file(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

func remove(path string)  {
	err := os.Remove(path)
	check(err)
}

