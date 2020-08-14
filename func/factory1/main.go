package main

import "fmt"

/*
   ①封装人，属性包括姓名、年龄、身高、体重、颜值、资产、性别、性取向
   ②给人封装结婚方法，参数是潜在的结婚对象：
         a、如果对方的性取向有问题，panic
       b、如果对方的颜值过低，返回错误
         c、否则返回满意程度
*/

type Gender int //性别取向

func (g Gender) String() string {
	return []string{"Male", "Female", "Bisexual"}[g]
}

//性别枚举
const (
	Male     = iota //男
	Female          //女
	Bisexual        //人妖
)

type Human struct {
	Name          string //名字
	Age           int    //年龄
	Height        int    //身高
	Weight        int    //体重
	Looking       int    //自己的颜值
	TargetLooking int    //期望的颜值
	Rmb           int    //人民币
	Sex           Gender //自己的性别
	TargetSex     Gender //结婚目标性别

}

//不理想配偶错误
type BadSouseError struct {
	why string
}

func (bse *BadSouseError) Error() string {
	return bse.why
}

//工厂方法
func CreateBadSpouseError(o *Human) string {
	bse := new(BadSouseError)
	if o.Rmb < 1000 {
		bse.why = "太穷"
	} else if o.Weight > 200 {
		bse.why = "太胖"
	} else if o.Age > 50 {
		bse.why = "太老"
	}
	return bse.Error()
}

func (h *Human) Marry(o *Human) (happiness int, err error) {

	//如果他的性别不等于你期望性别，报错
	if o.Sex != h.TargetSex {
		panic(&BadSouseError{"淡定，，，我不是你的菜"})
		return
	}
	if errS := CreateBadSpouseError(o); errS != "" {
		return
	}
	//如果颜值过低，返回错误
	if o.Looking < h.TargetLooking {
		panic(&BadSouseError{"sorry，颜值不匹配"})
		return
	}

	//计算幸福程度
	happiness = (o.Height * o.Looking) / (h.Weight * h.TargetLooking)
	return
}

func NewHuman(name string, age, height, weight, rmb, looking, targetlooking int, sex, targetsex Gender) *Human {
	human := new(Human)
	human.Name = name
	human.Age = age
	human.Height = height
	human.Weight = weight
	human.Rmb = rmb
	human.Looking = looking
	human.TargetLooking = targetlooking
	human.Sex = sex
	human.TargetSex = targetsex
	return human
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	cook := NewHuman("库克", 11, 180, 123, 123456, 60, 10, Male, Female)
	ySister := NewHuman("你妹", 20, 155, 150, 43210, 20, 90, Female, Male)
	happiness, err := cook.Marry(ySister)
	if err != nil {
		fmt.Println("牵手失败", err)
	} else {
		fmt.Println("牵手成功,幸福指数=", happiness)

	}

}
