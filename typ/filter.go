package typ //import "go.iondynamics.net/helpdesk/typ"

import (
	"strings"
	"time"
)

type AbstractFilter struct {
	Complement bool
	set        bool
}

func (vf *AbstractFilter) IsSet() bool {
	return vf.set
}

type ValueFilter struct {
	AbstractFilter
	value interface{}
}

func (vf *ValueFilter) Set(v interface{}) {
	vf.value = v
	vf.AbstractFilter.set = true
}

func (vf *ValueFilter) Unset() {
	vf.set = false
	vf.value = nil
}

func (vf *ValueFilter) Get() interface{} {
	return vf.value
}

func (vf *ValueFilter) Check(v interface{}) bool {
	if !vf.IsSet() {
		return true
	}

	return (vf.value == v) != vf.Complement
}

//

type BoolFilter struct {
	ValueFilter
}

func (bf *BoolFilter) Set(v bool) {
	bf.ValueFilter.Set(v)
}

func (bf *BoolFilter) Get() bool {
	return bf.ValueFilter.Get().(bool)
}

type StateFilter struct {
	ValueFilter
}

func (sf *StateFilter) Set(v State) {
	sf.ValueFilter.Set(v)
}

func (sf *StateFilter) Get() State {
	return sf.ValueFilter.Get().(State)
}

type RoleFilter struct {
	ValueFilter
}

func (rf *RoleFilter) Set(v Role) {
	rf.ValueFilter.Set(v)
}

func (rf *RoleFilter) Get() Role {
	return rf.ValueFilter.Get().(Role)
}

type TimeFilter struct {
	ValueFilter
}

func (tf *TimeFilter) Set(v time.Time) {
	tf.ValueFilter.Set(v)
}

func (tf *TimeFilter) Get() time.Time {
	return tf.ValueFilter.Get().(time.Time)
}

type StringFilter struct {
	ValueFilter
	Contains bool
}

func (sf *StringFilter) Set(v string) {
	sf.ValueFilter.Set(v)
}

func (sf *StringFilter) Get() string {
	return sf.ValueFilter.Get().(string)
}

func (sf *StringFilter) Check(str string) bool {
	if sf.Contains == false {
		return sf.ValueFilter.Check(str)
	}

	if !sf.IsSet() {
		return true
	}

	return strings.Contains(str, sf.Get()) != sf.Complement
}
