package data

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	model "ArquicturaSW/model"
)

func InsertData(db *gorm.DB) {
	log.Info("Inserting data... ")

	err := db.First(&model.User{}).Error

	if err != nil {
		db.Create(&model.User{Name: "Ian", LastName: "Horrocks", UserName: "ianho10", Email: "@ucc.edu.ar", Password: "hola123"})
		db.Create(&model.User{Name: "Nano", LastName: "Nallar", UserName: "nanonallar", Email: "qseyo@ucc.edu.ar", Password: "hola123"})
		db.Create(&model.User{Name: "Elian", LastName: "Stuyck", UserName: "elian123", Email: "qseyo@ucc.edu.ar", Password: "hola123"})
		db.Create(&model.User{Name: "Agus", LastName: "Ferrero", UserName: "agusferrero", Email: "qseyo@ucc.edu.ar", Password: "hola123"})
		db.Create(&model.User{Name: "Tomi", LastName: "Calvo", UserName: "tomicalvo", Email: "qseyo@ucc.edu.ar", Password: "hola123"})
	}

	err = db.First(&model.Category{}).Error

	if err != nil {
		db.Create(&model.Category{Name: "Electronics"})              //1
		db.Create(&model.Category{Name: "Computers"})                //2
		db.Create(&model.Category{Name: "Automotive"})               //3
		db.Create(&model.Category{Name: "Baby"})                     //4
		db.Create(&model.Category{Name: "Beauty and personal care"}) //5
		db.Create(&model.Category{Name: "Toys and Games"})           //6
		db.Create(&model.Category{Name: "Sports and Outdoors"})      //7
	}

	err = db.First(&model.Product{}).Error

	if err != nil {
		db.Create(&model.Product{CategoryID: 1, Name: "SAMSUNG 32-inch Class LED Smart FHD TV", Description: "Model 2018", Price: 40000, Stock: 15})
		db.Create(&model.Product{CategoryID: 1, Name: "TCL 32-inch Class 3-Series HD LED Smart Android TV - 32S334", Description: "Model 2021", Price: 42000, Stock: 8})
		db.Create(&model.Product{CategoryID: 1, Name: "TCL 32-inch 1080p Roku Smart LED TV - 32S327", Description: "Model 2019", Price: 41000, Stock: 10})
		db.Create(&model.Product{CategoryID: 1, Name: "SAMSUNG 32-Inch Class QLED", Description: "4K UHD Dual LED Quantum HDR Smart TV with Alexa Built-in", Price: 53000, Stock: 5})
		db.Create(&model.Product{CategoryID: 1, Name: "Apple AirPods", Description: "2nd Generation", Price: 10000, Stock: 4})
		db.Create(&model.Product{CategoryID: 1, Name: "SanDisk 128GB Ultra microSD", Description: "Memory Card with Adapter - 120MB/s, C10, U1, Full HD", Price: 6000, Stock: 25})
		db.Create(&model.Product{CategoryID: 1, Name: "SAMSUNG EVO Select Micro SD", Description: "Memory-Card + Adapter, 256GB microSDXC 130MB/s Full HD & 4K UHD", Price: 7000, Stock: 25})
		db.Create(&model.Product{CategoryID: 1, Name: "SanDisk 128GB microSD", Description: "Licensed for Nintendo-Switch - SDSQXAO-128G-GNCZN", Price: 5000, Stock: 10})
		db.Create(&model.Product{CategoryID: 2, Name: "Logitech Wireless Combo", Description: "MK270", Price: 12000, Stock: 10})
		db.Create(&model.Product{CategoryID: 2, Name: "Logitech MX Keys Advanced Wireless", Description: " Illuminated Keyboard, Tactile Responsive Typing, Backlighting, Bluetooth, USB-C", Price: 7000, Stock: 10})
		db.Create(&model.Product{CategoryID: 2, Name: "HP 15-inch Laptop", Description: "11th Generation Intel Core i5-1135G7, Intel Iris Xe Graphics, 8 GB RAM, 256 GB SSD, Windows 11 Home", Price: 75000, Stock: 15})
		db.Create(&model.Product{CategoryID: 2, Name: "Acer Aspire 5 A515-46-R3UB", Description: "15.6 Full HD IPS Display | AMD Ryzen 3", Price: 70000, Stock: 7})
		db.Create(&model.Product{CategoryID: 2, Name: "Lenovo Flex 5 Laptop", Description: "14.0 FHD Touch Display, AMD Ryzen 5 5500U, 16GB RAM, 256GB Storage", Price: 72000, Stock: 16})
		db.Create(&model.Product{CategoryID: 2, Name: "Samsung Galaxy Tab A8", Description: "Android Tablet, 10.5” LCD Screen, 32GB Storage, Long-Lasting Battery", Price: 35000, Stock: 8})
		db.Create(&model.Product{CategoryID: 2, Name: "Samsung Galaxy Tab S8", Description: "Android Tablet, 11” LCD Screen, 128GB Storage, DeX Productivity", Price: 30000, Stock: 4})
		db.Create(&model.Product{CategoryID: 2, Name: "Samsung Tab A7 Lite 8.7", Description: "Gray 32GB (SM-T220NZAAXAR)", Price: 36000, Stock: 4})
		db.Create(&model.Product{CategoryID: 2, Name: "Acer Nitro 5 AN515-55-53E5 Gaming Laptop", Description: "Intel Core i5-10300H | NVIDIA GeForce RTX 3050 Laptop GPU | 15.6 FHD 144Hz IPS Display", Price: 200000, Stock: 7})
		db.Create(&model.Product{CategoryID: 2, Name: "Acer Aspire 5 Slim Laptop", Description: "15.6 inches Full HD IPS Display, AMD Ryzen 3 3200U, Vega 3 Graphics, 4GB DDR4, 128GB SSD", Price: 78000, Stock: 3})
		db.Create(&model.Product{CategoryID: 2, Name: "Samsung Chromebook 4", Description: "11.6 Intel UHD Graphics 600, Intel Celeron Processor N4020, 4GB, 32GB, Wi-Fi", Price: 40000, Stock: 10})
		db.Create(&model.Product{CategoryID: 2, Name: "HP ChromeBook 11 G4 EE: 11.6-inch", Description: "Intel Celeron N2840 2.16GHz | 16GB eMMC SSD | 4GB RAM | Chrome OS - Black", Price: 35000, Stock: 6})
		db.Create(&model.Product{CategoryID: 2, Name: "Lenovo IdeaPad 3 Laptop", Description: "14.0 FHD Display, AMD Ryzen 5 5500U, 8GB RAM, 256GB Storage, AMD Radeon 7 Graphics", Price: 100000, Stock: 13})
		db.Create(&model.Product{CategoryID: 2, Name: "Lenovo IdeaPad Duet 5 Chromebook", Description: "OLED 13.3 FHD Touch Display, Snapdragon SC7180, 4GB RAM, 64GB Storage", Price: 130000, Stock: 17})
		db.Create(&model.Product{CategoryID: 2, Name: "ASUS Chromebook Detachable CM3", Description: "10.5 Touchscreen WUXGA 16:10 Display, MediaTek 8183 Processor 64GB Storage, 4GB RAM", Price: 120000, Stock: 3})
		db.Create(&model.Product{CategoryID: 2, Name: "Lenovo - Tab P11 Plus", Description: "2K Display - MediaTek Octa-Core Processor - 6GB Memory - 128GB Storage - Android 11", Price: 190000, Stock: 8})
		db.Create(&model.Product{CategoryID: 2, Name: "Lenovo - 2021 - IdeaPad 3 - Gaming Laptop", Description: "AMD Ryzen 5 5600H - 8GB RAM - 256GB Storage - NVIDIA GeForce GTX 1650 - 15.6 FHD Display ", Price: 300000, Stock: 5})
		db.Create(&model.Product{CategoryID: 2, Name: "MSI Pulse GL66 Gaming Laptop", Description: "15.6 144Hz FHD 1080p Display, Intel Core i7-11800H, NVIDIA GeForce RTX 3070, 16GB, 512GB SSD", Price: 600000, Stock: 12})
		db.Create(&model.Product{CategoryID: 2, Name: "2020 HP 14 HD (1366 x 768) Thin and Light Laptop PC", Description: "Intel Celeron N4020 Dual-Core Processor, 4GB DDR4 Memory, 64GB eMMC", Price: 45000, Stock: 6})
		db.Create(&model.Product{CategoryID: 2, Name: "Acer Predator Helios 300 PH315-54-760S Gaming Laptop", Description: "Intel i7-11800H | NVIDIA GeForce RTX 3060 Laptop GPU | 15.6 Full HD 144Hz 3ms IPS Display | 16GB DDR4 | 512GB SSD", Price: 300000, Stock: 7})
		db.Create(&model.Product{CategoryID: 2, Name: "MSI Stealth 15M Gaming Laptop", Description: "15.6 144Hz FHD 1080p Display, Intel Core i7-11375H, NVIDIA GeForce RTX 3060, 16GB, 512GB SSD", Price: 300000, Stock: 12})
		db.Create(&model.Product{CategoryID: 2, Name: "Samsung Galaxy Tab S6 Lite 10.4", Description: "64GB Wi-Fi Tablet Oxford Gray - SM-P610NZAAXAR - S Pen Included", Price: 30000, Stock: 20})
		db.Create(&model.Product{CategoryID: 2, Name: "ASUS TUF Gaming F15 Gaming Laptop", Description: "15.6” 144Hz FHD IPS-Type Display, Intel Core i5-10300H Processor, GeForce GTX 1650, 8GB DDR4 RAM, 512GB PCIe SSD", Price: 170000, Stock: 6})
		db.Create(&model.Product{CategoryID: 2, Name: "HP Pavilion 15 Laptop, 11th Gen", Description: "Intel Core i7-1165G7 Processor, 16 GB RAM, 512 GB SSD Storage, Full HD IPS micro-edge Display", Price: 200000, Stock: 14})
		db.Create(&model.Product{CategoryID: 3, Name: "Microfiber Cleaning Cloths", Description: "Non-Abrasive, Reusable and Washable - Pack of 24", Price: 1100, Stock: 100})
		db.Create(&model.Product{CategoryID: 3, Name: "Chenille Premium Scratch-Free Microfiber Wash Mitt", Description: "MIC493, Lime Green", Price: 600, Stock: 50})
		db.Create(&model.Product{CategoryID: 3, Name: "Ontel Car Wash Cannon", Description: "Foam Blaster Hose Nozzle Spray Gun", Price: 1500, Stock: 60})
		db.Create(&model.Product{CategoryID: 3, Name: "Armor All, AA255 ", Description: "2.5 Gallon 2 Peak HP Wet/Dry Utility Shop Vacuum , Orange", Price: 4000, Stock: 30})
		db.Create(&model.Product{CategoryID: 3, Name: "Black Magic BM23", Description: "Tire Wet, 23 oz", Price: 800, Stock: 130})
		db.Create(&model.Product{CategoryID: 3, Name: "Chemical Guys TVD11316 Tire Kicker Sprayable", Description: "Extra Glossy Tire Shine", Price: 1300, Stock: 65})
		db.Create(&model.Product{CategoryID: 3, Name: "Chemical Guys SPI_402_16", Description: "Heavy Metal Polish Restorer and Protectant, 16 Ounce", Price: 2500, Stock: 40})
		db.Create(&model.Product{CategoryID: 3, Name: "Torq Foam Blaster 6 Foam Wash Gun", Description: "Chemical Guys ACC_326", Price: 3500, Stock: 15})
		db.Create(&model.Product{CategoryID: 3, Name: "Large Tire Dressing Applicator Pad", Description: "Chemical Guys - ACC_300_2 Acc_3002 ", Price: 800, Stock: 150})
		db.Create(&model.Product{CategoryID: 3, Name: "Amazon Basics Deluxe Microfiber Car Wash Mitt", Description: "(2 Pack)", Price: 1300, Stock: 65})
		db.Create(&model.Product{CategoryID: 3, Name: "Step Wax & Dry-26 oz. Double Pack with Microfiber Towel", Description: "Turtle Wax 50834 ", Price: 2000, Stock: 40})
		db.Create(&model.Product{CategoryID: 3, Name: "Grips Wiper Blade Squeegee", Description: "OXO", Price: 1500, Stock: 45})
		db.Create(&model.Product{CategoryID: 3, Name: "Optimum No Rinse Wash & Shine", Description: "32 oz. Bottle, Interior and Exterior Car Cleaner, Optimum Rinseless Car Wash System, Pro Car Care Products", Price: 3000, Stock: 120})
		db.Create(&model.Product{CategoryID: 3, Name: "Flow-Through Wash Brush with Adjustable Handle", Description: "Camco 43633 RV", Price: 3500, Stock: 110})
		db.Create(&model.Product{CategoryID: 3, Name: "Lucas LUC10160 Oil 10160", Description: "Slick Mist Speed Wax - 24 Oz.", Price: 1200, Stock: 70})
		db.Create(&model.Product{CategoryID: 3, Name: "Marine UV Protectant for Vinyl, Plastic, Rubber, Fiberglass, Leather And More", Description: "Dust and Dirt Repellant - Non-Toxic, Matte Finish - 10 fl. oz", Price: 3500, Stock: 76})
		db.Create(&model.Product{CategoryID: 3, Name: "Chemical Guys ACC151 Secondary Container Dilution Bottle with Heavy Duty Sprayer", Description: "16 oz., 1 Pack", Price: 900, Stock: 120})
		db.Create(&model.Product{CategoryID: 4, Name: "Little Tikes Rocking Horse", Description: "Magenta", Price: 4500, Stock: 34})
		db.Create(&model.Product{CategoryID: 4, Name: "Baby Delight Alpine Deluxe Portable Bouncer", Description: "Charcoal Tweed , 28x18x21 Inch (Pack of 1)", Price: 10000, Stock: 15})
		db.Create(&model.Product{CategoryID: 4, Name: "Fisher-Price Baby Bouncer", Description: "Geo Meadow, Infant Soothing and Play Seat, Multi", Price: 4000, Stock: 24})
		db.Create(&model.Product{CategoryID: 4, Name: "Fisher-Price Animal Activity Jumperoo", Description: "Blue, One size", Price: 11000, Stock: 23})
		db.Create(&model.Product{CategoryID: 4, Name: "VTech Sit-To-Stand Learning Walker", Description: "(Frustration Free Packaging), Blue", Price: 6000, Stock: 30})
		db.Create(&model.Product{CategoryID: 4, Name: "Baby Einstein 4-in-1 Kickin'", Description: "Tunes Music and Language Play Gym and Piano Tummy Time Activity Mat", Price: 5500, Stock: 40})
		db.Create(&model.Product{CategoryID: 4, Name: "Fisher-Price Deluxe Kick 'n Play", Description: "Piano Gym, Green, Gender Neutral (Frustration Free Packaging)", Price: 4500, Stock: 17})
		db.Create(&model.Product{CategoryID: 4, Name: "Fisher-Price Sit-Me-Up", Description: "Floor Seat Pacific Pebble, portable baby chair with toys", Price: 5000, Stock: 26})
		db.Create(&model.Product{CategoryID: 4, Name: "Graco Affix Highback Booster", Description: "Seat with Latch System, Atomic", Price: 13000, Stock: 52})
		db.Create(&model.Product{CategoryID: 4, Name: "Graco Tranzitions 3 in 1", Description: " Harness Booster Seat, Proof", Price: 15000, Stock: 20})
		db.Create(&model.Product{CategoryID: 4, Name: "Graco® Pack ‘n Play®", Description: "On The Go™ Playard, Kaden", Price: 11000, Stock: 35})
		db.Create(&model.Product{CategoryID: 4, Name: "Soothie Pacifier", Description: "Philips AVENT 0-3 Months, Green, 4 Pack, SCF190/41", Price: 1100, Stock: 26})
		db.Create(&model.Product{CategoryID: 4, Name: "Comotomo Baby Bottle", Description: "Green, 8 Ounce (2 Count)", Price: 3000, Stock: 13})
		db.Create(&model.Product{CategoryID: 4, Name: "Water Teethers", Description: "Infantino 3-Pack", Price: 1000, Stock: 30})
		db.Create(&model.Product{CategoryID: 5, Name: "Essence | Lash Princess False Lash Waterproof Mascara", Description: "Cruelty Free (Pack of 1)", Price: 2500, Stock: 100})
		db.Create(&model.Product{CategoryID: 5, Name: "Essence | Lash Princess False Lash Effect Mascara", Description: "Gluten & Cruelty Free", Price: 1300, Stock: 110})
		db.Create(&model.Product{CategoryID: 5, Name: "Milani Make It Last 3-in-1", Description: "Setting Spray and Primer- Prime + Correct + Set (2.03 Fl. Oz.) Makeup Finishing Spray and Primer", Price: 2500, Stock: 80})
		db.Create(&model.Product{CategoryID: 5, Name: "Grande Cosmetics GrandeLASH-MD Lash Enhancing Serum", Description: "Promotes Appearance of Longer, Thicker Eyelashes, Cruelty Free", Price: 4000, Stock: 57})
		db.Create(&model.Product{CategoryID: 5, Name: "Poreless Putty Primer", Description: "Silky, Skin-Perfecting, Lightweight, Long Lasting, Smooths, Hydrates, Minimizes Pores", Price: 1200, Stock: 32})
		db.Create(&model.Product{CategoryID: 5, Name: "CeraVe AM Facial Moisturizing Lotion SPF 30", Description: "Oil-Free Face Moisturizer with Sunscreen | Non-Comedogenic | 3 Ounce", Price: 1600, Stock: 50})
		db.Create(&model.Product{CategoryID: 5, Name: "Dove Body Wash with Pump with Skin Natural Nourishers", Description: "for Instantly Soft Skin and Lasting Nourishment Deep Moisture Cleanser That Effectively Washes", Price: 1300, Stock: 65})
		db.Create(&model.Product{CategoryID: 5, Name: "BLUE LIZARD Sensitive Mineral Sunscreen with Zinc Oxide", Description: "SPF 50+, Water Resistant, UVA/UVB Protection with Smart Bottle Technology - Fragrance Free, 5 oz", Price: 2000, Stock: 30})
		db.Create(&model.Product{CategoryID: 5, Name: "CeraVe Hydrating Facial Cleanser", Description: "Moisturizing Non-Foaming Face Wash with Hyaluronic Acid, Ceramides and Glycerin | 16 Fluid Ounce", Price: 2300, Stock: 28})
		db.Create(&model.Product{CategoryID: 5, Name: "Nizoral Anti-Dandruff Shampoo", Description: "Basic, Fresh, 7 Fl Oz", Price: 1900, Stock: 150})
		db.Create(&model.Product{CategoryID: 5, Name: "Color Wow Dream Coat Supernatural Spray", Description: "Multi-award-winning anti-frizz spray keeps hair frizz-free for days", Price: 3800, Stock: 45})
		db.Create(&model.Product{CategoryID: 5, Name: "Reed Diffuser", Description: "NEST Fragrances Rose Noir & Oud", Price: 2600, Stock: 31})
		db.Create(&model.Product{CategoryID: 5, Name: "Long Lasting And Addictive Personal Perfume Oil Fragrance", Description: "SWISS ARABIAN Noora - Luxury Products From Dubai", Price: 3100, Stock: 50})
		db.Create(&model.Product{CategoryID: 5, Name: "Hybrid Electric Trimmer and Shaver", Description: "Philips Norelco OneBlade, Frustration Free Packaging, QP2520/90", Price: 5200, Stock: 54})
		db.Create(&model.Product{CategoryID: 5, Name: "Gillette Venus Sensitive Disposable Razors for Women with Sensitive Skin", Description: "3 Count, Delivers Close Shave with Comfort", Price: 1100, Stock: 130})
		db.Create(&model.Product{CategoryID: 5, Name: "Philips Norelco Shaver 2300", Description: "Rechargeable Electric Shaver with PopUp Trimmer, Black, 1 Count, S1211/81", Price: 5300, Stock: 95})
		db.Create(&model.Product{CategoryID: 5, Name: "Finishing Touch Flawless Women's Painless", Description: "Hair Remover, Blush/Rose Gold", Price: 2600, Stock: 68})
		db.Create(&model.Product{CategoryID: 5, Name: "Veet Sensitive Hair Remover Gel Cream Pink", Description: "13.5 Fl Oz", Price: 1400, Stock: 60})
		db.Create(&model.Product{CategoryID: 5, Name: "Panasonic Nose Hair Trimmer and Ear Hair Trimmer ER430K", Description: "Vacuum Cleaning System , Men's, Wet/Dry, Battery-Operated", Price: 3300, Stock: 40})
		db.Create(&model.Product{CategoryID: 5, Name: "Gillette Venus ComfortGlide Womens Razor Blade Refills", Description: "6 Count, White Tea Scented Gel Bar Protects Againist Skin Irritation", Price: 3600, Stock: 50})
		db.Create(&model.Product{CategoryID: 5, Name: "Reed Diffuser", Description: "NEST Fragrances Rose Noir & Oud", Price: 2600, Stock: 31})
		db.Create(&model.Product{CategoryID: 1, Name: "Optoma GT1090HDR Short Throw Laser Home Theater Projector", Description: "4K HDR Input | Lamp-Free Reliable Operation 30,000 hours | Bright 4,200 lumens for Day and Night | Short Throw", Price: 150000, Stock: 20})
		db.Create(&model.Product{CategoryID: 1, Name: "Optoma CinemaX P2 White Smart 4K UHD Laser Projector for Home Theater", Description: "3000 Lumens Superior Image with Laser & 6-Segment Color Wheel | Ultra-Short Throw | Built-in Soundbar | Works w/ Alexa & Google", Price: 210000, Stock: 20})
		db.Create(&model.Product{CategoryID: 1, Name: "BenQ TH585P 1080p Home Entertainment Projector", Description: "3500 Lumens | High Contrast Ratio | Loud 10W Speaker | Low Input Lag for Gaming | Stream Netflix & Prime Video", Price: 70000, Stock: 25})
		db.Create(&model.Product{CategoryID: 1, Name: "EPOS Audio H3PRO Hybrid Wireless Closed Acoustic Gaming Headset", Description: "Ghost White", Price: 11000, Stock: 58})
		db.Create(&model.Product{CategoryID: 1, Name: "Philips PH805 Active Noise Canceling (ANC) Over Ear Wireless Bluetooth Performance Headphones", Description: "w/Hi-Res Audio, Comfort Fit and 30 Hours of Playtime (TAPH805BK)", Price: 12500, Stock: 25})
		db.Create(&model.Product{CategoryID: 1, Name: "JBL CHARGE 5", Description: "Portable Bluetooth Speaker with IP67 Waterproof and USB Charge out - Black", Price: 12000, Stock: 60})
		db.Create(&model.Product{CategoryID: 1, Name: "JBL Clip 4", Description: "Portable Speaker with Bluetooth, Built-in Battery, Waterproof and Dustproof Feature - Black (JBLCLIP4BLKAM)", Price: 8000, Stock: 60})
		db.Create(&model.Product{CategoryID: 1, Name: "SAMSUNG Galaxy S22+ Smartphone", Description: "Android Cell Phone, 256GB, 8K Camera & Video, Brightest Display, Long Battery Life", Price: 180000, Stock: 15})
		db.Create(&model.Product{CategoryID: 6, Name: "Mr. Potato Head Disney/Pixar Toy Story 4", Description: "Classic Mr. Potato Head Figure Toy for Kids Ages 2 and Up", Price: 2000, Stock: 29})
		db.Create(&model.Product{CategoryID: 6, Name: "Transformers Toys Heroic Optimus Prime Action Figure", Description: "Timeless Large-Scale Figure, Changes into Toy Truck - Toys for Kids 6 and Up, 11-inch", Price: 3500, Stock: 20})
		db.Create(&model.Product{CategoryID: 6, Name: "Disney Pixar Toy Story 4 Woody and Bullseye 2-Character Pack", Description: "Movie-inspired Relative-Scale for Storytelling Play", Price: 3000, Stock: 40})
		db.Create(&model.Product{CategoryID: 6, Name: "Funko Pop! Marvel: WandaVision - The Scarlet Witch", Description: "Vinyl Collectible Figure", Price: 1400, Stock: 40})
		db.Create(&model.Product{CategoryID: 6, Name: "Jurassic World Dinosaur Snap Squad Collectibles for Display", Description: "Play and Snap On Feature for Attaching to Backpacks, Lunch Packs and More", Price: 5000, Stock: 35})
		db.Create(&model.Product{CategoryID: 6, Name: "Pop! Comic Cover", Description: "Marvel Venom Lethal Protector Previews Exclusive Vinyl Figure", Price: 4000, Stock: 34})
		db.Create(&model.Product{CategoryID: 6, Name: "Avengers Series Marvel Assemble Titan Hero Iron Man", Description: "12 Action Figure", Price: 3800, Stock: 15})
		db.Create(&model.Product{CategoryID: 6, Name: "Funko Pop! Marvel: Loki", Description: "Funko Pop! Marvel: Loki", Price: 1500, Stock: 26})
		db.Create(&model.Product{CategoryID: 6, Name: "Disney Pixar Jessie Large Action Figure 12 in", Description: "Highly Posable with Authentic Detail, Toy Story Movie Collectable Cowgirl, Ages 3 Years & Up", Price: 2000, Stock: 16})
		db.Create(&model.Product{CategoryID: 6, Name: "Transformers Playskool Heroes Rescue Bots Energize Bumblebee Figure", Description: "Official Licenced Product", Price: 3400, Stock: 35})
		db.Create(&model.Product{CategoryID: 6, Name: "Funko Pop! Marvel: Avengers Game - Iron Man ", Description: "Multicolor & Pop! Marvel: Avengers Game - Captain America (Stark Tech Suit)", Price: 1800, Stock: 36})
		db.Create(&model.Product{CategoryID: 6, Name: "Funko Pop! Marvel: Avengers Endgame - Captain America with Broken Shield", Description: "Multicolor,3.75 inches", Price: 1900, Stock: 12})
		db.Create(&model.Product{CategoryID: 6, Name: "LEGO Marvel Infinity Gauntlet 76191 Collectible Building Kit", Description: "LEGO Marvel Infinity Gauntlet 76191 Collectible Building Kit", Price: 13000, Stock: 10})
		db.Create(&model.Product{CategoryID: 6, Name: "LEGO Creator Mighty Dinosaurs 31058 Build It Yourself Dinosaur Set", Description: "Create a Pterodactyl, Triceratops and T Rex Toy (174 Pieces)", Price: 3400, Stock: 28})
		db.Create(&model.Product{CategoryID: 6, Name: "LEGO Star Wars Imperial TIE Fighter 75300 Building Kit", Description: "Awesome Construction Toy for Creative Kids, New 2021 (432 Pieces)", Price: 13000, Stock: 32})
		db.Create(&model.Product{CategoryID: 6, Name: "LEGO Friends Forest House 41679 Building Kit", Description: "Forest Toy with a Tree House; Great Gift for Kids Who Love Nature; New 2021 (326 Pieces)", Price: 2700, Stock: 12})
		db.Create(&model.Product{CategoryID: 7, Name: "SKLZ Pro Mini Basketball Hoop", Description: "Pro kit of Basketball", Price: 5600, Stock: 38})
		db.Create(&model.Product{CategoryID: 7, Name: "Franklin Sports Official Size Football", Description: "Professional Quality", Price: 5000, Stock: 35})
		db.Create(&model.Product{CategoryID: 7, Name: "Wilson NCAA Supreme Junior Football with Pump & Tee", Description: "Professional Quality", Price: 5600, Stock: 53})
		db.Create(&model.Product{CategoryID: 7, Name: "Professional Quality", Description: "RTP Pro Baseball Fielding Glove - Infield, Outfield Gloves", Price: 4000, Stock: 62})
		db.Create(&model.Product{CategoryID: 7, Name: "Franklin Sports MLB Jumbo Kids Plastic Baseball Bat", Description: "Backyard Baseball Bat with Large Barrel for Kids + Toddlers - Kids Fat Plastic Bat", Price: 2300, Stock: 52})
		db.Create(&model.Product{CategoryID: 7, Name: "Under Armour Womens Mini Athletic Headbands", Description: "6-Pack", Price: 1900, Stock: 32})
		db.Create(&model.Product{CategoryID: 7, Name: "Champion Sports Official Heavy Duty Rubber Cover Nylon Basketballs", Description: "Proffesional ball", Price: 4500, Stock: 10})
		db.Create(&model.Product{CategoryID: 7, Name: "Franklin Sports NHL Kids Mini Hockey Set", Description: "Includes 1 Knee Hockey Goal - 2 Mini Hockey Sticks + 2 Foam Balls - Indoor Toy Mini Hockey Goal + Sticks Set", Price: 6300, Stock: 13})
		db.Create(&model.Product{CategoryID: 7, Name: "Gatorade Stainless Steel Sport Bottle", Description: "26oz, Double-Wall Insulation", Price: 2200, Stock: 30})
		db.Create(&model.Product{CategoryID: 7, Name: "Franklin Sports Kids Baseball Pitching Machine", Description: "Pop A Pitch Baseball Batting Machine with Youth Bat + 3 Plastic Baseballs - Boys + Girls Baseball Toy", Price: 3400, Stock: 42})
		db.Create(&model.Product{CategoryID: 7, Name: "GAMMA Sports Kids Training (Transition) Balls", Description: "Yellow/Green Dot, 78 Green Dot, 12-Pack", Price: 4300, Stock: 21})
		db.Create(&model.Product{CategoryID: 7, Name: "Franklin Sports Mini Sponge Foam Football ", Description: "Grip-Tech Youth Football with Sift and Tacky, Easy Grip Cover - Perfect for Small Kids", Price: 3100, Stock: 24})
	}

	log.Info("Data inserted")
}