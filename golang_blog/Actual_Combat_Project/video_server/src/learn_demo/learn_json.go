package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	//对数组类型的json编码
	x:= [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	//对map数据类型的json编码
	map_s := make(map[string]float32)
	map_s["age"] = 28.5
	m_s, err := json.Marshal(map_s)
	if err != nil {
		panic(err)
	}
    fmt.Println(string(m_s))

	//对对象类型进行json编码
	student := Student{"louis", 28}
	student_json, err := json.Marshal(student)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(student_json))

	//对student_json进行json解码
	var std interface{}
	json.Unmarshal([]byte(student_json), &std)
	fmt.Printf("%v\n", std)

	//
	MD5Inst := md5.New()
	MD5Inst.Write([]byte("louis"))
	Result := MD5Inst.Sum([]byte(""))
	fmt.Printf("%x\n", Result)

}