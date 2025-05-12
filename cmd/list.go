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
	"text/tabwriter"
	"time"
)

// listCmd represents the 'list' command for the CLI application.
// This command displays all tasks in the todo.csv file in a formatted table.
// It shows the ID, task description, how long ago the task was added, and completion status.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in your todo list.",
	Long: `Display all tasks in your todo list in a formatted table.

This command reads your todo.csv file and prints each task's ID, description,
how long ago it was added, and whether it is complete.

Example:
    todo list
This will print all tasks in a readable table format.`,
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

		// If there are no tasks, notify the user.
		if len(records) == 0 {
			fmt.Println("No items found")
			return
		}

		// Set up a tab writer for pretty table output.
		w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
		_, err = fmt.Fprintln(w, "ID\tTASK\tAGE\tDONE")
		if err != nil {
			log.Fatal(err)
		}

		// Print each task (skip the header row).
		for i, row := range records {
			if i == 0 {
				continue // skip header
			}
			age := humanizeTime(row[2])
			_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", row[0], row[1], age, row[3])
			if err != nil {
				log.Fatal(err)
			}
		}

		// Flush the writer to ensure all output is written.
		err = w.Flush()
		if err != nil {
			log.Fatal(err)
		}
	},
}

// humanizeTime converts a time string in time.DateTime format to a human-friendly string,
// such as "5 minutes ago", "2 hours ago", or "3 days ago".
func humanizeTime(timeStr string) string {
	t, err := time.Parse(time.RFC822Z, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(t)
	fmt.Println(t)
	fmt.Println(timeStr)
	switch {
	case duration < time.Minute:
		return fmt.Sprintf("%d seconds ago", int(duration.Seconds()))
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	default:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	}
}

// init registers the listCmd with the root command.
func init() {
	rootCmd.AddCommand(listCmd)
}
