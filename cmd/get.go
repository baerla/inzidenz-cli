package cmd

import (
	"errors"
	"fmt"
	"github.com/baerla/inzidenz-cli/config"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var sorting string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [[city]|[url]]",
	Short: "Gets current covid incidence of a city",
	Long: `Gets current covid incidence of a city by extracting 
the current incidence out of their homepage.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetConfig()
		if len(args) == 0 {
			printIncidenceForCities(conf.Cities)
			return
		} else if len(args) == 1 {
			for _, city := range conf.Cities {
				if city.Name == args[0] {
					printIncidenceForCity(&city)
					return
				}
			}
			return
		} else if len(args) == 2 {
			printIncidenceForCity(&config.City{Name: args[0], URL: args[1]})
			return
		} else {
			fmt.Println("Too many arguments provided")
			return
		}
	},
}

func printIncidenceForCity(city *config.City) {
	num, err := getIncidenceForCity(city)
	if err != nil {
		panic(err)
	}
	fmt.Print(getFormattedRow(city.Name, num, len(city.Name), len(num)))
}

func getFormattedRow(name string, incidence string, paddingName int, paddingIncidence int) string {
	return fmt.Sprintf("%*s  %*s\n", paddingName, name, paddingIncidence, incidence)
}

func printIncidenceForCities(cities []config.City) {
	rows := make(map[string]string, len(cities))
	longestName := 0
	longestIncidence := 0
	for _, city := range cities {
		incidence, err := getIncidenceForCity(&city)
		if err != nil {
			panic(err)
		}
		rows[city.Name] = incidence

		if len(city.Name) > longestName {
			longestName = len(city.Name)
		}
		if len(incidence) > longestIncidence {
			longestIncidence = len(incidence)
		}
	}

	names := make([]string, len(cities))
	for name := range rows {
		names = append(names, name)
	}
	names = names[5:]

	switch sorting {
	case "incidence":
		sort.SliceStable(names, func(i, j int) bool {
			left, errL := strconv.ParseFloat(strings.ReplaceAll(rows[names[i]], ",", "."), 32)
			right, errR := strconv.ParseFloat(strings.ReplaceAll(rows[names[j]], ",", "."), 32)
			if errL == nil || errR == nil {
				// sort descending
				return left > right
			} else {
				panic("could not parse incidence" + errL.Error() + errR.Error())
			}
		})
	case "name":
		sort.Strings(names)
	}
	for _, name := range names {
		fmt.Print(getFormattedRow(name, rows[name], longestName, longestIncidence))
	}
}

func getIncidenceForCity(city *config.City) (string, error) {
	num := ""
	// extract incidence from webpage
	content := getContentOfWebpage(city.URL)

	pattern := `(?sU)(7[ ,-]+Tage[ ,-]+)?Inzidenz.+(?P<incidence>[\d]*\.?[\d]+,\d+)\s*<`

	// check if some incidence can be extracted
	matched, err := regexp.MatchString(pattern, content)
	if err != nil {
		panic(err)
	}
	if !matched {
		panic(errors.New("no incidence found for " + city.Name))
	}
	regex := regexp.MustCompile(pattern)
	match := regex.FindAllStringSubmatch(content, -1)
	num = match[0][2]

	return strings.ReplaceAll(num, ".", ""), nil
}

func getContentOfWebpage(url string) string {
	// extract content from webpage
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if response.StatusCode != 200 {
		panic(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getCmd.PersistentFlags().StringVarP(&sorting, "sort", "s", "incidence", "sort by [incidence] or [name]")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
