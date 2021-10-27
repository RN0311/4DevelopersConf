package request_models

type CreateHalloweenProductModel struct {
	Id      string  `json:"id"`
	BrandId string  `json:"brandId"`
	Rate    float64 `json:"rate"`
}
