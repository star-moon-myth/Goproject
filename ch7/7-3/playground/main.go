package main

import (
	"bytes"
)

type MyWriter struct {
	buffer bytes.Buffer
	
}

func(m *MyWriter)Write(p []byte)(int ,error){
	return m.buffer.Write(p)
}

func(mw *MyWriter)String()string {
	return mw.buffer.String()
}


type MyStringerType struct {
}

func (ms MyStringerType)Stringvalue()string {
	return "哈哈😄，这次是值接受哦😯"
}
func(ms *MyStringerType)StringPointer()string{
	return "好吧，这次是指针了➡️，😢"
}
