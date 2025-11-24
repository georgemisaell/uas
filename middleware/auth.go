package middleware

import (
	"strings"
	"uas/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Token akses diperlukan",
			})
		}

		// Extract token dari "Bearer TOKEN"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Format token tidak valid",
			})
		}

		// Validasi token
		claims, err := utils.ValidateToken(tokenParts[1])
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Token tidak valid atau expired",
			})
		}

		// Simpan informasi user di context
		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)
		c.Locals("role_name", claims.RoleName)

		return c.Next()
	}
}

// Middleware untuk memerlukan role admin 
func AdminOnly() fiber.Handler { 
    return func(c *fiber.Ctx) error { 
        role := c.Locals("role_name").(string) 
        if role != "Admin" { 
            return c.Status(403).JSON(fiber.Map{ 
                "error": "Akses ditolak. Hanya admin yang diizinkan", 
            }) 
        } 
        return c.Next() 
    } 
}

// Middleware untuk memerlukan role lecture 
func Lecture() fiber.Handler { 
    return func(c *fiber.Ctx) error { 
        role := c.Locals("role_name").(string) 
        if role != "Dosen Wali" { 
            return c.Status(403).JSON(fiber.Map{ 
                "error": "Akses ditolak. Hanya dosen dan admin yang diizinkan", 
            }) 
        } 
        return c.Next() 
    } 
}

// Middleware untuk memerlukan role student 
func Student() fiber.Handler { 
    return func(c *fiber.Ctx) error { 
        role := c.Locals("role_name").(string) 
        if role != "Mahasiswa" { 
            return c.Status(403).JSON(fiber.Map{ 
                "error": "Akses ditolak. Hanya mahasiswa dan admin yang diizinkan", 
            }) 
        } 
        return c.Next() 
    } 
}