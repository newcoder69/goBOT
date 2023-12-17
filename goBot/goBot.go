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
var (
	messageCount map[string]int
	moneyCount map[string]int
	userState map[string]string
	
)
	

func main(){
	messageCount = make(map[string]int)
	moneyCount = make(map[string]int)
	userState = make(map[string]int)
	//create the bot session
	dg, err := discordgo.New("Bot insert your bot token here")
	if(err != nil){
		fmt.Println("Error creating discord session", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(getStats)
	dg.AddHandler(voiceStateUpdate)
	dg.AddHandler(brainPower)
	dg.AddHandler(MugMode)
	dg.AddHandler(Troll)
	dg.AddHandler(penisSize)
	dg.AddHandler(coinFlip)
	//dg.AddHandler(hydrationReminder)
	dg.AddHandler(monkey)
	dg.AddHandler(casinoEntrance)
	dg.AddHandler(slotMachine)

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
	
func MugMode(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "mug" {
		s.ChannelMessageSend(m.ChannelID, "maniac")
		embed := &discordgo.MessageEmbed{
			Title:       "Mug Maniac",
			Description: "hes going mug mode",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://i5.walmartimages.com/seo/Mug-Root-Beer-Caffeine-Free-Soda-Pop-12-oz-24-Pack-Cans_4490f5a5-41ec-4f3b-9643-398eaeb0fd2c.10f661ed5389b6231c1b8026b76dbf30.jpeg?odnHeight=640&odnWidth=640&odnBg=FFFFFF",
			},
			Color: 0x3498db,
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			fmt.Println("could not embed message")
		}
	}

}
func penisSize(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!dicksize" {
		chance := rand.Intn(100)
		if chance <= 99 {
			size := rand.Intn(12) + 1
			dick := fmt.Sprintf("%s", strings.Repeat("=", size))

			s.ChannelMessageSend(m.ChannelID, "your dick is this big"+" Ɛ"+dick+">")

			if size <= 4 {
				s.ChannelMessageSend(m.ChannelID, "thats pretty big home slice")
			} else if size <= 8 {
				s.ChannelMessageSend(m.ChannelID, "woah, nice cock. Thick but not too flacid")
			} else if size <= 12 {
				s.ChannelMessageSend(m.ChannelID, "gyatt damn he's packing")
			} else {
				s.ChannelMessageSend(m.ChannelID, "where is my dick at")
			}

		} else if chance >= 100 {
			dick := "======================================================================================="
			s.ChannelMessageSend(m.ChannelID, "SUPER MEGA GYTATT ALERT!! :siren: siren:\n Ɛ"+dick+">")
		}

	}

}

func coinFlip(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!coinflip") {
		result := rand.Intn(2) + 1

		if result == 1 {
			s.ChannelMessageSend(m.ChannelID, "you got heads :coin:")
		} else if result == 2 {
			s.ChannelMessageSend(m.ChannelID, "you got tails :coin:")
		}

	}
}
func monkey(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "monkey" {
		embed := &discordgo.MessageEmbed{

			Title:       "monkey",
			Description: "mmmm monkey",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://media1.tenor.com/m/nd6rC1fcc3oAAAAC/effy.gif",
			},
			Color: 0x3498db,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

}

func casinoEntrance(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author.ID

	moneyCount[user] = 1000

	if strings.HasPrefix(m.Content, "!wallet") {
		s.ChannelMessageSend(m.ChannelID, "you have "+fmt.Sprint(moneyCount[user])+" dollars :dollar:")
	}

	if strings.HasPrefix(m.Content, "!casino") {
		user_state[user] = "inside casino"
		s.ChannelMessageSend(m.ChannelID, "Welcome to the casino! Your options are Slot Machine :slot_machine: (type !slotmachine), working on others after slot machine is finished")
	}

}

func slotMachine(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author.ID
	baseCost := 10
	multiplier := 0.0

	if strings.HasPrefix(m.Content, "!slotmachine") && user_state[user] == "inside casino" {
		s.ChannelMessageSend(m.ChannelID, "Welcome to the slot machine! Test your luck \n it costs 10 coins to play plus you can bet for bigger rewards \n would you like to play?")
		user_state[user] = "playing_slot"
	} else if strings.HasPrefix(m.Content, "!slotmachine") && user_state[user] != "inside casino" {
		{
			s.ChannelMessageSend(m.ChannelID, "you're not in the casino silly goose. Welcome in!")
			user_state[user] = "inside casino"
		}

	}
	if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "no") {
		s.ChannelMessageSend(m.ChannelID, "okay have fun in the casino!")
	} else if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "yes") {
		user_state[user] = "betting"
		s.ChannelMessageSend(m.ChannelID, "awesome! the betting options are \n 0, \n 50, \n 100, \n 150, \n and 200")
	}

	if user_state[user] == "betting" && strings.HasPrefix(m.Content, "0") {
		multiplier = 0
		user_state[user] = "playing"
		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	} else if user_state[user] == "betting" && strings.HasPrefix(m.Content, "50") {
		multiplier = .5
		user_state[user] = "playing"
		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	} else if user_state[user] == "betting" && strings.HasPrefix(m.Content, "100") {
		multiplier = 1.1
		user_state[user] = "playing"
		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	} else if user_state[user] == "betting" && strings.HasPrefix(m.Content, "150") {
		multiplier = 1.5
		user_state[user] = "playing"
		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	} else if user_state[user] == "betting" && strings.HasPrefix(m.Content, "200") {
		multiplier = 2.0
		user_state[user] = "playing"
		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	}

	if user_state[user] == "playing" && strings.HasPrefix(m.Content, "yes") {
		rand := rand.Intn(100) + 1

		if rand >= 1 && rand <= 10 {
			s.ChannelMessageSend(m.ChannelID, "Sorry champ you lost, better luck next time. Try to get your money up and your funny down. :sad:")
			loss := baseCost + (baseCost * int(multiplier))
			moneyCount[user] = moneyCount[user] - float32(loss)
		} else if rand >= 11 && rand <= 67 {
			s.ChannelMessageSend(m.ChannelID, "Welp, you broke even. good job :sunglasses:")
			moneyCount[user] = moneyCount[user]
		} else if rand >= 68 && rand <= 99 {
			s.ChannelMessageSend(m.ChannelID, "Congratulations, you win")
			gain := baseCost + (baseCost * int(multiplier))
			moneyCount[user] = moneyCount[user] + float32(gain)
		}

	}

}

