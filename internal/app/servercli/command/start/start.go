package start

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "Start the proctor server",
	Long:    `Start the proctor server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Proctor server started")
	},
}

func GetCmd() *cobra.Command {
	return startCmd
}
