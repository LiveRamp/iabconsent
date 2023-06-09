package iabconsent_test

import (
	"github.com/LiveRamp/iabconsent"
)

// Test fixtures can be created here: https://iabgpp.com/

var gppParsedConsentFixtures = map[string]map[int]*iabconsent.MspaParsedConsent{
	// Valid GPP w/ US National MSPA, No Subsection (is thesame as false GPC subsection.
	"DBABLA~BVVqAAEABCA": {7: usNationalConsentFixtures["BVVqAAEABCA.QA"]},
	// Valid GPP w/ US National MSPA, Subsection of GPC False.
	"DBABLA~BVVqAAEABCA.QA": {7: usNationalConsentFixtures["BVVqAAEABCA.QA"]},
	// Valid GPP w/ US National MSPA, Subsection of GPC True.
	"DBABLA~BVVqAAEABCA.YA": {7: usNationalConsentFixtures["BVVqAAEABCA.YA"]},
	// Valid GPP w/ US California MSPA, Subsection of GPC False.
	"DBABBg~BVoYYZoI": {8: usCAConsentFixtures["BVoYYZoI"]},
	// Valid GPP w/ US Virginia MSPA, Subsection of GPC False.
	"DBABRg~BVoYYYI": {9: usVAConsentFixtures["BVoYYYI"]},
	// Valid GPP w/ US Colorado MSPA, Subsection of GPC False.
	"DBABJg~BVoYYQg": {10: usCOConsentFixtures["BVoYYQg"]},
	// Valid GPP w/ US Utah MSPA, Subsection of GPC False.
	"DBABFg~BVaGGGCA": {11: usUTConsentFixtures["BVaGGGCA.QA"]},
	// Valid GPP w/ US Connecticut MSPA, Subsection of GPC False.
	"DBABVg~BVoYYYQg": {12: usCTConsentFixtures["BVoYYYQg"]},
	// Valid GPP w/ US US National and Virgina MSPA, Subsection of GPC False.
	"DBACLMA~BVVqAAEABCA~BVoYYYI": {
		7: usNationalConsentFixtures["BVVqAAEABCA.QA"],
		9: usVAConsentFixtures["BVoYYYI"],
	},
	// Valid GPP w/ US US National, California MSPA, Virgina MSPA, Colorado MSPA, and Utah Subsection of GPC False.
	"DBABrGA~BVVqAAEABCA~BVoYYZoI~BVoYYYI~BVoYYQg~BVaGGGCA~BVoYYYQg": {
		7:  usNationalConsentFixtures["BVVqAAEABCA.QA"],
		8:  usCAConsentFixtures["BVoYYZoI"],
		9:  usVAConsentFixtures["BVoYYYI"],
		10: usCOConsentFixtures["BVoYYQg"],
		11: usUTConsentFixtures["BVaGGGCA.QA"],
		12: usCTConsentFixtures["BVoYYYQg"],
	},
	// Valid GPP string w/ sections for EU TCF V2 and US Privacy
	// Since both are not supported, Consent fixture should be blank.
	"DBACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN": {},
	// Valid GPP w/ US National MSPA and US Privacy, but skip US Privacy until supported.
	"DBABzw~1YNN~BVVqAAEABCA.QA": {7: usNationalConsentFixtures["BVVqAAEABCA.QA"]},
}
