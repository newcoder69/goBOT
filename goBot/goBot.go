package main

import (
	// "container/list"
	// "encoding/json"
	// "time"

	"math/rand"
	"sync"
	"time"

	//"net/http"

	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	//"time"

	// "google.golang.org/api/googleapi/transport"
	// "google.golang.org/api/youtube/v3"

	"github.com/bwmarrin/discordgo"
)

var (
	messageCount    map[string]int
	moneyCount      map[string]int
	user_state      map[string]string
	bet_amount      map[string]int
	moneyCountMutex sync.Mutex
	targetChannelID string
	interval        time.Duration
)

func main() {
	messageCount = make(map[string]int)
	moneyCount = make(map[string]int)
	user_state = make(map[string]string)
	bet_amount = make(map[string]int)
	//create the bot session
	dg, err := discordgo.New("Bot insert your bot token here")

	if err != nil {
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
	dg.AddHandler(slotMachine)
	dg.AddHandler(wallet)
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection", err)
		return
	}

	fmt.Println("Bot is now running, press Ctrl + C to quit")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
	dg.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(m.Content, s.State.User.Mention()) {
		_, err := s.ChannelMessageSend(m.ChannelID, "Hello "+m.Author.Mention()+"!")
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")

		if err != nil {
			fmt.Println("could not pong:", err)
		}
	}

	if m.Content == "!join" {
		channelId, err := findUserVoiceChannel(s, m.Author.ID, m.GuildID)
		if err != nil {
			fmt.Print("there was an error finding the voice channel")
		}

		err = joinVoiceChannel(s, m.GuildID, channelId)
		if err != nil {
			fmt.Print("There was an error joining the voice channel")
		}

	}

	if m.Content == "!quit" {
		voiceConnection, err := s.ChannelVoiceJoin(m.GuildID, "", false, false)

		if err != nil {
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

func getStats(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author.ID
	messageCount[user]++

	if strings.HasPrefix(m.Content, "!mystats") {
		count := messageCount[user]
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("you have sent %d messages", count))
		if err != nil {
			fmt.Print("error sending stats", err)
		}

	}
	if strings.HasPrefix(m.Content, "!stats") {
		if len(m.Mentions) > 0 {
			mentionedUserId := m.Mentions[0].Username
			count2 := messageCount[mentionedUserId]

			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("@%s has sent %d messages", mentionedUserId, count2))
		}
	}

}

func voiceStateUpdate(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
	if vsu.UserID == s.State.User.ID && vsu.ChannelID != "" {
		fmt.Print("Bot has entered a voice channel")
	}
}

func findUserVoiceChannel(s *discordgo.Session, UserID, guildID string) (string, error) {
	_, err := s.Guild(guildID)
	if err != nil {
		return "", err
	}
	voiceState, err := s.State.VoiceState(guildID, UserID)
	if err != nil {
		return "", fmt.Errorf("user is not in voice channel")
	}
	return voiceState.ChannelID, nil
}

func joinVoiceChannel(s *discordgo.Session, guildID, channelID string) error {
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, false)
	if err != nil {
		return err
	}
	vc.Speaking(true)
	return nil
}

func brainPower(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "!brainpower" {
		channelId, err := findUserVoiceChannel(s, m.Author.ID, m.GuildID)
		if err != nil {
			fmt.Print("there was an error finding the voice channel")
		}

		vc, err := s.ChannelVoiceJoin(m.GuildID, channelId, false, false)
		if err != nil {
			fmt.Print("There was an error joining the voice channel")
		}
		vc.Speaking(true)
	}

	// playAudioFile(s, m.GuildID, AudioFilePath)

}

// func playAudioFile(s *discordgo.Session, guildID, AudioFilepath string) {
// 	options := dca.StdEncodeOptions
// 	options.RawOutput = true
// 	// encodeSession, err := dca.EncodeFile(AudioFilepath, options)
// 	// if err != nil {
// 	// 	fmt.Print("error encoding the file", err)
// 	// }

// 	vc, err := s.ChannelVoiceJoin(guildID, "", false, false)
// 	if err != nil {
// 		fmt.Print("error joining the voice channel", err)
// 	}

// 	vc.Speaking(true)

// }

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

func Troll(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == "" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, "haha you cant send anything")
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

// func hydrationReminder(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	targetChannelID := "1177746908035436614"
// 	// interval := 30 * time.Second
// 	var numReminder int

// 	ticker := time.NewTicker(30 * time.Second)

// 	for {
// 		select {
// 		case <-ticker.C:
// 			embed := &discordgo.MessageEmbed{

// 				Title:       "go drink water",
// 				Description: "hydration reminder" + fmt.Sprintf("%d", numReminder),
// 				Image: &discordgo.MessageEmbedImage{
// 					URL: "https://gifdb.com/images/high/drinking-water-498-x-498-gif-917a02pya269jgqt.gif",
// 				},
// 				Color: 0x3498db,
// 			}
// 			s.ChannelMessageSendEmbed(targetChannelID, embed)

// 		}
// 		numReminder++
// 	}

// }

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

func slotMachine(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author.ID
	baseCost := 10
	if _, exists := moneyCount[user]; !exists {
		moneyCount[user] = 1000
	}

	if strings.HasPrefix(m.Content, "!casino") {
		user_state[user] = "inside casino"
		s.ChannelMessageSend(m.ChannelID, "Welcome to the casino! Your options are Slot Machine :slot_machine: (type !slotmachine), working on others after slot machine is finished")
	}

	if strings.HasPrefix(m.Content, "!slotmachine") && user_state[user] == "inside casino" {
		s.ChannelMessageSend(m.ChannelID, "Welcome to the slot machine! Test your luck \n it costs 10 coins to play plus you can bet for bigger rewards \n would you like to play?")
		user_state[user] = "playing_slot"
	}
	if strings.HasPrefix(m.Content, "!slotmachine") && user_state[user] == "inside casino" && moneyCount[user] < baseCost {
		s.ChannelMessageSend(m.ChannelID, "get your broke ass out of here and get your money up")
	}

	if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "no") {
		s.ChannelMessageSend(m.ChannelID, "okay have fun in the casino!")
	} else if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "yes") {

		s.ChannelMessageSend(m.ChannelID, "awesome! the betting options are \n 0, \n 50, \n 100, \n 150, \n and 200")
	}

	if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "0") {
		bet_amount[user] = 0

		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")

	} else if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "50") {
		bet_amount[user] = 50

		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	} else if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "100") {
		bet_amount[user] = 100

		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	} else if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "150") {
		bet_amount[user] = 150

		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	} else if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "200") {
		bet_amount[user] = 200

		s.ChannelMessageSend(m.ChannelID, "alright bets are locked in. Ready to spin?")
	}

	if user_state[user] == "playing_slot" && strings.HasPrefix(m.Content, "yup") {
		bet := bet_amount[user]
		s.ChannelMessageSend(m.ChannelID, "your bet is"+fmt.Sprintf(" %d", bet))
		rand := rand.Intn(100) + 1
		moneyCountMutex.Lock()
		if rand >= 1 && rand <= 40 {
			s.ChannelMessageSend(m.ChannelID, "Sorry champ you lost, better luck next time. Try to get your money up and your funny down. :cry:")
			loss := baseCost + bet
			moneyCount[user] -= loss
			fmt.Println("money has been subtracted")
		} else if rand >= 40 && rand <= 60 {
			s.ChannelMessageSend(m.ChannelID, "Welp, you broke even. good job :sunglasses:")
		} else if rand >= 60 && rand <= 99 {
			s.ChannelMessageSend(m.ChannelID, "Congratulations, you win")
			gain := baseCost + bet
			moneyCount[user] += gain
			fmt.Println("money has been added")
			fmt.Printf("%d", moneyCount[user])

		}
		moneyCountMutex.Unlock()
		s.ChannelMessageSend(m.ChannelID, "you now have "+fmt.Sprintf("%d", moneyCount[user])+" dollars :dollar:")
	}

}

func wallet(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author.ID
	currentBallance := moneyCount[user]
	if strings.HasPrefix(m.Content, "!wallet") {
		s.ChannelMessageSend(m.ChannelID, "you have "+fmt.Sprintf("%d", currentBallance)+" dollars :dollar:")
	}
}
