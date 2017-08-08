package types

import ("testing"
"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestHash_MarshalJSON(t *testing.T) {
	var hash Hash = [32]byte{12,15,12}
	res, _ := json.Marshal(hash)
	expected, _ := json.Marshal(hash[:])
	assert.Equal(t, res, expected)

	notExpected, _ := json.Marshal([32]byte{12,15,12})
	assert.NotEqual(t, res, notExpected)
}
func CToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}
