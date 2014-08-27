package tests

import(
  "testing"
  "appengine"
  . "todolist/models"
  "appengine/aetest"
  "appengine/datastore"
)

func TestUserCreation(t *testing.T){
  c, _ := aetest.NewContext(nil)

  defer c.Close()
  user := MyUser{
    0,
    "Carlos",
    "coneramu@gmail.com",
  }

  key, _ := user.Save(c)

  userFromDb, _ := GetUser(c, key)
  if want, val := "Carlos", userFromDb.Name; val != want{
    t.Errorf("the name should be %s and was %s", want, val)
  }

}

func TestUserListing(t *testing.T){
  c, _ := aetest.NewContext(nil)

  defer c.Close()

  user1 := MyUser{0, "Carlos", "coneramu@gmail.com"}
  user2 := MyUser{0, "Heriberto", "somemail@gmail.com"}
  user3 := MyUser{0, "Test", "test@gmail.com"}

  user1.Save(c)
  user2.Save(c)
  user3.Save(c)

  users, err := ListUsers(c)

  if err != nil{
    t.Errorf("Error found: %s", err.Error())
  }

  if want, val := 3, len(users); want != val{
    t.Errorf("The amount of users should be %s, %s found", want, val)
  }
}

func TestUpdateUser(t *testing.T){
  c, _ := aetest.NewContext(nil)

  defer c.Close()
  user := MyUser{
    0,
    "Carlos",
    "something@gmail.com",
  }

  key, _:= user.Save(c)

  userFromDb, _ := GetUser(c, key)
  userFromDb.Email = "coneramu@gmail.com"
  err := user.Update(c)
  if err != nil {
    t.Errorf("Error found: %s", err.Error())
  }

  if want, val:= "coneramu@gmail.com", userFromDb.Email; want != val {
    t.Errorf("The User was not updated")
  }
}

func TestDeleteUser(t *testing.T) {
  c, _ := aetest.NewContext(nil)

  defer c.Close()
  user := MyUser{
    0,
    "Carlos",
    "something@gmail.com",
  }

  user.Save(c)

  users, err := ListUsers(c)

  if want, val := 1, len(users); want != val{
    t.Errorf("The amount of users should be %s, %s found", want, val)
  }

  err = user.Delete(c)
  if err != nil {
    t.Errorf("An error occured during the user deletion")
  }

  users, err = ListUsers(c)
  if want, val := 0, len(users); want != val{
    t.Errorf("The amount of users should be %s, %s found", want, val)
  }
}

func TestDeleteAllUsers(t *testing.T) {
  c, _ := aetest.NewContext(nil)

  defer c.Close()

  user1 := MyUser{0, "Carlos", "coneramu@gmail.com"}
  user2 := MyUser{0, "Heriberto", "somemail@gmail.com"}
  user3 := MyUser{0, "Test", "test@gmail.com"}

  user1.Save(c)
  user2.Save(c)
  user3.Save(c)

  users, err := ListUsers(c)
  if want, val := 3, len(users); want != val{
    t.Errorf("The amount of users should be %s, %s found", want, val)
  }

  err = DeleteAllUsers(c)
  if err != nil {
    t.Errorf("An error occured during the multi users deletion %s", err)
  }

  users, err = ListUsers(c)
  if want, val := 0, len(users); want != val{
    t.Errorf("The amount of users should be %s, %s found", want, val)
  }

}

//------------------Lists---------------------

func populateListTest(c appengine.Context) *datastore.Key{
  user := MyUser{0, "Name", "email@email.com"}
  parentKey, _ := user.Save(c)

  for i := 0; i < 101 ; i++ {
    listName := "Name [" + string(i) + "]"
    list := List{ 0, listName, "description"}
    list.Save(c, parentKey)
  }

  return parentKey
}

func TestListCreation(t *testing.T){

  c, _ := aetest.NewContext(nil)
  defer c.Close()

  user := MyUser{0, "Name", "email@email.com"}
  key, _ := user.Save(c)

  list := List{0, "Name", "description"}

  listKey, err := list.Save(c, key)

  listaDB, err := GetList(c, listKey)

  if err != nil {
    t.Errorf("Error found %s", err)
  }

  if listaDB.Name != "Name" {
    t.Errorf("Name doesn't match, got: %s expected %s", listaDB.Name, "Name")
  }

}

func TestListUpdate(t *testing.T){
  c, _ := aetest.NewContext(nil)
  defer c.Close()

  user := MyUser{0, "Name", "email@email.com"}
  parentKey, _ := user.Save(c)

  list := List{0, "Name", "description"}

  listKey, _ := list.Save(c, parentKey)

  listaDB, _ := GetList(c, listKey)

  listaDB.Name = "New Name"

  listaDB.Save(c, parentKey)

  listaDB, _ = GetList(c, listKey)

  if want, listName := listaDB.Name, "New Name"; want != listName{
    t.Errorf("Name doesn't match, got: %s expected %s", want, listName)
  }


}

func TestGetAll(t *testing.T){
  c, _ := aetest.NewContext(nil)

  defer c.Close()
  parentKey := populateListTest(c)
  lists := ListUserLists(c, parentKey)

  if want, totList :=  101, len(lists); want != totList {
    t.Errorf("len(list): %s is different to %s", len(lists), want)
  }

  //for _, list := range lists {
    //t.Error(list)
  //}
}

