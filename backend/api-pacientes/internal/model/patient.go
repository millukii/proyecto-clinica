package model

import (
	hl7fhir "api-pacientes/internal/model/hl7-fhir"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID                  primitive.ObjectID   `bson:"_id,omitempty"`
	MedicalRecordNumber hl7fhir.Identifier   `bson:"medicalRecordNumber,omitempty"`
	Identifiers         []hl7fhir.Identifier `bson:"identifiers,omitempty"`
	Contacts            []hl7fhir.Contact    `bson:"contact,omitempty"`
	Address             []hl7fhir.Address    `bson:"address,omitempty"`
	Telecom             []hl7fhir.Telecom    `bson:"telecom,omitempty"`
	Name                []hl7fhir.Name       `bson:"name,omitempty"`
}
