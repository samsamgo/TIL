package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
)

// User is a struct that represents a user.
type User struct {
    ID int
    Name string
    Email string
}

// GetUser gets a user by ID.
func GetUser(id int) (*User, error) {
    // Connect to the database.
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mydb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare the query.
    stmt, err := db.Prepare("SELECT id, name, email FROM users WHERE id = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Execute the query.
    rows, err := stmt.Query(id)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Iterate over the rows and get the user.
    if rows.Next() {
        var user User
        err = rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            log.Fatal(err)
        }
        return &user, nil
    } else {
        return nil, nil
    }
}

// GetUsers gets all users.
func GetUsers() ([]User, error) {
    // Connect to the database.
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mydb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare the query.
    stmt, err := db.Prepare("SELECT id, name, email FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Execute the query.
    rows, err := stmt.Query()
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Iterate over the rows and get the users.
    users := []User{}
    for rows.Next() {
        var user User
        err = rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }
    return users, nil
}

// CreateUser creates a new user.
func CreateUser(name, email string) error {
    // Connect to the database.
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mydb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare the query.
    stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Execute the query.
    result, err := stmt.Exec(name, email)
    if err != nil {
        log.Fatal(err)
    }

    // Get the ID of the newly created user.
    id, err := result.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }

    // Return the ID of the newly created user.
    return nil
}

// UpdateUser updates an existing user.
func UpdateUser(id int, name, email string) error {
    // Connect to the database.
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mydb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare the query.
    stmt, err := db.Prepare("UPDATE users SET name = ?
