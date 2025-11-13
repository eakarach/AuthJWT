# AuthJWT

## Install the dependencies:
    go mod download

## Clean up and synchronize your go.mod file:
    go mod tidy

## Run the application:
    go run main.go
The API should now be running on `http://localhost:3000`.



## MOD
### fiber
    go get github.com/gofiber/fiber/v2
### jwt
    go get -u github.com/golang-jwt/jwt/v5
### bcrypt
    golang.org/x/crypto/bcrypt
### gorm
    go get -u gorm.io/gorm
### sqlserver
    go get gorm.io/driver/sqlserver
### go-mssqldb
    go get github.com/microsoft/go-mssqldb
