package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url string = "http://222.243.213.119:8999/yyoa/common/js/menu/test.jsp?doType=101&S1=(select%20md5(1))"
const url2 string = "http://222.243.213.119:8999/yyoa/common/js/menu/test.jsp?"
//-GET方式---------------------------------------------------------------------------------------------
//①不带参数
func get1()  {
	b,err:= http.Get(url)
	if err != nil{
		panic(err)
	}
	defer b.Body.Close()
	context,err := ioutil.ReadAll(b.Body)
	if err !=nil{
		panic(err)
	}
	fmt.Println(string(context))
	fmt.Println("------------------Get法①------------------------------------")
}

//②使用client实现
func get2(){
	client := new(http.Client)
	resp,err := client.Get(url)
	if err !=nil{
		log.Println(err)
	}
	defer resp.Body.Close()
	context ,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Println(err)
	}
	fmt.Println(string(context))
	fmt.Println("------------------Get法②------------------------------------")
}

//③使用client实现 + 传入参数
func get3(){
	client,err := &http.Client{}


}


// -POST方式-----------------------------------------------------------------------------------------------
func post(){
	b,err:= http.Post(url,"application/json",nil)
	if err != nil{
		panic(err)
	}
	defer b.Body.Close()
	context,err := ioutil.ReadAll(b.Body)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("%s",context)
}

func main() {
	get1()                      //选择实现方式
}










