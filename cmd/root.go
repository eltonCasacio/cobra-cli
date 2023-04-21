/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"os"

	"github.com/eltonCasacio/cobra-cli/internal/database"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-cli",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
}

type RunEFunc func(cmd *cobra.Command, args []string) error

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "cobra.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetCategoryDB(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
