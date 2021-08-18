package hl7fhir

type Telecom struct {
	Use       string `bson:"use,omitempty"`
	System    string `bson:"system,omitempty"`
	Value     string `bson:"value,omitempty"`
	Rank      int    `bson:"rank,omitempty"`
	Gender    string `bson:"gender,omitempty"`
	BirthDate string `bson:"birthdate,omitempty"`
}
