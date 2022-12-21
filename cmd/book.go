package cmd

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	doggl "github.com/tubenhirn/doggl/lib"
)

func init() {
	rootCmd.AddCommand(bookCmd)
	viper.SetDefault("duration", 28800)
	viper.SetDefault("description", "Homeoffice")
	viper.SetDefault("created_with", "doggl")
	viper.SetDefault("start_hour", 7)
	viper.SetDefault("start_minute", 15)
}

var bookCmd = &cobra.Command{
	Use:   "book [duration]",
	Short: "book time. duration is optional.",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if len(args) > 0 {
			if _, err := strconv.Atoi(args[0]); err != nil {
				return err
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		duration := viper.GetInt64("duration")
		// check for duration parameter
		if len(args) > 0 {
			customDuration, _ := strconv.ParseInt(args[0], 10, 64)
			duration = customDuration
		}
		startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), viper.GetInt("start_hour"), viper.GetInt("start_minute"), 00, 0, time.Local)
		timeEntry := doggl.TimeEntry{
			Duration:    duration,
			Start:       startTime.Format(time.RFC3339),
			ProjectId:   viper.GetInt("project"),
			WorkspaceId: viper.GetInt("workspace"),
			Description: viper.GetString("description"),
			CreatedWith: viper.GetString("created_with"),
		}

		apiToken := viper.GetString("token")

		dogglClient := doggl.NewDefaultClient(apiToken)
		dogglClient.StartTimeEntry(timeEntry)
	},
}
