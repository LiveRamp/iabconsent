package iabconsent_test

import (
	"strings"

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

func (s *MspaSuite) TestParseMSPA(c *check.C) {
	for sid, sections := range mspaConsentFixtures {
		for section, result := range sections {

			c.Log(section)

			var gppSection = iabconsent.NewMspa(sid, section)
			var p, err = gppSection.ParseConsent()

			c.Check(err, check.IsNil)
			c.Check(p, check.DeepEquals, result)
		}
	}
}

func (s *MspaSuite) TestParseMSPAError(c *check.C) {
	var mspaTests = []struct {
		desc string
		key  string
		sid  int
	}{
		{
			desc: "US National",
			key:  "usnat",
			sid:  iabconsent.UsNationalSID,
		},
		{
			desc: "California",
			key:  "usca",
			sid:  iabconsent.UsCaliforniaSID,
		},
		{
			desc: "Virginia",
			key:  "usva",
			sid:  iabconsent.UsVirginiaSID,
		},
		{
			desc: "Colorado",
			key:  "usco",
			sid:  iabconsent.UsColoradoSID,
		},
		{
			desc: "Utan",
			key:  "usut",
			sid:  iabconsent.UsUtahSID,
		},
		{
			desc: "Connecticut",
			key:  "usct",
			sid:  iabconsent.UsConnecticutSID,
		},
		{
			desc: "Florida",
			key:  "usfl",
			sid:  iabconsent.UsFloridaSID,
		},
		{
			desc: "Montana",
			key:  "usmt",
			sid:  iabconsent.UsMontanaSID,
		},
		{
			desc: "Oregon",
			key:  "usor",
			sid:  iabconsent.UsOregonSID,
		},
		{
			desc: "Texas",
			key:  "ustx",
			sid:  iabconsent.UsTexasSID,
		},
	}
	var tcs = []struct {
		desc          string
		consentString string
		expected      string
	}{
		{
			desc:          "Wrong Version.",
			consentString: "DVVqAAEABA",
			expected:      "non-v1 string passed.",
		},
		{
			desc:          "Bad Decoding.",
			consentString: "$%&*(",
			expected:      "parse %s consent string: illegal base64 data at input byte 0",
		},
	}
	for _, s := range mspaTests {
		for _, t := range tcs {
			c.Log(s.desc + " - " + t.desc)

			var gppSection = iabconsent.NewMspa(s.sid, t.consentString)
			var p, err = gppSection.ParseConsent()

			c.Check(p, check.IsNil)
			var expected = strings.Replace(t.expected, "%s", s.key, 1)
			c.Check(err, check.ErrorMatches, errors.Errorf(expected).Error())
		}
	}
}
