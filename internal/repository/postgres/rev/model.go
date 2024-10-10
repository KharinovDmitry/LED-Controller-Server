package rev

type Rev struct {
	ID     int `db:"id"`
	Width  int `db:"width"`
	Height int `db:"height"`
}
