package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/multierr"
)

func getWeekDayAbbreviation(weekday time.Weekday) string {
	weekdayNames := map[time.Weekday]string{
		time.Sunday:    "Sonntag",
		time.Monday:    "Montag",
		time.Tuesday:   "Dienstag",
		time.Wednesday: "Mittwoch",
		time.Thursday:  "Donnerstag",
		time.Friday:    "Freitag",
		time.Saturday:  "Samstag",
	}

	result := weekdayNames[weekday]
	result, _ = strings.CutSuffix(result, "tag")
	return result
}

func DirNamesForSingleDate(date time.Time) []string {
	dateStr := date.Format("06 01 02")
	root := fmt.Sprintf("%s %s. Fotos", dateStr, getWeekDayAbbreviation(date.Weekday()))
	return []string{
		path.Join(root, fmt.Sprintf("%s Pingos", dateStr)),
		path.Join(root, fmt.Sprintf("%s Flash Navega", dateStr)),

		path.Join(root, fmt.Sprintf("%s Ecobie", dateStr), "Ecobie 1"),
		path.Join(root, fmt.Sprintf("%s Ecobie", dateStr), "Ecobie 2"),
		path.Join(root, fmt.Sprintf("%s Ecobie", dateStr), "Ecobie 3"),
		path.Join(root, fmt.Sprintf("%s Ecobie", dateStr), "Ecobie 4"),
	}
}

func DirNamesForEntireMonth(date time.Time) []string {
	firstOfThisMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local)
	var result []string
	for d := firstOfThisMonth; d.Month() == date.Month(); d = d.AddDate(0, 0, 1) {
		result = append(result, DirNamesForSingleDate(d)...)
	}
	return result
}

// The function "main" is always the start of a Go program
func main() {
	// Get the current time and store in the variable "now"
	now := time.Now()

	// Call the function DirNamesForEntireMonth passing a date which is on the
	// variable "now"
	//
	// This will return a list of all the directories that should be created.
	dirs := DirNamesForEntireMonth(now)

	// Declare a variable named "errs" of type "error" without setting a value in to
	// it (in other words, without initializing it). The default value when a
	// variable is not initialized is "nil". In Go, all types have a "zero" values.
	// For the numeric types it is really "0". For the type string, the "zero" value
	// is "" (empty string). For other types it is usually "nil", which is kind of a
	// "null", which is kind of a "non-value", in the same way as ∞ (infinity) is
	// neither "zero" nor any other number.
	//
	// This variable is declared to accumulate all the errors that could occur while
	// creating the directories.
	var errs error

	for _, dir := range dirs {
		fmt.Printf("Creating directory '%s'\n", dir)
		err := os.MkdirAll(dir, 0o755)
		multierr.AppendInto(&errs, err)
		if errs != nil {
			fmt.Fprintf(os.Stderr, "Failed with error: %s\n", err)
			return
		}
	}
	fmt.Println("Directories created successfully")
}
