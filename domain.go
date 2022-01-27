package main

import (
	"strings"
)

type Domain struct {
	Name        string `json:"name"`
	Account     string `json:"account"`
	Separate    string `json:"separate"`
	Prefix      string `json:"prefix"`
	Point       int    `json:"point"`
	Length      int    `json:"length"`
	OriginalStr string
}

const (
	DefaultSeparate = "_"
	DefaultPrefix   = "0X_"
	DefaultPoint    = 17
	DefaultLength   = 8
)

func NewDomain(name, account string) *Domain {
	return &Domain{
		Name:        name,
		Account:     account,
		Separate:    DefaultSeparate,
		Prefix:      DefaultPrefix,
		Point:       DefaultPoint,
		Length:      DefaultLength,
		OriginalStr: "",
	}
}

func (d *Domain) Gen(pwd string) (md5Str string, result string, err error) {
	rawStrArray := []string{d.Name, d.Account, pwd}

	if d.Separate == "" {
		d.Separate = DefaultSeparate
	}
	if d.Prefix == "" {
		d.Prefix = DefaultPrefix
	}
	if d.Point == 0 {
		d.Point = DefaultPoint
	}
	if d.Length == 0 {
		d.Length = DefaultLength
	}

	d.OriginalStr = strings.Join(rawStrArray, d.Separate)

	return Gen(d.OriginalStr, d.Prefix, d.Point, d.Length)
}
