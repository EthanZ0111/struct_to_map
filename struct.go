package main

import (
	"reflect"
)

type StructFiledMeta struct {
	OriName     string
	TagName     string
	Idx         []int
	Type        reflect.Kind
	SubTypeList []StructFiledMeta
}

func ParseStructType(i interface{}, tagName string) []StructFiledMeta {
	rt, ok := i.(reflect.Type)
	if !ok {
		rt = reflect.TypeOf(i)
	}
	metaList := make([]StructFiledMeta, 0)
	for i := 0; i < rt.NumField(); i++ {
		ti := rt.Field(i)
		if ti.Type.Kind() == reflect.Ptr {
			fieldType := ti.Type.Elem()
			subST := ParseStructType(fieldType, tagName)
			meta := StructFiledMeta{
				OriName:     ti.Name,
				TagName:     ti.Tag.Get(tagName),
				Idx:         []int{i},
				Type:        fieldType.Kind(),
				SubTypeList: subST,
			}
			metaList = append(metaList, meta)
		} else {
			if ti.Type.Kind() == reflect.Struct {
				subST := ParseStructType(ti.Type, tagName)
				meta := StructFiledMeta{
					OriName:     ti.Name,
					TagName:     ti.Tag.Get(tagName),
					Idx:         []int{i},
					Type:        ti.Type.Kind(),
					SubTypeList: subST,
				}
				metaList = append(metaList, meta)
			} else {
				fm := StructFiledMeta{
					OriName: ti.Name,
					TagName: ti.Tag.Get(tagName),
					Idx:     []int{i},
					Type:    ti.Type.Kind(),
				}
				metaList = append(metaList, fm)
			}
		}
	}
	return metaList
}
