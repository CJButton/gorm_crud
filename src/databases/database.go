package databases

// Database does stuff
type Database struct{}

// CreateDatabase makes a db in posgresql
func (Database) CreateDatabase() {
	// dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Tokyo"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
