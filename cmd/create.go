/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/eltonCasacio/cobra-cli/internal/database"
	"github.com/spf13/cobra"
)

func NewCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines`,
		RunE:  runCreate(categoryDb),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := NewCreateCmd(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(createCmd)
	categoryCmd.PersistentFlags().StringP("name", "n", "", "passe -n ou --name mais nome da categoria")
	categoryCmd.PersistentFlags().StringP("description", "d", "", "passe -d ou --description mais descricao da categoria")
	categoryCmd.MarkFlagsRequiredTogether("name", "description")
}
