package main

import (
	"fmt"

    pyffTeams "github.com/sam-delap/learn-go/pyff-rewrite/teams"
)

// const passTd int = 6
// const passYd float64 = 0.04
// const rushRecTd int = 6
// const rushRecYd float64 = 0.1
// const interception int = -2
// const rec int = 1

func main() {
	teamLoop("blah.csv", []string{"all"})
}

func teamLoop(filename string, teams []string) {
	// Can't create this outside of a function as a global constant in golang
	allTeams := [32]string{"crd", "atl", "rav", "buf", "car", "chi", "cin", "cle", "dal", "den", "det", "gnb", "htx", "clt", "jax", "kan", "rai", "sdg", "ram", "mia", "min", "nwe", "nor", "nyg", "nyj", "phi", "pit", "sfo", "sea", "tam", "oti", "was"}

	if teams[0] == "all" {
		teams = allTeams[:]
	}

	fmt.Println(teams)
	var userInput string
	var decision bool
	for _, teamName := range teams {
		fmt.Printf("Would you like to project team %s? ", teamName)
		fmt.Scan(&userInput)
		decision = userInput == "y"
		if !decision {
			fmt.Println("Skipping")
			continue
		}

		fmt.Printf("Do you need to do team-level projections for team %s? ", teamName)
		fmt.Scan(&userInput)
		decision = userInput == "y"
		if decision {
            pyffTeams.ProjectTeam(teamName)
		}

	}
}
