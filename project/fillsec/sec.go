package main
import(

	"log"
	"os"
//	"fmt"
	"encoding/csv"
)
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func createShoppingList(data [][]string) []ShoppingRecord {
	var shoppingList []ShoppingRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec ShoppingRecord
			for j, field := range line {
				if j == 0 {
					rec.ID = field
				} else if j == 1 {
					rec.Seccode = field
				}
			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}




