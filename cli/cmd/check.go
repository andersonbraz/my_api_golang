package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "agent check --service <<name>>",
	Long: `*************************************************************************
This command check enviroment variables for your MongoDB service.
*************************************************************************`,
	Run: func(cmd *cobra.Command, args []string) {
		checkMongoDB()
	},
}

func checkMongoDB() {

	for _, e := range viper.AllKeys() {
		keyVar := strings.ToUpper(e)
		valVar := viper.GetString(e)
		resp := checkEnv(keyVar, valVar)
		fmt.Println(resp)
	}

}

func checkEnv(key string, value string) string {
	checkIn, exists := os.LookupEnv(key)
	if !exists {
		os.Setenv(key, value)
		viper.AutomaticEnv()
		fmt.Println(key, value)
	}
	return checkIn
}

func init() {

	checkCmd.Flags().BoolP("mongodb", "m", true, "Check enviroment variables for MongoDB")
	rootCmd.AddCommand(checkCmd)

}
