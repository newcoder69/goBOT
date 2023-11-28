package main

import (
	// "container/list"
	// "encoding/json"
	// "time"
	// "math/rand"
	//"net/http"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	// "google.golang.org/api/googleapi/transport"
	// "google.golang.org/api/youtube/v3"

	
	"github.com/bwmarrin/discordgo"
)
var(
	youtubeAPIKey ="AIzaSyAp_g5OG5UUSpvc6Co_lasklGjFeqwjMIc"
	customSearchID ="https://cse.google.com/cse.js?cx=20e96b3a33999445a"
	googleAPIKey = "AIzaSyAp_g5OG5UUSpvc6Co_lasklGjFeqwjMIc"
	messageCount map[string]int
)
	

func main(){
	messageCount = make(map[string]int)
	//create the bot session
	dg, err := discordgo.New("Bot MTE3NzczNjc4NDAzNDEzNjE5NA.GDDePA.gI9fJhRe_EWI5gGl2F6gcS_Q9JOaSg6Ia32k5w")
	if(err != nil){
		fmt.Println("Error creating discord session", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(getStats)
	dg.AddHandler(voiceStateUpdate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	fmt.Println("Bot is now running, press Ctrl + C to quit")
	
	

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc
	dg.Close()
	
}



func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(m.Content, s.State.User.Mention()){
		_, err := s.ChannelMessageSend(m.ChannelID, "Hello " + m.Author.Mention()+"!")
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}

	if m.Content == "ping"{
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")

		if err != nil {
			fmt.Println("could not pong:", err)
		}
	}

	if m.Content == "!join"{
			channelId, err := findUserVoiceChannel(s, m.Author.ID, m.GuildID)
			if err != nil{
				fmt.Print("there was an error finding the voice channel")
			}

			err = joinVoiceChannel(s, m.GuildID, channelId)
			if err != nil{
				fmt.Print("There was an error joining the voice channel")
			}

			
	}

	if m.Content == "!quit"{
		voiceConnection, err := s.ChannelVoiceJoin(m.GuildID, "", false , false)
		
		if err != nil{
			fmt.Print("there was an error disconnecting from the voice channel")
		}
		
		err = voiceConnection.Disconnect()
	}
	

	// if (strings.HasPrefix(m.Content, "!picture")){
	// 		query := strings.TrimSpace(strings.TrimPrefix(m.Content, "!picture"))
	// 		if query == ""{
	// 			s.ChannelMessageSend(m.ChannelID, "please enter a query to search for")
	// 			return
	// 		}
			
	// 		response,err := http.Get(customSearchID + query)
	// 		if err != nil{
	// 			fmt.Print("error finding picture", err)
	// 		}
	// 		defer response.Body.Close()
	// 		if response.StatusCode == 200 {
	// 			//_, err = s.ChannelFileSend(m.ChannelID, query + ".png")
	// 			if err != nil {
					
	// 			}
	// 			return

	// 		}

	// 		// imgurURL := fmt.Sprintf("https://api.imgur.com/3/gallery/search/?q=%s",query)
	// 		// images, err := searchImgur(imgurURL)
	// }

		
	}

	func getStats(s*discordgo.Session, m*discordgo.MessageCreate){
		user := m.Author.ID
		messageCount[user]++
	
		if strings.HasPrefix(m.Content,"!mystats"){
			count := messageCount[user]
			_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("you have sent %d messages", count))
			if err != nil{
				fmt.Print("error sending stats", err)
			}
			
		}
		if strings.HasPrefix(m.Content, "!stats"){
			if (len(m.Mentions) > 0) {
				mentionedUserId := m.Mentions[0].Username
				count2 := messageCount[mentionedUserId]
	
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("@%s has sent %d messages", mentionedUserId, count2))
			}
		}
		
	}

	func voiceStateUpdate(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate){
		if vsu.UserID == s.State.User.ID && vsu.ChannelID != "" {
			fmt.Print("Bot has entered a voice channel")
		}
	}

	func findUserVoiceChannel(s *discordgo.Session, UserID, guildID string)(string, error){
		_, err := s.Guild(guildID)
		if err != nil{
			return "", err
		}
		voiceState, err := s.State.VoiceState(guildID, UserID)
		if err != nil{
			return "", fmt.Errorf("user is not in voice channel")
		}
		return voiceState.ChannelID, nil
	}

	func joinVoiceChannel(s *discordgo.Session, guildID, channelID string) error {
		_, err := s.ChannelVoiceJoin(guildID, channelID,false, false)
		if err != nil{
			return err
		}
		return nil
	}
	