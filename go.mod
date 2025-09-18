module github.com/tagaertner/job-stories

go 1.24.0

require (
	github.com/99designs/gqlgen v0.17.79
	github.com/joho/godotenv v1.5.1
	github.com/vektah/gqlparser/v2 v2.5.30
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.0
// add other shared deps here
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/text v0.29.0 // indirect
)

replace github.com/tagaertner/job-stories/services/users => ./services/users

// replace github.com/tagaertner/job-stories/services/orders => ./services/orders

replace github.com/tagaertner/job-stories/services/stories => ./services/stories
