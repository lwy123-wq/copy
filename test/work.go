package test

import (
	"fmt"
	"reflect"
)

type Student struct { 
mp string
kw string
sz string
ua int
kt string
kk bool
fw int
bi string
vv string
cf int
} 
func (f *Student) SetMp(mp string) {

f.mp= mp

}

func (f *Student) SetKw(kw string) {

f.kw= kw

}

func (f *Student) SetSz(sz string) {

f.sz= sz

}

func (f *Student) SetUa(ua int) {

f.ua= ua

}

func (f *Student) SetKt(kt string) {

f.kt= kt

}

func (f *Student) SetKk(kk bool) {

f.kk= kk

}

func (f *Student) SetFw(fw int) {

f.fw= fw

}

func (f *Student) SetBi(bi string) {

f.bi= bi

}

func (f *Student) SetVv(vv string) {

f.vv= vv

}

func (f *Student) SetCf(cf int) {

f.cf= cf

}



func (f Student) Mp() string {

return f.mp

}

func (f Student) Kw() string {

return f.kw

}

func (f Student) Sz() string {

return f.sz

}

func (f Student) Ua() int {

return f.ua

}

func (f Student) Kt() string {

return f.kt

}

func (f Student) Kk() bool {

return f.kk

}

func (f Student) Fw() int {

return f.fw

}

func (f Student) Bi() string {

return f.bi

}

func (f Student) Vv() string {

return f.vv

}

func (f Student) Cf() int {

return f.cf

}
func MyRef()  {
	aa:=new(Student)
aa.SetMp("bficbpxa")
aa.SetKw("bficbpxa")
aa.SetSz("bficbpxa")
aa.SetUa(1)
aa.SetKt("bficbpxa")
aa.SetKk(false)
aa.SetFw(1)
aa.SetBi("bficbpxa")
aa.SetVv("bficbpxa")
aa.SetCf(1)
s:=Student{
mp:aa.Mp(),
kw:aa.Kw(),
sz:aa.Sz(),
ua:aa.Ua(),
kt:aa.Kt(),
kk:aa.Kk(),
fw:aa.Fw(),
bi:aa.Bi(),
vv:aa.Vv(),
cf:aa.Cf(),
} 
t:=reflect.TypeOf(s)
	v:=reflect.ValueOf(s)
	for i:=0;i<t.NumField();i++{
		key:=t.Field(i)
		val:=v.Field(i)
		fmt.Println(key.Name,key.Type,val)
	}
}