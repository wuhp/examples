package handler

import (
    "fmt"
    "log"
    "net/http"
    "runtime/debug"
    "time"

    "github.com/gorilla/mux"
)

type Route struct {
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

var routes = []Route{
    Route{"GET", "/api/v1/ping", Ping},

    Route{"GET",  "/api/v1/user",                wrapper(authWrapper(GetMyProfile))},
    Route{"POST", "/api/v1/user/reset-password", wrapper(authWrapper(ResetPassword))},
    Route{"POST", "/api/v1/login",               wrapper(authWrapper(Login))},
    Route{"GET",  "/api/v1/user/orgs",           wrapper(authWrapper(GetMyOrgs))},

    Route{"GET",    "/api/v1/orgs/{orgname}/teams",            wrapper(authWrapper(orgWrapper(ListOrgTeam)))},
    Route{"POST",   "/api/v1/orgs/{orgname}/teams",            wrapper(authWrapper(orgWrapper(PostTeam)))},
    Route{"GET",    "/api/v1/orgs/{orgname}/teams/{teamname}", wrapper(authWrapper(orgWrapper(teamWrapper(GetTeam))))},
    Route{"PUT",    "/api/v1/orgs/{orgname}/teams/{teamname}", wrapper(authWrapper(orgWrapper(teamWrapper(PutTeam))))},
    Route{"DELETE", "/api/v1/orgs/{orgname}/teams/{teamname}", wrapper(authWrapper(orgWrapper(teamWrapper(DeleteTeam))))},

    Route{"GET",    "/api/v1/orgs/{orgname}/teams/{teamname}/accesses",            wrapper(authWrapper(orgWrapper(teamWrapper(ListTeamAccess))))},
    Route{"POST",   "/api/v1/orgs/{orgname}/teams/{teamname}/accesses",            wrapper(authWrapper(orgWrapper(teamWrapper(PostTeamAccess))))},
    Route{"GET",    "/api/v1/orgs/{orgname}/teams/{teamname}/accesses/{reponame}", wrapper(authWrapper(orgWrapper(teamWrapper(GetTeamAccess))))},
    Route{"PUT",    "/api/v1/orgs/{orgname}/teams/{teamname}/accesses/{reponame}", wrapper(authWrapper(orgWrapper(teamWrapper(PutTeamAccess))))},
    Route{"DELETE", "/api/v1/orgs/{orgname}/teams/{teamname}/accesses/{reponame}", wrapper(authWrapper(orgWrapper(teamWrapper(DeleteTeamAccess))))},

    Route{"GET",    "/api/v1/orgs/{orgname}/teams/{teamname}/members",            wrapper(authWrapper(orgWrapper(teamWrapper(ListTeamMember))))},
    Route{"POST",   "/api/v1/orgs/{orgname}/teams/{teamname}/members",            wrapper(authWrapper(orgWrapper(teamWrapper(PostTeamMember))))},
    Route{"DELETE", "/api/v1/orgs/{orgname}/teams/{teamname}/members/{username}", wrapper(authWrapper(orgWrapper(teamWrapper(DeleteTeamMember))))},

    Route{"GET",    "/api/v1/orgs/{orgname}/members",            wrapper(authWrapper(orgWrapper(ListOrgMember)))},
    Route{"POST",   "/api/v1/orgs/{orgname}/members",            wrapper(authWrapper(orgWrapper(PostOrgMember)))},
    Route{"DELETE", "/api/v1/orgs/{orgname}/members/{username}", wrapper(authWrapper(orgWrapper(DeleteOrgMember)))},

    Route{"GET",    "/api/v1/repos",                        wrapper(authWrapper(ListRepo))},
    Route{"POST",   "/api/v1/repos",                        wrapper(authWrapper(PostRepo))},
    Route{"GET",    "/api/v1/repos/{namespace}/{reponame}", wrapper(authWrapper(repoWrapper(GetRepo)))},
    Route{"PUT",    "/api/v1/repos/{namespace}/{reponame}", wrapper(authWrapper(repoWrapper(PutRepo)))},
    Route{"DELETE", "/api/v1/repos/{namespace}/{reponame}", wrapper(authWrapper(repoWrapper(DeleteRepo)))},

    Route{"GET",    "/api/v1/repos/{namespace}/{reponame}/tags",          wrapper(authWrapper(repoWrapper(ListTag)))},
    Route{"DELETE", "/api/v1/repos/{namespace}/{reponame}tags/{tagname}", wrapper(authWrapper(repoWrapper(DeleteTag)))},

    // Api only for Admin
    Route{"GET",    "/api/v1/info",             wrapper(authWrapper(adminWrapper(Info)))},
    Route{"GET",    "/api/v1/users",            wrapper(authWrapper(adminWrapper(ListUser)))},
    Route{"POST",   "/api/v1/users",            wrapper(authWrapper(adminWrapper(PostUser)))},
    Route{"GET",    "/api/v1/users/{username}", wrapper(authWrapper(adminWrapper(userWrapper(GetUser))))},
    Route{"PUT",    "/api/v1/users/{username}", wrapper(authWrapper(adminWrapper(userWrapper(PutUser))))},
    Route{"DELETE", "/api/v1/users/{username}", wrapper(authWrapper(adminWrapper(userWrapper(DeleteUser))))},
    Route{"GET",    "/api/v1/orgs",             wrapper(authWrapper(adminWrapper(ListOrg)))},
    Route{"POST",   "/api/v1/orgs",             wrapper(authWrapper(adminWrapper(PostOrg)))},
    Route{"DELETE", "/api/v1/orgs/{orgname}",   wrapper(authWrapper(adminWrapper(orgWrapper(DeleteOrg))))},
}

type InnerResponseWriter struct {
    statusCode int
    setted     bool
    http.ResponseWriter
}

func (i *InnerResponseWriter) WriteHeader(status int) {
    if !i.setted {
        i.statusCode = status
        i.setted = true
    }

    i.ResponseWriter.WriteHeader(status)
}

func (i *InnerResponseWriter) Write(b []byte) (int, error) {
    i.setted = true
    return i.ResponseWriter.Write(b)
}

func wrapper(inner http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        s := time.Now()
        wr := &InnerResponseWriter{
            statusCode:     200,
            setted:         false,
            ResponseWriter: w,
        }

        defer func() {
            if err := recover(); err != nil {
                debug.PrintStack()
                wr.WriteHeader(http.StatusInternalServerError)
                log.Printf("Panic: %v\n", err)
                fmt.Fprintf(w, fmt.Sprintln(err))
            }

            d := time.Now().Sub(s)
            log.Printf("%s %s %d %s\n", r.Method, r.RequestURI, wr.statusCode, d.String())
        }()

        inner.ServeHTTP(wr, r)
    })
}

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    for _, route := range routes {
        router.Methods(route.Method).Path(route.Pattern).HandlerFunc(route.HandlerFunc)
    }

    return router
}
