package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func addProduct(products []Product, product Product) []Product {
	return append(products, product)
}

func deleteProduct(products []Product, id int) ([]Product, error) {
	result := []Product{}
	founditem := false
	for _, p := range products {
		if id != p.ID {
			result = append(result, p)
		} else {
			founditem = true
		}
	}
	if founditem {
		return result, nil
	} else {
		return nil, errors.New("товар не найден")
	}
}

func findProduct(products []Product, id int) (Product, error) {
	for _, p := range products {
		if id == p.ID {
			return p, nil
		}
	}
	return Product{}, errors.New("товар не найден")
}

func changeProduct(products []Product, quantity int, name string) error {
	for i := range products {
		if name == products[i].Name {
			products[i].Quantity = quantity
			return nil
		}
	}
	return errors.New("товар не найден")
}

func calcAllProduct(products []Product) float64 {
	var result float64
	for _, p := range products {
		result += float64(p.Price) * float64(p.Quantity)
	}
	return result
}

func printAllProduct(products []Product) []Product {
	return products
}

func findMostExpenciveProduct(products []Product) (string, float64) {
	result := products[0].Price
	var strres string
	for i := 1; i < len(products); i++ {
		if products[i].Price > result {
			result = products[i].Price
			strres = products[i].Name
		}
	}
	return strres, float64(result)
}

func allProduct(products []Product) []Product {
	return products
}

func main() {
	products := []Product{}
	for {

		fmt.Println("| ============== СКЛАД ============== |")
		fmt.Println("| 1. Добавить товар                   |")
		fmt.Println("| 2. Удалить товар                    |")
		fmt.Println("| 3. Найти товар                      |")
		fmt.Println("| 4. Изменить количество              |")
		fmt.Println("| 5. Показать стоимость всех товаров  |")
		fmt.Println("| 6. Вывести все товары               |")
		fmt.Println("| 7. Самый дорогой товар              |")
		fmt.Println("| 8. Список всех товаров              |")
		fmt.Println("| 0. Выход                            |")

		var userans int
		var name string
		var price float64
		var quantity int
		fmt.Scan(&userans)
		if userans == 1 {
			fmt.Println("Введите имя: ")
			fmt.Scan(&name)
			fmt.Println("Введите цену: ")
			fmt.Scan(&price)
			fmt.Println("Введите количество товара: ")
			fmt.Scan(&quantity)
			id := 1
			for _, p := range products {
				if p.ID >= id {
					id = p.ID + 1
				}
			}
			product := Product{ID: id, Name: name, Price: price, Quantity: quantity}
			products = addProduct(products, product)
			fmt.Printf("Товару присвоен id: %v\n", id)
		}

		if userans == 2 {
			var idToDelete int
			var err error
			fmt.Println("Введите ID товара, который хотите удалить: ")
			fmt.Scan(&idToDelete)
			products, err = deleteProduct(products, idToDelete)
			if err != nil {
				fmt.Println("Ошибка: ", err)
			} else {
				fmt.Println("Товар удален", products)
			}

		}
		if userans == 3 {
			var idToFind int
			fmt.Println("Введите ID товара, который хотите найти: ")
			fmt.Scan(&idToFind)
			product, err := findProduct(products, idToFind)
			if err != nil {
				fmt.Println("Ошибка: ", err)
			} else {
				fmt.Println(product)
			}
		}
		if userans == 4 {
			var nameToFind string
			var newQ int
			fmt.Println("Введите наименование товара, количество которого хотите изменить: ")
			fmt.Scan(&nameToFind)
			fmt.Println("Введите новое количество товара: ")
			fmt.Scan(&newQ)
			err := changeProduct(products, newQ, nameToFind)
			if err != nil {
				fmt.Println("Ошибка: ", err)
			} else {
				fmt.Println("Количество товара изменено")
				fmt.Printf("Текущее количество товара: %v\n", newQ)
			}
		}
		if userans == 5 {
			fmt.Printf("Общая стоимость товаров: %v\n", calcAllProduct(products))
		}
		if userans == 6 {
			fmt.Println(printAllProduct(products))
		}
		if userans == 7 {
			name, price := findMostExpenciveProduct(products)
			fmt.Printf("Самый дорогой товар %s, его цена %g\n", name, price)
		}
		if userans == 8 {
			fmt.Printf("Список всех товаров: %v\n", allProduct(products))
		}
		if userans == 0 {
			fmt.Println("Вы вышли из программы")
			break
		}

	}
}
