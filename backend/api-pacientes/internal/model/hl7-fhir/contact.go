package hl7fhir

type Contact struct {
	RelationShip []RelationShip `bson:"relationship,omitempty"`
	Name         []Name         `bson:"name,omitempty"`
	Telecom      []Telecom      `bson:"telecom,omitempty"`
	Address      Address        `bson:"address,omitempty"`
	Gender       string         `bson:"gender,omitempty"`
	Period       Period         `bson:"period,omitempty"`
}
