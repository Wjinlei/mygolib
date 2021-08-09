package myini

import (
	"github.com/go-ini/ini"
)

type Obj struct {
	File *ini.File
	Path string
}

// 加载ini文件到内存中
func Loadfile(filepath string) (*Obj, error) {
	f, err := ini.Load(filepath)
	if err != nil {
		return nil, err
	}
	return &Obj{File: f, Path: filepath}, nil
}

// 获取Section
func (o Obj) GetSection(section string) (*Section, error) {
	sec, err := o.File.GetSection(section)
	if err != nil {
		return nil, err
	}
	return sec, nil
}

// 删除Section
func (o Obj) DelSection(section string) error {
	o.File.DeleteSection(section)
	if err := o.save(); err != nil {
		return err
	}
	return nil
}

// 在指定Section下,新增一个key
func (o Obj) AddKey(section, keyname, keyval string) error {
	_, err := o.File.Section(section).NewKey(keyname, keyval)
	if err != nil {
		return err
	}
	if err := o.save(); err != nil {
		return err
	}
	return nil
}

// 删除key
func (o Obj) DelKey(section, keyname string) error {
	o.File.Section(section).DeleteKey(keyname)
	if err := o.save(); err != nil {
		return err
	}
	return nil
}

// 设置key的值,如果这个key不存在,则会新创建一个并设置其值
func (o Obj) SetKey(section, keyname, keyval string) error {
	o.File.Section(section).Key(keyname).SetValue(keyval)
	if err := o.save(); err != nil {
		return err
	}
	return nil
}

// 获取指定Section下的key如果没有这个key则设置一个新的key且值为 "" 并返回
func (o Obj) GetKey(section, keyname string) *ini.Key {
	return o.File.Section(section).Key(keyname)
}

// 获取指定Section下所有的key
func (o Obj) GetKeys(section string) []*ini.Key {
	return o.File.Section(section).Keys()
}

// 判断指定key是否存在
func (o Obj) HasKey(section, keyname string) bool {
	return o.File.Section(section).HasKey(keyname)
}

// 保存到文件
func (o Obj) save() error {
	return o.File.SaveTo(o.Path)
}
