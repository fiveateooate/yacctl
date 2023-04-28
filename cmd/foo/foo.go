/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package foo

import (
	"fmt"
	"log"

	"github.com/fiveateooate/yacctl/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fooCmd represents the foo command
var fooCmd = &cobra.Command{
	Use:   "foo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("foo-flag: %s", viper.GetString("foo-flag"))
	},
}

func init() {
	cmd.RootCmd.AddCommand(fooCmd)
	fooCmd.PersistentFlags().StringP("foo-flag", "", "bar", "set the foo type")
	if err := viper.BindPFlag("foo-flag", fooCmd.PersistentFlags().Lookup("foo-flag")); err != nil {
		log.Fatal("failed to bind flag")
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fooCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fooCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
