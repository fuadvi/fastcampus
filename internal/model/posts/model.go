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
