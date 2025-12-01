package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct { // adding tags for json labels
	TaxRate           float64             `json:"tax_rate"`     // upper case to access outside package
	InputPrices       []float64           `input_prices`        // upper case to access outside package
	TaxIncludedPrices map[string]string   `tax_included_prices` // upper case to access outside package
	IOManager         iomanager.IOManager `json:"-"`            // "-" means to exclude
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) { // removed err
	err := job.LoadData()
	if err != nil {
		// return err
		errorChan <- err // send error to channel
		return
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	// filemanager.WriteResult(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)

	//	return job.IOManager.WriteResult(job) // not needed with GO routines
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job TaxIncludedPriceJob) LoadData() error {
	// ASSIGN THE FILENAME
	/*	file, err := os.Open("prices.txt")
		if err != nil {
			fmt.Println("Error opening prices file")
			fmt.Println()
			return
		}

		// READ IN THE LINES FROM THE FILE
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() { // read line, one at at time
			lines = append(lines, scanner.Text())
		}
		// CHECK FOR ERRORS
		err = scanner.Err()
		if err != nil {
			fmt.Println("Error Reading the File content failed")
			fmt.Println(err)
			file.Close()
			return
		}
		//	file.Close()
	*/
	// load data method
	// lines, err := filemanager.ReadLines("prices.txt")

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	// convert strings to floats
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil
}

/*
	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Error converting line %d to float64 \n", lineIndex+1)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
	}
	job.InputPrices = prices
*/
//
//		fmt.Sscanf(scanner.Text(), "%f", &price)
//		job.InputPrices = append(job.InputPrices, price)
