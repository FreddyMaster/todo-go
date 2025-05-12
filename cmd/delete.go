/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// deleteCmd represents the 'delete' command for the CLI application.
// This command removes a task from the todo.csv file by its ID.
// If the specified ID is not found, it notifies the user.
var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete a task from your todo list by its ID.",
	Long: `Delete a task from your todo list by specifying its unique ID.

This command removes the task with the given ID from the todo.csv file.
For example:
    todo delete 12345
This will remove the task with ID 12345 from your todo list.`,
	Args: cobra.ExactArgs(1), // Ensures exactly one argument (the ID) is provided
	Run: func(cmd *cobra.Command, args []string) {
		// Open the CSV file containing the todo list.
		file, err := os.OpenFile("todo.csv", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// Ensure the file is closed after reading.
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		// Read all records from the CSV file.
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		// Get the target ID to remove from the command arguments.
		removeID := args[0]
		targetIdFound := false

		// Prepare a new slice for the filtered records (excluding the removed task).
		var filtered [][]string
		for _, row := range records {
			if row[0] == removeID {
				targetIdFound = true
				continue // Skip this row (i.e., delete it)
			}
			filtered = append(filtered, row)
		}

		// Inform the user if the specified ID was not found.
		if !targetIdFound {
			fmt.Println("ID not found")
			return
		}

		// Overwrite the CSV file with the updated records.
		fOut, err := os.Create("todo.csv")
		if err != nil {
			log.Fatal(err)
		}
		// Ensure the output file is closed after writing.
		defer func(fOut *os.File) {
			err := fOut.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(fOut)

		// Write the filtered records back to the CSV file.
		writer := csv.NewWriter(fOut)
		err = writer.WriteAll(filtered)
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	},
}

// init registers the deleteCmd with the root command.
func init() {
	rootCmd.AddCommand(deleteCmd)
}
