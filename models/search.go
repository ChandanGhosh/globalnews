package models

type Search struct {
	SearchQuery string
	NextPage    int
	TotalPage   int
	Result      News
}
