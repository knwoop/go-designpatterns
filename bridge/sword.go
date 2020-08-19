package bridge

type sword struct {
	enchantment Enchantment
}

func NewSword(e Enchantment) Weapon {
	return &sword{
		enchantment: e,
	}
}

func (s *sword) Wield() {
	s.enchantment.OnActivate()
}

func (s *sword) Swing() {
	s.enchantment.Apply()
}

func (s *sword) Unwield() {
	s.enchantment.OnDetective()
}
