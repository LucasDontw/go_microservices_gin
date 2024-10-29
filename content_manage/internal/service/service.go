package service

import "github.com/google/wire"

// ProviderSet is service providers.
// 如果有多個要注入直接(NewAppService1, NewAppService2, ...)往後增加即可，
// 這邊不須要指定各Service裡面的(uc *biz.xxxUsecase)，只需要biz層將相對應的xxxUsecase註冊，
// Service註冊時即可自行去找尋相對應的xxxUsecase注入其中
var ProviderSet = wire.NewSet(NewAppService)
