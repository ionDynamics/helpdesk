package typ //import "go.iondynamics.net/helpdesk/typ"

import (
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
