package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/lib/pq"
)

type Configuration struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	BotToken string
}

// Global variable for Db connection
var db *sql.DB

func main() {

	var err error

	// Retrieve the credentials from config.json
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Pass variables to driver and start connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configuration.Host, configuration.Port, configuration.User, configuration.Password, configuration.DbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Postgres DB!")

	dg, err := discordgo.New("Bot " + configuration.BotToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if len(m.Content) >= 8 && m.Content[:8] == "!allChat" {

		userMessage := m.Content[8:]
		var ret = ":no_mouth:**Lastest 10 messages:no_mouth:**\n\n"

		if userMessage == "" {

			rows, err := db.Query("SELECT id, message, author_name FROM all_messages ORDER BY id DESC LIMIT $1", 10)
			if err != nil {
				// handle this error better than this
				panic(err)
			}
			defer rows.Close()
			for rows.Next() {
				var id string
				var message string
				var author_name string
				err = rows.Scan(&id, &message, &author_name)
				if err != nil {
					// handle this error
					panic(err)
				}
				ret += "__#" + string(id) + " - " + author_name + ":__ " + message + "\n"
			}
			s.ChannelMessageSend(m.ChannelID, ret)

			// get any error encountered during iteration
			err = rows.Err()
			if err != nil {
				panic(err)
			}
		} else {

			sqlStatement := `
			INSERT INTO all_messages (message, author_name, channel_name, guild_name)
			VALUES ($1, $2, $3, $4)`

			_, err := db.Exec(sqlStatement, userMessage, m.Author.Username, m.ChannelID, m.GuildID)
			if err != nil {
				panic(err)
			}
		}
	}

}
