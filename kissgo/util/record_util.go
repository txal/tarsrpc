package util

/*
	主要用于测试代码的执行时间
*/

import (
	"fmt"
	"time"
)

//获得毫秒数
func getMsec() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

//记录时间并输出（name非空时才打印结果）
func RecordQuick(name string, f func()) (use int64) {
	begin := getMsec()
	f()
	end := getMsec()
	use = end - begin

	if name != "" {
		fmt.Printf("RecordQuick(\"%s\") use %d ms\n", name, use)
	}
	return
}

//时间记录结构
type RecordInfo struct {
	Start int64    //开始毫秒数
	Last  int64    //最后毫秒数
	Names []string //每次记录名
	Marks []int64  //每次经过毫秒数
}

//申请时间记录
func RecordStart() (rt *RecordInfo) {
	rt = &RecordInfo{}
	rt.Reset()
	return
}

//时间记录重置
func (rt *RecordInfo) Reset() {
	rt.Start = getMsec()
	rt.Last = getMsec()
	rt.Names = []string{} //make([]string, 0, 10)
	rt.Marks = []int64{}  //make([]int64, 0, 10)
	return
}

//直接获取间隔
func (rt *RecordInfo) Stop() (pass int64) {
	pass = getMsec() - rt.Last
	return
}

//记录一次
func (rt *RecordInfo) Mark(name string) (pass int64) {
	now := getMsec()
	pass = now - rt.Last

	rt.Names = append(rt.Names, name)
	rt.Marks = append(rt.Marks, pass)
	rt.Last = now

	return
}

//获取总时间
func (rt *RecordInfo) TotalMsec() (ret int64) {
	ret = rt.Last - rt.Start
	return
}

//获取总记录数
func (rt *RecordInfo) TotalCount() (ret int64) {
	ret = int64(len(rt.Names))
	return
}

//打印结果
func (rt *RecordInfo) Print(title string) {
	fmt.Printf("===== [%s]count=%d use=%d ms =======\n", title, rt.TotalCount(), rt.TotalMsec())
	for i := 0; i < len(rt.Names); i++ {
		fmt.Printf("  %s = %d ms\n", rt.Names[i], rt.Marks[i])
	}
	fmt.Printf("=================\n")
	return
}
