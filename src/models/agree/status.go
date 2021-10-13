package agree

const (
	NEW Status = iota + 1
	ONCHECK
	VERIFIED
)

type Status int

var status = map[Status]string{
	NEW:      "NEW",
	ONCHECK:  "ONCHECK",
	VERIFIED: "VERIFIED",
}

func (s Status) String() string {
	return status[s]
}

