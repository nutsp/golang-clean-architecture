package datasource

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type DB interface {
	Debug() DB
	WithContext(ctx context.Context) DB
	Create(value interface{}) DB
	Where(query interface{}, args ...interface{}) DB
	Find(out interface{}, where ...interface{}) DB
	Updates(values interface{}) DB
	Error() error
	Commit() DB
	Rollback() DB
	Begin() DB
	Save(value interface{}) DB
	Delete(value interface{}, where ...interface{}) DB
	First(out interface{}, where ...interface{}) DB
	Last(out interface{}, where ...interface{}) DB
	Select(query interface{}, args ...interface{}) DB
	Pluck(column string, value interface{}) DB
	RawQuery(query string, args ...interface{}) (*sql.Rows, error)
	RawRow(query string, args ...interface{}) *sql.Row
}

type GormDB struct {
	*gorm.DB
}

func (db *GormDB) Debug() DB {
	return &GormDB{db.DB.Debug()}
}

func (db *GormDB) WithContext(ctx context.Context) DB {
	return &GormDB{db.DB.WithContext(ctx)}
}

func (db *GormDB) Create(value interface{}) DB {
	return &GormDB{db.DB.Create(value)}
}

func (db *GormDB) Where(query interface{}, args ...interface{}) DB {
	return &GormDB{db.DB.Where(query, args...)}
}

func (db *GormDB) Find(out interface{}, where ...interface{}) DB {
	return &GormDB{db.DB.Find(out, where...)}
}

func (db *GormDB) Updates(values interface{}) DB {
	return &GormDB{db.DB.Updates(values)}
}

func (db *GormDB) Error() error {
	return db.DB.Error
}

func (db *GormDB) Commit() DB {
	return &GormDB{db.DB.Commit()}
}

func (db *GormDB) Rollback() DB {
	return &GormDB{db.DB.Rollback()}
}

func (db *GormDB) Begin() DB {
	return &GormDB{db.DB.Begin()}
}

func (db *GormDB) Save(value interface{}) DB {
	return &GormDB{db.DB.Save(value)}
}

func (db *GormDB) Delete(value interface{}, where ...interface{}) DB {
	return &GormDB{db.DB.Delete(value, where...)}
}

func (db *GormDB) First(out interface{}, where ...interface{}) DB {
	return &GormDB{db.DB.First(out, where...)}
}

func (db *GormDB) Last(out interface{}, where ...interface{}) DB {
	return &GormDB{db.DB.Last(out, where...)}
}

func (db *GormDB) Select(query interface{}, args ...interface{}) DB {
	return &GormDB{db.DB.Select(query, args...)}
}

func (db *GormDB) Pluck(column string, value interface{}) DB {
	return &GormDB{db.DB.Pluck(column, value)}
}

func (db *GormDB) RawQuery(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Raw(query, args...).Rows()
}

func (db *GormDB) RawRow(query string, args ...interface{}) *sql.Row {
	return db.DB.Raw(query, args...).Row()
}
