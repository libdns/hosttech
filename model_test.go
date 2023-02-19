package hosttech

import (
	"github.com/libdns/libdns"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHosttechRecordToLibdnsRecord(t *testing.T) {
	zone := "example.com"
	input := map[string]struct {
		expectedResult libdns.Record
		data           HosttechRecord
	}{
		"ARecord Test": {
			expectedResult: libdns.Record{
				ID:       "14",
				Type:     "A",
				Name:     "sub",
				Value:    "192.168.68.1",
				TTL:      1800 * time.Second,
				Priority: 0,
			},
			data: ARecord{
				Base: Base{
					Id:      14,
					Type:    "A",
					TTL:     1800,
					Comment: "Some comment",
				},
				Name: "sub.example.com",
				IPV4: "192.168.68.1",
			},
		},
		"AAAARecord Test": {
			expectedResult: libdns.Record{
				ID:       "23",
				Type:     "AAAA",
				Name:     "sub",
				Value:    "2607:f0d0:1002:51::4",
				TTL:      1800 * time.Second,
				Priority: 0,
			},
			data: AAAARecord{
				Base: Base{
					Id:      23,
					Type:    "AAAA",
					TTL:     1800,
					Comment: "Some comment",
				},
				Name: "sub.example.com",
				IPV6: "2607:f0d0:1002:51::4",
			},
		},
		"NSRecord Test": {
			expectedResult: libdns.Record{
				ID:       "12",
				Type:     "NS",
				Name:     "sub",
				Value:    "ns1.example.com",
				TTL:      1900 * time.Second,
				Priority: 0,
			},
			data: NSRecord{
				Base: Base{
					Id:      12,
					Type:    "NS",
					TTL:     1900,
					Comment: "Some comment",
				},
				OwnerName:  "sub.example.com",
				TargetName: "ns1.example.com",
			},
		},
		"CNAMERecord Test": {
			expectedResult: libdns.Record{
				ID:       "143",
				Type:     "CNAME",
				Name:     "sub",
				Value:    "site.example.com",
				TTL:      1700 * time.Second,
				Priority: 0,
			},
			data: CNAMERecord{
				Base: Base{
					Id:      143,
					Type:    "CNAME",
					TTL:     1700,
					Comment: "Some comment",
				},
				Name:  "sub.example.com",
				Cname: "site.example.com",
			},
		},
		"MXRecord Test": {
			expectedResult: libdns.Record{
				ID:       "748",
				Type:     "MX",
				Name:     "sub",
				Value:    "mail.server.com",
				TTL:      1750 * time.Second,
				Priority: 10,
			},
			data: MXRecord{
				Base: Base{
					Id:      748,
					Type:    "MX",
					TTL:     1750,
					Comment: "Some comment",
				},
				Name:      "mail.server.com",
				OwnerName: "sub.example.com",
				Pref:      10,
			},
		},
		"TXTRecord Test": {
			expectedResult: libdns.Record{
				ID:       "178",
				Type:     "TXT",
				Name:     "sub",
				Value:    "Some cool text",
				TTL:      1690 * time.Second,
				Priority: 0,
			},
			data: TXTRecord{
				Base: Base{
					Id:      178,
					Type:    "TXT",
					TTL:     1690,
					Comment: "Some comment",
				},
				Name: "sub.example.com",
				Text: "Some cool text",
			},
		},
		"TLSARecord Test": {
			expectedResult: libdns.Record{
				ID:       "61",
				Type:     "TLSA",
				Name:     "sub",
				Value:    "TLSA text",
				TTL:      1700 * time.Second,
				Priority: 0,
			},
			data: TLSARecord{
				Base: Base{
					Id:      61,
					Type:    "TLSA",
					TTL:     1700,
					Comment: "Some comment",
				},
				Name: "sub.example.com",
				Text: "TLSA text",
			},
		},
	}

	for name, testStruct := range input {
		t.Run(name, func(t *testing.T) {
			output := testStruct.data.toLibdnsRecord(zone)

			assert.Equal(t, testStruct.expectedResult, output)
		})
	}
}
