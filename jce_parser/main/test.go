package main

import (
	"bytes"
	"fmt"
	"yytars/jce_parser/gojce"
	newjce "tars/jce_parser/jce/include/new/vvideo"
	oldjce "tars/jce_parser/jce/include/old/vvideo"
)

func main() {
	finfo := oldjce.FileListPlayInfo{
		M_vid:   "123",
		M_title: "123",
	}
	os := gojce.NewOutputStream()
	if err := finfo.WriteTo(os); err != nil {
		fmt.Printf("%v", err)
	}

	bs := os.ToBytes()
	//fmt.Printf("bs=\n%s\n", hex.Dump(bs))

	var finfo2 newjce.FileListPlayInfo
	is := gojce.NewInputStream(bs)
	if err := finfo2.ReadFrom(is); err != nil {
		fmt.Printf("%v", err)
	}

	buf := bytes.NewBuffer(nil)
	ds := gojce.NewDisplayer(buf, 0)
	finfo2.Display(ds)
	fmt.Println(buf.String())
}
