package models

//we will be exporting these products
//in go the products to be exported must be starting with a
// CAPITAL Letter
type Product struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Qty         int    `json:"qty"`
	LastUpdated string `json:"last_updated"`
}
