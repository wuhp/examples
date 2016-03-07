package handler

import (
    "net/http"
)

// GET /api/v1/info
//
func Info(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/users
//
func ListUser(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// POST /api/v1/users
//
func PostUser(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/users/{username}
//
func GetUser(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// PUT /api/v1/users/{username}
//
func PutUser(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// DELETE /api/v1/users/{username}
//
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/orgs
//
func ListOrg(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// POST /api/v1/orgs
//
func PostOrg(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// DELETE /api/v1/orgs/{orgname}
//
func DeleteOrg(w http.ResponseWriter, r *http.Request) {
    // TBD
}
