package models

import (
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

type (
	Accounts struct {
		Id  string `bson:"_id"`
		created  string
		updated  string
		email string
		wcd_guid string
		name  *Name
		country string
		status string
	}

 	Name  struct {
		full string
		first string
		last string
	}
)

//func init() {
//	orm.RegisterModel(new(Accounts))
//}

// AddAccounts insert a new Accounts into database and returns
// last inserted Id on success.
//func AddAccounts(m *Accounts) (id int64, err error) {
//	o := orm.NewOrm()
//	id, err = o.Insert(m)
//	return
//}

// GetAccountsById retrieves Accounts by Id. Returns error if
// Id doesn't exist
func GetAccountsById(id string) (v *Accounts) {
	session := singleton.sessions["dev"].DB("oz_development")
	var results Accounts{}
	c := session.C("accounts")
	err := c.Find(bson.M{"_id": id}).One(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
	return &results[0]
}

// GetAllAccounts retrieves all Accounts matches certain condition. Returns empty list if
// no records exist
//func GetAllAccounts(query map[string]string, fields []string, sortby []string, order []string,
//	offset int64, limit int64) (ml []interface{}, err error) {
//	o := orm.NewOrm()
//	qs := o.QueryTable(new(Accounts))
//	// query k=v
//	for k, v := range query {
//		// rewrite dot-notation to Object__Attribute
//		k = strings.Replace(k, ".", "__", -1)
//		qs = qs.Filter(k, v)
//	}
	// order by:
//	var sortFields []string
//	if len(sortby) != 0 {
//		if len(sortby) == len(order) {
//			// 1) for each sort field, there is an associated order
//			for i, v := range sortby {
//				orderby := ""
//				if order[i] == "desc" {
//					orderby = "-" + v
//				} else if order[i] == "asc" {
//					orderby = v
//				} else {
//					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
//				}
//				sortFields = append(sortFields, orderby)
//			}
//			qs = qs.OrderBy(sortFields...)
//		} else if len(sortby) != len(order) && len(order) == 1 {
//			// 2) there is exactly one order, all the sorted fields will be sorted by this order
//			for _, v := range sortby {
//				orderby := ""
//				if order[0] == "desc" {
//					orderby = "-" + v
//				} else if order[0] == "asc" {
//					orderby = v
//				} else {
//					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
//				}
//				sortFields = append(sortFields, orderby)
//			}
//		} else if len(sortby) != len(order) && len(order) != 1 {
//			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
//		}
//	} else {
//		if len(order) != 0 {
//			return nil, errors.New("Error: unused 'order' fields")
//		}
//	}
//
//	var l []Accounts
//	qs = qs.OrderBy(sortFields...)
//	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
//		if len(fields) == 0 {
//			for _, v := range l {
//				ml = append(ml, v)
//			}
//		} else {
//			// trim unused fields
//			for _, v := range l {
//				m := make(map[string]interface{})
//				val := reflect.ValueOf(v)
//				for _, fname := range fields {
//					m[fname] = val.FieldByName(fname).Interface()
//				}
//				ml = append(ml, m)
//			}
//		}
//		return ml, nil
//	}
//	return nil, err
//}
//
//// UpdateAccounts updates Accounts by Id and returns error if
//// the record to be updated doesn't exist
//func UpdateAccountsById(m *Accounts) (err error) {
//	o := orm.NewOrm()
//	v := Accounts{Id: m.Id}
//	// ascertain id exists in the database
//	if err = o.Read(&v); err == nil {
//		var num int64
//		if num, err = o.Update(m); err == nil {
//			fmt.Println("Number of records updated in database:", num)
//		}
//	}
//	return
//}
//
//// DeleteAccounts deletes Accounts by Id and returns error if
//// the record to be deleted doesn't exist
//func DeleteAccounts(id int64) (err error) {
//	o := orm.NewOrm()
//	v := Accounts{Id: id}
//	// ascertain id exists in the database
//	if err = o.Read(&v); err == nil {
//		var num int64
//		if num, err = o.Delete(&Accounts{Id: id}); err == nil {
//			fmt.Println("Number of records deleted in database:", num)
//		}
//	}
//	return
//}
