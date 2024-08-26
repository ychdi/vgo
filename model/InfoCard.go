package Model

type InfoCard struct {
	Model
	InfoId  uint      `gorm:"column:info_id" json:"info_id"`
	Name    string    `gorm:"column:name" json:"name"`
	CardLog []CardLog `gorm:"foreignKey:CardId"`
}
