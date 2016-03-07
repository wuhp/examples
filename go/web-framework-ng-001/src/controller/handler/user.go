package handler

import (
    "encoding/json"
    "net/http"

    "controller/model"
)

// GET /api/v1/user
//
func GetMyProfile(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(UserVars[r])
}

// POST /api/v1/user/reset-password
//
func ResetPassword(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    in := struct {
        Password string `json:"password"`
    }{}

    if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
        http.Error(w, RequestBodyDecodeError, http.StatusBadRequest)
        return
    }

    UserVars[r].ResetPassword(hashPassword(in.Password))
}

// POST /api/v1/login
//
func Login(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    in := struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }{}

    if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
        http.Error(w, RequestBodyDecodeError, http.StatusBadRequest)
        return
    }

    u := model.GetUserByName(in.Username)
    if u == nil || u.GetPassword() != hashPassword(in.Password) {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    out := struct {
        Key string `json:"key"`
    }{}

    out.Key = encodeUserToken(in.Username)
    json.NewEncoder(w).Encode(out)
}

// GET /api/v1/user/orgs
//
func GetMyOrgs(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(UserVars[r].Orgs())
}
