/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"time"
)

type TodoList struct {
	id       int
	task     string
	time     string
	complete bool
}

// addCmd represents the 'add' command for the CLI application.
// This command adds a new task to the todo.csv file with a unique ID, the task description,
// the current timestamp, and a default 'Complete' status of false.
var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to your todo list.",
	Long: `Add a new task to your todo list by specifying a task description.

This command appends a new task to the todo.csv file, assigning it a unique process-based ID,
the current timestamp, and marking it as incomplete by default.

Example:
    todo add "Buy groceries"
This will add a new task with the description "Buy groceries".`,
	Args: cobra.ExactArgs(1), // Ensures exactly one argument (the task description) is provided
	Run: func(cmd *cobra.Command, args []string) {
		// Open the CSV file for appending, create it if it doesn't exist.
		file, err := os.OpenFile("todo.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// Ensure the file is closed after writing.
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		// Check if the file is new (empty); if so, write the header row.
		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if fileInfo.Size() == 0 {
			// File is new, write the CSV header.
			err = writer.Write([]string{"id", "task", "time", "Complete"})
			if err != nil {
				log.Fatal(err)
			}
		}

		// Prepare the new task row:
		// - id: uses a UUID for uniqueness.
		// - task: the task description from the user.
		// - time: current timestamp in a standard format.
		// - Complete: set to false by default.
		row := []string{
			uuid.New().String(),
			args[0],
			time.Now().Format(time.RFC822Z),
			strconv.FormatBool(false),
		}

		// Write the new task row to the CSV file.
		err = writer.Write(row)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// init registers the addCmd with the root command.
func init() {
	rootCmd.AddCommand(addCmd)
}
