package partners

type Partner int

const (
	AIRBA Partner = iota + 1
)

var (
	partner = map[Partner]string{
		AIRBA:       "AIRBA",
	}
)

func (p Partner) String() string {
	return partner[p]
}
