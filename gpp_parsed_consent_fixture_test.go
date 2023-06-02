package iabconsent_test

import (
	"github.com/LiveRamp/iabconsent"
)

// Test fixtures can be created here: https://iabgpp.com/

var gppParsedConsentFixtures = map[string]map[int]*iabconsent.MspaParsedConsent{
	// Valid GPP w/ US National MSPA, No Subsection.
	"DBABL~BVVqAAEABAA": {7: {
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
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
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
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
		Gpc:                     false,
	},
	},
	// Valid GPP w/ US National MSPA, Subsection of GPC False.
	"DBABL~BVVqAAEABAA.QA": {7: {
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
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
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
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
		Gpc:                     false,
	},
	},
	// Valid GPP w/ US National MSPA, Subsection of GPC True.
	"DBABL~BVVqAAEABAA.YA": {7: {
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
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
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
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
		Gpc:                     true,
	},
	},
	// Valid GPP w/ US Virgina MSPA, Subsection of GPC False.
	"DBABRg~BVoYYYA": {9: {
		Version:                         1,
		SharingNotice:                   iabconsent.NoticeProvided,
		SaleOptOutNotice:                iabconsent.NoticeProvided,
		TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
		SaleOptOut:                      iabconsent.NotOptedOut,
		TargetedAdvertisingOptOut:       iabconsent.NotOptedOut,
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
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
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
		Gpc:                     false,
	},
	},
	// Valid GPP w/ US US National and Virgina MSPA, Subsection of GPC False.
	"DBACLMA~BVVqAAEABAA~BVoYYYA": {7: {
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
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
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
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
		Gpc:                     false,
	},
		9: {
			Version:                         1,
			SharingNotice:                   iabconsent.NoticeProvided,
			SaleOptOutNotice:                iabconsent.NoticeProvided,
			TargetedAdvertisingOptOutNotice: iabconsent.NoticeProvided,
			SaleOptOut:                      iabconsent.NotOptedOut,
			TargetedAdvertisingOptOut:       iabconsent.NotOptedOut,
			SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
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
			MspaServiceProviderMode: iabconsent.MspaNotApplicable,
			Gpc:                     false,
		},
	},
	// Valid GPP string w/ sections for EU TCF V2 and US Privacy
	// Since both are not supported, Consent fixture should be blank.
	"DBACNY~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA~1YNN": {},
	// Valid GPP w/ US National MSPA and US Privacy, but skip US Privacy until supported.
	"DBABzw~1YNN~BVVqAAEABAA.QA": {7: {
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
		SensitiveDataProcessing: map[int]iabconsent.MspaConsent{
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
		MspaServiceProviderMode: iabconsent.MspaNotApplicable,
	},
	},
}
