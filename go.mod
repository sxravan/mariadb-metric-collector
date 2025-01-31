module github.com/sxravan/mariadb-metric-collector

go 1.23.5

require (
	github.com/sxravan/mariadb-metric-collector/collector v0.0.2
	github.com/sxravan/mariadb-metric-collector/config v0.0.1
	github.com/sxravan/mariadb-metric-collector/db v0.0.3
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)

replace github.com/sxravan/mariadb-metric-collector/config => ./config

replace github.com/sxravan/mariadb-metric-collector/db => ./db

replace github.com/sxravan/mariadb-metric-collector/collector => ./collector
