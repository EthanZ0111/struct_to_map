package main

import "reflect"

func AssignMap(v interface{}, fieldList []StructFiledMeta, useTagName bool) map[string]interface{} {
	vt, ok := v.(reflect.Value)
	if !ok {
		vt = reflect.ValueOf(v)
	}
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	ret := make(map[string]interface{})
	for _, v := range fieldList {
		name := v.OriName
		if useTagName {
			name = v.TagName
		}
		ret[name] = vt.Field(v.Idx[0]).Interface()
		if v.Type == reflect.Struct {
			ret[name] = AssignMap(vt.Field(v.Idx[0]), v.SubTypeList, useTagName)
		}
	}
	return ret
}
