package migration

import (
	"fmt"
	db "github.com/Stefan923/go-estate-market/data/database"
	model2 "github.com/Stefan923/go-estate-market/data/model"
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

	addTable(database, &tables, model2.UserAccount{})
	addTable(database, &tables, model2.User{})
	addTable(database, &tables, model2.Role{})
	addTable(database, &tables, model2.UserRole{})

	addTable(database, &tables, model2.Country{})
	addTable(database, &tables, model2.State{})
	addTable(database, &tables, model2.City{})
	addTable(database, &tables, model2.Address{})

	addTable(database, &tables, model2.PropertyCategory{})
	addTable(database, &tables, model2.Property{})
	addTable(database, &tables, model2.PropertyDetails{})
	addTable(database, &tables, model2.PropertyPrice{})

	addTable(database, &tables, model2.Post{})
	addTable(database, &tables, model2.PostComment{})

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
	adminRole := model2.Role{Name: "ADMIN"}

	createRoleIfNotExists(database, &adminRole)
	createRoleIfNotExists(database, &model2.Role{Name: "USER"})
	createAdminUserInformation(database, adminRole)
}

func createRoleIfNotExists(database *gorm.DB, role *model2.Role) {
	exists := 0
	database.
		Model(&model2.Role{}).
		Select("1").
		Where("name = ?", role.Name).
		First(&exists)
	if exists == 0 {
		database.Create(&role)
	}
}

func createAdminUserInformation(database *gorm.DB, role model2.Role) {
	exists := 0
	database.
		Model(&model2.UserAccount{}).
		Select("1").
		Where("email = ?", adminUserEmail).
		First(&exists)
	if exists == 0 {
		encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(adminUserPassword), bcryptHashCost)
		if err != nil {
			log.Println("Error while encrypting user password: ", err)
		}
		encryptedPassword := fmt.Sprintf("%s", encryptedPasswordBytes)

		userAccount := model2.UserAccount{Email: adminUserEmail, Password: encryptedPassword}
		database.Create(&userAccount)

		user := model2.User{
			FirstName:     "Admin",
			LastName:      "Account",
			PhoneNumber:   "0555512345",
			Enabled:       true,
			UserAccountId: userAccount.Id,
		}
		database.Create(&user)

		userRole := model2.UserRole{
			UserId: user.Id,
			RoleId: role.Id,
		}
		database.Create(&userRole)
	}
}
