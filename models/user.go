package models

import (
  "appengine"
  "appengine/datastore"
)

type MyUser struct{
  Id int64
  Name string
  Email string
}

func (this *MyUser) key(c appengine.Context) *datastore.Key{
  if this.Id == 0 {
    return datastore.NewIncompleteKey(c, USER_NAME, getCatalogKey(c))
  }
  return datastore.NewKey(c, USER_NAME, "", this.Id, getCatalogKey(c))
}

func (this *MyUser) Save(c appengine.Context) (*datastore.Key, error){
  key := this.key(c)
  finalKey, err := datastore.Put(c, key, this)
  this.Id = finalKey.IntID()
  return finalKey, err
}

func GetUser(c appengine.Context, key *datastore.Key) (MyUser, error){
  var user MyUser
  err := datastore.Get(c, key, &user)
  if err != nil {
    return user, err
  }
  user.Id = key.IntID()
  return user, nil
}

func ListUsers(c appengine.Context) ([]MyUser, error){
  _, err, users := getUserKeys(c)
  if err != nil{
    return users, err
  }
  return users, nil
}

func (this *MyUser) Update(c appengine.Context) error {
  _, err :=datastore.Put(c, this.key(c), this)
  if err != nil {
    return err
  }

  return nil

}

func (this *MyUser) Delete(c appengine.Context)  error {
  return datastore.Delete(c, this.key(c))
}

func getUserKeys(c appengine.Context) ([]*datastore.Key , error, []MyUser){
  query := datastore.NewQuery(USER_NAME).Ancestor(getCatalogKey(c)).Limit(10)
  users := make([]MyUser,0,10)
  keys, err := query.GetAll(c,&users)
  return keys, err, users
}

func DeleteAllUsers(c appengine.Context) error {
  keys, err, _ := getUserKeys(c)
  if err != nil {
    return err
  }
  return datastore.DeleteMulti(c, keys)
}

