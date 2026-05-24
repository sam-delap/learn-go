/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"strings"
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pyff",
	Short: "A handy CLI utility for fantasy football projections",
	Long: `pyff is a lightweight CLI for doing fantasy football projections.
	It is used as a tester project when I write other languages, because it combines
	HTTP requests, data fetching/formatting, HTML parsing, and a basic command-line interface.

	pyff accepts a list of team 3-letter Pro-Football-Reference codes, separated by spaces.
	To project, say, the New England Patriots, one would say:

	pyff nwe

	To project all available teams, simply run:

	pyff all
	`,
	RunE: Project,
	Args: cobra.MinimumNArgs(1),
}

var allTeams = [32]string{ "crd", "atl", "rav", "buf", "car", "chi", "cin", "cle", "dal", "den", "det", "gnb", "htx", "clt", "jax", "kan", "rai", "sdg", "ram", "mia", "min", "nwe", "nor", "nyg", "nyj", "phi", "pit", "sfo", "sea", "tam", "oti", "was" }

func Project (cmd *cobra.Command, args []string) error {
	if args[0] == "all" {
		args = allTeams[:]
	}

	scanner := bufio.NewScanner(os.Stdin)

	for _, teamName := range args {
		fmt.Printf("Would you like to project team %s? [Y/n] ", teamName)
		scanner.Scan()
		response := strings.ToLower(scanner.Text())

		if response == "n" {
			fmt.Println("Curses, foiled again!")
			continue
		}

		fmt.Printf("Do you need to do team-level run/pass play projections for team %s? [y/N] ", teamName)
		scanner.Scan()
		response = strings.ToLower(scanner.Text())
		
		if response == "y" {
			// do team-level pass/run play projection
			// save those projections to the CSV file
		}
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pyff.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("filepath", "f", "projections.csv", "File path to projections CSV")
}


