package models

import "time"

type CreateTweet struct {
	UserId  string    `bson:"userId" json:"userId,omitempty"`
	Message string    `bson:"message" json:"message"`
	Date    time.Time `bson:"time" json:"time,omitempty"`
}
