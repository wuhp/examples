package model

import (
    "time"
)

type Repo struct {
    Id          int64  `json:"id"`
    Namespace   string `json:"namespace"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Readme      string `json:"readme"`
    IsPublic    bool   `json:"is_public"`
    CreateTs    int64  `json:"create_ts"`
}

type Tag struct {
    RepoId int64  `json:"repo_id"`
    Name   string `json:"name"`
}

type Hook struct {
    RepoId int64  `json:"repo_id"`
    Name   string `json:"name"`
    Url    string `json:"url"`
    Secret string `json:"secret"`
}

////////////////////////////////////////////////////////////////////////////////

func ListRepo(cs []*Condition, o *Order, p *Paging) []*Repo {
    where, vs := GenerateWhereSql(cs)
    order := GenerateOrderSql(o)
    limit := GenerateLimitSql(p)

    rows, err := db.Query(`
        SELECT
            id, namespace, name, description, readme, is_public, create_ts
        FROM
            repo
        ` + where + order + limit, vs...,
    )
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    l := make([]*Repo, 0)
    for rows.Next() {
        r := new(Repo)
        if err := rows.Scan(
            &r.Id, &r.Namespace, &r.Name, &r.Description, &r.Readme, &r.IsPublic, &r.CreateTs,
        ); err != nil {
            panic(err)
        }

        l = append(l, r)
    }

    return l
}

func GetRepoById(id int64) *Repo {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("id", "=", id))

    l := ListRepo(conditions, nil, nil)
    if len(l) == 0 {
        return nil
    }

    return l[0]
}

func GetRepoByName(namespace, name string) *Repo {
    conditions := make([]*Condition, 0)
    conditions = append(conditions, NewCondition("namespace", "=", namespace))
    conditions = append(conditions, NewCondition("name", "=", name))

    l := ListRepo(conditions, nil, nil)
    if len(l) == 0 {
        return nil
    }

    return l[0]
}

func (r *Repo) Save() {
    stmt, err := db.Prepare(`
        INSERT INTO repo(
            namespace, name, description, readme, is_public, create_ts
        )
        VALUES(?, ?, ?, ?, ?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    r.CreateTs = time.Now().UTC().Unix()

    result, err := stmt.Exec(
        r.Namespace, r.Name, r.Description, r.Readme, r.IsPublic, r.CreateTs,
    )
    if err != nil {
        panic(err)
    }

    r.Id, err = result.LastInsertId()
    if err != nil {
        panic(err)
    }
}

func (r *Repo) Update() {
    stmt, err := db.Prepare(`
        UPDATE
            repo
        SET
            description = ?,
            readme = ?,
            is_public = ?
        WHERE
            id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(
        r.Description,
        r.Readme,
        r.IsPublic,
    ); err != nil {
        panic(err)
    }
}

func (r *Repo) Delete() {
    stmt, err := db.Prepare(`
        DELETE FROM
            repo
        WHERE
            id = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(r.Id); err != nil {
        panic(err)
    }
}

////////////////////////////////////////////////////////////////////////////////

func (r *Repo) Tags() []*Tag {
    rows, err := db.Query(`
        SELECT
            repo_id, name
        FROM
            tag
        WHERE
            repo_id =  ?
        `, r.Id,
    )
    if err != nil {
        panic(err)
    }   
    defer rows.Close()
    
    l := make([]*Tag, 0)
    for rows.Next() { 
        t := new(Tag)
        if err := rows.Scan(
            &t.RepoId, &t.Name,
        ); err != nil {
            panic(err)
        }   
        
        l = append(l, t)
    }   
    
    return l
}

func (r *Repo) AddTag(t *Tag) {
    stmt, err := db.Prepare(`
        INSERT INTO tag(
            repo_id, name
        )
        VALUES(?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(
        r.Id, t.Name,
    )
    if err != nil {
        panic(err)
    }
}

func (r *Repo) RemoveTag(name string) {
    stmt, err := db.Prepare(`
        DELETE FROM
            tag
        WHERE
            repo_id = ? and name = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(r.Id, name); err != nil {
        panic(err)
    }
}

////////////////////////////////////////////////////////////////////////////////

func (r *Repo) Hooks() []*Hook {
    rows, err := db.Query(`
        SELECT
            repo_id, name, url, secret
        FROM
            hook
        WHERE
            repo_id =  ?
        `, r.Id,
    )
    if err != nil {
        panic(err)
    }   
    defer rows.Close()
    
    l := make([]*Hook, 0)
    for rows.Next() { 
        h := new(Hook)
        if err := rows.Scan(
            &h.RepoId, &h.Name, &h.Url, &h.Secret,
        ); err != nil {
            panic(err)
        }   
        
        l = append(l, h)
    }

    return l
}

func (r *Repo) AddHook(hook *Hook) {
    stmt, err := db.Prepare(`
        INSERT INTO hook(
            repo_id, name, url, secret
        )
        VALUES(?, ?, ?, ?)
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(
        r.Id, hook.Name, hook.Url, hook.Secret,
    )
    if err != nil {
        panic(err)
    }
}

func (r *Repo) RemoveHook(name string) {
    stmt, err := db.Prepare(`
        DELETE FROM
            hook
        WHERE
            repo_id = ? and name = ?
    `)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    if _, err := stmt.Exec(r.Id, name); err != nil {
        panic(err)
    }
}
