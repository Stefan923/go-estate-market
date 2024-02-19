package migration

import (
	"fmt"
	"github.com/Stefan923/go-estate-market/config"
	db "github.com/Stefan923/go-estate-market/data/database"
	model2 "github.com/Stefan923/go-estate-market/data/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type CountriesConfig struct {
	Countries []model2.Country
}

const (
	bcryptHashCost    int    = 10
	adminUserEmail    string = "admin@estate-market.com"
	adminUserPassword string = "admin"
)

func Run() {
	database := db.GetDatabase()

	createAllTables(database)
	createRolesAndAdminUser(database)
	loadAndCreateCountries(database)
	createPropertyCategories(database)
	createCurrenciesIfNotExist(database)
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
	addTable(database, &tables, model2.Currency{})

	addTable(database, &tables, model2.PropertyCategory{})
	addTable(database, &tables, model2.Property{})
	addTable(database, &tables, model2.PropertyDetail{})
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
	createRoleIfNotExists(database, &model2.Role{Name: "USER", DefaultRole: true})
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

func loadAndCreateCountries(database *gorm.DB) {
	fileReader := config.FileReader[CountriesConfig]{}
	countries := fileReader.GetContent("/app/config/countries")
	if countries == nil {
		return
	}

	for _, country := range countries.Countries {
		if err := createCountry(database, &country); err != nil {
			return
		}
	}
}

func createCountry(database *gorm.DB, country *model2.Country) error {
	exists := 0

	database.
		Model(&model2.Country{}).
		Select("1").
		Where("name = ?", country.Name).
		First(&exists)

	countryToCreate := model2.Country{
		Name: country.Name,
	}
	if exists == 0 {
		if err := database.Create(&countryToCreate).Error; err != nil {
			log.Println("Error while creating country \": ", countryToCreate.Name, "\": ", err)
			return err
		}
	} else {
		database.
			Where("name = ? and deleted_at is null", country.Name).
			First(&countryToCreate)
	}

	for _, state := range *country.States {
		state.CountryId = countryToCreate.Id
		if err := createState(database, &state); err != nil {
			return err
		}
	}

	return nil
}

func createState(database *gorm.DB, state *model2.State) error {
	exists := 0

	database.
		Model(&model2.State{}).
		Select("1").
		Where("name = ?", state.Name).
		First(&exists)

	stateToCreate := model2.State{
		Name:      state.Name,
		CountryId: state.CountryId,
	}
	if exists == 0 {
		if err := database.Create(&stateToCreate).Error; err != nil {
			log.Println("Error while creating state \"", stateToCreate.Name, "\": ", err)
			return err
		}
	} else {
		database.
			Where("name = ? and deleted_at is null", state.Name).
			First(&stateToCreate)
	}

	for _, city := range *state.Cities {
		city.StateId = stateToCreate.Id
		if err := createCity(database, &city); err != nil {
			return err
		}
	}

	return nil
}

func createCity(database *gorm.DB, city *model2.City) error {
	exists := 0

	database.
		Model(&model2.City{}).
		Select("1").
		Where("name = ?", city.Name).
		First(&exists)

	if exists == 0 {
		if err := database.Create(city).Error; err != nil {
			log.Println("Error while creating city \"", city.Name, "\": ", err)
			return err
		}
	}
	return nil
}

func createPropertyCategories(database *gorm.DB) {
	createPropertyCategoryIfNotExists(database, &model2.PropertyCategory{Name: "House", Icon: "category-house"})
	createPropertyCategoryIfNotExists(database, &model2.PropertyCategory{Name: "Apartment", Icon: "category-apartment"})
	createPropertyCategoryIfNotExists(database, &model2.PropertyCategory{Name: "Terrain", Icon: "category-terrain"})
}

func createPropertyCategoryIfNotExists(database *gorm.DB, propertyCategory *model2.PropertyCategory) {
	exists := 0
	database.
		Model(&model2.PropertyCategory{}).
		Select("1").
		Where("name = ?", propertyCategory.Name).
		First(&exists)
	if exists == 0 {
		database.Create(&propertyCategory)
	}
}

func createCurrenciesIfNotExist(database *gorm.DB) {
	createCurrencyIfNotExists(database, &model2.Currency{Code: "EUR"})
	createCurrencyIfNotExists(database, &model2.Currency{Code: "USD"})
	createCurrencyIfNotExists(database, &model2.Currency{Code: "GBP"})
	createCurrencyIfNotExists(database, &model2.Currency{Code: "RON"})
}

func createCurrencyIfNotExists(database *gorm.DB, currency *model2.Currency) {
	exists := 0
	database.
		Model(&model2.Currency{}).
		Select("1").
		Where("code = ?", currency.Code).
		First(&exists)
	if exists == 0 {
		database.Create(&currency)
	}
}
