package hl7fhir

type Participant struct {
	Type []ParticipantRole `bson:"type,omitempty"`
	//Patient, Practitioner,PractitionerRole,RelatedPerson,Device,HealthCareService,Location
	Actor string `bson:"actor,omitempty"`
	//required, optional, information-only
	Required string `bson:"required,omitempty"`
	//accpeted, declined, tentative, needs-action
	Status string `bson:"status,omitempty"`
	Period Period `bson:"period,omitempty"`
}
