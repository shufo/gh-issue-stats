package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/shufo/gh-issue-stats/internal/github"
	"github.com/shufo/gh-issue-stats/internal/stats"
	"github.com/shufo/gh-issue-stats/internal/utils"
	"github.com/spf13/cobra"
)

var (
	outputFile string
	statsFile  string
	format     string
	debug      bool
)

func Exec() {
	rootCmd := &cobra.Command{
		Use:   "gh issue-stats [repository]",
		Short: "Generate GitHub issue statistics",
		Long: `A GitHub CLI extension to analyze repository issues and generate statistics.
Provides detailed information about issues grouped by labels and overall statistics.

Examples:
  # Current repository
  gh issue-stats

  # Specific repository
  gh issue-stats owner/repo

  # With output format
  gh issue-stats owner/repo --format json`,
		Args:          cobra.MaximumNArgs(1),
		RunE:          runCommand,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file for raw issues data (optional)")
	rootCmd.Flags().StringVarP(&statsFile, "stats", "s", "", "Output file for statistics data (optional)")
	rootCmd.Flags().StringVarP(&format, "format", "f", "", "Output format: table (default), json, csv, or tsv")
	rootCmd.Flags().BoolVarP(&debug, "debug", "v", false, "Enable verbose debug output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runCommand(cmd *cobra.Command, args []string) error {
	utils.SetupLogger(debug)
	utils.SetDebug(debug)
	github.SetDebug(debug)

	var repository string
	if len(args) > 0 {
		repository = args[0]
		// Validate repository format
		if !isValidRepositoryFormat(repository) {
			return fmt.Errorf("invalid repository format. Expected format: owner/repo")
		}
	}
	// Fetch issues
	issues, err := github.FetchIssues(repository)
	if err != nil {
		return err
	}

	// Save issues if output file is specified
	if outputFile != "" {
		if err := utils.SaveToFile(issues, outputFile); err != nil {
			return err
		}
	}

	// Calculate statistics
	stats := stats.CalculateStatistics(issues)

	// Save statistics if stats file is specified
	if statsFile != "" {
		if err := utils.SaveToFile(stats, statsFile); err != nil {
			return err
		}
	}

	// Output based on format
	switch strings.ToLower(format) {
	case "json":
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(stats)
	case "csv":
		return utils.WriteDelimitedOutput(stats, ',')
	case "tsv":
		return utils.WriteDelimitedOutput(stats, '\t')
	default:
		utils.PrintStatistics(stats)
	}

	return nil
}

// isValidRepositoryFormat validates the repository argument format
func isValidRepositoryFormat(repo string) bool {
	parts := strings.Split(repo, "/")
	return len(parts) == 2 && parts[0] != "" && parts[1] != ""
}
