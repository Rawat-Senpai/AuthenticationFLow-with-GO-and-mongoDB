package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notes struct {
	NotesId          primitive.ObjectID `bson:"_id,omitempty" json:"NotesId"`
	NotesHeading     *string            `json:"heading"`
	NotesDescription *string            `json:"notesDescription"`
	TimeStamp        *string            `json:"timeStamp"`
	CreatedBy        string             `bson:"createdBy" json:"createdBy"`
}
