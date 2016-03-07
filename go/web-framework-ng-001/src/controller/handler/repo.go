package handler

import (
    "net/http"
)

// GET /api/v1/repos
//
func ListRepo(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// POST /api/v1/repos
//
func PostRepo(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/repos/{namespace}/{reponame}
//
func GetRepo(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// PUT /api/v1/repos/{namespace}/{reponame}
//
func PutRepo(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// DELETE /api/v1/repos/{namespace}/{reponame}
//
func DeleteRepo(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/repos/{namespace}/{reponame}/tags
//
func ListTag(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// DELETE /api/v1/repos/{namespace}/{reponame}tags/{tagname}
//
func DeleteTag(w http.ResponseWriter, r *http.Request) {
    // TBD
}
