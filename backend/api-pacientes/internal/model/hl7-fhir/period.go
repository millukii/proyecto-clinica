package hl7fhir

import "go.mongodb.org/mongo-driver/bson/primitive"

type Period struct {
	Start primitive.DateTime `bson:"start,omitempty"`
	End   primitive.DateTime `bson:"end,omitempty"`
}
