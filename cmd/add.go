/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Open mock db file or create it if it doesn't exist yet
		file, err := os.OpenFile("todo.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		// Check if file is new (empty), if so, write header
		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if fileInfo.Size() == 0 {
			// File is new, write header
			err = writer.Write([]string{"id", "task", "time", "Complete"})
			if err != nil {
				log.Fatal(err)
			}
		}

		// Write the new task as a row
		row := []string{
			strconv.Itoa(os.Getpid()),
			args[0],
			time.Now().Format(time.DateTime),
			strconv.FormatBool(false),
		}

		err = writer.Write(row)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
