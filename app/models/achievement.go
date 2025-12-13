package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateAchievementRequest struct {
	AchievementType string                 `json:"achievementType" validate:"required"`
	Title           string                 `json:"title" validate:"required"`
	Description     string                 `json:"description"`
	Details         map[string]interface{} `json:"details"`
	Tags            []string               `json:"tags"`
	EventDate       time.Time              `json:"eventDate"`
}

type AchievementMongo struct {
	ID              primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	StudentID       string                 `bson:"studentId" json:"student_id"`
	AchievementType string                 `bson:"achievementType" json:"achievement_type"`
	Title           string                 `bson:"title" json:"title"`
	Description     string                 `bson:"description" json:"description"`
	Details         map[string]interface{} `bson:"details" json:"details"`
	Tags            []string               `bson:"tags" json:"tags"`
	Points          int                    `bson:"points" json:"points"`
	CreatedAt       time.Time              `bson:"createdAt" json:"created_at"`
	UpdatedAt       time.Time              `bson:"updatedAt" json:"updated_at"`
}

type AchievementReference struct {
	ID                 string    `json:"id"`
	StudentID          string    `json:"student_id"`
	MongoAchievementID string    `json:"mongo_achievement_id"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
}