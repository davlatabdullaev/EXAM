package main

import "fmt"

// Task 1
// Barcha customerlarni umumiy mablag'ini
// va qancha summa harid qilganini hisoblang va ko'rsating
func task_1(customers []Customer) {
	totalCash := 0
	for _, customer := range customers {
		totalCash += customer.Cash
	}

	fmt.Printf("Umumiy mablag': %d\n", totalCash)
	for _, customer := range customers {
		fmt.Printf("%s %s harid qilgan summa: %d\n", customer.FirstName, customer.LastName, customer.Basket.Total)
	}

}

// Task 2
// Eng ko'p pul sarflagan customerni aniqlang va ko'rsating
func task_2(customers []Customer) {

	l := len(customers)
	katta := 0
	var name1 string
	var name2 string
	for i := 0; i < l; i++ {
		if katta < customers[i].Basket.Total {
			katta = customers[i].Basket.Total
			if katta == customers[i].Basket.Total {
				name1 = customers[i].FirstName
				name2 = customers[i].LastName
			}
		}
	}
	fmt.Printf("Eng ko'p pul sarf etgan mijoz: %s %s(%dso'm)\n", name1, name2, katta)
}

// Task 3
// To'plamdagi eng qimmat productni aniqlang
// va ko'rsating

func task_3(customers []Customer) {
	katta := 0
	var name string
	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			if product.Price > katta {
				katta = product.Price
				name = product.Name
			}
		}
	}
	fmt.Printf("Eng qimmat mahsulot:%s(%dso'm)", name, katta)
}

// Task 4
// Barcha productlar uchun o'rtacha narxni hisoblang va ko'rsating
func task_4(customers []Customer) {
	sum := 0
	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			sum += product.Price
		}
	}
	fmt.Printf("O'rtacha narx: %d so'm", sum)
}

// Task 5
// Eng arzon savdo qilgan customerni aniqlang va ko'rsating
func task_5(customers []Customer) {
	kichik := 65446435
	var ism string
	var familiya string
	for _, customer := range customers {
		if customer.Basket.Total < kichik {
			kichik = customer.Basket.Total
			ism = customer.FirstName
			familiya = customer.LastName
		}
	}
	fmt.Printf("Eng arzon savdo qilgan mijoz: %s %s (%d)\n", ism, familiya, kichik)
}

// Task 6
// Eng ko'p sotilgan productlar categoriyasini aniqlang va chiqaring:

func task_6(customers []Customer) {

	type SotilganKategoriya struct {
		Category string
		Sales    int
	}

	sotilgankategoriya := make(map[string]int)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			sotilgankategoriya[product.Category] += product.Quantity
		}
	}

	var KopSotilganKategoriya SotilganKategoriya
	for category, sales := range sotilgankategoriya {
		if sales > KopSotilganKategoriya.Sales {
			KopSotilganKategoriya = SotilganKategoriya{Category: category, Sales: sales}
		}
	}
	fmt.Printf("Eng ko'p sotilgan kategoriya: %s\n", KopSotilganKategoriya.Category)

}

// Task 7
// Eng ko'p va eng kam sotilgan productlar nomini chiqarish

func task_7(customers []Customer) {

	type SotilganProduct struct {
		Name  string
		Sotuv int
	}

	sotilganproduct := make(map[string]int)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			sotilganproduct[product.Name] += product.Quantity
		}
	}

	var KopSotilganProduct SotilganProduct
	for product, sales := range sotilganproduct {
		if sales > KopSotilganProduct.Sotuv {
			KopSotilganProduct = SotilganProduct{Name: product, Sotuv: sales}
		}
	}
	fmt.Printf("Eng ko'p sotilgan mahsulot nomi: %s(%d ta)\n", KopSotilganProduct.Name, KopSotilganProduct.Sotuv)

	var KamSotilganProduct SotilganProduct
	for product, sales := range sotilganproduct {
		if sales < KopSotilganProduct.Sotuv {
			KamSotilganProduct = SotilganProduct{Name: product, Sotuv: sales}
		}
	}
	fmt.Printf("Eng kam sotilgan mahsulot nomi: %s(%d ta)\n", KamSotilganProduct.Name, KamSotilganProduct.Sotuv)

}

// Task 8
// Har bir savdo uchun o'rtacha mahsulot miqdorini hisoblang va ko'rsating
func task_8(customers []Customer) {
	mahsulotMiqdori:=0
	var	OrtaMiqdor float64 = 1.1
	sum:=0
	for _, customer := range customers{
		for _, product := range customer.Basket.Products{
           mahsulotMiqdori+=product.Quantity
		   sum++
		}
	}
	fmt.Println(sum)
	OrtaMiqdor  = float64(mahsulotMiqdori)/float64(sum) 
	fmt.Println("O'rtacha mahsulot miqdori: ", OrtaMiqdor)
}

// Task 9
// Eng ko'p mahsulot sotib olishgan mijozni aniqlang va ko'rsating:
func task_9(customers []Customer) {
	engKoP := 0
	var customerName string
	var customerLastName string
	for _, customer := range customers {
		totalProducts := 0
		for _, products := range customer.Basket.Products {
			totalProducts += products.Quantity
			if totalProducts > engKoP {
				engKoP = totalProducts
				customerName = customer.FirstName
				customerLastName = customer.LastName
			}
		}
	}
	fmt.Printf("Eng ko'p mahsulot sotib olgan mijoz: %s %s (%d ta)\n", customerName, customerLastName, engKoP)
}

// Task 10
// Eng ko'p savdolarda ko'rinadigan mahsulotni aniqlang va ko'rsating
func task_10(customers []Customer) {
	MahsulotSoni := make(map[string]int)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			MahsulotSoni[product.Name]++
		}
	}
	var KopkoringanMahsulotNomi string
	KopkoringanMahsulot := 0
	for productName, frequency := range MahsulotSoni {
		if frequency > KopkoringanMahsulot {
			KopkoringanMahsulot = frequency
			KopkoringanMahsulotNomi = productName
		}
	}

	if KopkoringanMahsulot > 0 {
		fmt.Printf("Eng ko'p savdolarda ko'rinadigan mahsulot: %s (%d ta)\n", KopkoringanMahsulotNomi, KopkoringanMahsulot)
	} else {
		fmt.Println("Mijozlar savatchalarida hech qanday mahsulot topilmadi.")
	}
}

// Task 11
// O'rtacha savdo mablag'i bo'yicha eng ko'p mablag'ga ega customerni
// aniqlang va ko'rsating
func task_11(customers []Customer) {
	maxOrtachaXarajat := 0
	var engKopXarajatMijozIsmi, engKopXarajatMijozFamiliyasi string

	for _, mijoz := range customers {
		jamiXarajat := 0
		for _, mahsulot := range mijoz.Basket.Products {
			jamiXarajat += mahsulot.Price * mahsulot.Quantity
		}

		ortachaXarajat := jamiXarajat / len(mijoz.Basket.Products)
		if ortachaXarajat > maxOrtachaXarajat {
			maxOrtachaXarajat = ortachaXarajat
			engKopXarajatMijozIsmi = mijoz.FirstName
			engKopXarajatMijozFamiliyasi = mijoz.LastName
		}
	}

	fmt.Printf("O'rtacha savdo mablag'i: %d so'm\n", maxOrtachaXarajat)
	fmt.Printf("Eng kop xarajatga ega bo'lgan mijoz: %s %s\n", engKopXarajatMijozIsmi, engKopXarajatMijozFamiliyasi)
}

// Task 12
// Eng ko'p umumiy daromad (quantity * price) olishgan toifani aniqlang va ko'rsating
func task_12(customers []Customer) {
	daromadMap := make(map[string]int)

	for _, customer := range customers {

		for _, product := range customer.Basket.Products {
			daromad := product.Quantity * product.Price
			daromadMap[product.Category] += daromad
		}
	}

	maxDaromad := 0
	var Kategoriyasi string
	for category, daromad := range daromadMap {
		if daromad > maxDaromad {
			maxDaromad = daromad
			Kategoriyasi = category
		}
	}

	fmt.Printf("Eng ko'p daromad olishgan toifa: %s (%d so'm)\n", Kategoriyasi, maxDaromad)
}

// Task 13
// Eng qimmat productni o'z ichiga olgan savdoni aniqlang va ko'rsating
func task_13(customers []Customer) {
	var topProduct string
	var customerName string
	var customerLastName string
	for _, customer := range customers{
		maxPrice := 0
	for _, product := range customer.Basket.Products {
		if product.Price > maxPrice {
			maxPrice = product.Price
			topProduct = product.Name
			customerName = customer.FirstName
			customerLastName=customer.LastName
		}
	}
			fmt.Printf("Eng qimmat mahsulot: %s (%s %s - %d) \n", topProduct,customerName, customerLastName, maxPrice)
	}

}
// Task 14
// Har bir mijoz uchun eng ko'p xarid qilgan toifani aniqlang va ko'rsating
func task_14(customers []Customer) {
    var topCustomerName, topCustomerLastName, topProductName string

    for _, customer := range customers {
        maxTotalSpending := 0
        totalSpending := 0
        var productCategory string
		totalPrice:=0
        for _, product := range customer.Basket.Products {
            spending := product.Quantity
            totalSpending += spending
            if spending > maxTotalSpending {
				totalPrice = product.Quantity*product.Price
                maxTotalSpending = spending
                topCustomerName = customer.FirstName
                topCustomerLastName = customer.LastName
                productCategory = product.Category
            }
        }
        topProductName = productCategory
		fmt.Printf("Eng ko'p xarid qilgan toifa: %s %s (%s - %d)\n", topCustomerName, topCustomerLastName, topProductName, totalPrice)
    }
}
// Task 15
// Har bir productdan savdo paytida umumiy nechta sotilganini aniqlang va ko'rsating
func task_15(customers []Customer) {
	sum:=0
for _, customer := range customers{
	for _, product := range customer.Basket.Products{
    sum+=product.Quantity
	}
}  
fmt.Printf("Mahsulotlar umumiy %d ta sotilgan ", sum) 
}

