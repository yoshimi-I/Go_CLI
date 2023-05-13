/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

// disneyCmd represents the disney command
var disneyCmd = &cobra.Command{
	Use:   "d",
	Short: "Set up Disney+",
	Long: `YOu can use this command to set up Disney+.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if cmd.Flag("m").Value.String() == "true"{
			err = open.Run("https://www.disneyplus.com/ja-jp/brand/marvel")
			if err != nil {
				log.Fatalln("起動に失敗しました")
			}
		} else{
			err = open.Run("https://www.disneyplus.com/ja-jp/home")
			if err != nil {
				log.Fatalln("起動に失敗しました")
			}
		}
	},
}
func init() {
	rootCmd.AddCommand(disneyCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disneyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	disneyCmd.Flags().BoolP("m", "m", false, "Set up Marvel movie")
}
