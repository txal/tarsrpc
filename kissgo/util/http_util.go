/***************************************************
	http辅助功能 by huangzhibin
	------------------------------------------------
***************************************************/
package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//获取http的get参数
func HttpGetInt64(req *http.Request, name string) (ret int64) {
	str := req.URL.Query().Get(name)
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		ret = 0
	}
	return
}

//获取http的get参数
func HttpGetString(req *http.Request, name string) (ret string) {
	ret = req.URL.Query().Get(name)
	return
}

// HttpApiName获取api名字
func HttpApiName(req *http.Request) string {
	apiname := req.URL.Path

	if i := strings.LastIndex(apiname, "/"); i >= 0 {
		apiname = apiname[i+1:]
	}

	if apiname == "" {
		return "emptyuri"
	}

	return apiname
}

//json返回（js方式"callback"非空；移动端json"callback"为空）（data可以是结构体，可以为结构体指针，最好是结构体指针）
func HttpRespJson(rw http.ResponseWriter, req *http.Request, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")

	var jsonStr []byte
	if str, ok := data.([]byte); ok {
		jsonStr = str
	} else {
		jsonStr, _ = json.Marshal(data)
	}

	callback := req.URL.Query().Get("callback")
	if callback == "" {
		fmt.Fprintf(rw, "%s", jsonStr)
	} else {
		fmt.Fprintf(rw, "%s(%s)", callback, jsonStr)
	}
}

//生成key
func getHttpKey(req *http.Request) string {
	return req.Method + ":" + req.URL.String()
}

//使用缓存返回（返回true为使用缓存，false为找不到缓存）
func HttpCacheReturn(c *CacheInfo, rw http.ResponseWriter, req *http.Request) bool {
	key := getHttpKey(req)
	r, err := c.GetCacheOrWait(key)
	if err != nil {
		return false
	}
	callback := req.URL.Query().Get("callback")
	if callback == "" {
		fmt.Fprintf(rw, "%s", r)
	} else {
		fmt.Fprintf(rw, "%s(%s)", callback, r)
	}
	return true
}

//设置缓存
func HttpCacheRespJson(c *CacheInfo, rw http.ResponseWriter, req *http.Request, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")

	var jsonStr []byte
	if str, ok := data.([]byte); ok {
		jsonStr = str
	} else {
		jsonStr, _ = json.Marshal(data)
	}

	callback := req.URL.Query().Get("callback")
	if callback == "" {
		fmt.Fprintf(rw, "%s", jsonStr)
	} else {
		fmt.Fprintf(rw, "%s(%s)", callback, jsonStr)
	}

	key := getHttpKey(req)
	c.SetCache(key, jsonStr)
}

//下载文件（外部要return）
func HttpRespFile(rw http.ResponseWriter, req *http.Request, fileName string, bs []byte) {
	//log.Debug("downloadFile name=%s len=%v", fileName, len(bs))
	rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	rw.Write(bs)
}

//分页
type PageInfo struct {
	PageIndex int64 `json:"pageIndex"` //当前第几页（input）
	PageSize  int64 `json:"pageSize"`  //每页大小（input）
	PageMax   int64 `json:"pageMax"`   //最大页数（output）
	Number    int64 `json:"number"`    //最大数量（output）
}

//分页：初始化
func HttpGetPage(req *http.Request) (page *PageInfo) {
	page = &PageInfo{}
	page.PageIndex = HttpGetInt64(req, "pageIndex")
	page.PageSize = HttpGetInt64(req, "pageSize")
	if page.PageSize <= 0 {
		page.PageSize = 50
	}
	if page.PageIndex <= 0 {
		page.PageIndex = 1
	}
	return
}

//分页：设置最大数量
func (page *PageInfo) SetCount(count int64) {
	page.Number = count
	if count == 0 {
		page.PageMax = 1
		page.PageIndex = 1
	} else {
		page.PageMax = (count + page.PageSize - 1) / page.PageSize
		if page.PageIndex > page.PageMax {
			page.PageIndex = page.PageMax
		}
	}
}

//分页：跳过多少条记录
func (page *PageInfo) Skip() int {
	return int((page.PageIndex - 1) * page.PageSize)
}

//分页：最多获取多少条记录
func (page *PageInfo) Limit() int {
	return int(page.PageSize)
}
