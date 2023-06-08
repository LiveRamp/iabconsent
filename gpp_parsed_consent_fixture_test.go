package iabconsent_test

import (
	"github.com/LiveRamp/iabconsent"
)

// Test fixtures can be created here: https://iabgpp.com/

var gppParsedConsentFixtures = map[string]map[int]*iabconsent.MspaParsedConsent{
	// Valid GPP w/ US National MSPA, No Subsection (is thesame as false GPC subsection.
	"DBABL~BVVqAAEABAA": {7: usNationalConsentFixtures["BVVqAAEABAA.QA"]},
	// Valid GPP w/ US National MSPA, Subsection of GPC False.
	"DBABL~BVVqAAEABAA.QA": {7: usNationalConsentFixtures["BVVqAAEABAA.QA"]},
	// Valid GPP w/ US National MSPA, Subsection of GPC True.
	"DBABL~BVVqAAEABAA.YA": {7: usNationalConsentFixtures["BVVqAAEABAA.YA"]},
	// Valid GPP w/ US California MSPA, Subsection of GPC False.
	"DBABBg~BVoYYZoA": {8: usCAConsentFixtures["BVoYYZoA"]},
	// Valid GPP w/ US Virginia MSPA, Subsection of GPC False.
	"DBABRg~BVoYYYA": {9: usVAConsentFixtures["BVoYYYA"]},
	// Valid GPP w/ US Colorado MSPA, Subsection of GPC False.
	"DBABJg~BVoYYQA": {10: usCOConsentFixtures["BVoYYQA"]},
	// Valid GPP w/ US Utah MSPA, Subsection of GPC False.
	"DBABFg~BVaGGGAA": {11: usUTConsentFixtures["BVaGGGAA.QA"]},
	// Valid GPP w/ US US National and Virgina MSPA, Subsection of GPC False.
	"DBACLMA~BVVqAAEABAA~BVoYYYA": {
		7: usNationalConsentFixtures["BVVqAAEABAA.QA"],
		9: usVAConsentFixtures["BVoYYYA"],
	},
	// Valid GPP w/ US US National, California MSPA, Virgina MSPA, Colorado MSPA, and Utah Subsection of GPC False.
	"DBABrsA~BVVqAAEABAA~BVoYYZoA~BVoYYYA~BVoYYQA~BVaGGGAA": {
		7:  usNationalConsentFixtures["BVVqAAEABAA.QA"],
		8:  usCAConsentFixtures["BVoYYZoA"],
		9:  usVAConsentFixtures["BVoYYYA"],
		10: usCOConsentFixtures["BVoYYQA"],
		11: usUTConsentFixtures["BVaGGGAA.QA"],
	},
	// Valid GPP string w/ sections for EU TCF V2 and US Privacy
	// Since both are not supported, Consent fixture should be blank.
	"DBACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN": {},
	// Valid GPP w/ US National MSPA and US Privacy, but skip US Privacy until supported.
	"DBABzw~1YNN~BVVqAAEABAA.QA": {7: usNationalConsentFixtures["BVVqAAEABAA.QA"]},
}
