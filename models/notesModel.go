package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notes struct {
	NotesId          primitive.ObjectID `bson:"_notesId,,omitempty"`
	NotesHeading     *string            `json:"heading"`
	NotesDescription *string            `json:"notesDescription"`
	TimeStamp        *string            `json:"timeStamp"`
	CreatedBy        string             `json:"createdBy"`
}
