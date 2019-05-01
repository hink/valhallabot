package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func getServersStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	GetSquadValhallaStatus(s, m)
	GetPostScriptumValhallaStatus(s,m)
}

func GetSquadValhallaStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "Server error retrieving Squad server details"
	playerCount, err := getServerPlayerCount(1155175)
	if err == nil {
		// NO ERROR
		msg = fmt.Sprintf("There are currently %d players in Squad on Valhalla.", playerCount)
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}

func GetPostScriptumValhallaStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "Server error retrieving PostScriptum server details"
	playerCount, err := getServerPlayerCount(1155175)
	if err == nil {
		// NO ERROR
		msg = fmt.Sprintf("There are currently %d players in PostScriptum on Valhalla.", playerCount)
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}

func getServerPlayerCount(serverID int) (int, error) {
	server, err := bm.Server(serverID)
	if err != nil {
		return 0, err
	}

	return server.Attributes.Players, nil
}
