package dto

type ConsumerCreateReq struct {
	Code             string
	Slug             string
	Secret           string
	WhiteListMethods []string
}

type ConsumerUpdateReq struct {
	Code             string
	Slug             string
	Secret           string
	WhiteListMethods []string
}

type ConsumerValidateReq struct {
	ConsumerId string
	Method     string
}
