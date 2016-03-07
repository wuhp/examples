package model

import (
    "time"
)

type Org struct {
    Id          int64  `json:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    CreateTs    int64  `json:"create_ts"`
    Description string `json:"description"`
}

////////////////////////////////////////////////////////////////////////////////

func ListOrg(cs []*Condition, o *Order, p *Paging) []*Org {
    if cs == nil {
        cs = make([]*Condition, 0)
    }

    cs = append(cs, NewCondition("is_org", "=", true))

    where, vs := GenerateWhereSql(cs)
    order := GenerateOrderSql(o)
    limit := GenerateLimitSql(p)

    rows, err := db.Query(`
        SELECT
            id, name, display_name, create_ts, org_description
        FROM
            account
        ` + where + order + limit, vs...,
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

func GetOrgById(id int64) *Org {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("id", "=", id))

    l := ListOrg(conditions, nil, nil)
    if len(l) == 0 {
        return nil
    }

    return l[0]
}

func GetOrgByName(name string) *Org {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("name", "=", name))

    l := ListOrg(conditions, nil, nil)
    if len(l) == 0 {
        return nil
    }

    return l[0]
}

func (o *Org) Save() {
    stmt, err := db.Prepare(`
        INSERT INTO account(
            name, display_name, create_ts, org_description
        )
        VALUES(?, ?, ?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    o.CreateTs = time.Now().UTC().Unix()

    result, err := stmt.Exec(
        o.Name, o.DisplayName, o.CreateTs, o.Description,
    )
    if err != nil {
        panic(err)
    }

    o.Id, err = result.LastInsertId()
    if err != nil {
        panic(err)
    }
}

func (o *Org) Update() {
    stmt, err := db.Prepare(`
        UPDATE
            account
        SET
            display_name = ?,
            org_description = ?
        WHERE
            id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(
        o.DisplayName,
        o.Description,
        o.Id,
    ); err != nil {
        panic(err)
    }
}

func (o *Org) Delete() {
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

    if _, err := stmt.Exec(o.Id); err != nil {
        panic(err)
    }
}

////////////////////////////////////////////////////////////////////////////////

func (o *Org) Teams() []*Team {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("org", "=", o.Name))

    return ListTeam(conditions, nil, nil)
}

////////////////////////////////////////////////////////////////////////////////

func (o *Org) Members() []*User {
    rows, err := db.Query(`
        SELECT
            account.id, account.name, account.display_name, account.create_ts,
            account.user_type, account.user_email, account.user_global_permission
        FROM
            account, orgmember
        WHERE
            account.name = orgmember.member and orgmember.org = ?
        `, o.Name,
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

func (o *Org) AddMember(u *User) {
    stmt, err := db.Prepare(`
        INSERT INTO orgmember(
            org, member
        )
        VALUES(?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(
        o.Name, u.Name,
    )
    if err != nil {
        panic(err)
    }
}

func (o *Org) RemoveMember(u *User) {
    stmt, err := db.Prepare(`
        DELETE FROM
            orgmember
        WHERE
            org = ? and member = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(o.Name, u.Name); err != nil {
        panic(err)
    }
}
