package main

import (
	"fmt"

	id "github.com/llmariner/common/pkg/id"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func newGenCmd() *cobra.Command {
	var (
		prefix string
		n      int
	)

	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate a new ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := id.GenerateID(prefix, n)
			if err != nil {
				return fmt.Errorf("failed to generate ID: %s", err)
			}
			fmt.Println(id)

			return nil
		},
	}

	cmd.Flags().StringVar(&prefix, "prefix", "", "Prefix for the ID")
	cmd.Flags().IntVar(&n, "n", 1, "Number of random bytes to generate")
	if err := cmd.MarkFlagRequired("prefix"); err != nil {
		klog.Fatalf("Failed to mark flag required: %s", err)
	}
	if err := cmd.MarkFlagRequired("n"); err != nil {
		klog.Fatalf("Failed to mark flag required: %s", err)
	}

	return cmd
}

func newIDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "id",
		Short: "ID related commands",
	}

	cmd.SilenceUsage = true
	cmd.AddCommand(newGenCmd())

	return cmd
}
