package resource

import (
	"github.com/fonero-project/horizon/httpx"
	"github.com/fonero-project/horizon/render/hal"
	"github.com/fonero-project/horizon/txsub"
	"golang.org/x/net/context"
)

// Populate fills out the details
func (res *TransactionSuccess) Populate(ctx context.Context, result txsub.Result) {
	res.Hash = result.Hash
	res.Ledger = result.LedgerSequence
	res.Env = result.EnvelopeXDR
	res.Result = result.ResultXDR
	res.Meta = result.ResultMetaXDR

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	res.Links.Transaction = lb.Link("/transactions", result.Hash)
	return
}
