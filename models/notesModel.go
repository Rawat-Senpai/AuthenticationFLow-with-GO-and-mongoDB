package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notes struct {
	NotesId          primitive.ObjectID `bson:"_id,omitempty" json:"NotesId"`
	NotesHeading     *string            `bson:"heading"  json:"heading"`
	NotesDescription *string            `bson:"notesDescription" json:"notesDescription"`
	TimeStamp        *string            `bson:"timeStamp" json:"timeStamp"`
	CreatedBy        string             `bson:"createdBy" json:"createdBy"`
}
