package bridge

type soulEatingEnchantment struct{}

func NewSoulEatingEnchantment() Enchantment {
	return &soulEatingEnchantment{}
}

func (e *soulEatingEnchantment) OnActivate() {
}

func (e *soulEatingEnchantment) Apply() {

}

func (e *soulEatingEnchantment) OnDetective() {
}
