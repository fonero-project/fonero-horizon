package resource

import (
	"github.com/fonero-project/fonero-horizon/db2/core"
)

func (this *AccountFlags) Populate(row core.Account) {
	this.AuthRequired = row.IsAuthRequired()
	this.AuthRevocable = row.IsAuthRevocable()
}
