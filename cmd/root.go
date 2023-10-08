package cmd

import (
	. "WikipediaImage/parse"
	. "WikipediaImage/store"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "WikipediaImage",
	Short: "Parse Wiki Image.",
	Long:  "./WikipediaImage 2023 1\n./WikipediaImage 2023 2\n......",
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := strconv.Atoi(args[0])
		month, _ := strconv.Atoi(args[1])
		parse := Parse{Year: year, Month: month}
		result, err := parse.ParseImage()
		if err != nil {
			panic(err)
		}
		store := Store{}
		err = store.StoreImage(result)
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
