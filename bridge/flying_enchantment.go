package bridge

type flyingEnchantment struct{}

func NewFlyingEnchantment() Enchantment {
	return &flyingEnchantment{}
}

func (e *flyingEnchantment) OnActivate() {
}

func (e *flyingEnchantment) Apply() {

}

func (e *flyingEnchantment) OnDetective() {
}
