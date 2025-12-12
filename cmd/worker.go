package cmd

import (
	"api_shop/internal/worker"

	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "run asynq worker",
	Run: func(cmd *cobra.Command, args []string) {

		worker.StartWorker()

	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
