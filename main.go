package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Retrieve the Discord token from environment variables
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

	// Print that the bot is running
	fmt.Println("Bot is now running. Press Ctrl+C to exit.")

	// Start the HTTP server in a goroutine
	go startHTTPServer()

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
	if m.Content == "!test" {
		// Send a reply message
		s.ChannelMessageSend(m.ChannelID, "nigger")
	}
}

// Start the HTTP server on a specified port
func startHTTPServer() {
	// Define the port to listen on (can be set dynamically using command-line args or environment variables)
	port := ":8080" // You can change this to any available port number

	// Handle the root route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! The bot is running!")
	})

	// Start the HTTP server and listen on the specified port
	log.Printf("Starting HTTP server on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
