package iabconsent

import (
	"encoding/base64"
	"strings"

	"github.com/pkg/errors"
)

type MspaUsNational struct {
	GppSection
}

func NewMspaNationl(section string) *MspaUsNational {
	return &MspaUsNational{GppSection{sectionId: 7, sectionValue: section}}
}

func (m *MspaUsNational) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")
	// TODO(PX-2204): Re-usable subsections separated by '.', support for GPC will be added in.

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
	// TODO: Figure out if 0 is a valid value for MspaCoveredTransaction
	// If not, should we change to another value and continue?
	// Current documentation does not allow CoveredTx to be 0 value, but examples contradict this.
	//if p.MspaCoveredTransaction == MspaNotApplicable {
	//	// Value cannot be N/A, so just set to no to be conservative.
	//	p.MspaCoveredTransaction = MspaNo
	//}
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	// TODO(PX-2204): Parse remaining non-core reusable sections if they exist.
	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, _ = ParseGppSubSections(segments[1:])
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}
