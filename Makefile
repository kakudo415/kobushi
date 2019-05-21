main:
	go build -o bin/kobushi src/main.go
run: main
	./bin/kobushi
package:
	go get -u "github.com/gin-gonic/gin"
	go get -u "github.com/kakudo415/kid"
	go get -u "github.com/jinzhu/gorm"
	go get -u "github.com/go-sql-driver/mysql"