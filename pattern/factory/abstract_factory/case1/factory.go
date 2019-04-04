/*
 Abstract Factory 抽象工厂模式：
          提供一个创建一系列相关或者相互依赖对象的接口，而无需指定他们具体的类
 个人想法：工厂模式和抽象工厂模式：感觉抽象工厂可以叫集团模式，工厂模式下，
         是一个工厂下，对产品的每一个具体生成分配不同的流水线；
         集团模式：在集团下，有不同的工厂，可以生成不同的产品，
         每个工厂生产出来的同一个型号产品具体细节是不一样
 作者：   HCLAC
 日期：   20170306
*/
package factory

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func (u *User) GetId() int {
	if u == nil {
		return -1
	}
	return u.Id
}

func (u *User) SetId(id int) {
	if u == nil {
		return
	}
	u.Id = id
}

func (u *User) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *User) SetName(name string) {
	if u == nil {
		return
	}
	u.Name = name
}

type Department struct {
	Id   int
	Name string
}

func (d *Department) GetId() int {
	if d == nil {
		return -1
	}
	return d.Id
}
func (d *Department) SetId(id int) {
	if d == nil {
		return
	}
	d.Id = id
}
func (d *Department) GetName() string {
	if d == nil {
		return ""
	}
	return d.Name
}
func (d *Department) SetName(name string) {
	if d == nil {
		return
	}
	d.Name = name
}

type IUser interface {
	Insert(*User)
	GetUser(int) *User
}

type SqlServerUser struct {
}

func (s *SqlServerUser) Insert(u *User) {
	if s == nil {
		return
	}
	fmt.Println("往SqlServer的User表中插入一条User", u)
}

func (s *SqlServerUser) GetUser(id int) (u *User) {
	if s == nil {
		return nil
	}
	u = &User{id, "hclacS"}
	fmt.Println("从SqlServer的User表中获取一条User", *u)
	return
}

type AccessUser struct {
}

func (s *AccessUser) Insert(u *User) {
	if s == nil {
		return
	}
	fmt.Println("往AccessUser的User表中插入一条User", *u)
}

func (s *AccessUser) GetUser(id int) (u *User) {
	if s == nil {
		return nil
	}
	u = &User{id, "hclacA"}
	fmt.Println("从AccessUser的User表中获取一条User", *u)
	return
}

type IDepartment interface {
	Insert(*Department)
	GetDepartment(int) *Department
}

type SqlServerDepartment struct {
}

func (s *SqlServerDepartment) Insert(d *Department) {
	if s == nil {
		return
	}
	fmt.Println("往SqlServer的Department表中插入一条Department", *d)
}

func (s *SqlServerDepartment) GetDepartment(id int) (u *Department) {
	if s == nil {
		return nil
	}
	u = &Department{id, "hclacDS"}
	fmt.Println("从SqlServer的Department表中获取一条Department", *u)
	return
}

type AccessDepartment struct {
}

func (s *AccessDepartment) Insert(u *Department) {
	if s == nil {
		return
	}
	fmt.Println("往AccessDepartment的Department表中插入一条Department", *u)
}

func (s *AccessDepartment) GetDepartment(id int) (u *Department) {
	if s == nil {
		return nil
	}
	u = &Department{id, "hclacDA"}
	fmt.Println("从AccessDepartment的Department表中获取一条Department", *u)
	return
}

type Ifactory interface {
	createUser() IUser
	createDepartment() IDepartment
}

type SqlServerFactory struct {
}

func (s *SqlServerFactory) CreateUser() IUser {
	if s == nil {
		return nil
	}
	u := &SqlServerUser{}
	return u
}

func (s *SqlServerFactory) CreateDepartment() IDepartment {
	if s == nil {
		return nil
	}
	u := &SqlServerDepartment{}
	return u
}

type AccessFactory struct {
}

func (s *AccessFactory) CreateUser() IUser {
	if s == nil {
		return nil
	}
	u := &AccessUser{}
	return u
}

func (s *AccessFactory) CreateDepartment() IDepartment {
	if s == nil {
		return nil
	}
	u := &AccessDepartment{}
	return u
}

type DataAccess struct {
	db string
}

func (d *DataAccess) CreateUser(db string) IUser {
	if d == nil {
		return nil
	}

	var u IUser

	if db == "sqlserver" {
		u = new(SqlServerUser)
	} else if db == "access" {
		u = new(AccessUser)
	}
	return u
}

func (d *DataAccess) CreateDepartment(db string) IDepartment {
	if d == nil {
		return nil
	}

	var u IDepartment

	if db == "sqlserver" {
		u = new(SqlServerDepartment)
	} else if db == "access" {
		u = new(AccessDepartment)
	}
	return u
}
