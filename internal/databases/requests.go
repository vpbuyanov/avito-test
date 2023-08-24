package databases

const (
	CreateSegments        = `insert into public."Segments" (slug) values ($1) returning slug`
	DeleteSegments        = `delete from public."Segments" where slug = $1`
	SetUserSegments       = `update public."Users" set segments = $2 where user_id = $1`
	GetActiveUserSegments = `select segments from public."Users" where user_id = $1`
)
