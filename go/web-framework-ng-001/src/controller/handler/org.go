package handler

import (
    "encoding/json"
    "net/http"

    "controller/model"
)

// GET /api/v1/orgs/{orgname}/teams
//
func ListOrgTeam(w http.ResponseWriter, r *http.Request) {
    // TBD
    // json.NewEncoder(w).Encode(OrgVars[r].Teams())
}

// POST /api/v1/orgs/{orgname}/teams
//
func PostTeam(w http.ResponseWriter, r *http.Request) {
    org := OrgVars[r]
    user := LoginUserVars[r]
    if !user.IsAdminOfOrg(org) {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    defer r.Body.Close()

    in := struct {
        Name        string `json:"name"`
        Description string `json:"description"`
        Type        string `json:"type"`
    }{}
    if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
        http.Error(w, RequestBodyDecodeError, http.StatusBadRequest)
        return
    }

    if len(in.Name) == 0 {
        http.Error(w, RequestBodyError, http.StatusBadRequest)
        return
    }

    if len(in.Type) == 0 {
        in.Type = ManagedAuth
    }

    if model.GetTeamByName(org.Name, in.Name) != nil {
        http.Error(w, DuplicateResource, http.StatusBadRequest)
        return
    }

    t := new(model.Team)
    t.Org = org.Name
    t.Name = in.Name
    t.Description = in.Description
    t.Type = in.Type
    t.Save()

    w.WriteHeader(http.StatusCreated)
}

// GET /api/v1/orgs/{orgname}/teams/{teamname}
//
func GetTeam(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// PUT /api/v1/orgs/{orgname}/teams/{teamname}
//
func PutTeam(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// DELETE /api/v1/orgs/{orgname}/teams/{teamname}
//
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/orgs/{orgname}/teams/{teamname}/accesses
//
func ListTeamAccess(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// POST /api/v1/orgs/{orgname}/teams/{teamname}/accesses
//
func PostTeamAccess(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/orgs/{orgname}/teams/{teamname}/accesses/{reponame}
//
func GetTeamAccess(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// PUT /api/v1/orgs/{orgname}/teams/{teamname}/accesses/{reponame}
//
func PutTeamAccess(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// DELETE /api/v1/orgs/{orgname}/teams/{teamname}/accesses/{reponame}
//
func DeleteTeamAccess(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/orgs/{orgname}/teams/{teamname}/members
//
func ListTeamMember(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// POST /api/v1/orgs/{orgname}/teams/{teamname}/members
//
func PostTeamMember(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// DELETE /api/v1/orgs/{orgname}/teams/{teamname}/members/{username}
//
func DeleteTeamMember(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// GET /api/v1/orgs/{orgname}/members
//
func ListOrgMember(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// POST /api/v1/orgs/{orgname}/members
//
func PostOrgMember(w http.ResponseWriter, r *http.Request) {
    // TBD
}

// DELETE /api/v1/orgs/{orgname}/members/{username}
//
func DeleteOrgMember(w http.ResponseWriter, r *http.Request) {
    // TBD
}
