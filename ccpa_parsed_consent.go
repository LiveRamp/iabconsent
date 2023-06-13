package iabconsent

import "github.com/pkg/errors"

const (
	CCPAYes           = 'Y'
	CCPANo            = 'N'
	CCPANotApplicable = '-'
	CCPAVersion       = 1
)

// CcpaParsedConsent represents data extract from a California Consumer Privacy Act (CCPA) consent string.
// Format can be found here: https://github.com/InteractiveAdvertisingBureau/USPrivacy/blob/master/CCPA/US%20Privacy%20String.md
type CcpaParsedConsent struct {
	// The version of this string specification used to encode the string
	Version int

	// N = No, Y = Yes, - = Not Applicable
	Notice uint8

	// N = No, Y = Yes, - = Not Applicable; For use ONLY when CCPA does not apply.
	OptOutSale uint8

	// 0 = No, 1 = Yes, - = Not Applicable
	LSPACoveredTransaction uint8
}

func IsValidCCPAString(ccpaString string) (bool, error) {
	if len(ccpaString) != 4 {
		return false, errors.Wrap(nil, "invalid uspv consent string length")
	}

	if ccpaString[0]-'0' != CCPAVersion {
		return false, errors.Wrap(nil, "invalid uspv consent string version")
	}

	for i := 1; i < 4; i++ {
		if ccpaString[i] != CCPAYes && ccpaString[i] != CCPANo && ccpaString[i] != CCPANotApplicable {
			return false, errors.Wrap(nil, "invalid uspv consent string")
		}
	}

	return true, nil
}

func ParseCCPA(s string) (*CcpaParsedConsent, error) {
	if valid, err := IsValidCCPAString(s); !valid {
		return nil, err
	}

	return &CcpaParsedConsent{
		int(s[0] - '0'),
		s[1],
		s[2],
		s[3],
	}, nil
}
