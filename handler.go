package main

import (
	"fmt"
	"github.com/yoruba-codigy/goTelegram"
	"log"
	"strconv"
	"time"
)

func handler(update goTelegram.Update) {
	switch update.Type {
	case "text":
		processText(update)
	case "callback":
		processCallback(update)
	}
}

func processText(update goTelegram.Update) {
	if len(update.Command) == 0 {
		currentUserData := userList[update.Message.From.ID]
		currentReportData := reportList[update.Message.From.ID]

		if currentUserData != nil {
			if currentUserData.Data.UserID == update.Message.From.ID {
				switch currentUserData.CurrentStage {
				case 0:
					currentUserData.Data.Name = update.Message.Text
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete message for user name")
						log.Println(err)
					}
					currentUserData.CurrentStage++
					text := "What do you call your Secretary? (e.g. sir, Bro XYZ)"
					err = bot.SendMessage(text, update.Message.Chat)
					if err != nil {
						log.Println("can't send message for secretary")
						log.Println(err)
					}
				case 1:
					currentUserData.Data.Secretary = update.Message.Text
					currentUserData.CurrentStage++
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete message for secretary")
						log.Println(err)
					}
					text := "What is your secretary's WhatsApp number? (include country code; e.g. 2348012345678)"
					err = bot.SendMessage(text, update.Message.Chat)
					if err != nil {
						log.Println("can't send message for WA number")
						log.Println(err)
					}
				case 2:
					currentUserData.Data.WANumber = update.Message.Text
					currentUserData.CurrentStage++
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete message for whatsapp number")
						log.Println(err)
					}
					text := fmt.Sprintf("Please confirm these details.\n\nYour Name: %s\n"+
						"What to Call your Secretary: %s\n Secretary's WhatsApp Number: %s\n\n",
						currentUserData.Data.Name, currentUserData.Data.Secretary, currentUserData.Data.WANumber)

					text += "Send 'OK' to Proceed or 'Bail' to Cancel."
					err = bot.SendMessage(text, update.Message.Chat)
					if err != nil {
						log.Println("can't edit message for final stage of user")
						log.Println(err)
					}
				case currentUserData.Stages:
					if update.Message.Text == "OK" {
						createUser(currentUserData.Data)
						delete(userList, update.Message.From.ID)
						err = bot.SendMessage("You've been registered", update.Message.Chat)
						if err != nil {
							log.Println("Couldn't send registration message.\n", err)
						}
					}
					delete(userList, update.Message.From.ID)
					mainMenu(update)
				}
			}
		} else if currentReportData != nil {

			if currentReportData.Data.UserID == update.Message.From.ID {
				switch currentReportData.CurrentStage {
				case 0:
					currentReportData.Data.Hour, err = strconv.Atoi(update.Message.Text)
					if err != nil {
						log.Println("error in converting to int\n", err)
						err = bot.SendMessage("Please retry", update.Message.Chat)
						if err != nil {
							log.Println("Error when they entered hours\n", err)
							mainMenu(update)
						}
					}
					currentReportData.CurrentStage++
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete the hours sent\n", err)
					}
					text := "How many minutes?"
					currentReportData.Update, err = bot.EditMessage(currentReportData.Update, text)
					if err != nil {
						log.Println("couldn't edit message for minutes\n", err)
					}
				case 1:
					currentReportData.Data.Minute, err = strconv.Atoi(update.Message.Text)
					if err != nil {
						log.Println("error in converting to int\n", err)
						err = bot.SendMessage("Please retry", update.Message.Chat)
						if err != nil {
							log.Println("Error when they entered minutes\n", err)
							mainMenu(update)
						}
					}
					currentReportData.CurrentStage++
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete the minutes sent\n", err)
					}
					text := "How many Placements?"
					currentReportData.Update, err = bot.EditMessage(currentReportData.Update, text)
					if err != nil {
						log.Println("couldn't edit message for placements\n", err)
					}
				case 2:
					currentReportData.Data.Placement, err = strconv.Atoi(update.Message.Text)
					if err != nil {
						log.Println("error in converting to int\n", err)
						err = bot.SendMessage("Please retry", update.Message.Chat)
						if err != nil {
							log.Println("Error when they entered placement\n", err)
							mainMenu(update)
						}
					}
					currentReportData.CurrentStage++
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete the placement sent\n", err)
					}
					text := "How many Videos Shown?"
					currentReportData.Update, err = bot.EditMessage(currentReportData.Update, text)
					if err != nil {
						log.Println("couldn't edit message for videos\n", err)
					}
				case 3:
					currentReportData.Data.Video, err = strconv.Atoi(update.Message.Text)
					if err != nil {
						log.Println("error in converting to int\n", err)
						err = bot.SendMessage("Please retry", update.Message.Chat)
						if err != nil {
							log.Println("Error when they entered videos\n", err)
							mainMenu(update)
						}
					}
					currentReportData.CurrentStage++
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete the videos sent\n", err)
					}
					text := "How many Return Visits?"
					currentReportData.Update, err = bot.EditMessage(currentReportData.Update, text)
					if err != nil {
						log.Println("couldn't edit message for rv\n", err)
					}
				case 4:
					currentReportData.Data.ReturnVisit, err = strconv.Atoi(update.Message.Text)
					if err != nil {
						log.Println("error in converting to int\n", err)
						err = bot.SendMessage("Please retry", update.Message.Chat)
						if err != nil {
							log.Println("Error when they entered rv\n", err)
							mainMenu(update)
						}
					}
					currentReportData.CurrentStage++
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete the rv sent\n", err)
					}
					text := "How many Bible Studies?"
					currentReportData.Update, err = bot.EditMessage(currentReportData.Update, text)
					if err != nil {
						log.Println("couldn't edit message for bible studies\n", err)
					}
				case 5:
					currentReportData.Data.BibleStudy, err = strconv.Atoi(update.Message.Text)
					if err != nil {
						log.Println("error in converting to int\n", err)
						err = bot.SendMessage("Please retry", update.Message.Chat)
						if err != nil {
							log.Println("Error when they entered bible studies\n", err)
							mainMenu(update)
						}
					}
					currentReportData.CurrentStage++
					err = bot.DeleteMessage(update.Message)
					if err != nil {
						log.Println("couldn't delete the bible studies sent\n", err)
					}
					text := fmt.Sprintf("Please Review.\n\nHours: %d\nMinutes: %d\nPlacements: %d\nVideos: %d\n" +
						"Return Vists: %d\nBible Studies: %d\n\n", currentReportData.Data.Hour,
						currentReportData.Data.Minute, currentReportData.Data.Placement, currentReportData.Data.Video,
						currentReportData.Data.ReturnVisit, currentReportData.Data.BibleStudy)

					text += "Please tap OK to Continue. Or Cancel to, you know, cancel."

					bot.DeleteKeyboard()
					bot.AddButton("OK", "reportOK")
					bot.AddButton("Cancel", "bail")
					bot.MakeKeyboard(1)
					currentReportData.Update, err = bot.EditMessage(currentReportData.Update, text)
					if err != nil {
						log.Println("couldn't edit message for placements\n", err)
					}
				}
			}

		} else {
			err = bot.DeleteMessage(update.Message)
			if err != nil {
				log.Println("Can't delete useless text\n", err)
			}
		}
		return
	}

	switch update.Command {
	case "/start":
		var (
			newUser user
			count   int64
		)

		db.Where("user_id = ?", update.Message.From.ID).Find(&newUser).Count(&count)
		if count == 0 {
			newUser = user{
				UserID: update.Message.From.ID,
			}

			newProcessing := userPendingData{
				Stages:       3,
				Data:         newUser,
				CurrentStage: 0,
			}

			userList[newUser.UserID] = &newProcessing

			err := bot.SendMessage("Hi! Welcome!\nYou'd need to register to continue.\n\nWhat is your name?",
				update.Message.Chat)

			if err != nil {
				log.Println("Couldn't send registration message")
				log.Println(err)
			}
		} else {
			mainMenu(update)
		}
	}
}

func processCallback(update goTelegram.Update) {
	switch update.CallbackQuery.Data {
	case "main_menu":
		mainMenu(update)
	case "addReport":
		bot.DeleteKeyboard()
		bot.AddButton("Cancel", "bail")
		bot.MakeKeyboard(1)
		newReport := report{
			UserID: update.CallbackQuery.From.ID,
			Date:   time.Now(),
		}

		text := "Please reply with the following details.\nUse digits(e.g. 2, 10, 200) for all fields.\n\nHours:"
		message, err := bot.EditMessage(update.CallbackQuery.Message, text)
		if err != nil {
			log.Println("Can't edit message for adding hours\n", err)
		}

		reportToProcess := reportPendingData{
			Stages:       6,
			Data:         newReport,
			Update:       message,
			CurrentStage: 0,
		}

		reportList[update.CallbackQuery.From.ID] = &reportToProcess
	case "reportOK":
		data := reportList[update.CallbackQuery.From.ID]
		if data == nil {
			log.Println("couldn't fetch data in report ok")
			bot.DeleteKeyboard()
			bot.AddButton("Menu", "main_menu")
			bot.MakeKeyboard(1)
			_, err = bot.EditMessage(update.CallbackQuery.Message, "An error Occurred, please try again later.")
			if err != nil {
				log.Println("can't send error message when data is nil in reportOK\n", err)
			}
			return
		}
		createReport(data.Data)
		bot.DeleteKeyboard()
		bot.AddButton("Menu", "main_menu")
		bot.MakeKeyboard(1)
		_, err = bot.EditMessage(update.CallbackQuery.Message, "Report Recorded")
		if err != nil {
			log.Println("Couldn't tell that report was recorded", err)
		}
	case "viewReport":
	case "viewAllReport":
	case "submitLastReport":
		log.Println("chill")
	}
}

func mainMenu(update goTelegram.Update) {

	bot.DeleteKeyboard()
	text := fmt.Sprintf("Hi, ")
	bot.AddButton("Add Report", "addReport")
	bot.AddButton("View Report(Current Month)", "viewReport")
	bot.AddButton("View Report(All Records)", "viewAllReport")
	bot.AddButton("Collate & Send Last Month's Report", "submitLastReport")
	bot.MakeKeyboard(1)

	if update.Type == "text" {
		if update.Message.Chat.Type != "private" {
			err = bot.SendMessage("Can't use this in a group!", update.Message.Chat)
			if err != nil {
				log.Print("sending warning message failed\n", err)
				return
			}
			return
		}
		text += update.Message.From.Firstname + "."
		err = bot.SendMessage(text, update.Message.Chat)
		if err != nil {
			log.Println("can't send main menu message,\n", err)
		}
	} else {
		text += "Welcome, again."
		_, err = bot.EditMessage(update.CallbackQuery.Message, text)
		if err != nil {
			log.Println("can't edit message")
		}
	}
}