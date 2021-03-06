package horizon

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/fonero-project/fonero-horizon/test"
)

func TestDataActions_Show(t *testing.T) {
	ht := StartHTTPTest(t, "kahuna")
	defer ht.Finish()

	prefix := "/accounts/GCC2R6IXEKBAYK3CYANYRMX5MAZXN57G2ZTKDALAZEHGCUMUDRCYVYRO"
	// json

	w := ht.Get(prefix + "/data/name1")
	if ht.Assert.Equal(200, w.Code) {
		var result map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &result)
		ht.Require.NoError(err)
		decoded, err := base64.StdEncoding.DecodeString(result["value"])
		ht.Require.NoError(err)

		ht.Assert.Equal("0000", string(decoded))
	}

	// raw
	w = ht.Get(prefix+"/data/name1", test.RequestHelperRaw)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.Equal("0000", w.Body.String())
	}

	// missing
	w = ht.Get(prefix+"/data/missing", test.RequestHelperRaw)
	ht.Assert.Equal(404, w.Code)

	// regression: https://github.com/fonero-project/fonero-horizon/issues/325
	// names with special characters do not work
	w = ht.Get(prefix+"/data/name%20", test.RequestHelperRaw)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.Equal("its got spaces!", w.Body.String())
	}
}
