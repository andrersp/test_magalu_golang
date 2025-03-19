package entity

type Client struct {
	Name      string
	Email     string
	Favorites []Favorite
	ID        uint
}
