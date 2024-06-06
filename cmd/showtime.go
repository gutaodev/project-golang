/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
// cobra自定义的子命令，go run main.go show 会调用下面配置的Run
var showCmd = &cobra.Command{
	Use:   "showtime",
	Short: "A brief description of your command",
	Long: `
		++++++++++++++
		+ print time +
 		++++++++++++++
`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("show called")
		fmt.Println(time.Now())
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
