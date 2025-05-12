/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the 'complete' command for the CLI application.
// This command toggles the 'Complete' status of a task in the todo.csv file
// based on the provided task ID. If the task is marked as incomplete, it will
// be marked as complete, and vice versa.
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Toggle the completion status of a task by its ID.",
	Long: `Toggle the completion status of a task in your todo list.

You can use this command to mark a task as complete or incomplete by specifying its ID.
For example:
    todo complete 12345
This will flip the 'Complete' status of the task with ID 12345.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Open the CSV file containing the todo list.
		file, err := os.OpenFile("todo.csv", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// Ensure the file is properly closed after reading.
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

		// Get the target ID from the command arguments.
		completeID := args[0]
		targetIdFound := false

		// Prepare a new slice to hold the updated records.
		var filtered [][]string
		for _, row := range records {
			if row[0] == completeID {
				// Toggle the completion status (true <-> false) for the matching task.
				row[3] = strconv.FormatBool(!stringToBool(row[3]))
				targetIdFound = true
			}
			filtered = append(filtered, row)
		}

		// Inform the user if the specified ID was not found.
		if !targetIdFound {
			fmt.Println("ID not found")
		}

		// Overwrite the CSV file with the updated records.
		fOut, err := os.Create("todo.csv")
		if err != nil {
			log.Fatal(err)
		}
		// Ensure the output file is properly closed after writing.
		defer func(fOut *os.File) {
			err := fOut.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(fOut)

		// Write all updated records back to the CSV file.
		writer := csv.NewWriter(fOut)
		err = writer.WriteAll(filtered)
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	},
}

// init registers the completeCmd with the root command.
func init() {
	rootCmd.AddCommand(completeCmd)
}

// stringToBool converts a string to a boolean value.
// It logs a fatal error if the conversion fails.
func stringToBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
