package main

import (
	"fmt"
	"testing"
)

func TestCheetSheet(t *testing.T) {
	out := CheatSheet("animation")
	tmp := `//glasses
(•_•)
( •_•)>⌐■-■
(⌐■_■)
`
	fmt.Println("outb: ", []byte(tmp))
	// 原文多了一个 \n rune是10 也可以去out补齐，我这里删除了
	if out != `//glasses
(•_•)
( •_•)>⌐■-■
(⌐■_■)
` {
		t.Fail()
	}
}
