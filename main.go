package main

import (
	"fmt"
  "os"
	"github.com/bwmarrin/discordgo"
)

func main() {
	secret_token := os.Getenv("DISCORD_TOKEN")
	Token := secret_token // Replace with your bot's token

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate function as a callback for MessageCreate events
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	// Block until a signal is received
	select {}
}

// This function will be called whenever a new message is created in the server
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Don't let the bot reply to itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message is "!hi"
	if m.Content == "hi" {
		// Send a reply message
		s.ChannelMessageSend(m.ChannelID, "Hello!")
	}
}
