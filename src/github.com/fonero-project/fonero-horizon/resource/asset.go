package resource

import (
	"github.com/fonero-project/fonero-golang/xdr"
	"golang.org/x/net/context"
)

func (this *Asset) Populate(ctx context.Context, asset xdr.Asset) error {
	return asset.Extract(&this.Type, &this.Code, &this.Issuer)
}
