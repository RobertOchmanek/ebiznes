package database

import (
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func Init() {

	//Open sqlite connection and store data in ebiznes_go.db file
	db, err = gorm.Open("sqlite3", "ebiznes_go.db")

	if err != nil {
		panic("Error while connecting to DB")
	}

	//Migrate tables
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.OrderItem{})
	db.AutoMigrate(&model.Payment{})
	db.AutoMigrate(&model.Cart{})
	db.AutoMigrate(&model.CartItem{})

	initializeData(db)
}

func DbManager() *gorm.DB {
	return db
}

//Initialize database with mocked data
func initializeData(db *gorm.DB) {

	product1 := model.Product{
		CategoryId: 1,
		Name: "iPhone 13 Pro Max",
		Price: 1099.0,
		Image: "https://store.storeimages.cdn-apple.com/4668/as-images.apple.com/is/iphone-13-pro-max-blue-select?wid=940&hei=1112&fmt=png-alpha&.v=1645552346295",
	}
	db.Create(&product1)

	product2 := model.Product{
		CategoryId: 1,
		Name: "iPhone 13",
		Price: 799.0,
		Image: "https://store.storeimages.cdn-apple.com/4668/as-images.apple.com/is/iphone-13-family-hero?wid=940&hei=1112&fmt=png-alpha&.v=1645036276543",
	}
	db.Create(&product2)

	product3 := model.Product{
		CategoryId: 2,
		Name: "iPad Pro",
		Price: 899.0,
		Image: "https://store.storeimages.cdn-apple.com/4668/as-images.apple.com/is/ipad-pro-11-select-202104_FMT_WHH?wid=2000&hei=2000&fmt=png-alpha&.v=1617067382000",
	}
	db.Create(&product3)

	product4 := model.Product{
		CategoryId: 2,
		Name: "iPad Air",
		Price: 599.0,
		Image: "https://store.storeimages.cdn-apple.com/4668/as-images.apple.com/is/ipad-air-select-wifi-blue-202203?wid=940&hei=1112&fmt=png-alpha&.v=1645065732688",
	}
	db.Create(&product4)

	category1 := model.Category{
		Name: "Smartphone",
		Products: []model.Product{},
	}
	db.Create(&category1)

	category2 := model.Category{
		Name: "Tablet",
		Products: []model.Product{},
	}
	db.Create(&category2)
}