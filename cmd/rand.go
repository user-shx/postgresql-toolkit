/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"postgresql-toolkit/pkg/crypto"

	"github.com/spf13/cobra"
)

// randCmd represents the rand command
func newRandCmd() *cobra.Command {
	var length uint64

	cmd := &cobra.Command{
		Use:   "rand",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			password, err := crypto.Password(int(length))
			if err != nil {
				fmt.Errorf("%v\n", err)
			}
			fmt.Print(password)

		},
	}

	cmd.Flags().Uint64VarP(&length, "length", "l", 12, "Password length")

	return cmd
}
