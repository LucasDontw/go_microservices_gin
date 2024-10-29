package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
// 如果有多個要注入直接(NewContentUsecase1, NewContentUsecase2, ...)往後增加即可
var ProviderSet = wire.NewSet(NewContentUsecase)
