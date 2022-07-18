// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Programmer struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Picture *string  `json:"picture"`
	Company string   `json:"company"`
	Skills  []*Skill `json:"skills"`
}

type Skill struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Icon       *string `json:"icon"`
	Importance int     `json:"importance"`
}
