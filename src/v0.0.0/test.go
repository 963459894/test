package test

import "fmt"

type Student struct{
	Id int8  `json:"id"`
	Name string `json:"name"`
	Age int8 `json:"age"`
	Score int8 `json:"score"`
}

type List map[int8]*Student


//初始化结构体 返回指针
func NewStudent(id int8,name string, age int8,score int8)*Student{
	return &Student{
		Id:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
}

//获取列表
func (l List)GetList(){
	for k,v:=range l{
		fmt.Printf("%v=>%+v\n",k,*v)
	}
	fmt.Println("-----------------------------------------")
}

//获取详细信息
func (l List)GetDetail(index int8) {
	if _,ok:=l[index];!ok{
		fmt.Println("数据不存在",index)
		return
	}
	fmt.Printf("%v=>%+v\n",index,*l[index])
	fmt.Println("--------------------------------------")
}

//添加数据
func (l List)Add(student *Student){
	if _,ok:=l[student.Id];ok{
		fmt.Println("ID重复，添加失败")
		return
	}
	l[student.Id]=student
}

//修改数据
func (l List)update(id int8,student *Student){
	if _,ok:=l[id];!ok{
		fmt.Println("ID不存在，添加失败")
		return
	}

}

//删除数据
func (l List)delete(id int8){
	if _,ok:=l[id];!ok{
		fmt.Println("ID不存在，删除失败")
		return
	}
	delete(l,id)
}



