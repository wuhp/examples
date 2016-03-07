package model

import (
    "time"
)

type User struct {
    Id               int64  `json:"id"`
    Name             string `json:"name"`
    DisplayName      string `json:"display_name"`
    CreateTs         int64  `json:"create_ts"`
    Type             string `json:"type"`
    Email            string `json:"email"`
    GlobalPermission string `json:"global_permission"`
}

const (
    NoGlobalPermission = "none"
    GlobalPull         = "pull"
    GlobalPullPush     = "pull,push"
    GlobalAdmin        = "admin"
)

////////////////////////////////////////////////////////////////////////////////

func ListUser(cs []*Condition, o *Order, p *Paging) []*User {
    if cs == nil {
        cs = make([]*Condition, 0)
    }

    cs = append(cs, NewCondition("is_org", "=", false))

    where, vs := GenerateWhereSql(cs)
    order := GenerateOrderSql(o)
    limit := GenerateLimitSql(p)

    rows, err := db.Query(`
        SELECT
            id, name, display_name, create_ts, user_type, user_email, user_global_permission
        FROM
            account
        ` + where + order + limit, vs...,
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

func GetUserById(id int64) *User {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("id", "=", id))

    l := ListUser(conditions, nil, nil)
    if len(l) == 0 {
        return nil
    }

    return l[0]
}

func GetUserByName(name string) *User {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("name", "=", name))

    l := ListUser(conditions, nil, nil)
    if len(l) == 0 {
        return nil
    }

    return l[0]
}

func (u *User) Save() {
    stmt, err := db.Prepare(`
        INSERT INTO account(
            name, display_name, create_ts, user_type, user_email, user_global_permission
        )
        VALUES(?, ?, ?, ?, ?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    u.CreateTs = time.Now().UTC().Unix()

    result, err := stmt.Exec(
        u.Name, u.DisplayName, u.CreateTs, u.Type, u.Email, u.GlobalPermission,
    )
    if err != nil {
        panic(err)
    }

    u.Id, err = result.LastInsertId()
    if err != nil {
        panic(err)
    }
}

func (u *User) Update() {
    stmt, err := db.Prepare(`
        UPDATE
            account
        SET
            display_name = ?,
            user_email = ?,
            user_global_permission = ?
        WHERE
            id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(
        u.DisplayName,
        u.Email,
        u.GlobalPermission,
        u.Id,
    ); err != nil {
        panic(err)
    }
}

func (u *User) Delete() {
    stmt, err := db.Prepare(`
        DELETE FROM
            account
        WHERE
            id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(u.Id); err != nil {
        panic(err)
    }
}

////////////////////////////////////////////////////////////////////////////////

func (u *User) Orgs() []*Org {
    rows, err := db.Query(`
        SELECT
            account.id, account.name, account.display_name, account.create_ts,
            account.org_description,
        FROM
            account, orgmember
        WHERE
            account.name = orgmember.org and orgmember.member = ?
        `, u.Name,
    )
    if err != nil {
        panic(err)
    }   
    defer rows.Close()
    
    l := make([]*Org, 0)
    for rows.Next() { 
        o := new(Org)
        if err := rows.Scan(
            &o.Id, &o.Name, &o.DisplayName, &o.CreateTs, &o.Description,
        ); err != nil {
            panic(err)
        }
 
        l = append(l, o)
    }

    return l
}

func (u *User) Teams() []*Team {
    rows, err := db.Query(`
        SELECT
            team.id, team.org, team.name, team.description, team.type, team.create_ts
        FROM
            team, teammember
        WHERE
            team.id = teammember.team_id and teammember.member = ?
        `, u.Name,
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

func (u *User) HasGlobalAdmin() bool {
    return u.GlobalPermission == GlobalAdmin
}

func (u *User) HasGlobalPush() bool {
    return u.GlobalPermission == GlobalPullPush || u.GlobalPermission == GlobalAdmin
}

func (u *User) HasGlobalPull() bool {
    return u.GlobalPermission == GlobalPull || u.GlobalPermission == GlobalPullPush || u.GlobalPermission == GlobalAdmin
}

func (u *User) IsAdminOfOrg(o *Org) bool {
    if GetTeamByName(o.Name, "owner").HasMember(u) {
        return true
    }

    return u.HasGlobalAdmin()
}

func (u *User) GetPassword() string {
    rows, err := db.Query(`
        SELECT
            password
        FROM
            account
        WHERE
            name = ?
        `, u.Name,
    )
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var hp string
    for rows.Next() {
        if err := rows.Scan(&hp); err != nil {
            panic(err)
        }
    }

    return hp
}

func (u *User) ResetPassword(hashedPasswd string) {
    stmt, err := db.Prepare(`
        UPDATE
            account
        SET
            user_password = ?
        WHERE
            id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(
        hashedPasswd,
        u.Id,
    ); err != nil {
        panic(err)
    }
}
