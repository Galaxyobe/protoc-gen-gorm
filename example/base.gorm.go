// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base.proto

package example

import context "context"
import github_com_jinzhu_gorm "github.com/jinzhu/gorm"
import time "time"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/galaxyobe/protoc-gen-gorm/proto"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// new Base GORM controller with gorm.DB
func (m *Base) GORMController(db *github_com_jinzhu_gorm.DB) *BaseGORMController {
	return &BaseGORMController{
		DB: db,
		m:  m,
	}
}

// new Base GORM controller with gorm.DB
func NewBaseGORMController(db *github_com_jinzhu_gorm.DB) *BaseGORMController {
	return &BaseGORMController{DB: db, m: new(Base)}
}

type BaseGORMController struct {
	DB *github_com_jinzhu_gorm.DB
	m  *Base
}

func (g *BaseGORMController) M(m *Base) {
	g.m = m
}

func (g *BaseGORMController) M() *Base {
	return g.m
}

func (g *BaseGORMController) AutoMigrate() {
	g.DB.AutoMigrate(g.m)
}

func (g *BaseGORMController) Begin() error {
	if db := g.DB.Begin(); db.Error != nil {
		return db.Error
	} else {
		g.DB = db
	}

	return nil
}

func (g *BaseGORMController) Rollback() *github_com_jinzhu_gorm.DB {
	return g.DB.Rollback()
}

func (g *BaseGORMController) Commit() *github_com_jinzhu_gorm.DB {
	return g.DB.Commit()
}

func (g *BaseGORMController) Create() *github_com_jinzhu_gorm.DB {
	now := time.Now().Unix()
	g.m.CreateAt = now
	g.m.UpdateAt = now

	return g.DB.Create(g.m)
}

func (g *BaseGORMController) Delete() *github_com_jinzhu_gorm.DB {
	if g.m.Uuid == 0 {
		g.DB.Error = errors.New("the value of Uuid is not expected to be 0")
		return g.DB
	}

	return g.DB.Delete(g.m)
}

func (g *BaseGORMController) SoftDelete() *github_com_jinzhu_gorm.DB {
	g.m.DeleteAt = time.Now().Unix()

	return g.DB.Model(g.m).Select("DeleteAt").Updates(g.m)
}

func (g *BaseGORMController) Update() *github_com_jinzhu_gorm.DB {
	g.m.UpdateAt = time.Now().Unix()

	return g.DB.Model(g.m).Omit("Uuid", "CreateAt", "DeleteAt").Updates(g.m)
}

func (g *BaseGORMController) First() (*Base, error) {
	db := g.DB.First(g.m)

	return g.m, db.Error
}

// when where is empty string will be ignored
// when order is empty is null: nil,-1,"" will be ignored
// when offset and limit is null: nil,-1,"" will be ignored
func (g *BaseGORMController) Find(where, limit, offset, order interface{}) ([]*Base, error) {
	var array []*Base

	db := g.DB.Where(where).Order(order).Limit(limit).Offset(offset).Find(&array)

	return array, db.Error
}

// when where is empty string will be ignored
// when order is empty is null: nil,-1,"" will be ignored
// when offset and limit is null: nil,-1,"" will be ignored
func (g *BaseGORMController) Count(where, limit, offset, order interface{}) (int64, error) {
	var count int64 = 0

	db := g.DB.Model(g.m).Where(where).Order(order).Limit(limit).Offset(offset).Count(&count)

	return count, db.Error
}
