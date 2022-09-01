package main

import (
	"fmt"
	"github.com/yoruba-codigy/goTelegram"
	"log"
	"strings"
	"time"
)

func createReport(reportData report) {
	// if there is no report between the first of the month and the current date, fetch last month report and carry over the minutes to this month
	var reports []report
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	monthStart := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	previousMonthStart := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
	db.Where("user_id = ?", reportData.UserID).Where("date > ?", monthStart).Find(&reports)
	if len(reports) < 1 {
		// fetch all reports for last month and total it. Then add the minutes to this month
		reports = fetchReportBetweenDates(reportData.UserID, previousMonthStart, monthStart)
		var tempReport report
		for _, r := range reports {
			tempReport.Minute += r.Minute
		}
		if tempReport.Minute > 60 {
			tempReport.Minute = tempReport.Minute % 60
		}

		reportData.Minute += tempReport.Minute
	}
	db.Create(reportData)
}

func deleteAllReports(update goTelegram.Update) {
	db.Where("user_id = ?", update.CallbackQuery.From.ID).Delete(report{})
}

func viewCurrentPerReport(update goTelegram.Update) {
	var reports []report
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	monthStart := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	db.Where("user_id = ?", update.CallbackQuery.From.ID).Where("date > ?", monthStart).Find(&reports)

	var text string
	if len(reports) < 1 {
		text += "No reports found for this Month"
	} else {
		for _, r := range reports {
			text += "Date: " + r.Date.Format("January 02, 2006")
			text += fmt.Sprintf("\nHours: %d\nMinutes: %d\nPlacements: %d\nVideos: %d\nReturn Visits: %d"+
				"\nBible Studies: %d\n\n", r.Hour, r.Minute, r.Placement, r.Video, r.ReturnVisit, r.BibleStudy)
		}
	}

	bot.DeleteKeyboard()
	bot.AddButton("Back", "viewReport")
	bot.AddButton("Menu", "main_menu")
	bot.MakeKeyboard(2)
	_, err = bot.EditMessage(update.CallbackQuery.Message, text)
	if err != nil {
		log.Println("Ma stress mi. Somewhere inside report.go,\n", err)
	}
}

func currentMonthTotaled(update goTelegram.Update) {
	var reports []report
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	tempReport := report{
		Hour:        0,
		Minute:      0,
		Placement:   0,
		Video:       0,
		ReturnVisit: 0,
		BibleStudy:  0,
	}

	monthStart := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	db.Where("user_id = ?", update.CallbackQuery.From.ID).Where("date > ?", monthStart).Find(&reports)

	var text string
	if len(reports) < 1 {
		text += "No reports found for this Month"
	} else {
		for _, r := range reports {
			text = "Date: " + r.Date.Format("January 2006")
			tempReport.Hour += r.Hour
			tempReport.Minute += r.Minute
			tempReport.Placement += r.Placement
			tempReport.Video += r.Video
			tempReport.ReturnVisit += r.ReturnVisit
			tempReport.BibleStudy += r.BibleStudy
		}

		if tempReport.Minute > 60 {
			tempReport.Hour += tempReport.Minute / 60
			tempReport.Minute = tempReport.Minute % 60
		}
		text += fmt.Sprintf("\nHours: %d\nMinutes: %d\nPlacements: %d\nVideos: %d\nReturn Visits: %d"+
			"\nBible Studies: %d\n\n", tempReport.Hour, tempReport.Minute, tempReport.Placement,
			tempReport.Video, tempReport.ReturnVisit, tempReport.BibleStudy)
	}

	bot.DeleteKeyboard()
	bot.AddButton("Back", "viewReport")
	bot.AddButton("Menu", "main_menu")
	bot.MakeKeyboard(2)
	_, err = bot.EditMessage(update.CallbackQuery.Message, text)
	if err != nil {
		log.Println("Ma stress mi. Somewhere inside report.go,\n", err)
	}
}

func viewLastPerReport(update goTelegram.Update) {
	var reports []report
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	monthStart := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
	monthEnd := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	reports = fetchReportBetweenDates(update.CallbackQuery.From.ID, monthStart, monthEnd)
	var text string
	if len(reports) < 1 {
		text += "No reports found for Last Month"
	} else {
		for _, r := range reports {
			if r.Minute > 60 {
				r.Hour += r.Minute / 60
				r.Minute = r.Minute % 60
			}
			text += "Date: " + r.Date.Format("January 02, 2006")
			text += fmt.Sprintf("\nHours: %d\nMinutes: %d\nPlacements: %d\nVideos: %d\nReturn Visits: %d"+
				"\nBible Studies: %d\n\n", r.Hour, r.Minute, r.Placement, r.Video, r.ReturnVisit, r.BibleStudy)
		}
	}

	bot.DeleteKeyboard()
	bot.AddButton("Back", "viewReport")
	bot.AddButton("Menu", "main_menu")
	bot.MakeKeyboard(2)
	_, err = bot.EditMessage(update.CallbackQuery.Message, text)
	if err != nil {
		log.Println("Ma stress mi. Somewhere inside report.go,\n", err)
	}
}

func lastMonthTotaled(update goTelegram.Update) {
	var reports []report
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	tempReport := report{
		Hour:        0,
		Minute:      0,
		Placement:   0,
		Video:       0,
		ReturnVisit: 0,
		BibleStudy:  0,
	}

	monthStart := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
	monthEnd := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	reports = fetchReportBetweenDates(update.CallbackQuery.From.ID, monthStart, monthEnd)

	var text string
	if len(reports) < 1 {
		text += "No reports found for Last Month"
	} else {
		for _, r := range reports {
			text = "Date: " + r.Date.Format("January 2006")
			tempReport.Hour += r.Hour
			tempReport.Minute += r.Minute
			tempReport.Placement += r.Placement
			tempReport.Video += r.Video
			tempReport.ReturnVisit += r.ReturnVisit
			tempReport.BibleStudy += r.BibleStudy
		}

		if tempReport.Minute > 60 {
			tempReport.Hour += tempReport.Minute / 60
			tempReport.Minute = tempReport.Minute % 60
		}
	}

	text += fmt.Sprintf("\nHours: %d\nMinutes: %d\nPlacements: %d\nVideos: %d\nReturn Visits: %d"+
		"\nBible Studies: %d\n\n", tempReport.Hour, tempReport.Minute, tempReport.Placement,
		tempReport.Video, tempReport.ReturnVisit, tempReport.BibleStudy)

	bot.DeleteKeyboard()
	bot.AddButton("Back", "viewReport")
	bot.AddButton("Menu", "main_menu")
	bot.MakeKeyboard(2)
	_, err = bot.EditMessage(update.CallbackQuery.Message, text)
	if err != nil {
		log.Println("Ma stress mi. Somewhere inside report.go,\n", err)
	}
}

func allTotaled(update goTelegram.Update) {
	var reports []report

	tempReport := report{
		Hour:        0,
		Minute:      0,
		Placement:   0,
		Video:       0,
		ReturnVisit: 0,
		BibleStudy:  0,
	}

	db.Where("user_id = ?", update.CallbackQuery.From.ID).Find(&reports)

	var text string
	if len(reports) < 1 {
		text += "No reports found."
	} else {
		for _, r := range reports {
			tempReport.Hour += r.Hour
			tempReport.Minute += r.Minute
			tempReport.Placement += r.Placement
			tempReport.Video += r.Video
			tempReport.ReturnVisit += r.ReturnVisit
			tempReport.BibleStudy += r.BibleStudy
		}

		if tempReport.Minute > 60 {
			tempReport.Hour += tempReport.Minute / 60
			tempReport.Minute = tempReport.Minute % 60
		}

		text += fmt.Sprintf("\nHours: %d\nMinutes: %d\nPlacements: %d\nVideos: %d\nReturn Visits: %d"+
			"\nBible Studies: %d\n\n", tempReport.Hour, tempReport.Minute, tempReport.Placement,
			tempReport.Video, tempReport.ReturnVisit, tempReport.BibleStudy)
	}

	bot.DeleteKeyboard()
	bot.AddButton("Back", "viewReport")
	bot.AddButton("Menu", "main_menu")
	bot.MakeKeyboard(2)
	_, err = bot.EditMessage(update.CallbackQuery.Message, text)
	if err != nil {
		log.Println("Ma stress mi. Somewhere inside report.go,\n", err)
	}
}

func collateSendThisMonth(update goTelegram.Update) {
	var reports []report
	var who user
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	tempReport := report{
		Hour:        0,
		Minute:      0,
		Placement:   0,
		Video:       0,
		ReturnVisit: 0,
		BibleStudy:  0,
	}

	monthStart := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	db.Where("user_id = ?", update.CallbackQuery.From.ID).Where("date > ?", monthStart).Find(&reports)

	var (
		message string
		text    string
	)
	if len(reports) < 1 {
		text += "No reports found for this Month"
		message = text
	} else {
		for _, r := range reports {
			text = "This is my Report for " + r.Date.Format("January 2006")
			tempReport.Hour += r.Hour
			tempReport.Minute += r.Minute
			tempReport.Placement += r.Placement
			tempReport.Video += r.Video
			tempReport.ReturnVisit += r.ReturnVisit
			tempReport.BibleStudy += r.BibleStudy
		}
		if tempReport.Minute > 60 {
			tempReport.Hour += tempReport.Minute / 60
			tempReport.Minute = tempReport.Minute % 60
		}

		text += fmt.Sprintf("\nHours: %d\nMinutes: %d\nPlacements: %d\nVideos: %d\nReturn Visits: %d"+
			"\nBible Studies: %d\n\n", tempReport.Hour, tempReport.Minute, tempReport.Placement,
			tempReport.Video, tempReport.ReturnVisit, tempReport.BibleStudy)

		db.Where("user_id = ?", update.CallbackQuery.From.ID).Find(&who)
		text = fmt.Sprintf("Goodday %s,\n%s", who.Secretary, text)
		message = "Report to be Sent:\n\n" + text
		text = strings.ReplaceAll(text, " ", "+")
		text = strings.ReplaceAll(text, "\n", "%0A")

		message += fmt.Sprintf("\n\nClick the Link below to send your Report to your Secretary on Whatsapp.\n\n "+
			"https://wa.me/%s/?text=%s", who.WANumber, text)
	}
	bot.DeleteKeyboard()
	bot.AddButton("Back", "collate")
	bot.AddButton("Menu", "main_menu")
	bot.MakeKeyboard(2)
	_, err = bot.EditMessage(update.CallbackQuery.Message, message)
	if err != nil {
		log.Println("Ma stress mi. Somewhere inside report.go,\n", err)
	}
}

func collateSendLastMonth(update goTelegram.Update) {
	var reports []report
	var who user
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	tempReport := report{
		Hour:        0,
		Minute:      0,
		Placement:   0,
		Video:       0,
		ReturnVisit: 0,
		BibleStudy:  0,
	}

	monthStart := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
	monthEnd := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	db.Where("user_id = ?", update.CallbackQuery.From.ID).Where("date BETWEEN ? AND ?",
		monthStart, monthEnd).Find(&reports)

	var text string
	var message string
	if len(reports) < 1 {
		text += "No reports found for this Month"
		message = "You have no report for Last Month"
	} else {
		for _, r := range reports {
			text = "This is My Report for " + r.Date.Format("January 2006")
			tempReport.Hour += r.Hour
			tempReport.Minute += r.Minute
			tempReport.Placement += r.Placement
			tempReport.Video += r.Video
			tempReport.ReturnVisit += r.ReturnVisit
			tempReport.BibleStudy += r.BibleStudy
		}

		if tempReport.Minute > 60 {
			tempReport.Hour += tempReport.Minute / 60
			tempReport.Minute = tempReport.Minute % 60
		}
		text += fmt.Sprintf("\nHours: %d\nMinutes: %d\nPlacements: %d\nVideos: %d\nReturn Visits: %d"+
			"\nBible Studies: %d\n\n", tempReport.Hour, tempReport.Minute, tempReport.Placement,
			tempReport.Video, tempReport.ReturnVisit, tempReport.BibleStudy)
		message = "Report to be Sent:\n\n" + text
		text = strings.ReplaceAll(text, " ", "+")
		text = strings.ReplaceAll(text, "\n", "%0A")

		db.Where("user_id = ?", update.CallbackQuery.From.ID).Find(&who)
		message += fmt.Sprintf("\n\nClick the Link below to send your Report to your Secretary on Whatsapp.\n\n"+
			"Goodday %s, "+
			"https://wa.me/%s/?text=%s", who.Secretary, who.WANumber, text)
	}
	bot.DeleteKeyboard()
	bot.AddButton("Back", "collate")
	bot.AddButton("Menu", "main_menu")
	bot.MakeKeyboard(2)
	_, err = bot.EditMessage(update.CallbackQuery.Message, message)
	if err != nil {
		log.Println("Ma stress mi. Somewhere inside report.go,\n", err)
	}
}
