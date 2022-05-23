package cmd

import (
	"fmt"
	"learning-go/cmd/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		http.Cmd,
	)
}

// rootCmd ...
var rootCmd = &cobra.Command{
	Use:   "learning-go",
	Short: "learning-go",
	Long:  `learning-go`,
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
