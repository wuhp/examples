package main

import (
  "net/http"
  "regexp"
  "encoding/json"
  "io/ioutil"
  "strconv"
  "errors"
  "time"
)

/////////////////////////////////////////////////////////////////////////////////

type Router struct {
  Pattern string
  Method string
  Func func(http.ResponseWriter, *http.Request)
}

func Api(w http.ResponseWriter, r *http.Request) {
  // Handler server internal error by throwing exception
  defer func() {
    err := recover()
    if err != nil {
      apiInternalError(w, r, err.(error))
    }
  }()

  routers := []Router{
    { "/api/product$",                      "GET",     listProduct    },
    { "/api/product$",                      "POST",    createProduct  },
    { `/api/product/(\d+)$`,                "GET",     getProduct     },
    { `/api/product/(\d+)$`,                "PUT",     updateProduct  },
    { `/api/product/(\d+)$`,                "DELETE",  deleteProduct  },
    { `/api/product/(\d+)/release$`,        "GET",     listRelease    },
    { `/api/product/(\d+)/release$`,        "POST",    createRelease  },
    { `/api/product/(\d+)/release/(\d+)$`,  "GET",     getRelease     },
    { `/api/product/(\d+)/release/(\d+)$`,  "PUT",     updateRelease  },
    { `/api/product/(\d+)/release/(\d+)$`,  "DELETE",  deleteRelease  },
  }

  for _, rt := range routers {
    match, _ := regexp.MatchString(rt.Pattern, r.URL.Path)
    if match && r.Method == rt.Method {
      rt.Func(w, r)
      return
    }  
  }

  apiNotFound(w, r)
}

/////////////////////////////////////////////////////////////////////////////////

func parseProductId(path string) int64 {
  id, _ := strconv.ParseInt(
    regexp.MustCompile(
      `/api/product/(\d+)`,
    ).FindStringSubmatch(path)[1],
    10, 64,
  )
  return id
}

func parseReleaseId(path string) int64 {
  id, _ := strconv.ParseInt(
    regexp.MustCompile(
      `/api/product/\d+/release/(\d+)`,
    ).FindStringSubmatch(path)[1],
    10, 64,
  )
  return id
}

func now() string {
  return time.Now().Format("2006-01-02 15:04:05")
}

func isProductIdExist(id int64) bool {
  cond := make(map[string]interface{})
  cond["id"] = id
  return GetProductModel(cond) != nil
}

func isProductUniqueFieldsExist(name string) bool {
  cond := make(map[string]interface{})
  cond["name"] = name
  return GetProductModel(cond) != nil
}

func isReleaseIdExist(pid, rid int64) bool {
  cond := make(map[string]interface{})
  cond["product_id"] = pid
  cond["id"] = rid
  return GetReleaseModel(cond) != nil
}

func isReleaseUniqueFieldsExist(pid int64, version string) bool {
  cond := make(map[string]interface{})
  cond["product_id"] = pid
  cond["version"] = version
  return GetReleaseModel(cond) != nil
}

/////////////////////////////////////////////////////////////////////////////////

func apiNotFound(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusNotFound)
}

func apiBadRequest(w http.ResponseWriter, r *http.Request, e error) {
  w.WriteHeader(http.StatusBadRequest)
  w.Write([]byte(e.Error()))
}

func apiInternalError(w http.ResponseWriter, r *http.Request, e error) {
  w.WriteHeader(http.StatusInternalServerError)
  w.Write([]byte(e.Error()))
}

/////////////////////////////////////////////////////////////////////////////////

func listProduct(w http.ResponseWriter, r *http.Request) {
  l := ListProductModel(nil)
  body, _ := json.Marshal(l)
  w.WriteHeader(http.StatusOK)
  w.Write(body)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
  rbody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    apiBadRequest(w, r, err)
    return
  }

  m := new(ProductModel)
  if err := json.Unmarshal(rbody, m); err != nil {
    apiBadRequest(w, r, err)
    return
  }

  if err := ProductModelCheck(m); err != nil {
    apiBadRequest(w, r, err)
    return
  }

  if isProductUniqueFieldsExist(m.Name) {
    apiBadRequest(w, r, errors.New("Duplicated Product"))
    return
  }

  m.CreateTs = now()
  m.LastUpdateTs = m.CreateTs
  id := CreateProductModel(m)
  wbody, _ := json.Marshal(map[string]int64{"id":id})
  w.WriteHeader(http.StatusOK)
  w.Write(wbody)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
  cond := make(map[string]interface{})
  cond["id"] = parseProductId(r.URL.Path)

  m := GetProductModel(cond)
  if m == nil {
    apiNotFound(w, r)
    return
  }

  wbody, _ := json.Marshal(m)
  w.WriteHeader(http.StatusOK)
  w.Write(wbody)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
  id := parseProductId(r.URL.Path)
  if !isProductIdExist(id) {
    apiNotFound(w, r)
    return
  }

  DeleteProductModel(id)
  w.WriteHeader(http.StatusOK)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
  id := parseProductId(r.URL.Path)
  if !isProductIdExist(id) {
    apiNotFound(w, r)
    return
  }

  rbody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    apiBadRequest(w, r, err)
    return
  }

  in := make(map[string]interface{})
  if err := json.Unmarshal(rbody, &in); err != nil {
    apiBadRequest(w, r, err)
    return
  }

  if err := ProductModelUpdateCheck(&in); err != nil {
    apiBadRequest(w, r, err)
    return
  }

  in["last_update_ts"] = now()
  UpdateProductModel(id, in)
  w.WriteHeader(http.StatusOK)
}

/////////////////////////////////////////////////////////////////////////////////

func listRelease(w http.ResponseWriter, r *http.Request) {
  pid := parseProductId(r.URL.Path)
  if !isProductIdExist(pid) {
    apiNotFound(w, r)
    return
  }

  l := ListReleaseModel(map[string]interface{}{"product_id":pid})
  body, _ := json.Marshal(l)
  w.WriteHeader(http.StatusOK)
  w.Write(body)
}

func createRelease(w http.ResponseWriter, r *http.Request) {
  pid := parseProductId(r.URL.Path)
  if !isProductIdExist(pid) {
    apiNotFound(w, r)
    return
  }

  rbody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    apiBadRequest(w, r, err)
    return
  }

  m := new(ReleaseModel)
  if err := json.Unmarshal(rbody, m); err != nil {
    apiBadRequest(w, r, err)
    return
  }

  if err := ReleaseModelCheck(m); err != nil {
    apiBadRequest(w, r, err)
    return
  }

  if isReleaseUniqueFieldsExist(pid, m.Version) {
    apiBadRequest(w, r, errors.New("Duplicated Release"))
    return
  }

  m.CreateTs = now()
  m.LastUpdateTs = m.CreateTs
  id := CreateReleaseModel(pid, m)
  wbody, _ := json.Marshal(map[string]int64{"id":id})
  w.WriteHeader(http.StatusOK)
  w.Write(wbody)
}

func getRelease(w http.ResponseWriter, r *http.Request) {
  cond := make(map[string]interface{})
  cond["product_id"] = parseProductId(r.URL.Path)
  cond["id"] = parseReleaseId(r.URL.Path)

  m := GetReleaseModel(cond)
  if m == nil {
    apiNotFound(w, r)
    return
  }

  wbody, _ := json.Marshal(m)
  w.WriteHeader(http.StatusOK)
  w.Write(wbody)
}

func deleteRelease(w http.ResponseWriter, r *http.Request) {
  pid := parseProductId(r.URL.Path)
  rid := parseReleaseId(r.URL.Path)

  if !isReleaseIdExist(pid, rid) {
    apiNotFound(w, r)
    return
  }

  DeleteReleaseModel(pid, rid)
  w.WriteHeader(http.StatusOK)
}

func updateRelease(w http.ResponseWriter, r *http.Request) {
  pid := parseProductId(r.URL.Path)
  rid := parseReleaseId(r.URL.Path)

  if !isReleaseIdExist(pid, rid) {
    apiNotFound(w, r)
    return
  }

  rbody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    apiBadRequest(w, r, err)
    return
  }

  in := make(map[string]interface{})
  if err := json.Unmarshal(rbody, &in); err != nil {
    apiBadRequest(w, r, err)
    return
  }

  if err := ReleaseModelUpdateCheck(&in); err != nil {
    apiBadRequest(w, r, err)
    return
  }

  in["last_update_ts"] = now()
  UpdateReleaseModel(pid, rid, in)
  w.WriteHeader(http.StatusOK)
}
