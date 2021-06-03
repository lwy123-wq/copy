package work

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func RandSeq(n int) (string,string){
	str1:="qwertyuiopasdfghjklzxcvbnm"
	str2:="qwertyuiopasdfghjklzxcvbnm"
	byte1:=[]byte(str1)
	byte2:=[]byte(str2)
	result1:=[]byte{}
	result2:=[]byte{}
	b:=rand.New(rand.NewSource(time.Now().Unix()))
	for i:=0;i<n;i++ {
		result1=append(result1,byte1[b.Intn(len(byte1))])
	}
	for j:=0;j<n;j++ {
		result2=append(result2,byte2[b.Intn(len(byte2))])
	}
	return string(result1),string(result2)
}
func GetRandInt() int {
	r:=rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(2)
}
func GetRandBool() string {
	str:="true,false"
	var slice=strings.Split(str,",")
	var arr[2]string
	arr[0]=slice[0]
	arr[1]=slice[1]
	var arr1[1]string
	b:=rand.New(rand.NewSource(time.Now().Unix()))
	arr1[0]=arr[b.Intn(2)]
	return arr1[0]
}

func GetRand()string{
	str:="string,int,bool"
	var slice=strings.Split(str,",")
	var arr[3]string
	arr[0]=slice[0]
	arr[1]=slice[1]
	arr[2]=slice[2]
	//return arr[2]
	var arr1[3]string
	b:=rand.New(rand.NewSource(time.Now().Unix()))
	arr1[0]=arr[b.Intn(3)]
	return arr1[0]
}

func MyWrite()  {
	userFile:="/disk/go16pro/src/copy/aa.txt"
	out,err:=os.Create(userFile)
	defer out.Close()
	if err!=nil{
		fmt.Println(userFile,err)
		return
	}
	out.WriteString("package test \n\n")
	out.WriteString("type Student struct { \n")
	var arr[10] string
	var arr1[10] string
	var arr2[10] string
	//var once sync.Once
	for i := 0; i <10 ; i++ {
		a, b :=RandSeq(1)
		arr[i]=a
		arr1[i]=b
		c:=GetRand()
		arr2[i]=c
		out.WriteString(a+b+" "+c+"\n")
		/*once.Do(func() {
			out.WriteString("} "+"\n")
		})*/

		//out.WriteString("func (f *Student) Set"+a+b+"("+strings.ToLower(a)+b+" "+c+") {\n\nf."+strings.ToLower(a)+b+ "= "+strings.ToLower(a)+b+"\n\n}\n")
		time.Sleep(1*time.Second)
	}

	out.WriteString("} "+"\n")
	for i := 0; i <10; i++ {
		out.WriteString("func (f *Student) Set"+ strings.ToUpper(arr[i])+arr1[i]+"("+ arr[i]+arr1[i]+" "+arr2[i]+") {\n\nf."+ arr[i]+arr1[i]+ "= "+strings.ToLower(arr[i])+arr1[i]+"\n\n}\n\n")

	}
	for i := 0; i<10 ; i++ {
		out.WriteString("\n")
		out.WriteString("\n"+"func (f Student) "+strings.ToUpper(arr[i])+arr1[i]+"() "+arr2[i]+" {\n\nreturn f."+arr[i]+arr1[i]+"\n\n}")
	}
	out.WriteString("\n")
	out.WriteString("func MyRef()  {\n\taa:=new(Student)\n")
	for i := 0; i <10 ; i++ {
		if arr2[i]=="string" {
			a,b:=RandSeq(4)
			out.WriteString("aa.Set"+strings.ToUpper(arr[i])+arr1[i]+"("+string('"')+a+b+string('"')+")"+"\n")
		}
		if arr2[i]=="int"{
			var aa =GetRandInt()
			out.WriteString("aa.Set"+strings.ToUpper(arr[i])+arr1[i]+"("+strconv.Itoa(aa)+")"+"\n")
		}
		if arr2[i]=="bool" {
			out.WriteString("aa.Set"+strings.ToUpper(arr[i])+arr1[i]+"("+GetRandBool()+")"+"\n")
		}
	}
	out.WriteString("s:=Student{\n")
	for i := 0; i < 10; i++ {
		out.WriteString(arr[i]+arr1[i]+":aa."+strings.ToUpper(arr[i])+arr1[i]+"(),\n")
	}
	out.WriteString("} \n")
	out.WriteString("t:=reflect.TypeOf(s)\n\tv:=reflect.ValueOf(s)\n\tfor i:=0;i<t.NumField();i++{\n\t\tkey:=t.Field(i)\n\t\tval:=v.Field(i)\n\t\tfmt.Println(key.Name,key.Type,val)\n\t}\n}")

}
func Mycopyfile()  {
	in,err:=os.Open("/disk/go16pro/src/copy/aa.txt")
	if err!=nil{
		panic(err)
	}
	defer in.Close()
	fo,err:=os.Create("/disk/go16pro/src/copy/test/work.go")
	if err!=nil{
		panic(err)
	}
	defer fo.Close()
	buf:=make([]byte,1024)
	for{
		n,err:=in.Read(buf)
		if err!=nil &&err!=io.EOF{
			panic(err)
		}
		if n==0 {
			break
		}
		if n1,err:=fo.Write(buf[:n]);err!=nil{
			panic(err)
		}else if n1!=n {
			panic("error in writing")
		}
	}


}

