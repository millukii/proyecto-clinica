package hl7fhir

type Coding struct {
	System string `bson:"system,omitempty"`
	Code   string `bson:"code,omitempty"`
}
