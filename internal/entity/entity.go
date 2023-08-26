package entity

type ChangeSegmentsRequest struct {
	Slug string `json:"slug"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type ChangeUserSegmentsRequest struct {
	UserID  int64   `json:"user_id"`
	AddSlug []int64 `json:"add_slug"`
	DelSlug []int64 `json:"del_slug"`
}

type CheckSegmentsRequest struct {
	UserID int64 `json:"user_id"`
}

type ChangeSegmentsResponse struct {
	SegmentsId int64  `json:"segments_id"`
	Slug       string `json:"slug"`
	Status     string `json:"status"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
	Status string `json:"status"`
}

type ChangeUserSegmentsResponse struct {
	UserID   int64    `json:"user_id"`
	Segments []string `json:"segments"`
	Status   string   `json:"status"`
}

type CheckSegmentsResponse struct {
	Segments []string `json:"segments"`
	Status   string   `json:"status"`
}
