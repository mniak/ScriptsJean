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

func main() {
	dirs := DirNamesForEntireMonth(time.Now())

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
