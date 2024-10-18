package iabconsent_test

import (
	"github.com/LiveRamp/iabconsent"
)

// Test fixtures can be created here: https://iabgpp.com/
// TODO: (PXS-2413) Add test cases for the new states
var gppParsedConsentFixtures = map[string]map[int]*iabconsent.MspaParsedConsent{
	// Valid GPP w/ US National MSPA, No Subsection (is thesame as false GPC subsection.
	"DBABLA~BVVqAAEABCA": {iabconsent.UsNationalSID: mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"]},
	// Valid GPP w/ US National MSPA, Subsection of GPC False.
	"DBABLA~BVVqAAEABCA.QA": {iabconsent.UsNationalSID: mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"]},
	// Valid GPP w/ US National MSPA, Subsection of GPC True.
	"DBABLA~BVVqAAEABCA.YA": {iabconsent.UsNationalSID: mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.YA"]},
	// Valid GPP w/ US California MSPA, Subsection of GPC False.
	"DBABBg~BVoYYZoI": {iabconsent.UsCaliforniaSID: mspaConsentFixtures[iabconsent.UsCaliforniaSID]["BVoYYZoI"]},
	// Valid GPP w/ US Virginia MSPA, Subsection of GPC False.
	"DBABRg~BVoYYYI": {iabconsent.UsVirginiaSID: mspaConsentFixtures[iabconsent.UsVirginiaSID]["BVoYYYI"]},
	// Valid GPP w/ US Colorado MSPA, Subsection of GPC False.
	"DBABJg~BVoYYQg": {iabconsent.UsColoradoSID: mspaConsentFixtures[iabconsent.UsColoradoSID]["BVoYYQg"]},
	// Valid GPP w/ US Utah MSPA, Subsection of GPC False.
	"DBABFg~BVaGGGCA": {iabconsent.UsUtahSID: mspaConsentFixtures[iabconsent.UsUtahSID]["BVaGGGCA.QA"]},
	// Valid GPP w/ US Connecticut MSPA, Subsection of GPC False.
	"DBABVg~BVoYYYQg": {iabconsent.UsConnecticutSID: mspaConsentFixtures[iabconsent.UsConnecticutSID]["BVoYYYQg"]},
	// Valid GPP w/ US US National and Virgina MSPA, Subsection of GPC False.
	"DBACLMA~BVVqAAEABCA~BVoYYYI": {
		iabconsent.UsNationalSID: mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"],
		iabconsent.UsVirginiaSID: mspaConsentFixtures[iabconsent.UsVirginiaSID]["BVoYYYI"],
	},
	// Valid GPP w/ US US National, California MSPA, Virgina MSPA, Colorado MSPA, and Utah Subsection of GPC False.
	"DBABrGA~BVVqAAEABCA~BVoYYZoI~BVoYYYI~BVoYYQg~BVaGGGCA~BVoYYYQg": {
		iabconsent.UsNationalSID:    mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"],
		iabconsent.UsCaliforniaSID:  mspaConsentFixtures[iabconsent.UsCaliforniaSID]["BVoYYZoI"],
		iabconsent.UsVirginiaSID:    mspaConsentFixtures[iabconsent.UsVirginiaSID]["BVoYYYI"],
		iabconsent.UsColoradoSID:    mspaConsentFixtures[iabconsent.UsColoradoSID]["BVoYYQg"],
		iabconsent.UsUtahSID:        mspaConsentFixtures[iabconsent.UsUtahSID]["BVaGGGCA.QA"],
		iabconsent.UsConnecticutSID: mspaConsentFixtures[iabconsent.UsConnecticutSID]["BVoYYYQg"],
	},
	// Valid GPP string w/ sections for EU TCF V2 and US Privacy
	// Since both are not supported, Consent fixture should be blank.
	"DBACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN": {},
	// Valid GPP w/ US National MSPA and US Privacy, but skip US Privacy until supported.
	"DBABzw~1YNN~BVVqAAEABCA.QA": {7: mspaConsentFixtures[7]["BVVqAAEABCA.QA"]},
}
