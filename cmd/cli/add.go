package cli

import (
	"fmt"
	"folio/pkg/models"
	"folio/pkg/models/folioFile"
	"github.com/spf13/cobra"
	"path/filepath"
	"strconv"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a holding to your portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		// Parse amount
		amount, err := parseAmount(args[0])
		if err != nil {
			fmt.Println("Invalid amount")
		}

		// Parse coin from input string
		coinStr := args[1]
		coin, ok := allCoins[coinStr]
		if !ok {
			var coinId string
			coinMatches := match(allCoins, coinStr)
			switch len(coinMatches) {
			case 0:
				fmt.Println("Coin not found")
				return
			case 1:
				coinId = coinMatches[0]
			default:
				coinId, err = selectOption(coinMatches)
				if err != nil {
					fmt.Println("Invalid option")
					return
				}
			}
			coin = allCoins[coinId]
		}

		// Add to portfolio
		portfolio := folioFile.Portfolio(filepath.Join(homeDir, ".portfolio"))
		err = portfolio.Add(coin.Id, amount)
		if err != nil {
			fmt.Println("Error adding to portfolio")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

func parseAmount(s string) (float32, error) {
	amount64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	return float32(amount64), nil
}

func match(coins map[string]models.Coin, match string) []string {
	var ret []string
	for _, coin := range coins {
		if coin.Name == match || coin.Symbol == match {
			ret = append(ret, coin.Id)
		}
	}
	return ret
}

func selectOption(options []string) (string, error) {
	fmt.Println("Did you mean: ")
	for i, s := range options {
		fmt.Printf("\t%d. %s", i+1, s)
	}
	var in int
	_, err := fmt.Scanf("%d", &in)
	if err != nil || in > len(options) {
		return "", err
	}
	return options[in-1], nil
}

