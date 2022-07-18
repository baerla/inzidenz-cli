package cmd

import (
	"fmt"
	"github.com/baerla/inzidenz-cli/config"
	"net/url"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("add called, args: %v\n", args)
		u, err := url.ParseRequestURI(args[1])
		//fmt.Printf("u: %v ; err: %v \n", u, err)
		if err != nil {
			panic(err)
		}

		city := config.City{
			Name: args[0],
			URL:  u.String(),
		}
		conf := config.GetConfig()
		found := false
		for _, c := range conf.Cities {
			if c.URL == city.URL && c.Name == city.Name {
				found = true
				break
			}
		}
		if !found {
			conf.Cities = append(conf.Cities, city)
			conf.SaveToConfigFile()
			fmt.Printf("City %s with url: %s added", city.Name, city.URL)
		} else {
			fmt.Println("Already added")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
