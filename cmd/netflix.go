/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

// netflixCmd represents the netflix command
var netflixCmd = &cobra.Command{
	Use:   "n",
	Short: "Set up Netflix",
	Long: `You can use this command to set up Netflix.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := open.Run("https://www.netflix.com/browse")
		if err != nil {
			log.Fatalln("起動に失敗しました")
		}
	},
}

func init() {
	rootCmd.AddCommand(netflixCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netflixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netflixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
