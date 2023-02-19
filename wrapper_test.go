package hosttech

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHosttechRecordWrapper_UnmarshalJSON(t *testing.T) {
	input := map[string]struct {
		expectedResult HosttechRecordWrapper
		data           []byte
	}{
		"ARecord Test": {
			expectedResult: HosttechRecordWrapper{
				value: ARecord{
					Base: Base{
						Id:      10,
						Type:    "A",
						TTL:     3600,
						Comment: "my first record",
					},
					Name: "www",
					IPV4: "1.2.3.4",
				},
			},
			data: []byte(`{
      "id": 10,
      "type": "A",
      "name": "www",
      "ipv4": "1.2.3.4",
      "ttl": 3600,
      "comment": "my first record"
    }`),
		},
		"AAAARecord Test": {
			expectedResult: HosttechRecordWrapper{
				value: AAAARecord{
					Base: Base{
						Id:      11,
						Type:    "AAAA",
						TTL:     3600,
						Comment: "my first record",
					},
					Name: "www",
					IPV6: "2001:db8:1234::1",
				},
			},
			data: []byte(`{
      "id": 11,
      "type": "AAAA",
      "name": "www",
      "ipv6": "2001:db8:1234::1",
      "ttl": 3600,
      "comment": "my first record"
    }`),
		},
		"NSRecord Test": {
			expectedResult: HosttechRecordWrapper{
				value: NSRecord{
					Base: Base{
						Id:      14,
						Type:    "NS",
						TTL:     3600,
						Comment: "my first record",
					},
					OwnerName:  "sub",
					TargetName: "ns1.example.com",
				},
			},
			data: []byte(`{ "id": 14, "type": "NS", "ownername": "sub", "targetname": "ns1.example.com", "ttl": 3600, "comment": "my first record" }`),
		},
		"CNAMERecord Test": {
			expectedResult: HosttechRecordWrapper{
				value: CNAMERecord{
					Base: Base{
						Id:      13,
						Type:    "CNAME",
						TTL:     3600,
						Comment: "my first record",
					},
					Name:  "www",
					Cname: "site.example.com",
				}},
			data: []byte(`{ "id": 13, "type": "CNAME", "name": "www", "cname": "site.example.com", "ttl": 3600, "comment": "my first record" }`),
		},
		"MXRecord Test": {
			expectedResult: HosttechRecordWrapper{
				value: MXRecord{
					Base: Base{
						Id:      14,
						Type:    "MX",
						TTL:     3600,
						Comment: "my first record",
					},
					Name:      "mail.example.com",
					OwnerName: "owner name",
					Pref:      10,
				},
			},
			data: []byte(`{ "id": 14, "type": "MX", "ownername": "owner name", "name": "mail.example.com", "pref": 10, "ttl": 3600, "comment": "my first record" }`),
		},
		"TXTRecord Test": {
			expectedResult: HosttechRecordWrapper{
				value: TXTRecord{
					Base: Base{
						Id:      17,
						Type:    "TXT",
						TTL:     3600,
						Comment: "my first record",
					},
					Name: "txt name",
					Text: "v=spf1 ip4:1.2.3.4/32 -all",
				},
			},
			data: []byte(`{ "id": 17, "type": "TXT", "name": "txt name", "text": "v=spf1 ip4:1.2.3.4/32 -all", "ttl": 3600, "comment": "my first record" }`),
		},
		"TLSARecord Test": {
			expectedResult: HosttechRecordWrapper{
				value: TLSARecord{
					Base: Base{
						Id:      17,
						Type:    "TLSA",
						TTL:     3600,
						Comment: "my first record",
					},
					Name: "tlsa name",
					Text: "0 0 1 d2abde240d7cd3ee6b4b28c54df034b97983a1d16e8a410e4561cb106618e971",
				},
			},
			data: []byte(`{ "id": 17, "type": "TLSA", "name": "tlsa name", "text": "0 0 1 d2abde240d7cd3ee6b4b28c54df034b97983a1d16e8a410e4561cb106618e971", "ttl": 3600, "comment": "my first record" }`),
		},
	}

	for name, testStruct := range input {
		t.Run(name, func(t *testing.T) {
			output := HosttechRecordWrapper{}
			err := output.UnmarshalJSON(testStruct.data)

			if err != nil {
				t.Fail()
			}

			assert.Equal(t, testStruct.expectedResult, output)
		})
	}
}
