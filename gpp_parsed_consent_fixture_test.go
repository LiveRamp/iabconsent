package iabconsent_test

import (
	"github.com/LiveRamp/iabconsent"
)

// Need to figure out how to create these test fixtures more easily.

var gppParsedConsentFixtures = map[string]map[int]*iabconsent.MspaParsedConsent{
	// Valid GPP w/ US National MSPA
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
