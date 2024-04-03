package config

import (
	"cmp"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func AppAddr() string {
	return cmp.Or(os.Getenv("APP_ADDR"), ":8080")
}

func MySQL() *mysql.Config {
	c := mysql.NewConfig()

	c.User = cmp.Or(os.Getenv("DB_USER"), "root")
	c.Passwd = cmp.Or(os.Getenv("DB_PASS"), "pass")
	c.Net = cmp.Or(os.Getenv("DB_NET"), "tcp")
	c.Addr = fmt.Sprintf(
		"%s:%s",
		cmp.Or(os.Getenv("DB_HOST"), "localhost"),
		cmp.Or(os.Getenv("DB_PORT"), "3306"),
	)
	c.DBName = cmp.Or(os.Getenv("DB_NAME"), "app")
	c.Collation = "utf8mb4_general_ci"
	c.AllowNativePasswords = true

	return c
}
