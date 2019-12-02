package bridge

type Weapon interface {
	Wield() string
	Swing()
	Unwield()
}
