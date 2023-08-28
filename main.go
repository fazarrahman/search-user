/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"strings"

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

	var fetch, tags, isUseErrorApi string

	cmd := &cobra.Command{
		Use: "user",
		Run: func(ccmd *cobra.Command, args []string) {
		},
	}

	cmd.Flags().StringVar(&fetch, "fetch", "", "Fetch user list from API and write to users.csv")
	cmd.Flags().StringVar(&isUseErrorApi, "isUseErrorApi", "", "Fetch user list from 503 error API")
	cmd.Flags().StringVar(&tags, "tag", "", "User tag")
	cmd.Execute()

	if fetch == "true" {
		if isUseErrorApi == "true" {
			userHttpRepo := userHttp.New(lib.GetEnv("ERROR_503_URL"))
			svc = service.New(userHttpRepo)
		}
		svc.GetUserList()
	} else if strings.Trim(tags, " ") != "" {
		arr := strings.Split(tags, ",")
		svc.SearchUser(arr)
	}
}
