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
)

var CommandName string = "OpenRansim"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Short: "CLI for testing your Ransomware protection",
	Long:  `CLI for testing your Ransomware protection`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	var rootCmd = &cobra.Command{Use: CommandName}
	rootCmd.AddCommand(InsideCryptoCmd,
		MoverCmd,
		ReplacerCmd,
		StreamerCmd,
		StrongCryptorCmd,
		StrongCryptorFastCmd,
		StrongCryptorNetCmd,
		ThorCmd, LockyCmd, WeakCryptorCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//func init() {
//cobra.OnInitialize(initConfig)

// Here you will define your flags and configuration settings.
// Cobra supports Persistent Flags, which, if defined here,
// will be global for your application.

//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.calamar-cli.yaml)")

//RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
//RootCmd.PersistentFlags().IntVarP(&Order, "order", "o", 1, "order to execute")
//RootCmd.PersistentFlags().StringVarP(&EtherFilePath, "ether_file", "e", "", " The ether file path")
//Removed for next versions
//RootCmd.PersistentFlags().StringVarP(&ServerURL, "server", "s", "", "Server URL")
// Cobra also supports local flags, which will only run
// when this action is called directly.
//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}
