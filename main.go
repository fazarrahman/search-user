/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	userHttp "github.com/fazarrahman/search-user/domain/user/repository/http"
	"github.com/fazarrahman/search-user/lib"
	"github.com/fazarrahman/search-user/service"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userHttpRepo := userHttp.New(lib.GetEnv("URL"))
	svc := service.New(userHttpRepo)

	var tag string

	rootCmd := &cobra.Command{
		Use: "user",
		Run: func(ccmd *cobra.Command, args []string) {
		},
	}

	rootCmd.Flags().StringVar(&tag, "tag", "", "User tag")
	rootCmd.Execute()

	svc.GetUserList()
}
