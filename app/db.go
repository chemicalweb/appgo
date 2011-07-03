package app

import(
  "rand"
  "time"
  "reflect"
  "strings"
  "strconv"
  "appengine"
  "appengine/datastore"
)


//  DATASTORE MANAGER  -------------------------------------
type DSM struct {
  Context  appengine.Context
}

// USE:  customer.created = db.Now()
func (db *DSM) Now() int64{
  return time.Seconds()
}

// USE:  customer.id = db.Sequence()
func (db *DSM) Sequence() string {
  var nnn = rand.Intn(899) + 100  // from 100 to 999
  return strconv.Itoa64(time.Seconds())+strconv.Itoa(nnn)
}

// USE:  customer.key = db.NewKey("Customer")
func (db *DSM) NewKey(kind string) *datastore.Key{
  return datastore.NewIncompleteKey(kind)
}

// USE:  customer.keyname = db.KeyName("Customer","ALFKI")
func (db *DSM) KeyName(kind string, id string) *datastore.Key{
  return datastore.NewKey(kind, id, 0, nil)
}

// USE:  qry := db.Query("Customer").Order("-Date").Limit(10)
func (db *DSM) Query(entity string) *datastore.Query{
  return datastore.NewQuery(entity)
}

// USE:  db.Select(qry,&customers)
func (db *DSM) Select(qry *datastore.Query, recs interface{}) bool {
  _, err := qry.GetAll(db.Context, recs)
  return err==nil
}

// USE:  keys := db.SelectKeys(qry)
func (db *DSM) SelectKeys(qry *datastore.Query) ([]string, bool) {
  recs := new(interface{})
  keys, err := qry.KeysOnly().GetAll(db.Context, recs)
  skeys := []string{}
  if err==nil{
    for k := range keys {
      skeys = append(skeys,keys[k].StringID())
    }
  }
  return skeys, err==nil
}

// USE:  ok := db.Get("ALFKI",&customer)
func (db *DSM) Get(id string, rec interface{}) bool {
  kin := db.getKind(rec)
  key := db.KeyName(kin, id)
  err := datastore.Get(db.Context, key, rec)
  return err==nil
}

// USE:  ok := db.GetByKey(key, &customer)
func (db *DSM) GetByKey(key *datastore.Key, rec interface{}) bool {
  err := datastore.Get(db.Context, key, rec)
  return err==nil
}

// USE:  ok := db.New(&customer)
func (db *DSM) New(rec interface{}) bool {
  kin := db.getKind(rec)
  key := datastore.NewIncompleteKey(kin)
  _, err := datastore.Put(db.Context, key, rec)
  return err==nil
}

// USE:  ok := db.Put("ALFKI",&customer)
func (db *DSM) Put(id string, rec interface{}) bool {
  kin := db.getKind(rec)
  key := db.KeyName(kin, id)
  _, err := datastore.Put(db.Context, key, rec)
  return err==nil
}

// USE:  ok := db.PutByKey(key, &customer)
func (db *DSM) PutByKey(key *datastore.Key, rec interface{}) bool {
  _, err := datastore.Put(db.Context, key, rec)
  return err==nil
}

// USE: DB:Delete("Customer","ALFKI")
func (db *DSM) Delete(kind string, id string) bool {
  key := db.KeyName(kind, id)
  err := datastore.Delete(db.Context, key)
  return err==nil
}

// USE: DB:Delete(key)
func (db *DSM) DeleteByKey(key *datastore.Key) bool {
  err := datastore.Delete(db.Context, key)
  return err==nil
}

// INTERNAL USE:  kind := getKind(rec)
func (db *DSM) getKind(rec interface{}) string {
  var kind string
  typ := reflect.TypeOf(rec).Elem().String()
  ind := strings.LastIndex(typ,".")
  if ind<0 {
    kind = typ
  } else {
    kind = typ[ind+1:]
  }
  return kind
}
