package representation

type HalloweenProduct struct {
	Id      string  `json:"id"`
	BrandId string  `json:"brandId"`
	Rate    float64 `json:"rate"`
}
