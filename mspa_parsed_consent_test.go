package iabconsent_test

import (
	"github.com/go-check/check"
	"github.com/pkg/errors"

	"github.com/LiveRamp/iabconsent"
)

type MspaSuite struct{}

var _ = check.Suite(&MspaSuite{})

func (s *MspaSuite) TestReadMspaNotice(c *check.C) {
	var r = iabconsent.NewConsentReader([]byte{0b00011011})
	var mns = []iabconsent.MspaNotice{
		iabconsent.NoticeNotApplicable,
		iabconsent.NoticeProvided,
		iabconsent.NoticeNotProvided,
		iabconsent.InvalidNoticeValue,
	}
	for _, i := range mns {
		var mn, err = r.ReadMspaNotice()
		c.Check(err, check.IsNil)
		c.Check(mn, check.Equals, i)
	}
}

func (s *MspaSuite) TestReadMspaOptOut(c *check.C) {
	var r = iabconsent.NewConsentReader([]byte{0b00011011})
	var mos = []iabconsent.MspaOptout{
		iabconsent.OptOutNotApplicable,
		iabconsent.OptedOut,
		iabconsent.NotOptedOut,
		iabconsent.InvalidOptOutValue,
	}
	for _, i := range mos {
		var mo, err = r.ReadMspaOptOut()
		c.Check(err, check.IsNil)
		c.Check(mo, check.Equals, i)
	}
}

func (s *MspaSuite) TestReadMspaConsent(c *check.C) {
	var r = iabconsent.NewConsentReader([]byte{0b00011011})
	var mcs = []iabconsent.MspaConsent{
		iabconsent.ConsentNotApplicable,
		iabconsent.NoConsent,
		iabconsent.Consent,
		iabconsent.InvalidConsentValue,
	}
	for _, i := range mcs {
		var mc, err = r.ReadMspaConsent()
		c.Check(err, check.IsNil)
		c.Check(mc, check.Equals, i)
	}
}

func (s *MspaSuite) TestReadMspaBitfieldConsent(c *check.C) {
	var tcs = []struct {
		testBytes      []byte
		bitfieldLength uint
		expected       map[int]iabconsent.MspaConsent
	}{
		{testBytes: []byte{0b00000000},
			bitfieldLength: 1,
			expected:       map[int]iabconsent.MspaConsent{0: iabconsent.ConsentNotApplicable}},
		{testBytes: []byte{0b00000000},
			bitfieldLength: 4,
			expected: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.ConsentNotApplicable,
				2: iabconsent.ConsentNotApplicable,
				3: iabconsent.ConsentNotApplicable}},
		{testBytes: []byte{0b00011011},
			bitfieldLength: 4,
			expected: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.NoConsent,
				2: iabconsent.Consent,
				3: iabconsent.InvalidConsentValue}},
	}

	for _, t := range tcs {
		var r = iabconsent.NewConsentReader(t.testBytes)
		var bc, err = r.ReadMspaBitfieldConsent(t.bitfieldLength)
		c.Check(err, check.IsNil)
		c.Check(bc, check.DeepEquals, t.expected)
	}
}

func (s *MspaSuite) TestReadMspaNaYesNo(c *check.C) {
	var r = iabconsent.NewConsentReader([]byte{0b00011011})
	var mcs = []iabconsent.MspaNaYesNo{
		iabconsent.MspaNotApplicable,
		iabconsent.MspaYes,
		iabconsent.MspaNo,
		iabconsent.InvalidMspaValue,
	}
	for _, i := range mcs {
		var mc, err = r.ReadMspaNaYesNo()
		c.Check(err, check.IsNil)
		c.Check(mc, check.Equals, i)
	}
}

func (s *MspaSuite) TestParseUsNational(c *check.C) {
	for k, v := range usNationalConsentFixtures {
		c.Log(k)

		var gppSection = iabconsent.NewMspa(7, k)
		var p, err = gppSection.ParseConsent()

		c.Check(err, check.IsNil)
		c.Check(p, check.DeepEquals, v)
	}
}

func (s *MspaSuite) TestParseUsNationalError(c *check.C) {
	var tcs = []struct {
		desc        string
		usNatString string
		expected    error
	}{
		{
			desc:        "Wrong Version.",
			usNatString: "DVVqAAEABA",
			expected:    errors.New("non-v1 string passed."),
		},
		{
			desc:        "Bad Decoding.",
			usNatString: "$%&*(",
			expected:    errors.New("parse usnat consent string: illegal base64 data at input byte 0"),
		},
	}
	for _, t := range tcs {
		c.Log(t.desc)

		var gppSection = iabconsent.NewMspa(7, t.usNatString)
		var p, err = gppSection.ParseConsent()

		c.Check(p, check.IsNil)
		c.Check(err, check.ErrorMatches, t.expected.Error())
	}
}

func (s *MspaSuite) TestParseUsCA(c *check.C) {
	for k, v := range usCAConsentFixtures {
		c.Log(k)

		var gppSection = iabconsent.NewMspa(8, k)
		var p, err = gppSection.ParseConsent()

		c.Check(err, check.IsNil)
		c.Check(p, check.DeepEquals, v)
	}
}

func (s *MspaSuite) TestParseUsCAError(c *check.C) {
	var tcs = []struct {
		desc       string
		usCAString string
		expected   error
	}{
		{
			desc:       "Wrong Version.",
			usCAString: "CVoYYZoA",
			expected:   errors.New("non-v1 string passed."),
		},
		{
			desc:       "Bad Decoding.",
			usCAString: "$%&*(",
			expected:   errors.New("parse usca consent string: illegal base64 data at input byte 0"),
		},
	}
	for _, t := range tcs {
		c.Log(t.desc)

		var gppSection = iabconsent.NewMspa(8, t.usCAString)
		var p, err = gppSection.ParseConsent()

		c.Check(p, check.IsNil)
		c.Check(err, check.ErrorMatches, t.expected.Error())
	}
}

func (s *MspaSuite) TestParseUsVA(c *check.C) {
	for k, v := range usVAConsentFixtures {
		c.Log(k)

		var gppSection = iabconsent.NewMspa(9, k)
		var p, err = gppSection.ParseConsent()

		c.Check(err, check.IsNil)
		c.Check(p, check.DeepEquals, v)
	}
}

func (s *MspaSuite) TestParseUsVAError(c *check.C) {
	var tcs = []struct {
		desc       string
		usVAString string
		expected   error
	}{
		{
			desc:       "Wrong Version.",
			usVAString: "DVoYYYA",
			expected:   errors.New("non-v1 string passed."),
		},
		{
			desc:       "Bad Decoding.",
			usVAString: "$%&*(",
			expected:   errors.New("parse usva consent string: illegal base64 data at input byte 0"),
		},
	}
	for _, t := range tcs {
		c.Log(t.desc)

		var gppSection = iabconsent.NewMspa(9, t.usVAString)
		var p, err = gppSection.ParseConsent()

		c.Check(p, check.IsNil)
		c.Check(err, check.ErrorMatches, t.expected.Error())
	}
}

func (s *MspaSuite) TestParseUsCO(c *check.C) {
	for k, v := range usCOConsentFixtures {
		c.Log(k)

		var gppSection = iabconsent.NewMspa(10, k)
		var p, err = gppSection.ParseConsent()

		c.Check(err, check.IsNil)
		c.Check(p, check.DeepEquals, v)
	}
}

func (s *MspaSuite) TestParseUsCOError(c *check.C) {
	var tcs = []struct {
		desc       string
		usCOString string
		expected   error
	}{
		{
			desc:       "Wrong Version.",
			usCOString: "DVVqAAEABA",
			expected:   errors.New("non-v1 string passed."),
		},
		{
			desc:       "Bad Decoding.",
			usCOString: "$%&*(",
			expected:   errors.New("parse usco consent string: illegal base64 data at input byte 0"),
		},
	}
	for _, t := range tcs {
		c.Log(t.desc)

		var gppSection = iabconsent.NewMspa(10, t.usCOString)
		var p, err = gppSection.ParseConsent()

		c.Check(p, check.IsNil)
		c.Check(err, check.ErrorMatches, t.expected.Error())
	}
}
