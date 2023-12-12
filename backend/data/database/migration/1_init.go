package migration

import (
	db "backend/data/database"
	"backend/data/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

const (
	bcryptHashCost    int    = 10
	adminUserEmail    string = "admin@estate-market.com"
	adminUserPassword string = "admin"
)

func Run() {
	database := db.GetDatabase()

	createAllTables(database)
	createRolesAndAdminUser(database)
}

func createAllTables(database *gorm.DB) {
	var tables []interface{}

	addTable(database, &tables, model.Country{})

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		log.Println("Error while running `1_init` migration: " + err.Error())
		return
	}
	log.Println("Migration `1_init` has successfully run.")
}

func addTable(database *gorm.DB, tables *[]interface{}, model interface{}) {
	if !database.Migrator().HasTable(model) {
		*tables = append(*tables, model)
	}
}

func createRolesAndAdminUser(database *gorm.DB) {
	adminRole := model.Role{Name: "ADMIN"}

	createRoleIfNotExists(database, &adminRole)
	createRoleIfNotExists(database, &model.Role{Name: "USER"})
	createAdminUserInformation(database, adminRole)
}

func createRoleIfNotExists(database *gorm.DB, role *model.Role) {
	exists := 0
	database.
		Model(&model.Role{}).
		Select("1").
		Where("name = ?", role.Name).
		First(&exists)
	if exists == 0 {
		database.Create(role)
	}
}

func createAdminUserInformation(database *gorm.DB, role model.Role) {
	encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(adminUserPassword), bcryptHashCost)
	if err != nil {
		log.Println("Error while encrypting user password: ", err)
	}

	encryptedPassword := fmt.Sprintf("%s", encryptedPasswordBytes)
	exists := 0
	database.
		Model(&model.UserAccount{}).
		Select("1").
		Where("email = ?", adminUserEmail).
		First(&exists)
	if exists == 0 {
		userAccount := model.UserAccount{Email: adminUserEmail, Password: encryptedPassword}
		database.Create(userAccount)

		user := model.User{
			FirstName:   "Admin",
			LastName:    "Account",
			PhoneNumber: "0555512345",
			Enabled:     true,
			UserAccount: userAccount,
		}
		database.Create(user)

		userRole := model.UserRole{
			User: user,
			Role: role,
		}
		database.Create(userRole)
	}
}
