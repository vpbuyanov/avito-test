package entity

type ChangeSegmentsRequest struct {
	Slug string `json:"slug"`
}

type ChangeUserSegmentsRequest struct {
	UserID  int64    `json:"user_id"`
	AddSlug []string `json:"add_slug"`
	DelSlug []string `json:"del_slug"`
}

type CheckSegmentsRequest struct {
	UserID int64 `json:"user_id"`
}

type ChangeSegmentsResponse struct {
	Slug   string `json:"slug"`
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
