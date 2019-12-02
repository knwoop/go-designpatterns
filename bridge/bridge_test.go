package bridge_test

import (
	"github.com/kntaka/go_design_pattern/bridge"
	"testing"
)

func TestBridge(t *testing.T) {
	s := bridge.NewSword(bridge.NewSoulEatingEnchantment())
	expect := "+---+\n|AAA|\n+---+\n"
	result := s.Wield()
	if (result != expect) {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
	s.Wield()
	s.Swing()
	s.Unwield()

	h := bridge.NewHammer(bridge.NewFlyingEnchantment())
	h.Wield()
	h.Swing()
	h.Unwield()
}