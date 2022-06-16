/*
Copyright © 2022 TeoDev1611

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Flags Helpers!

var (
	name, branch, editor string
	speed, force         bool
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Do you want add or install a new config? Use this command!",
	Long: `With this command you can add a new configuration!
	
Flags avaliable for use with this command:
	name: Change the alias name for the shell usage and the folder name in the main save repo!
	editor [ REQUIRED ]: The editor to install the config!
	branch: The branch name to install default main!
	speed: Clone the repo with a depth of 999 default true!
	force: Force to clone if exists the repo in the config default false!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
	Example: "chevify install https://github.com/TeoDev1611/astro.nvim --name astro ",
	Aliases: []string{"i", "clone", "add"},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	installCmd.Flags().StringVarP(&name, "name", "n", "", "Change the alias name for the configuration!")
	installCmd.Flags().StringVarP(&editor, "editor", "e", "", "The editor to install the config!")
	installCmd.Flags().StringVarP(&branch, "branch", "b", "main", "Change the branch for download the repo!")
	installCmd.Flags().BoolVarP(&speed, "speed", "s", true, "Clone the repo with the depth 999 value :p!")
	installCmd.Flags().BoolVarP(&force, "force", "F", false, "Force the clone if exists the repo in the config recopilation!")
}
