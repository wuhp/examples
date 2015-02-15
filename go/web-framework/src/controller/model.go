package main

import (
  "database/sql"
  "fmt"
  "strings"

  _ "github.com/go-sql-driver/mysql"
)


var gdb *sql.DB = nil

/////////////////////////////////////////////////////////////////////////////////

func ConnectDB(address string, port int, user string, passwd string, db string) error {
  if gdb == nil {
    connLink := fmt.Sprintf(
      "%s:%s@tcp(%s:%d)/%s",
      user, passwd, address, port, db,
    )

    var err error
    if gdb, err = sql.Open("mysql", connLink); err != nil {
      return err
    }

    if err = gdb.Ping(); err != nil {
      gdb.Close()
      return err
    }
  }

  return nil
}

func DisconnectDB() {
  if gdb != nil {
    gdb.Close()
    gdb = nil
  }
}

/////////////////////////////////////////////////////////////////////////////////

type ProductModel struct {
  Id          int64   `json:"id"`
  Name        string  `json:"name"`
  Description string  `json:"description"`
  CreateTs    string  `json:"create_ts"`
  LastUpdateTs string `json:"last_update_ts"`
}

func ListProductModel(con map[string]interface{}) []*ProductModel {
  ks := make([]string, 0)
  vs := make([]interface{}, 0)
  for k, v := range con {
    ks = append(ks, k+"=?")
    vs = append(vs, v)
  }

  where := ""
  if con != nil {
    where = "WHERE " + strings.Join(ks, " and ")
  }

  rows, err := gdb.Query(
    `
    SELECT
      id, name, description, create_ts, last_update_ts
    FROM
      product
    ` + where,
    vs..., 
  )
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  l := make([]*ProductModel, 0)
  for rows.Next() {
    m := new(ProductModel)
    if err := rows.Scan(
      &m.Id, &m.Name, &m.Description,
      &m.CreateTs, &m.LastUpdateTs,
    ); err != nil {
      panic(err)
    }

    l = append(l, m)
  }

  return l
}

func CreateProductModel(m *ProductModel) int64 {
  stmt, err := gdb.Prepare(`
    INSERT INTO
      product(name, description, create_ts, last_update_ts)
    VALUES(?, ?, ?, ?)
  `)
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  result, err := stmt.Exec(
    &m.Name, &m.Description,
    &m.CreateTs, &m.LastUpdateTs,
  )
  if err != nil {
    panic(err)
  }

  id, err := result.LastInsertId()
  if err != nil {
    panic(err)
  }

  return id
}

func GetProductModel(con map[string]interface{}) *ProductModel {
  m := new(ProductModel)
  ks := make([]string, 0)
  vs := make([]interface{}, 0)
  for k, v := range con {
    ks = append(ks, k+"=?")
    vs = append(vs, v)
  }
  err := gdb.QueryRow(
    `
    SELECT
      id, name, description, create_ts, last_update_ts
    FROM
      product
    WHERE
    ` + strings.Join(ks, " and "),
    vs...,
  ).Scan(&m.Id, &m.Name, &m.Description, &m.CreateTs, &m.LastUpdateTs)

  switch {
  case err == sql.ErrNoRows:
    return nil
  case err != nil:
    panic(err)
  }

  return m
}

func DeleteProductModel(id int64) {
  stmt, err := gdb.Prepare(`
    DELETE FROM
      product
    WHERE
      id=?
  `)
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  if _, err := stmt.Exec(id); err != nil {
    panic(err) 
  }
}

func UpdateProductModel(id int64, con map[string]interface{}) {
  ks := make([]string, 0)
  vs := make([]interface{}, 0)
  for k, v := range con {
    ks = append(ks, k+"=?")
    vs = append(vs, v)
  }

  vs = append(vs, id)

  stmt, err := gdb.Prepare(
    "UPDATE product SET " +
    strings.Join(ks, ", ") +
    " WHERE id=?",
  )
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  if _, err := stmt.Exec(vs...); err != nil {
    panic(err)
  }
}

func ProductModelCheck(pm *ProductModel) error {
  // TBD
  return nil
}

func ProductModelUpdateCheck(con *map[string]interface{}) error {
  // TBD
  // check & filter, filter invalid key
  return nil
}

/////////////////////////////////////////////////////////////////////////////////

type ReleaseModel struct {
  Id          int64   `json:"id"`
  ProductId   int64   `json:"product_id"`
  Version     string  `json:"version"`
  Description string  `json:"description"`
  CreateTs    string  `json:"create_ts"`
  LastUpdateTs string `json:"last_update_ts"`
}

func ListReleaseModel(con map[string]interface{}) []*ReleaseModel {
  ks := make([]string, 0)
  vs := make([]interface{}, 0)
  for k, v := range con {
    ks = append(ks, k+"=?")
    vs = append(vs, v)
  }

  where := ""
  if con != nil {
    where = "WHERE " + strings.Join(ks, " and ")
  }

  rows, err := gdb.Query(
    `
    SELECT
      id, product_id, version, description, create_ts, last_update_ts
    FROM
      release_t
    ` + where,
    vs..., 
  )
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  l := make([]*ReleaseModel, 0)
  for rows.Next() {
    m := new(ReleaseModel)
    if err := rows.Scan(
      &m.Id, &m.ProductId, &m.Version,
      &m.Description, &m.CreateTs, &m.LastUpdateTs,
    ); err != nil {
      panic(err)
    }

    l = append(l, m)
  }

  return l
}

func CreateReleaseModel(pid int64, m *ReleaseModel) int64 {
  stmt, err := gdb.Prepare(`
    INSERT INTO
      release_t(product_id, version, description, create_ts, last_update_ts)
    VALUES(?, ?, ?, ?, ?)
  `)
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  m.ProductId = pid
  result, err := stmt.Exec(
    &m.ProductId, &m.Version, &m.Description,
    &m.CreateTs, &m.LastUpdateTs,
  )
  if err != nil {
    panic(err)
  }

  id, err := result.LastInsertId()
  if err != nil {
    panic(err)
  }

  return id
}

func GetReleaseModel(con map[string]interface{}) *ReleaseModel {
  m := new(ReleaseModel)
  ks := make([]string, 0)
  vs := make([]interface{}, 0)
  for k, v := range con {
    ks = append(ks, k+"=?")
    vs = append(vs, v)
  }

  err := gdb.QueryRow(
    `
    SELECT
      id, product_id, version, description, create_ts, last_update_ts
    FROM
      release_t
    WHERE
    ` + strings.Join(ks, " and "),
    vs...,
  ).Scan(
    &m.Id, &m.ProductId, &m.Version,
    &m.Description, &m.CreateTs, &m.LastUpdateTs,
  )

  switch {
  case err == sql.ErrNoRows:
    return nil
  case err != nil:
    panic(err)
  }

  return m
}

func DeleteReleaseModel(pid, rid int64) {
  stmt, err := gdb.Prepare(`
    DELETE FROM
      release_t
    WHERE
      product_id=? and id=?
  `)
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  if _, err := stmt.Exec(pid, rid); err != nil {
    panic(err) 
  }
}

func UpdateReleaseModel(pid, rid int64, con map[string]interface{}) {
  ks := make([]string, 0)
  vs := make([]interface{}, 0)
  for k, v := range con {
    ks = append(ks, k+"=?")
    vs = append(vs, v)
  }

  vs = append(vs, pid)
  vs = append(vs, rid)

  stmt, err := gdb.Prepare(
    "UPDATE release_t SET " +
    strings.Join(ks, ", ") +
    " WHERE product_id=? and id=?",
  )
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  if _, err := stmt.Exec(vs...); err != nil {
    panic(err)
  }
}

func ReleaseModelCheck(pm *ReleaseModel) error {
  // TBD
  return nil
}

func ReleaseModelUpdateCheck(con *map[string]interface{}) error {
  // TBD
  // check & filter, filter invalid key
  return nil
}
