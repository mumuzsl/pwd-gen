package main

import (
	"fmt"
	"testing"
)

func TestDomain_Gen(t *testing.T) {
	var a = NewDomain("apply", "1444")
	md5Str, result, err := a.Gen("aaa")

	fmt.Println(md5Str, result, err)

	if a.OriginalStr == "" {
		t.Fail()
	}
}

func TestDomain_Gen2(t *testing.T) {
	d := &Domain{
		Name:    "apply",
		Account: "1444",
	}
	md5Str, result, err := d.Gen("aaa")

	fmt.Println(md5Str, result, err)

	if d.OriginalStr == "" {
		t.Fail()
	}
}
