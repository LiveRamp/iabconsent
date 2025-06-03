package iabconsent

import (
	"encoding/base64"
	"fmt"
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

type MspaUsFL struct {
	GppSection
}

type MspaUsMT struct {
	GppSection
}

type MspaUsOR struct {
	GppSection
}

type MspaUsTX struct {
	GppSection
}

type MspaUsDE struct {
	GppSection
}

type MspaUsIA struct {
	GppSection
}

type MspaUsNE struct {
	GppSection
}

type MspaUsNH struct {
	GppSection
}

type MspaUsNJ struct {
	GppSection
}

type MspaUsTN struct {
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
	case UsFloridaSID:
		return &MspaUsFL{GppSection{sectionId: UsFloridaSID, sectionValue: section}}
	case UsMontanaSID:
		return &MspaUsMT{GppSection{sectionId: UsMontanaSID, sectionValue: section}}
	case UsOregonSID:
		return &MspaUsOR{GppSection{sectionId: UsOregonSID, sectionValue: section}}
	case UsTexasSID:
		return &MspaUsTX{GppSection{sectionId: UsTexasSID, sectionValue: section}}
	case UsDelawareSID:
		return &MspaUsDE{GppSection{sectionId: UsDelawareSID, sectionValue: section}}
	case UsIowaSID:
		return &MspaUsIA{GppSection{sectionId: UsIowaSID, sectionValue: section}}
	case UsNebraskaSID:
		return &MspaUsNE{GppSection{sectionId: UsNebraskaSID, sectionValue: section}}
	case UsNewHampshireSID:
		return &MspaUsNH{GppSection{sectionId: UsNewHampshireSID, sectionValue: section}}
	case UsNewJerseySID:
		return &MspaUsNJ{GppSection{sectionId: UsNewJerseySID, sectionValue: section}}
	case UsTennesseeSID:
		return &MspaUsTN{GppSection{sectionId: UsTennesseeSID, sectionValue: section}}
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
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/main/Sections/US-National/IAB%20Privacy%E2%80%99s%20Multi-State%20Privacy%20Agreement%20(MSPA)%20US%20National%20Technical%20Specification.md
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	// Support both v1 and v2
	if p.Version != 1 && p.Version != 2 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1 and v2
	// length of v1 is 60 bits padded with 4 bits of 0s making valid length 64 bits
	// length of v2 is 70 bits padded with 2 bits of 0s making valid length 72 bits
	if p.Version == 1 && r.Size() != 64 {
		return nil, errors.New("invalid consent string length for v1")
	} else if p.Version == 2 && r.Size() != 72 {
		return nil, errors.New("invalid consent string length for v2")
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

	// see spec in IAB GPP repo for differences between v1 and v2
	if p.Version == 1 {
		p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(12)
		p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(2)
	} else if p.Version == 2 {
		p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(16)
		p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	}
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
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
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 46 bits padded with 2 bits of 0s making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
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
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
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
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 40 bits no padding needed
	if p.Version == 1 && r.Size() != 40 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
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
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 38 bits with 2 bits of padding making valid length 40 bits
	if p.Version == 1 && r.Size() != 40 {
		return nil, errors.New("invalid consent string length for v1")
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
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
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
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 42 bits with 6 bits of padding making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
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
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
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
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 44 bits with 4 bits of padding making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
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
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsFL) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usfl consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/FL
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 46 bits with 2 bits of padding making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsMT) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usmt consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/MT
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 46 bits with 2 bits of padding making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsOR) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usor consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/OR
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 52 bits with 4 bits of padding making valid length 56 bits
	if p.Version == 1 && r.Size() != 56 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(11)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsTX) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse ustx consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/TX
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 42 bits with 6 bits of padding making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsDE) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usde consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/DE
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 52 bits with 4 bits of padding making valid length 56 bits
	if p.Version == 1 && r.Size() != 56 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(9)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(5)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

// Fix Iowa implementation
func (m *MspaUsIA) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usia consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/IA
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 42 bits with 6 bits of padding making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
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
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsNE) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usne consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/NE
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 42 bits with 6 bits of padding making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsNH) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usnh consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/NH
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 46 bits with 2 bits of padding making valid length 48 bits
	if p.Version == 1 && r.Size() != 48 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsNJ) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usnj consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/NJ
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 54 bits with 2 bits of padding making valid length 56 bits
	if p.Version == 1 && r.Size() != 56 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(10)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(5)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsTN) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse ustn consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/TN
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	// length of v1 is 40 bits with no padding making valid length 40 bits
	if p.Version == 1 && r.Size() != 40 {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}
