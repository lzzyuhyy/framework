package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var rootCmd = &cobra.Command{
	Use:   "gen-proto",
	Short: "a short cmd to generate pb code",
	Long:  `gen-proto is a short cmd to generate pb code, just need pb file's path'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("flag used error, [./***.exe] -h to get help")
			return
		}
		fmt.Println("generate pb by", args[0])
		GenerateProtoFile(args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GenerateProtoFile(filepath string) {
	if filepath == "" {
		fmt.Println("don't set proto file path")
		os.Exit(1)
	}

	cmd := exec.Command("protoc",
		"--go_out=.",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=.",
		"--go-grpc_opt=paths=source_relative",
		filepath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running protoc: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Protobuf code generated successfully.")
}
