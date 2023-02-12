package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Offer struct {
	OfferID   string `json:"offer_id"`
	MarketSku int    `json:"market_sku"`
	Price     int    `json:"price"`
}

type Fid struct {
	Offers []Offer `json:"offers"`
}

func main() {
	var oneFid, allFids Fid
	var first = true

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		jsonStr := scanner.Text()
		if first {
			first = false
			continue
		}

		err := json.Unmarshal([]byte(jsonStr), &oneFid)
		if err != nil {
			// обработка ошибки
			os.Exit(1)
		}

		allFids.Offers = append(allFids.Offers, oneFid.Offers...)
	}

	sort.Slice(allFids.Offers, func(i, j int) bool {
		if allFids.Offers[i].Price < allFids.Offers[j].Price {
			return true
		} else if allFids.Offers[i].Price == allFids.Offers[j].Price {
			if allFids.Offers[i].OfferID < allFids.Offers[j].OfferID {
				return true
			}
		}
		return false
	})

	result, err := json.Marshal(allFids)
	if err != nil {
		// обработка ошибки
		os.Exit(1)
	}

	fmt.Println(string(result))
}
