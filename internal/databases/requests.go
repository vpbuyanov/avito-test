package databases

const (
	CreateSegments        = `insert into public."Segments" (slug) values ($1) returning segments_id, slug`
	DeleteSegments        = `delete from public."Segments" where slug = $1 returning segments_id, slug`
	CreateUser            = `insert into public."Users" (name) values ($1) returning user_id`
	SetUserSegments       = `insert into public."Belong" (user_id, segments_id) values($1, $2) returning user_id, segments_id`
	DeleteUserSegments    = `delete from public."Belong" where user_id = ($1) and (segments_id = $2) returning user_id, segments_id`
	GetActiveUserSegments = `select slug from public."Segments" inner join public."Belong" B on "Segments".segments_id = B.segments_id and user_id = $1`
)
