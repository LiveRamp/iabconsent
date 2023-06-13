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

type MspaUsVA struct {
	GppSection
}

type MspaUsCA struct {
	GppSection
}

type MspaUsCO struct {
	GppSection
}

type MspaUsCT struct {
	GppSection
}

type MspaUsUT struct {
	GppSection
}

type TCFEU struct {
	GppSection
}

type TCFCA struct {
	GppSection
}

type USPV struct {
	GppSection
}

type NotSupported struct {
	GppSection
}

func NewTCFEU(section string) *TCFEU {
	return &TCFEU{GppSection{sectionId: SectionIDEUTCFv2, sectionValue: section}}
}

func NewTCFCA(section string) *TCFCA {
	return &TCFCA{GppSection{sectionId: SectionIDCANTCF, sectionValue: section}}
}

func NewUSPV(section string) *USPV {
	return &USPV{GppSection{sectionId: SectionIDUSPV1, sectionValue: section}}
}

func NewMspaNational(section string) *MspaUsNational {
	return &MspaUsNational{GppSection{sectionId: SectionIDUSNAT, sectionValue: section}}
}

func NewMspaCA(section string) *MspaUsCA {
	return &MspaUsCA{GppSection{sectionId: SectionIDUSCA, sectionValue: section}}
}

func NewMspaVA(section string) *MspaUsVA {
	return &MspaUsVA{GppSection{sectionId: SectionIDUSVA, sectionValue: section}}
}

func NewMspaCO(section string) *MspaUsCO {
	return &MspaUsCO{GppSection{sectionId: SectionIDUSCO, sectionValue: section}}
}

func NewMspaUT(section string) *MspaUsUT {
	return &MspaUsUT{GppSection{sectionId: SectionIDUSUT, sectionValue: section}}
}

func NewMspaCT(section string) *MspaUsCT {
	return &MspaUsCT{GppSection{sectionId: SectionIDUSCT, sectionValue: section}}
}

func NewNotSupported(section string, sectionID int) *NotSupported {
	return &NotSupported{GppSection{sectionId: sectionID, sectionValue: section}}
}

func (n *NotSupported) ParseConsent() (GppParsedConsent, error) {
	return nil, errors.New(fmt.Sprintf("Section ID %d is not supported", n.sectionId))
}

func (t *TCFEU) ParseConsent() (GppParsedConsent, error) {
	return ParseV2(t.sectionValue)
}

func (t *TCFCA) ParseConsent() (GppParsedConsent, error) {
	return ParseV2(t.sectionValue)
}

func (u *USPV) ParseConsent() (GppParsedConsent, error) {
	return ParseCCPA(u.sectionValue)

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
	p.SensitiveDataProcessing, _ = r.ReadMspaBitfieldConsent(12)
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
	p.SensitiveDataProcessing, _ = r.ReadMspaBitfieldConsent(8)
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
		return nil, errors.Wrap(err, "parse usva consent string")
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
	p.SensitiveDataProcessing, _ = r.ReadMspaBitfieldConsent(7)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
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

func (m *MspaUsUT) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usva consent string")
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
	p.SensitiveDataProcessing, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
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

func (m *MspaUsCT) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usva consent string")
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
	p.SensitiveDataProcessing, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
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

func (m *MspaUsCA) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usva consent string")
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

	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.SharingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataLimitUseNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.SharingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessing, _ = r.ReadMspaBitfieldConsent(9)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(2)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
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
