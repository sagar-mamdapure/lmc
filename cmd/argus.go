/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// argusCmd represents the argus command
var argusCmd = &cobra.Command{
	Use:    "argus",
	Hidden: true,
	Short:  "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("argus called: ", cmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(argusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// argusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// argusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
