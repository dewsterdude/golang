package cmdmanager

// THIS HAS SAME METHODS AS FILEMANAGER, BUT INTERACTS VIA CONSOLE INSTEAD OF FILES

import "fmt"

type CMDManager struct { // stay empty just to attach some methods

}

func (cmd CMDManager) ReadLines() ([]string, error) {

	var prices []string
	var price string

	fmt.Println("Please enter your prices. Confirm every price with Enter OR type 0 to finish")

	for {
		fmt.Print("Price: ")
		fmt.Scan(&price)
		prices = append(prices, price)
		if price == "0" {
			break
		}
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{} // return command manager struct
}
