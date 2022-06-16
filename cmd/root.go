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
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chevify",
	Short: "Tool designed to manage several code editors configurations",
	Long: `Chevify is a tool designed to make your life easier by being able to manage several code editors
configurations simultaneously and allow you to fully seamlessly transition between them.

Editors supported: Neovim Vim soon work in progreess!

https://github.com/TeoDev1611/chevify`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
