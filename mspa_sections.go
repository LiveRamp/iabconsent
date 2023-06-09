package iabconsent

import (
	"encoding/base64"
	"strings"

	"github.com/pkg/errors"
)

type MspaUsNational struct {
	GppSection
}

type MspaUsCA struct {
	GppSection
}

type MspaUsVA struct {
	GppSection
}

type MspaUsCO struct {
	GppSection
}

type MspaUsUT struct {
	GppSection
}

type MspaUsCT struct {
	GppSection
}

// NewMspa returns a supported parser given a GPP Section ID.
// If the SID is not yet supported, it will be null.
func NewMspa(sid int, section string) GppSectionParser {
	switch sid {
	case UsNationalSID:
		return &MspaUsNational{GppSection{sectionId: UsNationalSID, sectionValue: section}}
	case UsCaliforniaSID:
		return &MspaUsCA{GppSection{sectionId: UsCaliforniaSID, sectionValue: section}}
	case UsVirginiaSID:
		return &MspaUsVA{GppSection{sectionId: UsVirginiaSID, sectionValue: section}}
	case UsColoradoSID:
		return &MspaUsCO{GppSection{sectionId: UsColoradoSID, sectionValue: section}}
	case UsUtahSID:
		return &MspaUsUT{GppSection{sectionId: UsUtahSID, sectionValue: section}}
	case UsConnecticutSID:
		return &MspaUsCT{GppSection{sectionId: UsConnecticutSID, sectionValue: section}}
	}
	// Skip if no matching struct, as Section ID is not supported yet.
	// Any newly supported Section IDs should be added as cases here.
	return nil
}

func (m *MspaUsNational) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usnat consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-National#core-segment
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("non-v1 string passed.")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.SharingOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataProcessingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataLimitUseNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.SharingOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(12)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(2)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, _ = ParseGppSubSections(segments[1:])
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsCA) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usca consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/CA
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("non-v1 string passed.")
	}

	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.SharingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataLimitUseNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.SharingOptOut, _ = r.ReadMspaOptOut()
	// SensitiveDataProcessingOptOuts, as opposed to Consent.
	p.SensitiveDataProcessingOptOuts, _ = r.ReadMspaBitfieldOptOut(9)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(2)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, _ = ParseGppSubSections(segments[1:])
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsVA) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usva consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/VA
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("non-v1 string passed.")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	// This has a shorter length than UsNational.
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	// While an array in UsNational, we can just use an array of 1 for a single value.
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, _ = ParseGppSubSections(segments[1:])
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsCO) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usco consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/CO
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("non-v1 string passed.")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(7)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, _ = ParseGppSubSections(segments[1:])
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsUT) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usut consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/UT
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("non-v1 string passed.")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataProcessingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingOptOuts, _ = r.ReadMspaBitfieldOptOut(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, _ = ParseGppSubSections(segments[1:])
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsCT) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usct consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/CT
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("non-v1 string passed.")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, _ = ParseGppSubSections(segments[1:])
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}
