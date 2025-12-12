package cmd

import (
	"api_shop/internal/http"
	"fmt"

	"github.com/spf13/cobra"
)

// serverCmd represents the hi command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run http server",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("server will start ")
		http.StartServer()

		fmt.Println("after start server")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
