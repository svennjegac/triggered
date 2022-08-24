package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "trigg",
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		str := strings.Join(args, " ")

		strTrigg := ""
		lowerCase := true
		for _, r := range str {
			var s string
			if lowerCase {
				s = strings.ToLower(string(r))
			} else {
				s = strings.ToUpper(string(r))
			}
			strTrigg += s
			lowerCase = !lowerCase
		}

		fmt.Println(strTrigg)
	},
}
