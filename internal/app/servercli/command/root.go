package command

import (
	"fmt"
	"os"
	"proctor-command/internal/app/servercli/command/start"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "proctord",
	Short: "proctord - Handle executing jobs and maintaining their configuration",
	Long:  `proctord - Handle executing jobs and maintaining their configuration`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(start.GetCmd())
}
