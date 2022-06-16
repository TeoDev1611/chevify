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
	log "github.com/TeoDev1611/chevify/internal/utils/logger"

	"github.com/TeoDev1611/chevify/internal/utils/browser"
	"github.com/spf13/cobra"
)

// bugsCmd represents the bugs command
var bugsCmd = &cobra.Command{
	Use:   "bugs",
	Short: "You found a bug ? Here you can report!",
	Long: `With this command you can report the bugs of chevify

Valid arguments:
    report: Create the bug report to github with a Issue!
    open: You can open the repo url for search in the issues!

DETAILING THE FUNCTIONALITY AND ARGUMENTS

REPORT:
    With this you can report some bugs of the chevify cli
    ARGS: [ <none> <default> ]

OPEN:
    With this you can open the repo issues for search
    ARGS: [ <none> ]
`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "report":
			log.Info("Opening the Issue Template here you can write a new issue!")
			if len(args) == 2 {
				if args[1] == "default" {
					browser.OpenUrlBrowser("https://github.com/TeoDev1611/chevify/issues/new")
					break
				}
			}
			browser.OpenUrlBrowser("https://github.com/TeoDev1611/chevify/issues/new?assignees=&labels=bug&template=bug_report.md&title=%5BBUG%5D")
		case "open":
			log.Info("Opening the Issue URL here you can search all issues avaliable!")
			browser.OpenUrlBrowser("https://github.com/TeoDev1611/chevify/issues")
		}
	},
	Example:   "chevify bugs report",
	Aliases:   []string{"bug"},
	ValidArgs: []string{"report", "open"},
}

func init() {
	rootCmd.AddCommand(bugsCmd)
}
