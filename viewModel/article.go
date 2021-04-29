package viewModel

type ViewArticle struct {
	BaseModel
	Title string
	Content string
	Author string
	Channel string
	Published bool
	ReadAmount int
	LikeAmount int
}


