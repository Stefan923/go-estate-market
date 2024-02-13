package dto

type PostDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PropertyId  uint   `json:"propertyId"`
}

type PostCreationDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostCommentDto struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
}

type PostCommentCreationDto struct {
	Description string `json:"description"`
	PostId      uint   `json:"postId"`
}

type AnnounceCreationDto struct {
	Property PropertyCreationDto    `json:"property"`
	Post     PostCommentCreationDto `json:"post"`
}
