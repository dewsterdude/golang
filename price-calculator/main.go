package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

// prices that include tax
// some that have tax rates

func main() {
	// oldPrices := []float64{10, 20, 30} // hardcoded for the moment
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates)) // length of taxrates
	errorChans := make([]chan bool, len(taxRates))

	// result := make(map[float64][]float64) // prices for every tax rate

	for index, taxRate := range taxRates { // loop goes thru all tax rates
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index]) // start GO routine

		// or
		// cmdm := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)

		/*	err := priceJob.Process()

			if err != nil {
				fmt.Println("Could not process job")
				fmt.Println(err)
			}
		*/
	}
	// when only one channel will emit an error, try the select statement

	for index, _ taxRates := range taxRates {
		select {
		case err := <-errorChans[index]
			if err != nil 	{
				fmt.Println(err)
			}	
			}

	}


	for _, doneChan := range doneChans {
		<-doneChan // enure you go thru all channels before being done
	}
	// wont really workt this way
	for _, errorChan := range errorChans {
		<-errorChan // enure you go thru all channels before being done
	}

	//		prices[stringIndex] = floatPrice

	/* 	for _, taxRate := range taxRates { // loop goes thru all tax rates
	   		taxIncludedPrices := make([]float64, len(prices)) // len = empty slots
	   		for priceIndex, price := range prices {
	   			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
	   		}
	   		result[taxRate] = taxIncludedPrices
	   	}
	   	fmt.Println(result)
	*/
}
