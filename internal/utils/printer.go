package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shufo/gh-issue-stats/pkg/types"
)

var header = []string{"Label", "Open", "Closed", "Total", "Open %", "Average Time to close (days)", "Median Time to close (days)"}

func PrintStatistics(stats types.Statistics) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)

	// Configure table style
	t.Style().Format.Header = text.FormatTitle
	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateHeader = true
	t.Style().Options.SeparateRows = false

	// Set header
	t.AppendHeader(table.Row{header})

	// Add label statistics rows
	for _, stat := range stats.LabelStats {
		t.AppendRow(table.Row{
			stat.Name,
			stat.Open,
			stat.Closed,
			stat.Total,
			fmt.Sprintf("%.2f%%", stat.OpenPercentage),
			fmt.Sprintf("%.0f", math.Round((stat.AvgCloseTime/(24*60*60))*100/100)),
			fmt.Sprintf("%.0f", stat.MedianCloseTime),
		})
	}

	// Add separator and total row
	t.AppendSeparator()
	t.AppendRow(table.Row{
		"Total",
		stats.OverallStats.Open,
		stats.OverallStats.Closed,
		stats.OverallStats.Total,
		fmt.Sprintf("%.2f%%", stats.OverallStats.OpenPercentage),
		fmt.Sprintf("%.0f", math.Round((stats.OverallStats.AvgCloseTime/(24*60*60))*100/100)),
		fmt.Sprintf("%.0f", stats.OverallStats.MedianCloseTime),
	})

	// Render the table
	t.Render()
}

func SaveToFile(data interface{}, filename string) error {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Saving to %s...", filename)
	if !debug {
		s.Start()
	}

	StartSpinner(fmt.Sprintf(" Saving to %s...", filename))

	file, err := os.Create(filename)
	if err != nil {
		if !debug {
			StopSpinner()
		}
		return fmt.Errorf("failed to create file %s: %v", filename, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		if !debug {
			s.Stop()
		}
		return fmt.Errorf("failed to write to file %s: %v", filename, err)
	}

	if !debug {
		StopSpinner()
	}

	DebugPrintf("Data saved to %s", filename)

	return nil
}

func WriteDelimitedOutput(stats types.Statistics, delimiter rune) error {
	writer := csv.NewWriter(os.Stdout)
	writer.Comma = delimiter

	// Write header
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header: %v", err)
	}

	// Write label statistics
	for _, stat := range stats.LabelStats {
		row := []string{
			stat.Name,
			strconv.Itoa(stat.Open),
			strconv.Itoa(stat.Closed),
			strconv.Itoa(stat.Total),
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing row: %v", err)
		}
	}

	// Write total row
	totalRow := []string{
		"Total",
		strconv.Itoa(stats.OverallStats.Open),
		strconv.Itoa(stats.OverallStats.Closed),
		strconv.Itoa(stats.OverallStats.Total),
	}
	if err := writer.Write(totalRow); err != nil {
		return fmt.Errorf("error writing total row: %v", err)
	}

	writer.Flush()
	return writer.Error()
}
