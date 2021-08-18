package hl7fhir

type Identifier struct {
	Use      string   `bson:"use,omitempty"`
	Type     Type     `bson:"type,omitempty"`
	System   string   `bson:"system,omitempty"`
	Value    string   `bson:"value,omitempty"`
	Period   Period   `bson:"period,omitempty"`
	Assigner Assigner `bson:"assigner,omitempty"`
}
