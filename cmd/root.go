package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Execute() {
	root := rootCmd()
	root.AddCommand(lanaCmd())

	if err := root.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "app",
		Short:   "Application Description",
		Example: "you can us the follow commands: create/add/remove/checkout",
	}
}
