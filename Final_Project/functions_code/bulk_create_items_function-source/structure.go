package helloworld

// structure definition
type Grocery_item struct {
	ID             string `json:"Id" firestore:"ID"`
	Name           string `json:"name" firestore:"Product_Name"`
	Price          int    `json:"price" firestore:"Price (Rs)"`
	Category       string `json:"category" firestore:"Category"`
	Weight         string `json:"weight" firestore:"Weight"`
	Veg            bool   `json:"veg" firestore:"Vegetarian"`
	Brand          string `json:"brand" firestore:"Brand"`
	Quantity       int    `json:"quantity" firestore:"Item_Package_Quantity"`
	Pack_info      string `json:"pack_info" firestore:"Package_Information"`
	Manufacturer   string `json:"manufacturer" firestore:"Manufacturer"`
	Country_origin string `json:"country_origin" firestore:"Country_of_Origin"`
	Availability   bool   `json:"availability,omitempty" firestore:"Product_Availability,omitempty"`
	Discount       bool   `json:"discount,omitempty" firestore:"Discount_Applicable,omitempty"`
	Offers         bool   `json:"offers,omitempty" firestore:"Offers_Applicable,omitempty"`
}