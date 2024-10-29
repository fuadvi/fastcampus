package memberships

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type (
	UserModel struct {
		ID        int64     `db:"id"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		Email     string    `db:"email"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}

	LoginResponse struct {
		AccessToken string `json:"access_token"`
	}
)
