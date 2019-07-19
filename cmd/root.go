package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
	"github.com/zzh8829/kuci/pkg/kuci"
)

var rootCmd = &cobra.Command{
	Use:   "kuci",
	Short: "kubernetes ci",
	Run: func(cmd *cobra.Command, args []string) {
		c := kuci.NewController()
		c.Start()
	},
}

// Execute function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
