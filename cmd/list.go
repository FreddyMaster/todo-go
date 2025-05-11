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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Open mock db file
		file, err := os.Open("todo.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		if len(records) == 0 {
			fmt.Println("No items found")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
		_, err = fmt.Fprintln(w, "ID\tTASK\tTIME\tDONE")
		if err != nil {
			log.Fatal(err)
		}

		for i, row := range records {
			// skip header
			if i == 0 {
				continue
			}

			age := humanizeTime(row[2])
			_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", row[0], row[1], age, row[3])
			if err != nil {
				log.Fatal(err)
			}
		}

		err = w.Flush()
		if err != nil {
			log.Fatal(err)
		}

	},
}

// humanizeTime converts a RFC3339 time string to "x minutes ago" etc.
func humanizeTime(timeStr string) string {
	t, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	duration := time.Since(t)
	switch {
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	default:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
