package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all procs",
	Long:  `List of all procs available for execution`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func GetCmd() *cobra.Command {
	return listCmd
}

func init() {
	listCmd.Flags().BoolP("listtoggle", "t", false, "Help message for toggle")
}
