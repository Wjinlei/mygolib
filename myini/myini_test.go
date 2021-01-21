package myini

import (
	"fmt"
	"testing"
)

func TestLoadfile(t *testing.T) {
	_, err := Loadfile("test.ini")
	if err != nil {
		t.Error(err)
	}
}

func TestAddKey(t *testing.T) {
	f, err := Loadfile("test.ini")
	if err != nil {
		t.Error(err)
	}
	if err := f.AddKey("main", "path1", "/wwwroot/path1"); err != nil {
		t.Error(err)
	}
}

func TestDelKey(t *testing.T) {
	f, err := Loadfile("test.ini")
	if err != nil {
		t.Error(err)
	}
	if err := f.DelKey("main", "path1"); err != nil {
		t.Error(err)
	}
}

func TestSetKey(t *testing.T) {
	f, err := Loadfile("test.ini")
	if err != nil {
		t.Error(err)
	}
	if err := f.SetKey("main", "path1", "/wwwroot/path1"); err != nil {
		t.Error(err)
	}
}

func TestGetKey(t *testing.T) {
	f, err := Loadfile("test.ini")
	if err != nil {
		t.Error(err)
	}
	key := f.GetKey("main", "port")
	fmt.Println(key.String())
}

func TestGetKeys(t *testing.T) {
	f, err := Loadfile("test.ini")
	if err != nil {
		t.Error(err)
	}
	keys := f.GetKeys("main")
	for _, key := range keys {
		fmt.Println(key.String())
	}
}

func TestHasKey(t *testing.T) {
	f, err := Loadfile("test.ini")
	if err != nil {
		t.Error(err)
	}
	if !f.HasKey("main", "path1") {
		t.Error("path1 不存在")
	}
}
