package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id              int       `orm:"column(id);auto"`
	StudentId       string    `orm:"column(student_id);size(16)"`
	StudentPassword string    `orm:"column(student_password);size(64)"`
	Password        string    `orm:"column(password);size(64)"`
	Phone           string    `orm:"column(phone);size(11)"`
	Photo           string    `orm:"column(photo);size(256);null"`
	Nickname        string    `orm:"column(nickname);size(64);null"`
	RealName        string    `orm:"column(real_name);size(64);null"`
	Sex             string    `orm:"column(sex);size(1);null"`
	Birth           time.Time `orm:"column(birth);type(date);null"`
	Email           string    `orm:"column(email);size(64);null"`
	School          string    `orm:"column(school);size(64);null"`
	SchoolPart      string    `orm:"column(school_part);size(64);null"`
	College         string    `orm:"column(college);size(64);null"`
	Dormitory       string    `orm:"column(dormitory);size(64);null"`
	Specialty       string    `orm:"column(specialty);size(64);null"`
	Education       string    `orm:"column(education);size(64);null"`
	SchoolYear      string    `orm:"column(school_year);size(4);null"`
	Identity        string    `orm:"column(identity);size(20);null"`
	Hobby           string    `orm:"column(hobby);size(512);null"`
	IsSingle        string    `orm:"column(is_single);size(1);null"`
	AncestralHome   string    `orm:"column(ancestral_home);size(64);null"`
	Lvl             int       `orm:"column(lvl)"`
	Exp             int       `orm:"column(exp)"`
	UserType        int8      `orm:"column(user_type)"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserByStudentId retrieves User by StudentId. Returns error if
// any error occurs
func GetUserByStudentId(studentId string) (v *User, err error) {

	o := orm.NewOrm()
	v = &User{}
	if err = o.QueryTable("user").Filter("student_id", studentId).One(v); err == nil {
		return v, nil
	}
	return nil, err

}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []User
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
