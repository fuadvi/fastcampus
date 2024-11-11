package posts

import "time"

type (
	CreateRequestPost struct {
		PostTitle    string   `json:"post_title"`
		PostContent  string   `json:"post_content"`
		PostHashTags []string `json:"post_hash_tag"`
	}
)

type (
	PostModel struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		PostTitle    string    `db:"post_title"`
		PostContent  string    `db:"post_content"`
		PostHashTags string    `db:"post_hash_tag"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
		CreatedBy    string    `db:"created_by"`
		UpdatedBy    string    `db:"updated_by"`
	}
)

type (
	GetAllPost struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID           int64    `json:"id"`
		UserID       int64    `json:"user_id"`
		Username     string   `json:"username"`
		PostTitle    string   `json:"post_title"`
		PostContent  string   `json:"post_content"`
		PostHashTags []string `json:"post_hash_tag"`
		IsLiked      bool     `json:"is_liked"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}

	GetPostResponse struct {
		PostDetail Post      `json:"post_detail"`
		LikeCount  int       `json:"like_count"`
		Comments   []Comment `json:"comments"`
	}

	Comment struct {
		ID             int64  `json:"id"`
		UserID         int64  `json:"user_id"`
		Username       string `json:"username"`
		CommentContent string `json:"comment_content"`
	}
)
