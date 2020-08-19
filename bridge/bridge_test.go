package bridge_test

import (
	"testing"

	"github.com/knwoop/go-designpatterns/bridge"
)

func TestBridge(t *testing.T) {
	s := bridge.NewSword(bridge.NewSoulEatingEnchantment())
	s.Wield()
	s.Swing()
	s.Unwield()

	h := bridge.NewHammer(bridge.NewFlyingEnchantment())
	h.Wield()
	h.Swing()
	h.Unwield()
}
