package dbsync

import "fmt"

func ConnStr(DBName string, Host string, Port int, Username string, Password string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, Username, Password, DBName)
}
