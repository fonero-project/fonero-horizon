package horizon

import (
	"testing"

	"github.com/fonero-project/fonero-horizon/test"
)

func TestEffectActions_Index(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	w := ht.Get("/effects?limit=20")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(11, w.Body)
	}

	// test streaming, regression for https://github.com/fonero-project/fonero-horizon/issues/147
	w = ht.Get("/effects?limit=2", test.RequestHelperStreaming)
	ht.Assert.Equal(200, w.Code)

	// filtered by ledger
	w = ht.Get("/ledgers/1/effects")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(0, w.Body)
	}

	w = ht.Get("/ledgers/2/effects")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(9, w.Body)
	}

	w = ht.Get("/ledgers/3/effects")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(2, w.Body)
	}

	// filtered by account
	w = ht.Get("/accounts/GCC2R6IXEKBAYK3CYANYRMX5MAZXN57G2ZTKDALAZEHGCUMUDRCYVYRO/effects")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(3, w.Body)
	}

	w = ht.Get("/accounts/GCC2R6IXEKBAYK3CYANYRMX5MAZXN57G2ZTKDALAZEHGCUMUDRCYVYRO/effects")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(2, w.Body)
	}

	w = ht.Get("/accounts/GCC2R6IXEKBAYK3CYANYRMX5MAZXN57G2ZTKDALAZEHGCUMUDRCYVYRO/effects")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(3, w.Body)
	}

	// filtered by transaction
	w = ht.Get("/transactions/GCC2R6IXEKBAYK3CYANYRMX5MAZXN57G2ZTKDALAZEHGCUMUDRCYVYRO/effects")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(3, w.Body)
	}

	// filtered by operation
	w = ht.Get("/operations/1/effects")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(3, w.Body)
	}

	// before history
	ht.ReapHistory(1)
	w = ht.Get("/effects?order=desc&cursor=1-1")
	ht.Assert.Equal(410, w.Code)
	ht.Logger.Error(w.Body.String())
}
