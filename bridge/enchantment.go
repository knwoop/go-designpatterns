package bridge

type Enchantment interface {
	OnActivate()
	Apply()
	OnDetective()
}
