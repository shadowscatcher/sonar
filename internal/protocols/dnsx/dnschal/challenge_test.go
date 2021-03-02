package dnschal_test

import (
	"net"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bi-zone/sonar/internal/protocols/dnsx"
	"github.com/bi-zone/sonar/internal/protocols/dnsx/dnschal"
	"github.com/bi-zone/sonar/internal/protocols/dnsx/dnsrec"
	"github.com/bi-zone/sonar/internal/testutils"
)

var (
	rec *dnsrec.Records
	srv *dnsx.Server

	g = testutils.Globals(
		testutils.DNSDefaultRecords(&rec),
		testutils.DNSX([](func() dnsx.Handler){
			func(r **dnsrec.Records) func() dnsx.Handler {
				return func() dnsx.Handler {
					return *r
				}
			}(&rec),
		}, func(net.Addr, []byte, map[string]interface{}) {}, &srv),
	)
)

func TestMain(m *testing.M) {
	testutils.TestMain(m, g)
}

func TestDNSChallenge(t *testing.T) {
	provider := &dnschal.Provider{Records: rec}

	for _, name := range []string{
		"_acme-challenge.sonar.local.",
		"_aCme-chAlLEnge.soNAr.lOcal.",
	} {

		err := provider.Present("sonar.local", "", "key1")
		require.NoError(t, err)

		err = provider.Present("sonar.local", "", "key2")
		require.NoError(t, err)

		msg := new(dns.Msg)
		msg.Id = dns.Id()
		msg.RecursionDesired = true
		msg.Question = make([]dns.Question, 1)
		msg.Question[0] = dns.Question{
			Name:   name,
			Qtype:  dns.TypeTXT,
			Qclass: dns.ClassINET,
		}

		c := &dns.Client{}
		in, _, err := c.Exchange(msg, "127.0.0.1:1053")
		require.NoError(t, err)
		require.Len(t, in.Answer, 2)

		for i, txt := range []string{
			"gXQJloeiZiH04s3XzAOz2s7bP7liJVsar9Azyr6DFTA",
			"sQJTdkyLIz-zdULiNAHHtFDlpvl1HztaAU9vZ-i8mZ0",
		} {
			a, ok := in.Answer[i].(*dns.TXT)
			require.True(t, ok)
			require.Len(t, a.Txt, 1)
			assert.Equal(t, txt, a.Txt[0])
		}

		err = provider.CleanUp("sonar.local", "", "")
		require.NoError(t, err)

		in, _, err = c.Exchange(msg, "127.0.0.1:1053")
		require.NoError(t, err)
		require.NotNil(t, in)
		require.Len(t, in.Answer, 0)
	}
}