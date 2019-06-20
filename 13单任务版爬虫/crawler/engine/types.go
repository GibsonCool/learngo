package engine

/*
	将整个爬虫架构数据体信息抽象

*/
// 请求内容
type Request struct {
	Url        string                   //  输入url
	ParserFunc func([]byte) ParseResult //通过给定 解析函数 解析出结果
}

type ParseResult struct {
	Requests []Request     //返回下一级 请求内容
	Items    []interface{} //这一级的爬取信息
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
