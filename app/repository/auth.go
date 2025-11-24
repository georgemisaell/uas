package repository

import (
	"uas/app/models"
	"uas/database"
)

// GetUserByLoginInput mencari user & role berdasarkan username/email
func Login(loginInput string) (models.User, error) {
    var user models.User

    // Query kita pindahkan ke sini
    // Note: Kita scan password_hash ke user.PasswordHash
    // Pastikan struct User field PasswordHash ada tag `json:"-"` biar aman
    query := `
        SELECT 
            u.id, u.username, u.email, u.password_hash, u.full_name, 
            u.role_id, r.name as role_name, u.is_active, u.created_at
        FROM users u
        JOIN roles r ON u.role_id = r.id
        WHERE u.username = $1 OR u.email = $1
    `

    err := database.ConnectDB().QueryRow(query, loginInput).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.PasswordHash, // Scan hash langsung ke struct (hidden field)
        &user.FullName,
        &user.RoleID,
        &user.RoleName,
        &user.IsActive,
        &user.CreatedAt,
    )

    return user, err
}