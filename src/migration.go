package src

//AutoMigration update the database schema
func AutoMigration() {
	db := dbConnect()
	defer db.Close() // ultimo comando da função, sera executado depois do db.automigrate

	db.AutoMigrate(user{}, debt{})
	db.Model(&debt{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
