package teams

import (
    "time"
    "fmt"
)

type Team struct {
    TeamName string
    CurrentYear time.Time
}

func ProjectTeam(teamName string) {
    team := Team{
        TeamName: teamName,
        CurrentYear: time.Now(),
    }

    fmt.Println(team)
}
