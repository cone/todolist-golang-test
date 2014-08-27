package models

import (
  "appengine"
  "appengine/datastore"
)
const CATALOG_NAME = "Catalog"
const TODOLIST_NAME = "todolist"
const USER_NAME = "User"
const LIST_NAME = "List"

func getParentKey(c appengine.Context, entityName string, schema string) *datastore.Key{
  return datastore.NewKey(c, entityName, schema, 0, nil)
}

func getCatalogKey(c appengine.Context) *datastore.Key{
  return getParentKey(c, CATALOG_NAME, TODOLIST_NAME)
}

