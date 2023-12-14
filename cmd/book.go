package cmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	doggl "github.com/tubenhirn/doggl/lib"
)

// is key for context storage
const apiTokenKey doggl.ContextKey = "api_token"

var date string
var description string
var duration string

func init() {
	rootCmd.AddCommand(bookCmd)

	bookCmd.Flags().StringVarP(&date, "date", "d", "", "A custom date for your booking (format 2022-01-15).")
	bookCmd.Flags().StringVarP(&description, "description", "e", "Homeoffice", "A custom description for the time entry.")
	bookCmd.Flags().StringVarP(&duration, "duration", "u", "8h", "The duration of the time entry.")

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
	RunE: func(cmd *cobra.Command, args []string) error {
		// pass the apitoken to a context
		apiTokenVal := viper.GetString("api_token")
		if apiTokenVal == "" {
			return errors.New("api_token not set.")
		}
		ctx := context.WithValue(context.Background(), apiTokenKey, apiTokenVal)

		// create a new httpClient with the context and its params
		dogglClient := doggl.NewDefaultClient(ctx)

		startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), viper.GetInt("start_hour"), viper.GetInt("start_minute"), 00, 0, time.Local)
		// check for customStartDate flag and parse it to the right time format
		if date != "" {
			customStartDate, err := time.Parse("2006-01-02", date)
			if err != nil {
				return err
			}
			// set the customStartDate as startTime for the timeEntry
			startTime = time.Date(customStartDate.Year(), customStartDate.Month(), customStartDate.Day(), viper.GetInt("start_hour"), viper.GetInt("start_minute"), 00, 0, time.Local)
		}

		durationInt, err := timeStringToDuration(duration)
		if err != nil {
			return err
		}

		// prepare the timeEntry struct
		timeEntry := doggl.TimeEntry{
			Duration:    durationInt,
			Start:       startTime.Format(time.RFC3339),
			ProjectId:   viper.GetInt("project"),
			WorkspaceId: viper.GetInt("workspace"),
			Description: description,
			CreatedWith: viper.GetString("created_with"),
		}

		// add the timentry
		_, resErr := dogglClient.StartTimeEntry(timeEntry)
		if resErr != nil {
			return resErr
		}

		fmt.Printf("new time entry created. starting at %v for %s.", startTime.Format(time.TimeOnly), duration)

		return nil
	},
}

func timeStringToDuration(timeString string) (int64, error) {
	durationNano, err := time.ParseDuration(timeString)
	if err != nil {
		return 0, err
	}
	durationSec := durationNano.Seconds()
	return int64(durationSec), nil
}
