package iabconsent_test

import (
	"github.com/LiveRamp/iabconsent"
)

// Test fixtures can be created here: https://iabgpp.com/
var gppParsedConsentFixtures = map[string]map[int]*iabconsent.MspaParsedConsent{
	//// Valid GPP w/ US National MSPA, No Subsection (is the same as false GPC subsection).
	//"DBABLA~BVVqAAEABCA": {iabconsent.UsNationalSID: mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"]},
	//// Valid GPP w/ US National MSPA, Subsection of GPC False.
	//"DBABLA~BVVqAAEABCA.QA": {iabconsent.UsNationalSID: mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"]},
	//// Valid GPP w/ US National MSPA, Subsection of GPC True.
	//"DBABLA~BVVqAAEABCA.YA": {iabconsent.UsNationalSID: mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.YA"]},
	//// Valid GPP w/ US California MSPA, Subsection of GPC False.
	//"DBABBg~BVoYYZoI": {iabconsent.UsCaliforniaSID: mspaConsentFixtures[iabconsent.UsCaliforniaSID]["BVoYYZoI"]},
	//// Valid GPP w/ US Virginia MSPA, Subsection of GPC False.
	//"DBABRg~BVoYYYI": {iabconsent.UsVirginiaSID: mspaConsentFixtures[iabconsent.UsVirginiaSID]["BVoYYYI"]},
	//// Valid GPP w/ US Colorado MSPA, Subsection of GPC False.
	//"DBABJg~BVoYYQg": {iabconsent.UsColoradoSID: mspaConsentFixtures[iabconsent.UsColoradoSID]["BVoYYQg"]},
	//// Valid GPP w/ US Utah MSPA, Subsection of GPC False.
	//"DBABFg~BVaGGGCA": {iabconsent.UsUtahSID: mspaConsentFixtures[iabconsent.UsUtahSID]["BVaGGGCA.QA"]},
	//// Valid GPP w/ US Connecticut MSPA, Subsection of GPC False.
	//"DBABVg~BVoYYYQg": {iabconsent.UsConnecticutSID: mspaConsentFixtures[iabconsent.UsConnecticutSID]["BVoYYYQg"]},
	//// Valid GPP w/ US US National and Virgina MSPA, Subsection of GPC False.
	//"DBACLMA~BVVqAAEABCA~BVoYYYI": {
	//	iabconsent.UsNationalSID: mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"],
	//	iabconsent.UsVirginiaSID: mspaConsentFixtures[iabconsent.UsVirginiaSID]["BVoYYYI"],
	//},
	//// Valid GPP w/ US US National, California MSPA, Virgina MSPA, Colorado MSPA, and Utah Subsection of GPC False.
	//"DBABrGA~BVVqAAEABCA~BVoYYZoI~BVoYYYI~BVoYYQg~BVaGGGCA~BVoYYYQg": {
	//	iabconsent.UsNationalSID:    mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"],
	//	iabconsent.UsCaliforniaSID:  mspaConsentFixtures[iabconsent.UsCaliforniaSID]["BVoYYZoI"],
	//	iabconsent.UsVirginiaSID:    mspaConsentFixtures[iabconsent.UsVirginiaSID]["BVoYYYI"],
	//	iabconsent.UsColoradoSID:    mspaConsentFixtures[iabconsent.UsColoradoSID]["BVoYYQg"],
	//	iabconsent.UsUtahSID:        mspaConsentFixtures[iabconsent.UsUtahSID]["BVaGGGCA.QA"],
	//	iabconsent.UsConnecticutSID: mspaConsentFixtures[iabconsent.UsConnecticutSID]["BVoYYYQg"],
	//},
	//// Valid GPP w/ US US National, California MSPA, Virgina MSPA, Colorado MSPA, Utah MSPA, Conneticut MSPA, Florida MSPA and Montana MSPA Subsection of GPC False.
	//"DBABrWA~BVVqAAEABCA~BVoYYZoI~BVoYYYI~BVoYYQg~BVaGGGCA~BVoYYYQg~Bqqqqqqo~Bqqqqqqo": {
	//	iabconsent.UsNationalSID:    mspaConsentFixtures[iabconsent.UsNationalSID]["BVVqAAEABCA.QA"],
	//	iabconsent.UsCaliforniaSID:  mspaConsentFixtures[iabconsent.UsCaliforniaSID]["BVoYYZoI"],
	//	iabconsent.UsVirginiaSID:    mspaConsentFixtures[iabconsent.UsVirginiaSID]["BVoYYYI"],
	//	iabconsent.UsColoradoSID:    mspaConsentFixtures[iabconsent.UsColoradoSID]["BVoYYQg"],
	//	iabconsent.UsUtahSID:        mspaConsentFixtures[iabconsent.UsUtahSID]["BVaGGGCA.QA"],
	//	iabconsent.UsConnecticutSID: mspaConsentFixtures[iabconsent.UsConnecticutSID]["BVoYYYQg"],
	//	iabconsent.UsFloridaSID:     mspaConsentFixtures[iabconsent.UsFloridaSID]["Bqqqqqqo"],
	//	iabconsent.UsMontanaSID:     mspaConsentFixtures[iabconsent.UsMontanaSID]["Bqqqqqqo"],
	//},
	////// Valid GPP string w/ sections for EU TCF V2 and US Privacy
	////// Since both are not supported, Consent fixture should be blank.
	////"DBACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN": {},
	////// Valid GPP w/ US National MSPA and US Privacy, but skip US Privacy until supported.
	////"DBABzw~1YNN~BVVqAAEABCA.QA": {7: mspaConsentFixtures[7]["BVVqAAEABCA.QA"]},
	//// Valid GPP w/ US Florida MSPA, Subsection of GPC False.
	//"DBABAw~Bqqqqqqo": {iabconsent.UsFloridaSID: mspaConsentFixtures[iabconsent.UsFloridaSID]["Bqqqqqqo"]},
	//// Valid GPP w/ US Montana MSPA, Subsection of GPC False.
	//"DBABQw~Bqqqqqqo": {iabconsent.UsMontanaSID: mspaConsentFixtures[iabconsent.UsMontanaSID]["Bqqqqqqo"]},
	//// Valid GPP w/ US Oregon MSPA, Subsection of GPC False.
	//"DBABIw~BqqqqqqqoA": {iabconsent.UsOregonSID: mspaConsentFixtures[iabconsent.UsOregonSID]["BqqqqqqqoA"]},
	//// Valid GPP w/ US Texas MSPA, Subsection of GPC False.
	//"DBABEw~BqqqqqqA": {iabconsent.UsTexasSID: mspaConsentFixtures[iabconsent.UsTexasSID]["BqqqqqqA"]},
	//// Valid GPP w/ US Delaware MSPA, Subsection of GPC False.
	//"DBABUw~BqqqqqqqoA": {iabconsent.UsDelawareSID: mspaConsentFixtures[iabconsent.UsDelawareSID]["BqqqqqqqoA"]},
	//// Valid GPP w/ US Iowa MSPA, Subsection of GPC False.
	//"DBABCw~BVVVVVVA": {iabconsent.UsIowaSID: mspaConsentFixtures[iabconsent.UsIowaSID]["BVVVVVVA"]},
	//// Valid GPP w/ US Nebraska MSPA, Subsection of GPC False.
	//"DBABSw~BmaqqqqA": {iabconsent.UsNebraskaSID: mspaConsentFixtures[iabconsent.UsNebraskaSID]["BmaqqqqA"]},
	//// Valid GPP w/ US New Hampshire MSPA, Subsection of GPC False.
	//"DBABKw~Bpmqqqqo": {iabconsent.UsNewHampshireSID: mspaConsentFixtures[iabconsent.UsNewHampshireSID]["Bpmqqqqo"]},
	//// Valid GPP w/ US New Jersey MSPA, Subsection of GPC False.
	//"DBABAYA~BlWqqqmaqA": {iabconsent.UsNewJerseySID: mspaConsentFixtures[iabconsent.UsNewJerseySID]["BlWqqqmaqA"]},
	//// Valid GPP w/ US Tennessee MSPA, Subsection of GPC False.
	//"DBABQYA~Bqqqqqo": {iabconsent.UsTennesseeSID: mspaConsentFixtures[iabconsent.UsTennesseeSID]["Bqqqqqo"]},

	//"DBABM~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA": {iabconsent.EuropeTCFv2SID: mspaConsentFixtures[iabconsent.EuropeTCFv2SID]["CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA"]},
}

var tcfeuParsedConsentFixtures = map[string]map[int]*iabconsent.V2ParsedConsent{
	"DBABM~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA": {iabconsent.EuropeTCFv2SID: tcfeuV2ConsentFixtures[iabconsent.EuropeTCFv2SID]["CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA"]},
	"DBABMA~CPzyhEAPzyhEAEXdtAENDXCwAP_AAH_AACiQI9AB4C5EQCFDcHJNAIoUAAQDQIhAAAAgAAABgYAACBoAAIwAAAAwAAAAAAoCAAAAIABAAAEAAAAAAAEAAAAAAAEAAEAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAiAAAAAIAEEAAAAACAAEAAAAAABAAAgAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAQQUgSAALAAqABcADIAHAAQAAkABkADQAHIAPAAfQBEAEUAJgATwApABfADMAGgAPwAhIBlAGWAOeAdwB3gEDgIOAhABFgCngF1AXmAyYBlgDPgGqgP3AgoAAAAA.YAAAAAAAAAAA": {iabconsent.EuropeTCFv2SID: tcfeuV2ConsentFixtures[iabconsent.EuropeTCFv2SID]["CPzyhEAPzyhEAEXdtAENDXCwAP_AAH_AACiQI9AB4C5EQCFDcHJNAIoUAAQDQIhAAAAgAAABgYAACBoAAIwAAAAwAAAAAAoCAAAAIABAAAEAAAAAAAEAAAAAAAEAAEAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAiAAAAAIAEEAAAAACAAEAAAAAABAAAgAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAQQUgSAALAAqABcADIAHAAQAAkABkADQAHIAPAAfQBEAEUAJgATwApABfADMAGgAPwAhIBlAGWAOeAdwB3gEDgIOAhABFgCngF1AXmAyYBlgDPgGqgP3AgoAAAAA.YAAAAAAAAAAA"]},
}

var ccpaParsedConsentFixtures = map[string]map[int]*iabconsent.CcpaParsedConsent{
	"DBABTA~1---": {iabconsent.UsPVSID: ccpaConsentFixtures[iabconsent.UsPVSID]["1---"]},
	"DBABTA~1YYY": {iabconsent.UsPVSID: ccpaConsentFixtures[iabconsent.UsPVSID]["1YYY"]},
	"DBABTA~1NYN": {iabconsent.UsPVSID: ccpaConsentFixtures[iabconsent.UsPVSID]["1NYN"]},
	"DBABTA~1NNN": {iabconsent.UsPVSID: ccpaConsentFixtures[iabconsent.UsPVSID]["1NNN"]},
}

var tcfcaParsedConsentFixtures = map[string]map[int]*iabconsent.V2CAParsedConsent{
	"DBABDA~CPzyhEAPzyhEAEXdtAENDXCgAf-AAP-AAAj0AHgLkRAIUNwck0AihQABANAiEAAICAAAAGBgAAIGgAAjAAAADAAAAAACgIAAAAgAEAAAQAAAAAAAQAAAAAAAQAAQAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIAAAAAgAQQAAAAAIAAQAACAAAEAACAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAABBBSBKAAsACoAFwAMgAcABAACQAGQANAAcgA8AB9AEQARQAmABPACkAF8AMwAaAA_ACEAFLAMoAywBzwDuAO8AgcBBwEIAIsAU8AuoC8wGTAMsAZ8A1UB-4EFA": {iabconsent.CanadaTCFSID: tcfcaConsentFixtures[iabconsent.CanadaTCFSID]["CPzyhEAPzyhEAEXdtAENDXCgAf-AAP-AAAj0AHgLkRAIUNwck0AihQABANAiEAAICAAAAGBgAAIGgAAjAAAADAAAAAACgIAAAAgAEAAAQAAAAAAAQAAAAAAAQAAQAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIAAAAAgAQQAAAAAIAAQAACAAAEAACAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAABBBSBKAAsACoAFwAMgAcABAACQAGQANAAcgA8AB9AEQARQAmABPACkAF8AMwAaAA_ACEAFLAMoAywBzwDuAO8AgcBBwEIAIsAU8AuoC8wGTAMsAZ8A1UB-4EFA"]},
}
