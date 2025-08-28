package cmd

import (
	"WikipediaImage/index"
	"WikipediaImage/parse"
	"WikipediaImage/store"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "WikipediaImage",
	Short: "Parse Wiki Image.",
	Long:  "./WikipediaImage 2023 1\n./WikipediaImage 2023 2\n......",
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := strconv.Atoi(args[0])
		month, _ := strconv.Atoi(args[1])
		parse := parse.Parse{Year: year, Month: month}
		result, err := parse.ParseImage()
		if err != nil {
			panic(err)
		}

		store := store.Store{}
		rets, err := store.StoreImage(result)
		if err != nil {
			panic(err)
		}

		index := index.IndexResult{Year: year, Month: month, ImageResults: rets}
		err = index.WriteIndex()
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
