module example.com/myproject

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gin-gonic/gin v1.7.5
	github.com/joho/godotenv v1.4.0
)

replace (
	github.com/joho/godotenv => github.com/subosito/gotenv => v2.2.0 %!s(MISSING)
	github.com/go-sql-driver/mysql => github.com/go-sql-driver/mysql => v1.7.0 %!s(MISSING)
)

// exclude directive for the conflicting dependency
exclude github.com/some/problematic/v2 v2.1.0
