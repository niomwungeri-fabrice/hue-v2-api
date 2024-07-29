package main

import (
	"fmt"
	"github.com/niomwungeri-fabrice/hue-v2-api/cmd"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	rootCmd := &cobra.Command{
		Use:   "hue-v2-api",
		Short: "Hue v2 API CLI",
		Long:  "CLI tool to interact with the Philips Hue v2 API.",
	}
	rootCmd.AddCommand(cmd.DevicesCmd)
	//rootCmd.AddCommand(cmd.InitLight())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
