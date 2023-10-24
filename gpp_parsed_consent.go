package iabconsent

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	EuropeTCFv2SID = 2
	CanadaTCFSID   = iota + 4
	UsPVSID
	UsNationalSID
	UsCaliforniaSID
	UsVirginiaSID
	UsColoradoSID
	UsUtahSID
	UsConnecticutSID
)

// GppHeader is the first section of a GPP Consent String.
// See ParseGppHeader for in-depth format.
type GppHeader struct {
	Type     int
	Version  int
	Sections []int
}

// GppParsedConsent is an empty interface since GPP will need to handle more consent structs
// than just the Multi-state Privacy Agreement structs.
type GppParsedConsent interface {
}

// GppSection contains the specific Section ID (important to match up correct parsing).
// and pre-parsed Section Value, including all subsections.
type GppSection struct {
	sectionId    int
	sectionValue string
}

type GppSectionParser interface {
	ParseConsent() (GppParsedConsent, error)
	GetSectionId() int
	GetSectionValue() string
}

// GetSectionId returns the Section ID for a given GppSection.
func (g *GppSection) GetSectionId() int {
	return g.sectionId
}

// GetSectionValue returns the Section Value for a given GppSection.
func (g *GppSection) GetSectionValue() string {
	return g.sectionValue
}

type GppSubSection struct {
	// Global Privacy Control (GPC) is signaled and set.
	Gpc bool
}

type GppSubSectionTypes int

const (
	SubSectCore GppSubSectionTypes = iota
	SubSectGpc
)

// ParseGppHeader parses the first (and required) part of any GPP Consent String.
// It is used to read the Type, Version, and which sections are contained in the following string(s).
// Format is:
// Type	    Int(6)	Fixed to 3 as “GPP Header field”
// Version	Int(6)	Version of the GPP spec (version 1, as of Jan. 2023)
// Sections	Range(Fibonacci)	List of Section IDs that are contained in the GPP string.
func ParseGppHeader(s string) (*GppHeader, error) {
	// IAB's base64 conversion means a 6 bit grouped value can be converted to 8 bit bytes.
	// Any leftover bits <8 would be skipped in normal base64 decoding.
	// Therefore, pad with 6 '0's w/ `A` to ensure that all bits are decoded into bytes.

	gap := 3 - len(s)%4
	if gap != 0 {
		s += strings.Repeat("A", gap)
	}

	var b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.Wrap(err, "parse gpp header consent string")
	}

	var r = NewConsentReader(b)

	var g = &GppHeader{}
	g.Type, _ = r.ReadInt(6)
	if g.Type != 3 {
		return nil, errors.New("wrong gpp header type " + fmt.Sprint(g.Type))
	}
	g.Version, _ = r.ReadInt(6)
	if g.Version != 1 {
		return nil, errors.New("unsupported gpp version " + fmt.Sprint(g.Version))
	}
	g.Sections, _ = r.ReadFibonacciRange()
	return g, r.Err
}

// MapGppSectionToParser takes a base64 Raw URL Encoded string which represents a GPP v1 string
// of the format {gpp header}~{section 1}[.{sub-section}][~{section n}]
// and returns each pair of section value and parsing function that should be used.
// The pairs are returned to allow more control over how parsing functions are applied.
func MapGppSectionToParser(s string) ([]GppSectionParser, error) {
	var gppHeader *GppHeader
	var err error
	// ~ separated fields. with the format {gpp header}~{section 1}[.{sub-section}][~{section n}]
	var segments = strings.Split(s, "~")
	if len(segments) < 2 {
		return nil, errors.New("not enough gpp segments")
	}

	gppHeader, err = ParseGppHeader(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "read gpp header")
	} else if len(segments[1:]) != len(gppHeader.Sections) {
		// Return early if sections in header do not match sections passed.
		return nil, errors.New("mismatch number of sections")
	}
	// Go through each section and add parsing function and section value to returned value.
	var gppSections = make([]GppSectionParser, 0)
	for i := 1; i < len(segments); i++ {
		var gppSection GppSectionParser
		switch sid := gppHeader.Sections[i-1]; sid {
		case EuropeTCFv2SID:
			gppSection = NewTCFEU(segments[i])
		case CanadaTCFSID:
			gppSection = NewTCFCA(segments[i])
		case UsPVSID:
			gppSection = NewUSPV(segments[i])
		case UsNationalSID, UsCaliforniaSID, UsVirginiaSID, UsColoradoSID, UsUtahSID, UsConnecticutSID:
			gppSection = NewMspa(sid, segments[i])
		default:
			gppSection = NewNotSupported(segments[i], sid)
		}

		gppSections = append(gppSections, gppSection)
	}
	return gppSections, nil
}

// ParseGppConsent takes a base64 Raw URL Encoded string which represents a GPP v1 string and
// returns a map of Section ID to ParsedConsents with consent parsed via a consecutive parsing.
func ParseGppConsent(s string) (map[int]GppParsedConsent, error) {
	var gppSections []GppSectionParser
	var err error
	gppSections, err = MapGppSectionToParser(s)
	if err != nil {
		return nil, err
	}
	var gppConsents = make(map[int]GppParsedConsent, len(gppSections))
	// Consecutively, go through each section and try to parse.
	for _, gpp := range gppSections {
		var consent GppParsedConsent
		var consentErr error
		consent, consentErr = gpp.ParseConsent()
		if consentErr != nil {
			return nil, consentErr
		} else {
			gppConsents[gpp.GetSectionId()] = consent
		}
	}
	return gppConsents, nil
}

// ParseGppSubSections parses the subsections that may be appended to GPP sections after a `.`
// Currently, GPC is the only subsection, so we only have a single Subsection parsing function.
// In the future, Section IDs may need their own SubSection parser.
func ParseGppSubSections(subSections []string) (*GppSubSection, error) {
	var gppSub = new(GppSubSection)
	// There could be >1 subsection, but we will only return a single GppSubSection result.
	for _, s := range subSections {
		// Actual base64 encoded data, so no need to add extra `0`s.
		var b, err = base64.RawURLEncoding.DecodeString(s)
		if err != nil {
			return nil, errors.Wrap(err, "parse gpp subsection string")
		}
		var r = NewConsentReader(b)

		var subType int
		subType, err = r.ReadInt(2)
		if err != nil {
			return nil, errors.Wrap(err, "parse gpp subsection type")
		}
		// Check for specific SubSection Type, and then parse subsection correctly.
		switch GppSubSectionTypes(subType) {
		case SubSectGpc:
			var gppValue bool
			gppValue, err = ParseGpcSubsection(r)
			if err != nil {
				return nil, errors.Wrap(err, "parse gpp subsection gpc bool")
			}
			// Only override if not set to true already, as we want the most restrictive value
			// if > 1 GPC subsection.
			if gppSub.Gpc != true {
				gppSub.Gpc = gppValue
			}
		}
	}
	return gppSub, nil
}

// ParseGpcSubsection reads the next bit as a bool, and returns the result of value.
// Info about GPC subsection here: https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/606b99efc16b649c5c1f8f1d2eb0d0d3258c4a2d/Sections/US-National/IAB%20Privacy%E2%80%99s%20National%20Privacy%20Technical%20Specification.md#gpc-sub-section
func ParseGpcSubsection(r *ConsentReader) (bool, error) {
	var gppValue bool
	var err error
	gppValue, err = r.ReadBool()
	if err != nil {
		return false, errors.Wrap(err, "parse gpp subsection gpc bool")
	}
	return gppValue, err
}
