package types

// UserContributionFollow フォロー
type UserContributionFollow struct {
	ID                 uint    `db:"id" json:"id"`
	UserContributionID int     `db:"user_contribution_id" json:"userContributionId"`
	CreatedAt          string  `db:"created_at" json:"createdAt"`
	UpdatedAt          string  `db:"updated_at" json:"updatedAt"`
	DeletedAt          *string `db:"deleted_at" json:"deletedAt"`
}
