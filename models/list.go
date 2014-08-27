package models

import (
  "appengine"
  "appengine/datastore"
)

type List struct{
  Id int64
  Name string
  Description string
}

func (this *List) key(c appengine.Context, parentKey *datastore.Key) *datastore.Key{
  if this.Id == 0 {
    return datastore.NewIncompleteKey(c, LIST_NAME, parentKey)
  }
  return datastore.NewKey(c, LIST_NAME, "", this.Id, parentKey)
}

func (this *List) Save(c appengine.Context, parentKey *datastore.Key)(*datastore.Key, error){
  key := this.key(c, parentKey)
  finalKey, err := datastore.Put(c, key, this)
  this.Id = finalKey.IntID()
  return finalKey, err
}

func GetList(c appengine.Context, key *datastore.Key) (List, error){
  var list List
  err := datastore.Get(c, key, &list)
  if err != nil {
    return list, err
  }
  list.Id = key.IntID()
  return list, nil
}

func ListUserLists(c appengine.Context, parentKey *datastore.Key) []List{
  query := datastore.NewQuery(LIST_NAME).Ancestor(parentKey)
  lists := []List{}
  query.GetAll(c, &lists)
  return lists 
}
