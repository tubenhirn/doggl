package cmd

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	doggl "github.com/tubenhirn/doggl/lib"
)

const apiToken doggl.ContextKey = "api_token"
var date string

func init() {
	rootCmd.AddCommand(bookCmd)

	bookCmd.Flags().StringVarP(&date, "date", "d", "", "A custom date for your booking (format 2022-01-15).")

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
		// load the api token from the env|configfile
		apiTokenVal := viper.GetString("api_token")
		// pass the token to a context
		ctx := context.WithValue(context.Background(), apiToken, apiTokenVal)

		// create a new httpClient with the context and its params
		dogglClient := doggl.NewDefaultClient(ctx)

		duration := viper.GetInt64("duration")
		// check for duration parameter and parse it to int if present
		if len(args) > 0 {
			customDuration, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				panic(err)
			}
			// set the customDuration as duration for the timeEntry
			duration = customDuration
		}

		startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), viper.GetInt("start_hour"), viper.GetInt("start_minute"), 00, 0, time.Local)
		// check for customStartDate flag and parse it to the right time format
		if date != "" {
			customStartDate, err := time.Parse("2006-01-02", date)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			// set the customStartDate as startTime for the timeEntry
			startTime = time.Date(customStartDate.Year(), customStartDate.Month(), customStartDate.Day(), viper.GetInt("start_hour"), viper.GetInt("start_minute"), 00, 0, time.Local)
		}

		// prepare the timeEntry struct
		timeEntry := doggl.TimeEntry{
			Duration:    duration,
			Start:       startTime.Format(time.RFC3339),
			ProjectId:   viper.GetInt("project"),
			WorkspaceId: viper.GetInt("workspace"),
			Description: viper.GetString("description"),
			CreatedWith: viper.GetString("created_with"),
		}

		// add the timentry
		_, resErr := dogglClient.StartTimeEntry(timeEntry)
		if resErr != nil {
			panic(resErr)
		} else {
			fmt.Println("new time entry created.")
		}

	},
}
