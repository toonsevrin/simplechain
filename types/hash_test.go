package types

import ("testing"
"encoding/json"
	"github.com/stretchr/testify/assert"
	"bytes"
	"fmt"
)

func TestHash_MarshalJSON(t *testing.T) {
	var hash Hash = [32]byte{12,15,12}
	res, _ := json.Marshal(hash)
	fmt.Println(string(res))
	expected, _ := json.Marshal(hash[:])
	assert.Equal(t, res, expected)

	notExpected, _ := json.Marshal([32]byte{12,15,12})
	assert.NotEqual(t, res, notExpected)
}
func TestHash_UnmarshalJSON(t *testing.T) {
	hashJson := `"DA8MAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="`
	actual := Hash{}
	json.Unmarshal([]byte(hashJson), &actual)
	expected := [32]byte{12,15,12}
	assert.True(t, bytes.Equal(actual[:], expected[:]))
}
