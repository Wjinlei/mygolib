package public

import "testing"

func TestChkDNFormat(t *testing.T) {
	dn := ChkDNFormat("www.test.com")
	if !dn {
		t.Error("check failed")
	}
}

func TestFilterSql(t *testing.T) {
	sql := FilterSql("()][].")
	if sql != "_" {
		t.Error(sql)
	}
}

func TestIntToString(t *testing.T) {
	src := []int8{65, 66, 67}
	dst := IntToString(src)
	if dst != "ABC" {
		t.Error("could not convert")
	}
}

func TestUIntToString(t *testing.T) {
	src := []uint8{65, 66, 67}
	dst := UintToString(src)
	if dst != "ABC" {
		t.Error("could not convert")
	}
}

func TestByteToString(t *testing.T) {
	src := []byte{65, 66, 67}
	dst := ByteToString(src)
	if dst != "ABC" {
		t.Error("could not convert")
	}
	src = []byte{0, 65, 66, 67}
	dst = ByteToString(src)
	if dst != "ABC" {
		t.Error("could not convert")
	}
}

func TestHexToUint32(t *testing.T) {
	if HexToUint32("FFFFFFFF") != 4294967295 {
		t.Error("Could not convert")
	}
}

func TestParseInt32(t *testing.T) {
	ret := ParseInt32("11111")
	if ret != int32(11111) {
		t.Error("could not parse")
	}
}

func TestParseUint64(t *testing.T) {
	ret := ParseUint64("11111")
	if ret != uint64(11111) {
		t.Error("could not parse")
	}
}

func TestParseFloat64(t *testing.T) {
	ret := ParseFloat64("11111.11")
	if ret != float64(11111.11) {
		t.Error("could not parse")
	}
	ret = ParseFloat64("11111")
	if ret != float64(11111) {
		t.Error("could not parse")
	}
}
