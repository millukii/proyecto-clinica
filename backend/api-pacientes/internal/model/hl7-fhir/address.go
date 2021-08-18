package hl7fhir

type Address struct {
	Use        string   `bson:"use,omitempty"`
	Type       string   `bson:"type,omitempty"`
	Text       string   `bson:"text,omitempty"`
	Line       []string `bson:"line,omitempty"`
	City       string   `bson:"city,omitempty"`
	District   string   `bson:"district,omitempty"`
	State      string   `bson:"state,omitempty"`
	PostalCode string   `bson:"postalCode,omitempty"`
	period     Period   `bson:"period,omitempty"`
}
