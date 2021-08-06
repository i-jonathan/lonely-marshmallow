package main

import (
	"github.com/yoruba-codigy/goTelegram"
	"time"
)

type user struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Secretary string `json:"secretary"`
	WANumber  string `json:"wa_number"`
}

type report struct {
	UserID      int `json:"user_id"`
	Hour        int `json:"hour"`
	Minute      int `json:"minute"`
	Placement   int `json:"placement"`
	Video       int `json:"video"`
	ReturnVisit int `json:"return_visit"`
	BibleStudy  int `json:"bible_study"`
	Date        time.Time
}

type userPendingData struct {
	Stages       int
	Data         user
	Message      goTelegram.Message
	CurrentStage int
}

type reportPendingData struct {
	Stages       int
	Data         report
	Update       goTelegram.Message
	CurrentStage int
}
