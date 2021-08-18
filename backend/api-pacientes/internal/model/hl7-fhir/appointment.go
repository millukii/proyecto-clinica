package hl7fhir

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	MedicalRecordNumber Identifier         `bson:"medicalRecordNumber,omitempty"`
	Identifiers         []Identifier       `bson:"identifiers,omitempty"`
	//proposed, pending, booked, arrived, fulfilled, cancelled, noshow, entered-in-error, checked-in, waitlist
	Status            string            `bson:"status,omitempty"`
	CancelationReason CancelationReason `bson:"cancelationReason,omitempty"`
	ServiceCategories []ServiceCategory `bson:"serviceCategories,omitempty"`
	ServiceTypes      []ServiceType     `bson:"serviceTypes,omitempty"`
	Specialty         []Specialty       `bson:"specialty,omitempty"`
	AppoinmentType    AppointmentType   `bson:"appointmentType,omitempty"`
	ReasonCode        []ReasonCode      `bson:"reasonCode,omitempty"`
	//Condition, Procedure, Observation, ImmunizationRecommendations
	ReasonReference     []string           `bson:"reasonReference,omitempty"`
	Priority            uint               `bson:"priority,omitempty"`
	Start               primitive.DateTime `bson:"start,omitempty"`
	End                 primitive.DateTime `bson:"end,omitempty"`
	MinutesDuration     int                `bson:"minutesDuration,omitempty"`
	Slot                string             `bson:"slot,omitempty"`
	Created             primitive.DateTime `bson:"created,omitempty"`
	Comment             string             `bson:"comment,omitempty"`
	PatientInstructions string             `bson:"patientInstructions,omitempty"`
	Participants        []Participant      `bson:"participants,omitempty"`
	RequestedPeriod     []Period           `bson:"requestedPeriod,omitempty"`
}
