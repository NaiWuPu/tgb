package main

import (
	"QingGoUCenter/models/shop"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sort"
)

type ShopCategoryOne struct {
	Id         int64
	Name       string
	ShowName   string
	ChildCount int
	IsUse	   bool
	Level      int8
	Children   []*ShopCategoryOne
}

type ReturnShopCategoryOne struct {
	Id         int64
	Name       string
	ShowName   string
}

func main() {
	//使用分类
	ids := []int64{
		1375,1373,1376,1453,1454,1441,1480,1236,1521,1522,1523,1524,1525,1527,1528,1529,1530,1474,1536,1537,1538,1539,1540,1541,1542,1543,1544,1545,1546,1857,1772,1773,1533,1534,1535,1846,1395,1396,1397,1398,1400,1401,1402,1403,1404,1405,1406,1447,1475,1739,1494,1505,1506,1507,1508,1509,1510,1511,1512,1513,1514,1515,1516,1517,1518,1519,1520,1778,1588,1589,1590,1591,1592,1593,1594,1855,1575,1452,1473,1531,1379,1587,1386,
	}
	//全部分类
	var shopCategoryList []*shop.ShopCategory
	ShopCategoryOneMap := make(map[int64]*ShopCategoryOne)
	err := orm.RegisterDataBase("default", "mysql", "jcc:jcc@tcp(172.16.0.203:3306)/city")
	if err != nil {
		log.Fatal(err)
		return
	}
	o := orm.NewOrm()
	total, err := o.QueryTable(&shop.ShopCategory{}).OrderBy("Level").All(&shopCategoryList)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(total)

	//组装map，填充key，做排序
	var cateIds []int
	for _, sc := range shopCategoryList {
		ShopCategoryOneMap[sc.Id] = &ShopCategoryOne{
			Id:       sc.Id,
			Name:     sc.Name,
			ShowName: sc.Name,
			Level: sc.Level,
		}
		cateIds = append(cateIds, int(sc.Id))
	}

	sort.Slice(cateIds, func(i, j int) bool {
		return cateIds[i] > cateIds[j]
	})
	//填充子集
	for _, sc := range shopCategoryList {
		//填充showName
		if ShopCategoryOneMap[sc.Parentid] == nil {
			continue
		}
		ShopCategoryOneMap[sc.Id].ShowName = ShopCategoryOneMap[sc.Parentid].ShowName + "/" + ShopCategoryOneMap[sc.Id].ShowName
		ShopCategoryOneMap[sc.Parentid].Children = append(ShopCategoryOneMap[sc.Parentid].Children, ShopCategoryOneMap[sc.Id])
		//标记无用分类,不在数组中
		for _, useId := range ids{
			if useId == sc.Id {
				ShopCategoryOneMap[sc.Parentid].ChildCount ++
				ShopCategoryOneMap[sc.Id].IsUse = true
				ShopCategoryOneMap[sc.Parentid].IsUse = false
			}
		}
	}

	for _, id := range cateIds {
		sc := ShopCategoryOneMap[int64(id)]
		//如果子集全用了，那子集都不可用，三级数据不参与
		if sc.ChildCount >= len(sc.Children) && sc.Level < 3 && len(sc.Children) > 0{
			for _, child := range sc.Children{
				child.IsUse = false
			}
			sc.IsUse = true
		}
	}

	//删除标记和子集长度相同数据
	var okShopCategoryList []*ReturnShopCategoryOne
	for _, id := range cateIds {
		sc := ShopCategoryOneMap[int64(id)]
		if sc.IsUse{
			okShopCategoryList = append(okShopCategoryList, &ReturnShopCategoryOne{
				Id:       sc.Id,
				Name:     sc.Name,
				ShowName: sc.ShowName,
			})
		}
	}

	for id, sc := range okShopCategoryList {
		fmt.Println("id:", id, "  Id:", sc.Id, "  Name:", sc.Name, "  ShowName:", sc.ShowName)
	}

}
