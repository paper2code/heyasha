package models

import (
	"time"
)

type Message struct  {
	ID              uint64
	FlexID          string
	//FlexIDType      FlexIDType
	Sentence        string
	User            *User
	//StructuredInput *StructuredInput
	Stems           []string
	Plugin          string
	CreatedAt       *time.Time
	//determines if msg is from the user or Asha
	AshaSent      bool
	NeedsTraining bool
	Trained       bool
	// Tokens breaks the sentence into words. Tokens like ,.' are treated as
	// individual words.
	Tokens []string
	Route  string
}