package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type Account struct {
	Id          int       `orm:"column(id);auto"`
	Serverid    *Server   `orm:"column(serverid);rel(fk)"`
	Containerid string    `orm:"column(containerid);size(255)"`
	Port        int       `orm:"column(port)"`
	Password    string    `orm:"column(password);size(255)"`
	Userid      *User     `orm:"column(userid);rel(fk)"`
	Createtime  time.Time `orm:"column(createtime);auto_now_add;type(datetime)"`
	Expiretime  time.Time `orm:"column(expiretime);type(datetime);null"`
	Cycle       int8      `orm:"column(cycle)"`
	Active      int8      `orm:"column(active)"`
	Firstprice	float64   `orm:"column(firstprice);null;digits(12);decimals(2)"`
	Recurringprice float64 `orm:"column(recurringprice);null;digits(12);decimals(2)"`
}

type SimpleAccount struct {
	Id              int
	Containerid     string
	Port        	int
	Password        string
}

func init() {
	orm.RegisterModel(new(Account))
}

// AddAccount insert a new Account into database and returns
// last inserted Id on success.
func AddAccount(m *Account) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAccountById retrieves Account by Id. Returns error if
// Id doesn't exist
func GetAccountById(id int) (v *Account, err error) {
	o := orm.NewOrm()
	v = &Account{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAccountDetailById(id int) (v orm.Params, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT account.id, account.userid, account.port, account.password, account.active, server.ip, server.title FROM account, server WHERE account.id = ? and account.serverid=server.id", strconv.Itoa(id)).Values(&maps)
	if err!=nil{
		return nil, err
	}
	if num>0{
		return maps[0],nil
	}
	return nil, errors.New("404")
}


// GetAllAccount retrieves all Account matches certain condition. Returns empty list if
// no records exist
func GetAllAccount(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Account))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Account
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// GetAllActiveAccount retrieves all Server matches certain condition. Returns empty list if
// no records exist
func GetAllActiveAccount(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (l []Account, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Account))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	//var l []Account
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		return l, nil
	}
	return nil, err
}

//返回全部active的account
func GetAllActiveAccountNew() ([]Account, error) {
	fmt.Println("Get all active accounts")
	o := orm.NewOrm()
	var accounts []Account
	num, err := o.QueryTable("account").Filter("active", 1).All(&accounts)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return accounts, err
}

func GetAllAccountByUserId(uid string) ([]orm.Params, error){
	var maps []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw("SELECT account.id, account.expiretime, account.cycle, account.active, account.recurringprice, server.title FROM account,server WHERE account.userid = ? and server.id=account.serverid order by account.id desc", uid).Values(&maps)
	if err!=nil {
		return nil,err
	}
	return maps,nil
}

// UpdateAccount updates Account by Id and returns error if
// the record to be updated doesn't exist
func UpdateAccountById(m *Account) (err error) {
	o := orm.NewOrm()
	v := Account{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAccount deletes Account by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAccount(id int) (err error) {
	o := orm.NewOrm()
	v := Account{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Account{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
