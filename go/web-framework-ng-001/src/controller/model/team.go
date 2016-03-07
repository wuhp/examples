package model

import (
    "time"
)

type Team struct {
    Id          int64  `json:"id"`
    Org         string `json:"org"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Type        string `json:"type"`
    CreateTs    int64  `json:"create_ts"`
}

////////////////////////////////////////////////////////////////////////////////

func ListTeam(cs []*Condition, o *Order, p *Paging) []*Team {
    where, vs := GenerateWhereSql(cs)
    order := GenerateOrderSql(o)
    limit := GenerateLimitSql(p)

    rows, err := db.Query(`
        SELECT
            id, org, name, description, type, create_ts
        FROM
            team
        ` + where + order + limit, vs...,
    )
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    l := make([]*Team, 0)
    for rows.Next() {
        t := new(Team)
        if err := rows.Scan(
            &t.Id, &t.Org, &t.Name, &t.Description, &t.Type, &t.CreateTs,
        ); err != nil {
            panic(err)
        }

        l = append(l, t)
    }

    return l
}

func GetTeamById(id int64) *Team {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("id", "=", id))

    l := ListTeam(conditions, nil, nil)
    if len(l) == 0 {
        return nil
    }

    return l[0]
}

func GetTeamByName(org, name string) *Team {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("org", "=", org))
    conditions = append(conditions, NewCondition("name", "=", name))

    l := ListTeam(conditions, nil, nil)
    if len(l) == 0 {
        return nil
    }

    return l[0]
}

func (t *Team) Save() {
    stmt, err := db.Prepare(`
        INSERT INTO team(
            org, name, description, type, create_ts
        )
        VALUES(?, ?, ?, ?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    t.CreateTs = time.Now().UTC().Unix()

    result, err := stmt.Exec(
        t.Org, t.Name, t.Description, t.Type, t.CreateTs,
    )
    if err != nil {
        panic(err)
    }

    t.Id, err = result.LastInsertId()
    if err != nil {
        panic(err)
    }
}

func (t *Team) Update() {
    stmt, err := db.Prepare(`
        UPDATE
            team
        SET
            description = ?
        WHERE
            id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(
        t.Description,
    ); err != nil {
        panic(err)
    }
}

func (t *Team) Delete() {
    stmt, err := db.Prepare(`
        DELETE FROM
            team
        WHERE
            id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(t.Id); err != nil {
        panic(err)
    }
}

/////////////////////////////////////////////////////////////////////////////////

func (t *Team) Members() []*User {
    rows, err := db.Query(`
        SELECT
            account.id, account.name, account.display_name, account.create_ts,
            account.user_type, account.user_email, account.user_global_permission
        FROM
            account, teammember
        WHERE
            account.name = teammember.member and teammember.team_id = ?
        `, t.Id,
    )
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    l := make([]*User, 0)
    for rows.Next() {
        u := new(User)
        if err := rows.Scan(
            &u.Id, &u.Name, &u.DisplayName, &u.CreateTs, &u.Type, &u.Email, &u.GlobalPermission,
        ); err != nil {
            panic(err)
        }

        l = append(l, u)
    }
    
    return l
}

func (t *Team) HasMember(u *User) bool {
    for _, m := range t.Members() {
        if m.Name == u.Name {
            return true
        }
    }

    return false
}

func (t *Team) AddMember(u *User) {
    stmt, err := db.Prepare(`
        INSERT INTO teammember(
            team_id, member
        )
        VALUES(?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(
        t.Id, u.Name,
    )
    if err != nil {
        panic(err)
    }
}

func (t *Team) RemoveMember(u *User) {
    stmt, err := db.Prepare(`
        DELETE FROM
            teammember
        WHERE
            team_id = ? and member = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(t.Id, u.Name); err != nil {
        panic(err)
    }
}

///////////////////////////////////////////////////////////////////////////////

func (t *Team) Accesses() ([]*Repo, []string) {
    rows, err := db.Query(`
        SELECT
            repo.id, repo.namespace, repo.name, repo.description, repo.readme,
            repo.is_public, repo.create_ts, teamaccess.access
        FROM
            repo, teamaccess
        WHERE
            repo.id = teamaccess.repo_id and teamaccess.team_id = ?
        `, t.Id,
    )
    if err != nil {
        panic(err)
    }   
    defer rows.Close()
    
    lr := make([]*Repo, 0)
    la := make([]string, 0)
    for rows.Next() {
        var s string
        r := new(Repo)
        if err := rows.Scan(
            &r.Id, &r.Namespace, &r.Name, &r.Description, &r.Readme,
            &r.IsPublic, &r.CreateTs, &s,
        ); err != nil {
            panic(err)
        }   
        
        lr = append(lr, r)
        la = append(la, s)
    }   
    
    return lr, la
}

func (t *Team) AddAccess(r *Repo, acc string) {
    stmt, err := db.Prepare(`
        INSERT INTO teamaccess(
            team_id, repo_id, access
        )
        VALUES(?, ?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(
        t.Id, r.Id, acc,
    )
    if err != nil {
        panic(err)
    }
}

func (t *Team) SetAccess(r *Repo, acc string) {
    stmt, err := db.Prepare(`
        UPDATE
            teamaccess
        SET
            access = ?
        WHERE
            team_id = ? and repo_id = ?
    `)  
    if err != nil {
        panic(err)
    }   
    defer stmt.Close()
    
    _, err = stmt.Exec(
        acc, t.Id, r.Id,
    )   
    if err != nil {
        panic(err)
    }   
}

func (t *Team) RemoveAccess(r *Repo) {
    stmt, err := db.Prepare(`
        DELETE FROM
            teamaccess
        WHERE
            team_id = ? and repo_id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(t.Id, r.Id); err != nil {
        panic(err)
    }
}
