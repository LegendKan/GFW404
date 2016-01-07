package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Bill struct {
	Id         int       `orm:"column(id);auto"`
	Accountid  int       `orm:"column(accountid)"`
	Price      float64   `orm:"column(price);digits(12);decimals(2)"`
	Createtime time.Time `orm:"column(createtime);auto_now_add;type(datetime)"`
	Expiretime time.Time `orm:"column(expiretime);type(datetime)"`
	Ispaid     int8      `orm:"column(ispaid)"`
	Active     int8      `orm:"column(active)"`
	Payno      string    `orm:"column(payno)";size(255)`
}

func init() {
	orm.RegisterModel(new(Bill))
}

// AddBill insert a new Bill into database and returns
// last inserted Id on success.
func AddBill(m *Bill) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBillById retrieves Bill by Id. Returns error if
// Id doesn't exist
func GetBillById(id int) (v *Bill, err error) {
	o := orm.NewOrm()
	v = &Bill{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBill retrieves all Bill matches certain condition. Returns empty list if
// no records exist
func GetAllBill(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []Bill, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Bill))
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

	var l []Bill
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		// if len(fields) == 0 {
		// 	for _, v := range l {
		// 		ml = append(ml, v)
		// 	}
		// } else {
		// 	// trim unused fields
		// 	for _, v := range l {
		// 		m := make(map[string]interface{})
		// 		val := reflect.ValueOf(v)
		// 		for _, fname := range fields {
		// 			m[fname] = val.FieldByName(fname).Interface()
		// 		}
		// 		ml = append(ml, m)
		// 	}
		// }
		return l, nil
	}
	return nil, err
}

func GetAllUnpaidBills(uid string) ([]orm.Params, error) {
	var maps []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw("SELECT bill.id, bill.accountid, bill.price, bill.createtime, bill.expiretime FROM account,bill WHERE account.userid = ? and bill.accountid=account.id and TO_DAYS(NOW()) - TO_DAYS(bill.expiretime) <= 5 and bill.ispaid=0 and bill.active=1 order by bill.id desc", uid).Values(&maps)
	if err!=nil {
		return nil,err
	}
	return maps,nil
}

func GetUnpaidBillByAccount(account Account) (Bill, bool) {
	var b Bill
	o := orm.NewOrm()
	b.Accountid=account.Id
	b.Active=1
	b.Expiretime=account.Expiretime
	err:= o.Read(&b)
	if err == orm.ErrNoRows {
	    fmt.Println("Err No Rows")
	    return b, false
	} else if err == orm.ErrMissPK {
	    fmt.Println("Err Miss PK")
	    return b, false
	} else {
	    return b, true
	}
}

// UpdateBill updates Bill by Id and returns error if
// the record to be updated doesn't exist
func UpdateBillById(m *Bill) (err error) {
	o := orm.NewOrm()
	v := Bill{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBill deletes Bill by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBill(id int) (err error) {
	o := orm.NewOrm()
	v := Bill{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Bill{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
