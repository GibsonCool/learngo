package model

type Profile struct {
	Id                string
	Name              string
	Gender            string //性别
	CurrentResidence  string //居住地
	Age               int
	SalaryOrEducation string //月薪或学历  男：月薪   女：学历
	WeddingStatus     string //婚况
	Height            int
	Remarks           string
	InfoLink          string //具体详情链接
}
