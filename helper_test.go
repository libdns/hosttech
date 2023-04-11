package hosttech

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveTrailingDot(t *testing.T) {
	input := map[string]struct {
		expectedResult string
		data           string
	}{
		"Should remove dot": {
			expectedResult: "test.com",
			data:           "test.com.",
		},
		"Should change nothing": {
			expectedResult: "test.de",
			data:           "test.de",
		},
	}

	for name, testStruct := range input {
		t.Run(name, func(t *testing.T) {
			output := RemoveTrailingDot(testStruct.data)

			assert.Equal(t, testStruct.expectedResult, output)
		})
	}
}
