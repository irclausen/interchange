package keeper

import (
	"github.com/irclausen/username/interchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
