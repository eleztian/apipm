package mongodb

import "gopkg.in/mgo.v2"

var (
	mgoSession *mgo.Session
	dbName     string
)

// Init init mongodb.url is used to connect mongodb.
func Init(url string, db string) error {
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	err = session.Ping()
	if err != nil {
		return err
	}
	mgoSession = session
	dbName = db
	return nil
}

type DB struct {
	session    *mgo.Session
	collection string
}

func NewCollection(collection string) *DB {
	if mgoSession == nil {
		panic("Please init mongodb")
	}
	return &DB{session: mgoSession.Copy(), collection: collection}
}

func (d *DB) Create() error {
	return mgoSession.DB(dbName).C(d.collection).Create(nil)
}

func (d *DB) Insert(docs ...interface{}) error {
	return mgoSession.DB(dbName).C(d.collection).Insert(docs...)
}

func (d *DB) Find(query string, resPoint interface{}) error {
	return d.session.DB(dbName).C(d.collection).Find(query).All(resPoint)
}

func (d *DB) FindByID(id string, resPoint interface{}) error {
	return d.session.DB(dbName).C(d.collection).FindId(id).One(resPoint)
}

func (d *DB) Update(selector interface{}, update interface{}) error {
	return d.session.DB(dbName).C(d.collection).Update(selector, update)
}

func (d *DB) UpdateByID(id string, update interface{}) error {
	return d.session.DB(dbName).C(d.collection).UpdateId(id, update)
}

func (d *DB) Close() {
	if d.session != nil {
		d.session.Close()
		d.session = nil
	}
}
