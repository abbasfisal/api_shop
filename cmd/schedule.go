package cmd

import (
	"api_shop/internal/schedule"

	"github.com/spf13/cobra"
)

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "run schedule server ",

	Run: func(cmd *cobra.Command, args []string) {

		schedule.StartScheduleServer()

	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
