package mbalign

import (
	"testing"
)

func TestAlignLeft(t *testing.T) {
	result := Align("helloハロー", 12, '.', LeftAlign)
	if result != "helloハロー." {
		t.Error(result)
	}

	longResult := Align("hogeほげ", 7, '.', LeftAlign)
	if longResult != "hogeほ." {
		t.Error(longResult)
	}

	onlyHalf := Align("hello, world", 13, '.', LeftAlign)
	if onlyHalf != "hello, world." {
		t.Error(onlyHalf)
	}
	onlyHalfLong := Align("hello, world", 8, '.', LeftAlign)
	if onlyHalfLong != "hello, w" {
		t.Error(onlyHalfLong)
	}

	onlyFull := Align("あいうえお", 12, '.', LeftAlign)
	if onlyFull != "あいうえお.." {
		t.Error(onlyFull)
	}

	onlyFullLong := Align("あいうえお", 9, '.', LeftAlign)
	if onlyFullLong != "あいうえ." {
		t.Error(onlyFullLong)
	}
}

func TestAlignRight(t *testing.T) {
	result := Align("helloハロー", 12, '.', RightAlign)

	if result != ".helloハロー" {
		t.Error(result)
	}

	longResult := Align("hogeほげ", 7, '.', RightAlign)
	if longResult != ".hogeほ" {
		t.Error(longResult)
	}

	onlyHalf := Align("hello, world", 14, '.', RightAlign)
	if onlyHalf != "..hello, world" {
		t.Error(onlyHalf)
	}

	onlyHalfLong := Align("hello, world", 8, '.', RightAlign)
	if onlyHalfLong != "hello, w" {
		t.Error(onlyHalfLong)
	}

	onlyFull := Align("あいうえお", 12, '.', RightAlign)
	if onlyFull != "..あいうえお" {
		t.Error(onlyFull)
	}
	onlyFullLong := Align("あいうえお", 9, '.', RightAlign)
	if onlyFullLong != ".あいうえ" {
		t.Error(onlyFullLong)
	}
}

func TestGetWidth(t *testing.T) {
	result := GetWidth("hoge")
	if result != 4 {
		t.Error(result)
	}

	result = GetWidth("あいう")
	if result != 6 {
		t.Error(result)
	}

	result = GetWidth("あいうabc")
	if result != 9 {
		t.Error(result)
	}
}
