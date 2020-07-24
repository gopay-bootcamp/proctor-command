package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List of all procs",
		Long:  `List of all procs available for execution`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list called")
		},
	}
}
