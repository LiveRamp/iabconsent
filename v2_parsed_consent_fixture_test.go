package iabconsent_test

import (
	"time"

	"github.com/LiveRamp/iabconsent"
)

var v2TestTime = time.Unix(1583436280, 9*nsPerDs).UTC()

// Consent fixtures have been generated using the tool at https://iabtcf.com/encode.
// To inspect a consent string used in the test u can enter it into: https://iabtcf.com/#/decode.

var v2ConsentFixtures = map[string]*iabconsent.V2ParsedConsent{
	// COvzTO5OvzTO5B7ABCENAPCYAKdAADkAAIqIFhwBAAGAAXAFGAsMAhYAgAMAAegBYAEKAAA.IFoEUQQgAIQwgIwQABAEAAAAOIAACAIAAAAQAIAgEAACEAAAAAgAQBAAAAAAAGBAAgAAAAAAAFAAECAAAgAAQARAEQAAAAAJAAIAAgAAAYQEAAAQmAgBC3ZAYzUw.QE5QAwCvgHyATkA
	"COvzTO5OvzTO5B7ABCENAPCYAKdAADkAAIqIFhwBAAGAAXAFGAsMAhYAgAMAAegBYAEKAAA.IFoEUQQgAIQwgIwQABAEAAAAOIAACAIAAAAQAIAgEAACEAAAAAgAQBAAAAAAAGBAAgAAAAAAAFAAECAAAgAAQARAEQAAAAAJAAIAAgAAAYQEAAAQmAgBC3ZAYzUw.QE5QAwCvgHyATkA": {
		Version:              2,
		Created:              v2TestTime,
		LastUpdated:          v2TestTime,
		CMPID:                123,
		CMPVersion:           1,
		ConsentScreen:        2,
		ConsentLanguage:      "EN",
		VendorListVersion:    15,
		TCFPolicyVersion:     2,
		IsServiceSpecific:    false,
		UseNonStandardStacks: true,
		SpecialFeaturesOptIn: map[int]bool{1: true},
		PurposesConsent: map[int]bool{
			1:  true,
			3:  true,
			6:  true,
			7:  true,
			8:  true,
			10: true,
		},
		PurposesLITransparency: map[int]bool{
			3: true,
			4: true,
			5: true,
			8: true,
		},
		PurposeOneTreatment:    true,
		PublisherCC:            "FR",
		MaxConsentVendorID:     707,
		IsConsentRangeEncoding: true,
		ConsentedVendors:       nil,
		NumConsentEntries:      4,
		ConsentedVendorsRange: []*iabconsent.RangeEntry{
			{StartVendorID: 12, EndVendorID: 12},
			{StartVendorID: 23, EndVendorID: 23},
			{StartVendorID: 163, EndVendorID: 163},
			{StartVendorID: 707, EndVendorID: 707},
		},
		MaxInterestsVendorID:     133,
		IsInterestsRangeEncoding: true,
		InterestsVendors:         nil,
		NumInterestsEntries:      4,
		InterestsVendorsRange: []*iabconsent.RangeEntry{
			{StartVendorID: 48, EndVendorID: 48},
			{StartVendorID: 61, EndVendorID: 61},
			{StartVendorID: 88, EndVendorID: 88},
			{StartVendorID: 133, EndVendorID: 133},
		},
		NumPubRestrictions:    0,
		PubRestrictionEntries: make([]*iabconsent.PubRestrictionEntry, 0),
		OOBDisclosedVendors: &iabconsent.OOBVendorList{
			SegmentType:     1,
			MaxVendorID:     720,
			IsRangeEncoding: false,
			Vendors: map[int]bool{
				2:   true,
				6:   true,
				8:   true,
				12:  true,
				18:  true,
				23:  true,
				37:  true,
				42:  true,
				47:  true,
				48:  true,
				53:  true,
				61:  true,
				65:  true,
				66:  true,
				72:  true,
				88:  true,
				98:  true,
				127: true,
				128: true,
				129: true,
				133: true,
				153: true,
				163: true,
				192: true,
				205: true,
				215: true,
				224: true,
				243: true,
				248: true,
				281: true,
				294: true,
				304: true,
				350: true,
				351: true,
				358: true,
				371: true,
				422: true,
				424: true,
				440: true,
				447: true,
				467: true,
				486: true,
				498: true,
				502: true,
				512: true,
				516: true,
				553: true,
				556: true,
				571: true,
				587: true,
				612: true,
				613: true,
				618: true,
				626: true,
				648: true,
				653: true,
				656: true,
				657: true,
				665: true,
				676: true,
				681: true,
				683: true,
				684: true,
				686: true,
				687: true,
				688: true,
				690: true,
				691: true,
				694: true,
				702: true,
				703: true,
				707: true,
				708: true,
				711: true,
				712: true,
				714: true,
				716: true,
				719: true,
				720: true,
			},
		},
		OOBAllowedVendors: &iabconsent.OOBVendorList{
			SegmentType:     2,
			MaxVendorID:     626,
			IsRangeEncoding: true,
			NumEntries:      3,
			VendorEntries: []*iabconsent.RangeEntry{
				{StartVendorID: 351, EndVendorID: 351},
				{StartVendorID: 498, EndVendorID: 498},
				{StartVendorID: 626, EndVendorID: 626},
			},
		},
		PublisherTCEntry: nil,
	},
	// COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFAAA
	// COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFADBQAQA9hAAAcAA
	"COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFADBQAQA9hAAAcAA": {
		Version:                2,
		Created:                v2TestTime,
		LastUpdated:            v2TestTime,
		CMPID:                  81,
		CMPVersion:             0,
		ConsentScreen:          0,
		ConsentLanguage:        "EN",
		VendorListVersion:      15,
		TCFPolicyVersion:       2,
		IsServiceSpecific:      true,
		UseNonStandardStacks:   false,
		SpecialFeaturesOptIn:   map[int]bool{1: true},
		PurposesConsent:        map[int]bool{1: true, 3: true, 4: true, 7: true},
		PurposesLITransparency: map[int]bool{3: true, 4: true, 5: true},
		PurposeOneTreatment:    false,
		PublisherCC:            "AA",
		MaxConsentVendorID:     61,
		IsConsentRangeEncoding: true,
		NumConsentEntries:      3,
		ConsentedVendorsRange: []*iabconsent.RangeEntry{
			{StartVendorID: 2, EndVendorID: 2},
			{StartVendorID: 37, EndVendorID: 37},
			{StartVendorID: 61, EndVendorID: 61},
		},
		MaxInterestsVendorID:     8,
		IsInterestsRangeEncoding: false,
		InterestsVendors: map[int]bool{
			2: true,
			6: true,
			8: true,
		},
		NumPubRestrictions: 3,
		PubRestrictionEntries: []*iabconsent.PubRestrictionEntry{
			{
				PurposeID:       1,
				RestrictionType: iabconsent.RequireConsent,
				NumEntries:      1,
				RestrictionsRange: []*iabconsent.RangeEntry{
					{
						StartVendorID: 123,
						EndVendorID:   123,
					},
				},
			},
			{
				PurposeID:         2,
				RestrictionType:   iabconsent.PurposeFlatlyNotAllowed,
				RestrictionsRange: []*iabconsent.RangeEntry{},
			},
			{
				PurposeID:         3,
				RestrictionType:   iabconsent.RequireLegitimateInterest,
				RestrictionsRange: []*iabconsent.RangeEntry{},
			},
		},
	},
	// COvzTO5OvzTO5BZAFMENAPCgAAAAAAAAAAwIFoEUQQgAIQwgIwQABAEAAAAOIAACAIAAAAQAIAgEAACEAAAAAgAQBAAAAAAAGBAAgAAAAAAAFAAECAAAgAAQARAEQAAAAAJAAIAAgAAAYQEAAAQmAgBC3ZAYzUwLQIoghAAQhhARggACAIAAAAcQAAEAQAAAAgAQBAIAAEIAAAABAAgCAAAAAAAMCABAAAAAAAAKAAIEAABAAAgAiAIgAAAAASAAQABAAAAwgIAAAhMBACFuyAxmpgAA
	"COvzTO5OvzTO5BZAFMENAPCgAAAAAAAAAAwIFoEUQQgAIQwgIwQABAEAAAAOIAACAIAAAAQAIAgEAACEAAAAAgAQBAAAAAAAGBAAgAAAAAAAFAAECAAAgAAQARAEQAAAAAJAAIAAgAAAYQEAAAQmAgBC3ZAYzUwLQIoghAAQhhARggACAIAAAAcQAAEAQAAAAgAQBAIAAEIAAAABAAgCAAAAAAAMCABAAAAAAAAKAAIEAABAAAgAiAIgAAAAASAAQABAAAAwgIAAAhMBACFuyAxmpgAA": {
		Version:                2,
		Created:                v2TestTime,
		LastUpdated:            v2TestTime,
		CMPID:                  89,
		CMPVersion:             5,
		ConsentScreen:          12,
		ConsentLanguage:        "EN",
		VendorListVersion:      15,
		TCFPolicyVersion:       2,
		IsServiceSpecific:      true,
		UseNonStandardStacks:   false,
		SpecialFeaturesOptIn:   map[int]bool{},
		PurposesConsent:        map[int]bool{},
		PurposesLITransparency: map[int]bool{},
		PurposeOneTreatment:    false,
		PublisherCC:            "GB",
		MaxConsentVendorID:     720,
		IsConsentRangeEncoding: false,
		ConsentedVendors: map[int]bool{
			2:   true,
			6:   true,
			8:   true,
			12:  true,
			18:  true,
			23:  true,
			37:  true,
			42:  true,
			47:  true,
			48:  true,
			53:  true,
			61:  true,
			65:  true,
			66:  true,
			72:  true,
			88:  true,
			98:  true,
			127: true,
			128: true,
			129: true,
			133: true,
			153: true,
			163: true,
			192: true,
			205: true,
			215: true,
			224: true,
			243: true,
			248: true,
			281: true,
			294: true,
			304: true,
			350: true,
			351: true,
			358: true,
			371: true,
			422: true,
			424: true,
			440: true,
			447: true,
			467: true,
			486: true,
			498: true,
			502: true,
			512: true,
			516: true,
			553: true,
			556: true,
			571: true,
			587: true,
			612: true,
			613: true,
			618: true,
			626: true,
			648: true,
			653: true,
			656: true,
			657: true,
			665: true,
			676: true,
			681: true,
			683: true,
			684: true,
			686: true,
			687: true,
			688: true,
			690: true,
			691: true,
			694: true,
			702: true,
			703: true,
			707: true,
			708: true,
			711: true,
			712: true,
			714: true,
			716: true,
			719: true,
			720: true,
		},
		MaxInterestsVendorID:     720,
		IsInterestsRangeEncoding: false,
		InterestsVendors: map[int]bool{
			2:   true,
			6:   true,
			8:   true,
			12:  true,
			18:  true,
			23:  true,
			37:  true,
			42:  true,
			47:  true,
			48:  true,
			53:  true,
			61:  true,
			65:  true,
			66:  true,
			72:  true,
			88:  true,
			98:  true,
			127: true,
			128: true,
			129: true,
			133: true,
			153: true,
			163: true,
			192: true,
			205: true,
			215: true,
			224: true,
			243: true,
			248: true,
			281: true,
			294: true,
			304: true,
			350: true,
			351: true,
			358: true,
			371: true,
			422: true,
			424: true,
			440: true,
			447: true,
			467: true,
			486: true,
			498: true,
			502: true,
			512: true,
			516: true,
			553: true,
			556: true,
			571: true,
			587: true,
			612: true,
			613: true,
			618: true,
			626: true,
			648: true,
			653: true,
			656: true,
			657: true,
			665: true,
			676: true,
			681: true,
			683: true,
			684: true,
			686: true,
			687: true,
			688: true,
			690: true,
			691: true,
			694: true,
			702: true,
			703: true,
			707: true,
			708: true,
			711: true,
			712: true,
			714: true,
			716: true,
			719: true,
			720: true,
		},
		NumPubRestrictions:    0,
		PubRestrictionEntries: make([]*iabconsent.PubRestrictionEntry, 0),
	},
}

var v2CAConsentFixtures = map[string]*iabconsent.V2CAParsedConsent{
	"CPb_z8APb_z8AEXahAENDDCgAf-AAP-AAAjhArgAUABcADQAOAArABcAGQAOAAgABIAC0AGgAOoAegB8AEWAJgAmgBQACkAFsAMIAaIBBgEIAI4AUoArQBbgDKAHaAPEAg4BJQCdgFNAKeAdQBAAC8wGDgMZAZYA00BrgDdwH1ARmAjgBpsB4ABYADgAKgAXAAyABwAEAAJAAaAA-ACKAEwAKQAaAA_ACEAEcAKUAW4AygCDgEWAKeAa8A6gCwgF5gMsAZeA00A.YAAAAAAAAAA": {
		Version:                       2,
		Created:                       time.Date(2022, time.July, 12, 0, 0, 0, 0, time.UTC),
		LastUpdated:                   time.Date(2022, time.July, 12, 0, 0, 0, 0, time.UTC),
		CMPID:                         279,
		CMPVersion:                    1697,
		ConsentScreen:                 0,
		ConsentLanguage:               "EN",
		VendorListVersion:             195,
		TCFPolicyVersion:              2,
		UseNonStandardStacks:          true,
		SpecialFeatureExpressConsent:  map[int]bool{},
		PurposesExpressConsent:        map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true},
		PurposesImpliedConsent:        map[int]bool{2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true},
		MaxExpressConsentVendorID:     1136,
		IsExpressConsentRangeEncoding: true,
		ExpressConsentedVendors:       map[int]bool(nil),
		NumExpressConsentEntries:      43,
		VendorExpressConsent: []*iabconsent.RangeEntry{
			{StartVendorID: 10, EndVendorID: 11},
			{StartVendorID: 13, EndVendorID: 14},
			{StartVendorID: 21, EndVendorID: 21},
			{StartVendorID: 23, EndVendorID: 25},
			{StartVendorID: 28, EndVendorID: 28},
			{StartVendorID: 32, EndVendorID: 32},
			{StartVendorID: 36, EndVendorID: 36},
			{StartVendorID: 45, EndVendorID: 45},
			{StartVendorID: 52, EndVendorID: 52},
			{StartVendorID: 58, EndVendorID: 58},
			{StartVendorID: 61, EndVendorID: 62},
			{StartVendorID: 69, EndVendorID: 69},
			{StartVendorID: 76, EndVendorID: 77},
			{StartVendorID: 80, EndVendorID: 80},
			{StartVendorID: 82, EndVendorID: 82},
			{StartVendorID: 91, EndVendorID: 91},
			{StartVendorID: 97, EndVendorID: 97},
			{StartVendorID: 104, EndVendorID: 104},
			{StartVendorID: 131, EndVendorID: 132},
			{StartVendorID: 142, EndVendorID: 142},
			{StartVendorID: 165, EndVendorID: 165},
			{StartVendorID: 173, EndVendorID: 173},
			{StartVendorID: 183, EndVendorID: 183},
			{StartVendorID: 202, EndVendorID: 202},
			{StartVendorID: 237, EndVendorID: 237},
			{StartVendorID: 241, EndVendorID: 241},
			{StartVendorID: 263, EndVendorID: 263},
			{StartVendorID: 293, EndVendorID: 293},
			{StartVendorID: 315, EndVendorID: 315},
			{StartVendorID: 333, EndVendorID: 333},
			{StartVendorID: 335, EndVendorID: 335},
			{StartVendorID: 468, EndVendorID: 468},
			{StartVendorID: 512, EndVendorID: 512},
			{StartVendorID: 755, EndVendorID: 755},
			{StartVendorID: 775, EndVendorID: 775},
			{StartVendorID: 793, EndVendorID: 793},
			{StartVendorID: 812, EndVendorID: 812},
			{StartVendorID: 845, EndVendorID: 845},
			{StartVendorID: 860, EndVendorID: 860},
			{StartVendorID: 887, EndVendorID: 887},
			{StartVendorID: 1002, EndVendorID: 1002},
			{StartVendorID: 1126, EndVendorID: 1126},
			{StartVendorID: 1136, EndVendorID: 1136},
		},
		MaxImpliedConsentVendorID:     845,
		IsImpliedConsentRangeEncoding: true,
		ImpliedConsentedVendors:       map[int]bool(nil),
		NumImpliedConsentEntries:      30,
		VendorImpliedConsent: []*iabconsent.RangeEntry{{StartVendorID: 11, EndVendorID: 11},
			{StartVendorID: 14, EndVendorID: 14},
			{StartVendorID: 21, EndVendorID: 21},
			{StartVendorID: 23, EndVendorID: 23},
			{StartVendorID: 25, EndVendorID: 25},
			{StartVendorID: 28, EndVendorID: 28},
			{StartVendorID: 32, EndVendorID: 32},
			{StartVendorID: 36, EndVendorID: 36},
			{StartVendorID: 52, EndVendorID: 52},
			{StartVendorID: 62, EndVendorID: 62},
			{StartVendorID: 69, EndVendorID: 69},
			{StartVendorID: 76, EndVendorID: 76},
			{StartVendorID: 82, EndVendorID: 82},
			{StartVendorID: 104, EndVendorID: 104},
			{StartVendorID: 126, EndVendorID: 126},
			{StartVendorID: 132, EndVendorID: 132},
			{StartVendorID: 142, EndVendorID: 142},
			{StartVendorID: 165, EndVendorID: 165},
			{StartVendorID: 183, EndVendorID: 183},
			{StartVendorID: 202, EndVendorID: 202},
			{StartVendorID: 263, EndVendorID: 263},
			{StartVendorID: 278, EndVendorID: 278},
			{StartVendorID: 335, EndVendorID: 335},
			{StartVendorID: 431, EndVendorID: 431},
			{StartVendorID: 468, EndVendorID: 468},
			{StartVendorID: 706, EndVendorID: 706},
			{StartVendorID: 755, EndVendorID: 755},
			{StartVendorID: 812, EndVendorID: 812},
			{StartVendorID: 815, EndVendorID: 815},
			{StartVendorID: 845, EndVendorID: 845},
		},
		PublisherTCEntry: (*iabconsent.PublisherTCEntry)(nil),
	},
}
