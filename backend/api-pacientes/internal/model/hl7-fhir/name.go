package hl7fhir

type Name struct {
	Use    string   `bson:"use,omitempty"`
	Family string   `bson:"family,omitempty"`
	Given  []string `bson:"use,omitempty"`
	Period Period   `bson:"period,omitempty"`
}
