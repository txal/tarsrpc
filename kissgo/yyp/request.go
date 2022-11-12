// @author kordenlu
// @创建时间 2017/09/12 15:20
// 功能描述:

package yyp

// http://cherry.yy.com/ypress/page.action?id=164

type Header struct {
	Length uint32
	URI    uint32
	Status uint16
}

type Message struct {
	Header
	Content []byte
}
