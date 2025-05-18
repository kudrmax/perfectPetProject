package database

import "fmt"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func (c *Config) ToDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DbName,
	)
}
