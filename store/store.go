package store

type DataBase struct {
	Data map[string]string
}

func (db *DataBase) Set(key, value string) {
	db.Data[key] = value
}

func (db *DataBase) Get(key string) string {
	return db.Data[key]
}

func Create() *DataBase {
	db := &DataBase{
		Data: make(map[string]string),
	}
	return db
}
