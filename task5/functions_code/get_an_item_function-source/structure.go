package helloworld

// structure definition
type Grocery_item struct {
	ID             string `json:"ID" firestore:"ID"`
	Name           string `json:"Product_Name" firestore:"Product_Name"`
	Price          int    `json:"Price (Rs)" firestore:"Price (Rs)"`
	Category       string `json:"Category" firestore:"Category"`
	Image          string `json:"Image" firestore:"Image"`
	Weight         string `json:"Weight" firestore:"Weight"`
	Veg            bool   `json:"Vegetarian" firestore:"Vegetarian"`
	Brand          string `json:"Brand" firestore:"Brand"`
	Quantity       int    `json:"Item_Package_Quantity" firestore:"Item_Package_Quantity"`
	Pack_info      string `json:"Package_Information" firestore:"Package_Information"`
	Manufacturer   string `json:"Manufacturer" firestore:"Manufacturer"`
	Country_origin string `json:"Country_of_Origin" firestore:"Country_of_Origin"`
	Availability   bool   `json:"Product_Availability,omitempty" firestore:"Product_Availability,omitempty"`
	Discount       bool   `json:"Discount_Applicable,omitempty" firestore:"Discount_Applicable,omitempty"`
	Offers         bool   `json:"Offers_Applicable,omitempty" firestore:"Offers_Applicable,omitempty"`
}
