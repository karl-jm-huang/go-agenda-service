package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// dmCmd represents the dm command
var lmCmd = &cobra.Command{
	Use:   "listallmeetings",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get("http://localhost:8080/v1/meetings")

		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		if res.StatusCode == http.StatusOK {
			fmt.Println(string(body))
		} else {
			fmt.Println("Get failed")
		}
	},
}

func init() {
	RootCmd.AddCommand(lmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
