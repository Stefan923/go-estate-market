package dto

type PostDto struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PropertyId  uint   `json:"propertyId"`
}

type PostCreationDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostCreationWithIdDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PropertyId  uint   `json:"propertyId"`
}

type PostCommentDto struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
}

type PostCommentCreationDto struct {
	Description string `json:"description"`
	PostId      uint   `json:"postId"`
}
