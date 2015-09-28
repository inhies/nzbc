package main

import "math"

func outputXML() {
	/*
		output, err := xml.MarshalIndent(i, "  ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	*/
}
func testLocal() {
	// This section is for using the test data at the end of this file
	/*
		var i newznab.SearchResponse
		err := xml.Unmarshal(testdata, &i)
		if err != nil {
			log.Fatal(err)
		}
		pretty.Print(i.Channel)
		return

		for _, n := range i.Channel.NZBs {
			pretty.Print(n.Attributes)
		}
		_ = output
		//os.Stdout.Write(output)

		return
	*/
}

// A rounding function for float64
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(roundOn, val)
	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
