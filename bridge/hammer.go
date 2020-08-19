package bridge

type hammer struct {
	enchantment Enchantment
}

func NewHammer(e Enchantment) Weapon {
	return &hammer{
		enchantment: e,
	}
}

func (h *hammer) Wield() {
	h.enchantment.OnActivate()
}

func (h *hammer) Swing() {
	h.enchantment.Apply()
}

func (h *hammer) Unwield() {
	h.enchantment.OnDetective()
}
