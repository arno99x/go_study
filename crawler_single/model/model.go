package model

var XinZus = []string{"白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "魔羯座", "水瓶座", "双鱼座"}
var Educations = []string{"小学", "初中", "高中", "中专", "大专", "本科", "研究生", "硕士", "博士", "博士后", "院士"}
var Marriages = []string{"未婚", "离异", "丧偶", "已婚"}

type Profile struct {
	//姓名
	Name string
	//性别
	Gender string
	//年龄
	Age int
	//身高
	Height int
	//体重
	Weight int
	//收入
	Income string
	//婚烟状况
	Marriage string
	//教育水平
	Education string
	//职业
	Occupation string
	//户口
	HuKou string
	//星座
	XinZuo string
	//是否有房
	House string
	//是否有车
	Car string
	//其它未识别信息
	Other string
}
