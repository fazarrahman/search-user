# search-user

### Command to fetch user list from API
go run main.go --fetch=true

### Command to fetch user list using error 503 url API
go run main.go --fetch=true --isUseErrorApi=true

### Command to search users by tags (with example tags param)
go run main.go --tag="cupidatat,dolor"

### Command to run test
go test -short -coverprofile coverage.out -v ./domain/user/repository/http