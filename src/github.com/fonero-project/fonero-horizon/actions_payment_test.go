package horizon

import (
	"testing"
	"time"

	"github.com/fonero-project/horizon/db2/history"
	"github.com/fonero-project/horizon/resource/operations"
)

func TestPaymentActions(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	w := ht.Get("/payments")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(4, w.Body)
	}

	// filtered by ledger
	w = ht.Get("/ledgers/1/payments")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(0, w.Body)
	}

	w = ht.Get("/ledgers/3/payments")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
	}

	// filtered by account
	w = ht.Get("/accounts/GCC2R6IXEKBAYK3CYANYRMX5MAZXN57G2ZTKDALAZEHGCUMUDRCYVYRO/payments")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
	}

	// switch scenarios
	ht.T.Scenario("pathed_payment")

	// filtered by transaction
	w = ht.Get("/transactions/GCC2R6IXEKBAYK3CYANYRMX5MAZXN57G2ZTKDALAZEHGCUMUDRCYVYRO/payments")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(0, w.Body)
	}

	w = ht.Get("/transactions/GCC2R6IXEKBAYK3CYANYRMX5MAZXN57G2ZTKDALAZEHGCUMUDRCYVYRO/payments")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
	}
}

func TestPayment_CreatedAt(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	w := ht.Get("/ledgers/3/payments")
	records := []operations.Base{}
	ht.UnmarshalPage(w.Body, &records)

	l := history.Ledger{}
	hq := history.Q{Session: ht.HorizonSession()}
	ht.Require.NoError(hq.LedgerBySequence(&l, 3))

	ht.Assert.WithinDuration(l.ClosedAt, records[0].LedgerCloseTime, 1*time.Second)
}
