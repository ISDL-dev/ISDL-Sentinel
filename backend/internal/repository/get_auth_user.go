package repository

import (
	"log"

	"github.com/ISDL-dev/ISDL_Sentinel/backend/internal/schema"
)

func GetAuthUser() (map[string]schema.PostUserRequest, error) {
    users := make(map[string]schema.PostUserRequest)

    rows, err := db.Query("SELECT mail_address, password, name FROM users")//table名古い
    if err != nil {
        log.Printf("Error querying database: %v", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var mailAddress, password, name string
        err := rows.Scan(&mailAddress, &password, &name)
        if err != nil {
            return nil, err
        }

        users[mailAddress] = schema.PostUserRequest{
            MailAddress: mailAddress,
            Password:    password,
            Name:        name,
        }
    }

    return users, nil
}