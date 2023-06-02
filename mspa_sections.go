package iabconsent

import (
	"encoding/base64"
	"strings"

	"github.com/pkg/errors"
)

type MspaUsNational struct {
	GppSection
}

type MspaUsVA struct {
	GppSection
}

func NewMspaNational(section string) *MspaUsNational {
	return &MspaUsNational{GppSection{sectionId: 7, sectionValue: section}}
}

func NewMspaVA(section string) *MspaUsVA {
	return &MspaUsVA{GppSection{sectionId: 9, sectionValue: section}}
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
