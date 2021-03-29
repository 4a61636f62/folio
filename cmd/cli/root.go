package cli

import (
	"folio/pkg/models"
	"folio/pkg/models/folioFile"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"path/filepath"
)

var homeDir string
var allCoins map[string]models.Coin

var RootCmd = &cobra.Command{
	Use: "folio",
	Short: "Folio is a CLI cryptocurrency portfolio tracker",
}

func init() {
	var err error
	homeDir, _  = homedir.Dir()
	allCoins, err = folioFile.Coins(filepath.Join(homeDir, ".coinlist")).All()
	if err != nil {
		panic(err)
	}
}

