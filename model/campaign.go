package model

type Campaign struct {
	ID     string `json:"cid"`
	Img    string `json:"img"`
	CTA    string `json:"cta"`
	Status string `json:"-"`
}
