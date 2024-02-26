package models

type Node struct {
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	OwnComponent bool    `json:"ownComponent,omitempty"`
	Nodes        []*Node `json:"nodes"`
}
