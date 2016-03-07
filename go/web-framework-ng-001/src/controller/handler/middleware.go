package handler

import (
    "net/http"

    "controller/model"
)

var LoginUserVars map[*http.Request]*model.User
var AdminVars     map[*http.Request]bool
var OrgVars       map[*http.Request]*model.Org
var TeamVars      map[*http.Request]*model.Team
var RepoVars      map[*http.Request]*model.Repo
var UserVars      map[*http.Request]*model.User

func init() {
    LoginUserVars = make(map[*http.Request]*model.User)
    AdminVars     = make(map[*http.Request]bool)
    OrgVars       = make(map[*http.Request]*model.Org)
    TeamVars      = make(map[*http.Request]*model.Team)
    RepoVars      = make(map[*http.Request]*model.Repo)
    UserVars      = make(map[*http.Request]*model.User)
}

func authWrapper(inner http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TBD
    })
}

func adminWrapper(inner http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TBD
    })
}

func orgWrapper(inner http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TBD
    })
}

func teamWrapper(inner http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TBD
    })
}

func repoWrapper(inner http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TBD
    })
}

func userWrapper(inner http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TBD
    })
}
