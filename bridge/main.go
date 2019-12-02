package bridge

// main is program entry point
func main() {
	s := NewSword(NewSoulEatingEnchantment())
	s.Wield()
	s.Swing()
	s.Unwield()

	h := NewHammer(NewFlyingEnchantment())
	h.Wield()
	h.Swing()
	h.Unwield()
}
