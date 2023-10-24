package iabconsent_test

import (
	"github.com/StackAdapt/iabconsent"
)

// Test fixtures can be created here: https://iabgpp.com/

// var mspaConsentFixtures = map[string]map[string]int{"yes": {"dog": 1}}

var mspaConsentFixtures = map[int]map[string]*iabconsent.MspaParsedConsent{
	// UsNational
	iabconsent.UsNationalSID: {
		// usnat without false GPC subsection.
		"BVVqAAEABCA.QA": {
			Version:                             1,
			SharingNotice:                       iabconsent.NoticeProvided,
			SaleOptOutNotice:                    iabconsent.NoticeProvided,
			SharingOptOutNotice:                 iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice:     iabconsent.NoticeProvided,
			SensitiveDataProcessingOptOutNotice: iabconsent.NoticeProvided,
			SensitiveDataLimitUseNotice:         iabconsent.NoticeProvided,
			SaleOptOut:                          iabconsent.NotOptedOut,
			SharingOptOut:                       iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:           iabconsent.NotOptedOut,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0:  iabconsent.ConsentNotApplicable,
				1:  iabconsent.ConsentNotApplicable,
				2:  iabconsent.ConsentNotApplicable,
				3:  iabconsent.ConsentNotApplicable,
				4:  iabconsent.ConsentNotApplicable,
				5:  iabconsent.ConsentNotApplicable,
				6:  iabconsent.ConsentNotApplicable,
				7:  iabconsent.NoConsent,
				8:  iabconsent.ConsentNotApplicable,
				9:  iabconsent.ConsentNotApplicable,
				10: iabconsent.ConsentNotApplicable,
				11: iabconsent.ConsentNotApplicable,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.ConsentNotApplicable,
			},
			PersonalDataConsents:    iabconsent.NoConsent,
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
		},
		// usnat with true GPC subsection.
		"BVVqAAEABCA.YA": {
			Version:                             1,
			SharingNotice:                       iabconsent.NoticeProvided,
			SaleOptOutNotice:                    iabconsent.NoticeProvided,
			SharingOptOutNotice:                 iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice:     iabconsent.NoticeProvided,
			SensitiveDataProcessingOptOutNotice: iabconsent.NoticeProvided,
			SensitiveDataLimitUseNotice:         iabconsent.NoticeProvided,
			SaleOptOut:                          iabconsent.NotOptedOut,
			SharingOptOut:                       iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:           iabconsent.NotOptedOut,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0:  iabconsent.ConsentNotApplicable,
				1:  iabconsent.ConsentNotApplicable,
				2:  iabconsent.ConsentNotApplicable,
				3:  iabconsent.ConsentNotApplicable,
				4:  iabconsent.ConsentNotApplicable,
				5:  iabconsent.ConsentNotApplicable,
				6:  iabconsent.ConsentNotApplicable,
				7:  iabconsent.NoConsent,
				8:  iabconsent.ConsentNotApplicable,
				9:  iabconsent.ConsentNotApplicable,
				10: iabconsent.ConsentNotApplicable,
				11: iabconsent.ConsentNotApplicable,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.ConsentNotApplicable,
			},
			PersonalDataConsents:    iabconsent.NoConsent,
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     true,
		},
		// usnat without subsection.
		"BqqAqqqqqqA": {
			Version:                             1,
			SharingNotice:                       iabconsent.NoticeNotProvided,
			SaleOptOutNotice:                    iabconsent.NoticeNotProvided,
			SharingOptOutNotice:                 iabconsent.NoticeNotProvided,
			TargetedAdvertisingOptOutNotice:     iabconsent.NoticeNotProvided,
			SensitiveDataProcessingOptOutNotice: iabconsent.NoticeNotProvided,
			SensitiveDataLimitUseNotice:         iabconsent.NoticeNotProvided,
			SaleOptOut:                          iabconsent.OptOutNotApplicable,
			SharingOptOut:                       iabconsent.OptOutNotApplicable,
			TargetedAdvertisingOptOut:           iabconsent.OptOutNotApplicable,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0:  iabconsent.Consent,
				1:  iabconsent.Consent,
				2:  iabconsent.Consent,
				3:  iabconsent.Consent,
				4:  iabconsent.Consent,
				5:  iabconsent.Consent,
				6:  iabconsent.Consent,
				7:  iabconsent.Consent,
				8:  iabconsent.Consent,
				9:  iabconsent.Consent,
				10: iabconsent.Consent,
				11: iabconsent.Consent,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.Consent,
				1: iabconsent.Consent,
			},
			PersonalDataConsents:    iabconsent.Consent,
			MspaCoveredTransaction:  iabconsent.MspaNo,
			MspaOptOutOptionMode:    iabconsent.MspaNo,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     false,
		},
	},
	// California
	iabconsent.UsCaliforniaSID: {
		// usca with subsection of GPC True.
		"BVoYYZoI.YA": {
			Version:                     1,
			SaleOptOutNotice:            iabconsent.NoticeProvided,
			SharingOptOutNotice:         iabconsent.NoticeProvided,
			SensitiveDataLimitUseNotice: iabconsent.NoticeProvided,
			SaleOptOut:                  iabconsent.NotOptedOut,
			SharingOptOut:               iabconsent.NotOptedOut,
			SensitiveDataProcessingOptOuts: map[int]iabconsent.MspaOptout{
				0: iabconsent.OptOutNotApplicable,
				1: iabconsent.OptedOut,
				2: iabconsent.NotOptedOut,
				3: iabconsent.OptOutNotApplicable,
				4: iabconsent.OptedOut,
				5: iabconsent.NotOptedOut,
				6: iabconsent.OptOutNotApplicable,
				7: iabconsent.OptedOut,
				8: iabconsent.NotOptedOut,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.NoConsent,
				1: iabconsent.Consent,
			},
			PersonalDataConsents:    iabconsent.Consent,
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     true,
		},
		// usca without subsection.
		"BVoYYZoI": {
			Version:                     1,
			SaleOptOutNotice:            iabconsent.NoticeProvided,
			SharingOptOutNotice:         iabconsent.NoticeProvided,
			SensitiveDataLimitUseNotice: iabconsent.NoticeProvided,
			SaleOptOut:                  iabconsent.NotOptedOut,
			SharingOptOut:               iabconsent.NotOptedOut,
			SensitiveDataProcessingOptOuts: map[int]iabconsent.MspaOptout{
				0: iabconsent.OptOutNotApplicable,
				1: iabconsent.OptedOut,
				2: iabconsent.NotOptedOut,
				3: iabconsent.OptOutNotApplicable,
				4: iabconsent.OptedOut,
				5: iabconsent.NotOptedOut,
				6: iabconsent.OptOutNotApplicable,
				7: iabconsent.OptedOut,
				8: iabconsent.NotOptedOut,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.NoConsent,
				1: iabconsent.Consent,
			},
			PersonalDataConsents:    iabconsent.Consent,
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     false,
		},
	},
	// Virginia
	iabconsent.UsVirginiaSID: {
		// usva with subsection of GPC True.
		"BVoYYYI.YA": {
			Version:                         1,
			SharingNotice:                   iabconsent.NoticeProvided,
			SaleOptOutNotice:                iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                      iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:       iabconsent.NotOptedOut,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.NoConsent,
				2: iabconsent.Consent,
				3: iabconsent.ConsentNotApplicable,
				4: iabconsent.NoConsent,
				5: iabconsent.Consent,
				6: iabconsent.ConsentNotApplicable,
				7: iabconsent.NoConsent,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.Consent,
			},
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     true,
		},
		// usva without subsection.
		"BVoYYYI": {
			Version:                         1,
			SharingNotice:                   iabconsent.NoticeProvided,
			SaleOptOutNotice:                iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                      iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:       iabconsent.NotOptedOut,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.NoConsent,
				2: iabconsent.Consent,
				3: iabconsent.ConsentNotApplicable,
				4: iabconsent.NoConsent,
				5: iabconsent.Consent,
				6: iabconsent.ConsentNotApplicable,
				7: iabconsent.NoConsent,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.Consent,
			},
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     false,
		},
	},
	// Colorado
	iabconsent.UsColoradoSID: {
		// usco without GPC Subsection
		"BVoYYQg": {
			Version:                         1,
			SharingNotice:                   iabconsent.NoticeProvided,
			SaleOptOutNotice:                iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                      iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:       iabconsent.NotOptedOut,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.NoConsent,
				2: iabconsent.Consent,
				3: iabconsent.ConsentNotApplicable,
				4: iabconsent.NoConsent,
				5: iabconsent.Consent,
				6: iabconsent.ConsentNotApplicable,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.NoConsent,
			},
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     false,
		},
		// usco with subsection of GPC True.
		"BVoYYQg.YA": {
			Version:                         1,
			SharingNotice:                   iabconsent.NoticeProvided,
			SaleOptOutNotice:                iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                      iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:       iabconsent.NotOptedOut,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.NoConsent,
				2: iabconsent.Consent,
				3: iabconsent.ConsentNotApplicable,
				4: iabconsent.NoConsent,
				5: iabconsent.Consent,
				6: iabconsent.ConsentNotApplicable,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.NoConsent,
			},
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     true,
		},
	},
	// Utah
	iabconsent.UsUtahSID: {
		// usut with subsection of GPC False.
		"BVaGGGCA.QA": {
			Version:                             1,
			SharingNotice:                       iabconsent.NoticeProvided,
			SaleOptOutNotice:                    iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice:     iabconsent.NoticeProvided,
			SensitiveDataProcessingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                          iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:           iabconsent.NotOptedOut,
			SensitiveDataProcessingOptOuts: map[int]iabconsent.MspaOptout{
				0: iabconsent.OptOutNotApplicable,
				1: iabconsent.OptedOut,
				2: iabconsent.NotOptedOut,
				3: iabconsent.OptOutNotApplicable,
				4: iabconsent.OptedOut,
				5: iabconsent.NotOptedOut,
				6: iabconsent.OptOutNotApplicable,
				7: iabconsent.OptedOut,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.Consent,
			},
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     false,
		},
		// usut with subsection of GPC True.
		"BVaGGGCA.YA": {
			Version:                             1,
			SharingNotice:                       iabconsent.NoticeProvided,
			SaleOptOutNotice:                    iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice:     iabconsent.NoticeProvided,
			SensitiveDataProcessingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                          iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:           iabconsent.NotOptedOut,
			SensitiveDataProcessingOptOuts: map[int]iabconsent.MspaOptout{
				0: iabconsent.OptOutNotApplicable,
				1: iabconsent.OptedOut,
				2: iabconsent.NotOptedOut,
				3: iabconsent.OptOutNotApplicable,
				4: iabconsent.OptedOut,
				5: iabconsent.NotOptedOut,
				6: iabconsent.OptOutNotApplicable,
				7: iabconsent.OptedOut,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.Consent,
			},
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     true,
		},
	},
	// Connecticut
	iabconsent.UsConnecticutSID: {
		// usct with subsection of GPC False.
		"BVoYYYQg": {
			Version:                         1,
			SharingNotice:                   iabconsent.NoticeProvided,
			SaleOptOutNotice:                iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                      iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:       iabconsent.NotOptedOut,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.NoConsent,
				2: iabconsent.Consent,
				3: iabconsent.ConsentNotApplicable,
				4: iabconsent.NoConsent,
				5: iabconsent.Consent,
				6: iabconsent.ConsentNotApplicable,
				7: iabconsent.NoConsent,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.Consent,
				1: iabconsent.ConsentNotApplicable,
				2: iabconsent.NoConsent,
			},
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     false,
		},
		// usct with subsection of GPC True.
		"BVoYYYQg.YA": {
			Version:                         1,
			SharingNotice:                   iabconsent.NoticeProvided,
			SaleOptOutNotice:                iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                      iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:       iabconsent.NotOptedOut,
			SensitiveDataProcessingConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.NoConsent,
				2: iabconsent.Consent,
				3: iabconsent.ConsentNotApplicable,
				4: iabconsent.NoConsent,
				5: iabconsent.Consent,
				6: iabconsent.ConsentNotApplicable,
				7: iabconsent.NoConsent,
			},
			KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
				0: iabconsent.Consent,
				1: iabconsent.ConsentNotApplicable,
				2: iabconsent.NoConsent,
			},
			MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
			MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
			MspaServiceProviderMode: iabconsent.MspaNo,
			Gpc:                     true,
		},
	},
}

var usCAConsentFixtures = map[string]*iabconsent.MspaParsedConsent{
	//"BVoAAACQ.QA": {
	//	Version:                     1,
	//	SaleOptOutNotice:            iabconsent.NoticeProvided,
	//	SharingOptOutNotice:         iabconsent.NoticeProvided,
	//	SensitiveDataLimitUseNotice: iabconsent.NoticeProvided,
	//	SaleOptOut:                  iabconsent.NotOptedOut,
	//	SharingOptOut:               iabconsent.NotOptedOut,
	//	SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
	//		0: iabconsent.ConsentNotApplicable,
	//		1: iabconsent.ConsentNotApplicable,
	//		2: iabconsent.ConsentNotApplicable,
	//		3: iabconsent.ConsentNotApplicable,
	//		4: iabconsent.ConsentNotApplicable,
	//		5: iabconsent.ConsentNotApplicable,
	//		6: iabconsent.ConsentNotApplicable,
	//		7: iabconsent.ConsentNotApplicable,
	//		8: iabconsent.ConsentNotApplicable,
	//	},
	//	KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
	//		0: iabconsent.ConsentNotApplicable,
	//		1: iabconsent.ConsentNotApplicable,
	//	},
	//	PersonalDataConsents:    iabconsent.ConsentNotApplicable,
	//	MspaCoveredTransaction:  iabconsent.MspaNo,
	//	MspaOptOutOptionMode:    iabconsent.MspaYes,
	//	MspaServiceProviderMode: iabconsent.MspaNotApplicable,
	//
	//	Gpc: false,
	//},
	//"BVoBAARQ.QA": {
	//	Version:                     1,
	//	SaleOptOutNotice:            iabconsent.NoticeProvided,
	//	SharingOptOutNotice:         iabconsent.NoticeProvided,
	//	SensitiveDataLimitUseNotice: iabconsent.NoticeProvided,
	//	SaleOptOut:                  iabconsent.NotOptedOut,
	//	SharingOptOut:               iabconsent.NotOptedOut,
	//	SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
	//		0: iabconsent.ConsentNotApplicable,
	//		1: iabconsent.ConsentNotApplicable,
	//		2: iabconsent.ConsentNotApplicable,
	//		3: iabconsent.NoConsent,
	//		4: iabconsent.ConsentNotApplicable,
	//		5: iabconsent.ConsentNotApplicable,
	//		6: iabconsent.ConsentNotApplicable,
	//		7: iabconsent.ConsentNotApplicable,
	//		8: iabconsent.ConsentNotApplicable,
	//	},
	//	KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
	//		0: iabconsent.ConsentNotApplicable,
	//		1: iabconsent.NoConsent,
	//	},
	//	PersonalDataConsents:    iabconsent.ConsentNotApplicable,
	//	MspaCoveredTransaction:  iabconsent.MspaYes,
	//	MspaOptOutOptionMode:    iabconsent.MspaYes,
	//	MspaServiceProviderMode: iabconsent.MspaNotApplicable,
	//},
	"BUoAAAAA": {
		Version:                     1,
		SaleOptOutNotice:            iabconsent.NoticeProvided,
		SharingOptOutNotice:         iabconsent.NoticeProvided,
		SensitiveDataLimitUseNotice: iabconsent.NoticeNotApplicable,
		SaleOptOut:                  iabconsent.NotOptedOut,
		SharingOptOut:               iabconsent.NotOptedOut,
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
			1: iabconsent.ConsentNotApplicable,
			2: iabconsent.ConsentNotApplicable,
			3: iabconsent.ConsentNotApplicable,
			4: iabconsent.ConsentNotApplicable,
			5: iabconsent.ConsentNotApplicable,
			6: iabconsent.ConsentNotApplicable,
			7: iabconsent.ConsentNotApplicable,
			8: iabconsent.ConsentNotApplicable,
		},
		KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
			1: iabconsent.ConsentNotApplicable,
		},
		PersonalDataConsents:    iabconsent.ConsentNotApplicable,
		MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
		MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
	},
}

var usCOConsentFixtures = map[string]*iabconsent.MspaParsedConsent{
	"BAAAAAA.QA": {
		Version:                         1,
		SharingNotice:                   iabconsent.NoticeNotApplicable,
		SaleOptOutNotice:                iabconsent.NoticeNotApplicable,
		TargetedAdvertisingOptOutNotice: iabconsent.NoticeNotApplicable,
		SaleOptOut:                      iabconsent.OptOutNotApplicable,
		TargetedAdvertisingOptOut:       iabconsent.OptOutNotApplicable,
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
			1: iabconsent.ConsentNotApplicable,
			2: iabconsent.ConsentNotApplicable,
			3: iabconsent.ConsentNotApplicable,
			4: iabconsent.ConsentNotApplicable,
			5: iabconsent.ConsentNotApplicable,
			6: iabconsent.ConsentNotApplicable,
		},
		KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
		},
		MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
		MspaOptOutOptionMode:    iabconsent.MspaNotApplicable,
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
		Gpc:                     false,
	},
	"BAABAFA.QA": {
		Version:                         1,
		SharingNotice:                   iabconsent.NoticeNotApplicable,
		SaleOptOutNotice:                iabconsent.NoticeNotApplicable,
		TargetedAdvertisingOptOutNotice: iabconsent.NoticeNotApplicable,
		SaleOptOut:                      iabconsent.OptOutNotApplicable,
		TargetedAdvertisingOptOut:       iabconsent.OptOutNotApplicable,
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
			1: iabconsent.ConsentNotApplicable,
			2: iabconsent.ConsentNotApplicable,
			3: iabconsent.NoConsent,
			4: iabconsent.ConsentNotApplicable,
			5: iabconsent.ConsentNotApplicable,
			6: iabconsent.ConsentNotApplicable,
		},
		KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
		},
		MspaCoveredTransaction:  iabconsent.MspaYes,
		MspaOptOutOptionMode:    iabconsent.MspaYes,
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
		Gpc:                     false,
	},
}

var usCTConsentFixtures = map[string]*iabconsent.MspaParsedConsent{
	"BBABAACQ.YA": {
		Version:                         1,
		SharingNotice:                   iabconsent.NoticeNotApplicable,
		SaleOptOutNotice:                iabconsent.NoticeNotApplicable,
		TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
		SaleOptOut:                      iabconsent.OptOutNotApplicable,
		TargetedAdvertisingOptOut:       iabconsent.OptOutNotApplicable,
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
			1: iabconsent.ConsentNotApplicable,
			2: iabconsent.ConsentNotApplicable,
			3: iabconsent.NoConsent,
			4: iabconsent.ConsentNotApplicable,
			5: iabconsent.ConsentNotApplicable,
			6: iabconsent.ConsentNotApplicable,
			7: iabconsent.ConsentNotApplicable,
		},
		KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
			1: iabconsent.ConsentNotApplicable,
			2: iabconsent.ConsentNotApplicable,
		},
		MspaCoveredTransaction:  iabconsent.MspaNotApplicable,
		MspaOptOutOptionMode:    iabconsent.MspaNo,
		MspaServiceProviderMode: iabconsent.MspaYes,
		Gpc:                     true,
	},
}

var usUTConsentFixtures = map[string]*iabconsent.MspaParsedConsent{
	"BGBEQRkA": {
		Version:                             1,
		SharingNotice:                       iabconsent.NoticeNotApplicable,
		SaleOptOutNotice:                    iabconsent.NoticeProvided,
		TargetedAdvertisingOptOutNotice:     iabconsent.NoticeNotProvided,
		SensitiveDataProcessingOptOutNotice: iabconsent.NoticeNotApplicable,
		SaleOptOut:                          iabconsent.OptOutNotApplicable,
		TargetedAdvertisingOptOut:           iabconsent.OptedOut,
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
			0: iabconsent.ConsentNotApplicable,
			1: iabconsent.NoConsent,
			2: iabconsent.ConsentNotApplicable,
			3: iabconsent.NoConsent,
			4: iabconsent.ConsentNotApplicable,
			5: iabconsent.ConsentNotApplicable,
			6: iabconsent.NoConsent,
			7: iabconsent.ConsentNotApplicable,
		},
		KnownChildSensitiveDataConsents: map[int]iabconsent.MspaConsent{0: iabconsent.NoConsent},
		MspaCoveredTransaction:          iabconsent.MspaNo,
		MspaOptOutOptionMode:            iabconsent.MspaYes,
		MspaServiceProviderMode:         iabconsent.MspaNotApplicable,
	},
}
