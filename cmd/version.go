package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the current version number of doggl.",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Doggl - simple toggl cli. " + AppVersion)
  },
}
