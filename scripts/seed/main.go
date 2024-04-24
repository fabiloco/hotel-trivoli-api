package main

import (
	"fabiloco/hotel-trivoli-api/api/config"
	"fabiloco/hotel-trivoli-api/api/utils"
	"fabiloco/hotel-trivoli-api/pkg/entities"
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func seed(db *gorm.DB) {
 // Drop tables
  db.Migrator().DropTable(
    &entities.Product{}, 
    &entities.ProductType{}, 
    &entities.Service{},
    &entities.Room{},
    &entities.RoomHistory{},
    &entities.Receipt{},
    &entities.IndividualReceipt{},
    &entities.Role{}, 
    &entities.Person{}, 
    &entities.User{},
  )
  db.AutoMigrate(
    &entities.Product{}, 
    &entities.ProductType{}, 
    &entities.Service{},
    &entities.Room{},
    &entities.RoomHistory{},
    &entities.Receipt{},
    &entities.IndividualReceipt{},
    &entities.Role{}, 
    &entities.Person{}, 
    &entities.User{}, 
	)

	// Create ProductTypes
	productTypes := []entities.ProductType{
		{Name: "Bebidas"},
		{Name: "Mecatos"},
		{Name: "Preservativos"},
		{Name: "Medicamentos"},
		{Name: "Aseo personal"},
		{Name: "Misceláneo"},
	}

	for _, pt := range productTypes {
		db.Create(&pt)
	}

	// Create Products
	products := []entities.Product{
		{Name: "Cerveza Lata", Stock: 99, Price: 6000.0, Type: []entities.ProductType{productTypes[0]}, Img: "https://copservir.vtexassets.com/arquivos/ids/1171256-800-auto?v=638418762835070000&width=800&height=auto&aspect=true"},
		{Name: "Gaseosas", Stock: 99, Price: 5000.0, Type: []entities.ProductType{productTypes[0]}, Img: "https://simonparrilla.com.co/wp-content/uploads/2022/06/SIMON-PARRILLA-BEBIDAS-GASEOSA.webp"},
		{Name: "Speed Max", Stock: 99, Price: 3000.0, Type: []entities.ProductType{productTypes[0]}, Img: "https://exitocol.vtexassets.com/arquivos/ids/20045747/Bebida-SPEED-MAX-473-ml-3469968_a.jpg?v=638338334845430000"},
		{Name: "JP Chenet Rosada", Stock: 99, Price: 12000.0, Type: []entities.ProductType{productTypes[0]}, Img: "https://jumbocolombiaio.vtexassets.com/arquivos/ids/206164/3500610093708.jpg?v=637814201879370000"},
		{Name: "Agua Cristalina", Stock: 99, Price: 2000.0, Type: []entities.ProductType{productTypes[0]}, Img: "https://mercayahorra.com/wp-content/uploads/2020/06/Agua-Cristalina-x300Ml.jpg"},
		{Name: "Smirnoff Ice", Stock: 99, Price: 15000.0, Type: []entities.ProductType{productTypes[0]}, Img: "https://mistiendas.com.co/2413-large_default/smirnoff-ice-original-botella.jpg"},
		{Name: "Serv. Aguardiente", Stock: 99, Price: 45000.0, Type: []entities.ProductType{productTypes[0]}, Img: "https://copservir.vtexassets.com/arquivos/ids/1170449/DTE-SS-COOL-NIGHT-GEL_F.jpg?v=638418477793370000"},
		{Name: "Serv. Ron 8 Años", Stock: 99, Price: 90000.0, Type: []entities.ProductType{productTypes[0]}, Img: "https://copservir.vtexassets.com/arquivos/ids/1170449/DTE-SS-COOL-NIGHT-GEL_F.jpg?v=638418477793370000"},

		{Name: "Manicero", Stock: 99, Price: 2000.0, Type: []entities.ProductType{productTypes[1]}, Img: "https://bodegalaesperanza.com/images/bodegalaesperanza/products/6032aeb97202f.jpeg"},
		{Name: "Papa Margarita", Stock: 99, Price: 3000.0, Type: []entities.ProductType{productTypes[1]}, Img: "https://www.confipuma.com/app_data_archivos/confipuma.soomi.co/productos/producto_ba268b32e52ef9e1233086f1d7f3388b877b860a1642763324.png"},
		{Name: "Chocolatina", Stock: 99, Price: 1000.0, Type: []entities.ProductType{productTypes[1]}, Img: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSKYrsRIS9dAg-LoLgZa_jNb__L5Dtvzp4yXUrTFk3Eaw&s"},
		{Name: "Lecherita", Stock: 99, Price: 2000.0, Type: []entities.ProductType{productTypes[1]}, Img: "https://copservir.vtexassets.com/arquivos/ids/1170449/DTE-SS-COOL-NIGHT-GEL_F.jpg?v=638418477793370000"},

		{Name: "Preservativo Today", Stock: 99, Price: 5000.0, Type: []entities.ProductType{productTypes[2]}, Img: "https://dcdn.mitiendanube.com/stores/001/307/250/products/condones-today-lubricado-x-1-nonita-co1-920cf66f3dc363c58216383955336130-480-0.jpg"},

		{Name: "Sildenafil", Stock: 99, Price: 3000.0, Type: []entities.ProductType{productTypes[3]}, Img: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSf0w0uCyf1x-z41skSL03FVURW5mX95KAXwWfxXCutSA&s"},
		{Name: "Fisiomax lubricante", Stock: 99, Price: 3000.0, Type: []entities.ProductType{productTypes[3]}, Img: "https://pasteurio.vtexassets.com/arquivos/ids/187396/Cuidado-Personal-Lubricantes_Fisiomax_Pasteur_155205_sobre_01.jpg?v=638188879961730000"},
		{Name: "Aspirina tableta", Stock: 99, Price: 2000.0, Type: []entities.ProductType{productTypes[3]}, Img: "https://www.dfs.net.co/3494-large_default/aspirina-ultra-x-20-tabletas.jpg"},

		{Name: "Desodorante Sobre", Stock: 99, Price: 3000.0, Type: []entities.ProductType{productTypes[4]}, Img: "https://copservir.vtexassets.com/arquivos/ids/1170449/DTE-SS-COOL-NIGHT-GEL_F.jpg?v=638418477793370000"},
		{Name: "Talco Mexana", Stock: 99, Price: 7000.0, Type: []entities.ProductType{productTypes[4]}, Img: "https://jumbocolombiaio.vtexassets.com/arquivos/ids/189129/7702123008859.jpg?v=637813990505970000"},
		{Name: "Maquina de afeitar", Stock: 99, Price: 3000.0, Type: []entities.ProductType{productTypes[4]}, Img: "https://copservir.vtexassets.com/arquivos/ids/1254323/MAQUINA-PRESTOBARBA-3_F.png?v=638455411739930000"},
		{Name: "Crema dental", Stock: 99, Price: 3000.0, Type: []entities.ProductType{productTypes[4]}, Img: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRlNjPsWmfeC2bBZaLqufFKau8IBlPG9YUtSJyUUpB9Hw&s"},
		{Name: "Protectores diarios", Stock: 99, Price: 1000.0, Type: []entities.ProductType{productTypes[4]}, Img: "https://exitocol.vtexassets.com/arquivos/ids/18220766/Protectores-Diarios-Largos-X-50-Unidades-233488_a.jpg?v=638186717751600000"},
		{Name: "Cepillo dental", Stock: 99, Price: 2000.0, Type: []entities.ProductType{productTypes[4]}, Img: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTq0LoDlHGH4B9gXp4SPSV5CGxfzCudD-L0IyRpM1l_Qw&s"},
		{Name: "Toalla Higienica", Stock: 99, Price: 2000.0, Type: []entities.ProductType{productTypes[4]}, Img: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTBcLR5KowTZtdolBKuRFntosfBgTE_G9PKATy15g1IHA&s"},
		{Name: "Shampoo", Stock: 99, Price: 2000.0, Type: []entities.ProductType{productTypes[4]}, Img: "https://lavaquita.co/cdn/shop/products/supermercados_la_vaquita_supervaquita_shampoo_h_s_375c_2en1_limpieza_renovador_cuidado_capilar_700x700.jpg?v=1620655513"},

		{Name: "Mechera", Stock: 99, Price: 2000.0, Type: []entities.ProductType{productTypes[5]}, Img: "https://carulla.vtexassets.com/arquivos/ids/6057290/encendedor-mechera-swiss-texas-x-24-unidades.jpg?v=637680294582130000"},
		// Add more mock data as needed
	}

	for _, p := range products {
		db.Create(&p)
	}

	// Create Services
	services := []entities.Service{
		{Name: "Motel moto", Price: 26000.0},
		{Name: "Motel carro", Price: 28000.0},
		{Name: "Motel tercer piso", Price: 22000.0},
		{Name: "Motel tina", Price: 42000.0},
		{Name: "Hotel tercer piso solo", Price: 27000.0},
		{Name: "Hotel tercer piso pareja", Price: 58000.0},
		// Add more mock data as needed
	}

	for _, s := range services {
		db.Create(&s)
	}

	// Create Rooms
	rooms := []entities.Room{
		{Number: 101},
		{Number: 102},
		{Number: 103},
		{Number: 104},
		{Number: 105},
		{Number: 106},
		{Number: 107},
		{Number: 108},
		{Number: 109},
		{Number: 110},
		{Number: 301},
		{Number: 302},
		{Number: 303},
		// Add more mock data as needed
	}

	for _, r := range rooms {
		db.Create(&r)
	}

  var now_date = time.Now()
  var now_date2 = time.Now()
	// Create RoomHistories
	roomHistories := []entities.RoomHistory{
		{
			StartDate: time.Now().Add(-24 * time.Hour),
			EndDate:   &now_date,
			RoomID:    rooms[0].ID,
			ServiceID: services[0].ID,
		},
		{
			StartDate: time.Now().Add(-48 * time.Hour),
			EndDate:   &now_date2,
			RoomID:    rooms[1].ID,
			ServiceID: services[1].ID,
		},
		// Add more mock data as needed
	}

	for _, rh := range roomHistories {
		db.Create(&rh)
	}

	// Create Receipts
	receipts := []entities.Receipt{
		{
			TotalPrice:  150.0,
			TotalTime:   24 * time.Hour,
			Products:    []entities.Product{products[0], products[1]},
			ServiceID:   services[0].ID,
			RoomID:      rooms[0].ID,
		},
		{
			TotalPrice:  80.0,
			TotalTime:   24 * time.Hour,
			Products:    []entities.Product{products[2], products[3]},
			ServiceID:   services[1].ID,
			RoomID:      rooms[1].ID,
		},
		// Add more mock data as needed
	}

	for _, rec := range receipts {
		db.Create(&rec)
	}

	// Create Role
	role := []entities.Role{
		{
      Name: "USER",
		},
		{
      Name: "ADMIN",
		},
		// Add more mock data as needed
	}

	for _, rec := range role {
		db.Create(&rec)
	}


	// Create User
	person := []entities.Person{
		{
      Firstname: "admin",
      Lastname: "admin",
      Identification: "12345",
      Birthday: now_date.String(),
		},
		{
      Firstname: "normal",
      Lastname: "user",
      Identification: "54321",
      Birthday: now_date.String(),
		},
		// Add more mock data as needed
	}

	for _, rec := range person {
		db.Create(&rec)
	}


  adminPasswordHashed, _ := utils.HashPassword("admin")
  userPasswordHashed, _ := utils.HashPassword("normaluser")

	// Create User
	user := []entities.User{
		{
      PersonID: 1,
      Username: "admin",
      Password: adminPasswordHashed,
      RoleID: 2,
		},
		{
      PersonID: 2,
      Username: "user",
      Password: userPasswordHashed,
      RoleID: 1,
		},
		// Add more mock data as needed
	}

	for _, rec := range user {
		db.Create(&rec)
	}
}

func main() {
  var DB *gorm.DB
  p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

  if err != nil {
    println("Error parsing DB_PORT variable.")
  }

  dns := fmt.Sprintf(
    "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
    config.Config("DB_USER"), 
    config.Config("DB_PASSWORD"), 
    config.Config("DB_HOST"), 
    port,
    config.Config("DB_NAME"),
  )

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	seed(DB)
}
