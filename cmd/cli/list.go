package cli

import (
	"fmt"
	"folio/pkg/coingecko"
	"folio/pkg/models/folioFile"
	"github.com/spf13/cobra"
	"path/filepath"
	"sort"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all holdings in your portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		// Parse vsCurrencies
		var vsCurrencies []string
		if len(args) == 0 {
			vsCurrencies = []string{"usd"}
		} else {
			vsCurrencies = append(vsCurrencies, args...)
		}
		baseCurr := vsCurrencies[0]

		// Load PortfolioFile holdings
		portfolio := folioFile.Portfolio(filepath.Join(homeDir, ".portfolio"))
		holdings, err := portfolio.Holdings()
		if err != nil {
			fmt.Println("Error loading portfolio holdings")
		}

		// Fetch Prices
		ids := holdingIds(holdings)
		prices, err := coingecko.SimplePrice(ids, vsCurrencies)
		if err != nil {
			fmt.Println("Error fetching prices")
			return
		}

		// Sort holding ids based on value of holdings in descending order
		sort.Slice(ids, func(i int, j int) bool {
			return (holdings[ids[i]] * prices[ids[i]][baseCurr]) > (holdings[ids[j]] * prices[ids[j]][baseCurr])
		})

		// Display holdings and total
		totals := make(map[string]float32)
		for _, coinId := range ids {
			holdingAmount := holdings[coinId]
			fmt.Printf("%s : %f %s ", allCoins[coinId].Name, holdingAmount, allCoins[coinId].Symbol)
			for currency, price := range prices[coinId] {
				currencyAmount :=  price * holdingAmount
				totals[currency] += currencyAmount
				fmt.Printf("= %f %s ", currencyAmount, currency)
			}
			fmt.Println()
		}
		fmt.Println("=== TOTAL ===")
		fmt.Printf("%f %s ", totals[baseCurr], baseCurr)
		for _, currency := range vsCurrencies[1:] {
			fmt.Printf("= %f %s ", totals[currency], currency)
		}
		fmt.Println()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func holdingIds(m map[string]float32) []string {
	var ret []string
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}
