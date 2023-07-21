package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "screeps2go",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := GenerateSimpleType(Creep{}, `
			func (c Creep) Store() Store {
				return Store{CreepName: c.Name}
			}`); err != nil {
			return err
		}

		if err := GenerateSimpleType(StructureSpawn{}, ""); err != nil {
			return err
		}

		if err := GenerateSimpleType(Room{}, ""); err != nil {
			return err
		}

		if err := GenerateSimpleType(Store{}, ""); err != nil {
			return err
		}

		return nil
	},
}

func main() {
	rootCmd.Execute()
}
