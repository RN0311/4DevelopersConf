package halloweenproduct

type HalloweenProduct struct {
	Id      string  `bson:"Id"`
	BrandId string  `bson:"BrandId"`
	Rate    float64 `bson:"Rate"`
}
