package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer) // 如果有多個要注入直接(NewGRPCServer1, NewGRPCServer2, ...)往後增加即可
