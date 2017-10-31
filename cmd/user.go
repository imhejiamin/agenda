// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// usersCmd represents the users command
// var usersCmd = &cobra.Command{
// 	Use:   "users",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:
//
// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("users called")
// 	},
// }
func printError(error string) {
	fmt.Fprint(os.Stderr, error)
	os.Exit(1)
}
func checkEmpty(key, value string) {
	if value == "" {
		printError(key + " can't be empty!\n")
	}
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register user.",
	Long:  `You need to provide username and password to register, and the username can't be the same as others.`,
	Run: func(com *cobra.Command, args []string) {
		username, _ := com.Flags().GetString("user")
		checkEmpty("username", username)

		password, _ := com.Flags().GetString("password")
		checkEmpty("password", password)

		mail, _ := com.Flags().GetString("mail")
		checkEmpty("mail", mail)

		phone, _ := com.Flags().GetString("phone")
		checkEmpty("phone", phone)

		//	cmd.Register(username, password, mail, phone)   // 还没实现
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Long:  ``,
	Run: func(com *cobra.Command, args []string) {
		username, _ := com.Flags().GetString("user")
		checkEmpty("username", username)

		password, _ := com.Flags().GetString("password")
		checkEmpty("password", password)

		//cmd.Login(username, password) //还没实现
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long:  ``,
	Run: func(com *cobra.Command, args []string) {
		//cmd.Logout() //还没实现
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,
	Run: func(com *cobra.Command, args []string) {
		//cmd.ShowUsers()   // 还没实现
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your account.",
	Long: `Once you have deleted your account, you have no way to get it back!!!
And all of information about you will be erased! That's you are dead!!!`,
	Run: func(com *cobra.Command, args []string) {
		//cmd.DeleteUser()  //还没实现
	},
}

func init() {
	registerCmd.Flags().StringP("user", "u", "", "Username")
	registerCmd.Flags().StringP("password", "p", "", "Help message for username")
	registerCmd.Flags().StringP("mail", "m", "", "email.")
	registerCmd.Flags().StringP("phone", "t", "", "Phone")

	loginCmd.Flags().StringP("user", "u", "", "Input username")
	loginCmd.Flags().StringP("password", "p", "", "Input password")

	RootCmd.AddCommand(registerCmd)
	RootCmd.AddCommand(loginCmd)
	RootCmd.AddCommand(logoutCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}