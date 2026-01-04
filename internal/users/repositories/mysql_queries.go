package repositories

const (
	userInsertQR = `INSERT INTO users (
		identifier, 
		external_id, 
		first_name, 
		last_name, 
		username, 
		email, 
		email_verified, 
		created_at, 
		updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	userSelectByIdentifierQR = `SELECT 
		id, 
		identifier, 
		external_id, 
		first_name, 
		last_name, 
		username, 
		email, 
		email_verified, 
		created_at, 
		updated_at 
		FROM users WHERE identifier = ?;`
)
