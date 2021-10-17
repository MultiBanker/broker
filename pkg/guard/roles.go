package guard

type role struct {
	name   roleName
	parent *role
}

type roleName string

const (
	HeadAdmin          roleName = "head_admin"
	HeadSupportManager roleName = "head_support_manager"
	SupportManager     roleName = "support_manager"
	Market             roleName = "market"
	Partner            roleName = "partner"
)

var (
	roles         = all()
)

func all() []*role {
	head := role{
		name:   HeadAdmin,
		parent: nil,
	}
	market := role{
		name:   Market,
		parent: nil,
	}

	partner := role{
		name:   Partner,
		parent: nil,
	}

	headSupport := role{name: HeadSupportManager, parent: &head}
	supportManager := role{name: SupportManager, parent: &headSupport}

	return []*role{&supportManager, &market, &partner}
}

func check(allow roleName, names ...roleName) bool {
	if containsString(names, HeadAdmin) {
		return true
	}

	var res bool
	for i := range names {
		res = res || checkRole(allow, names[i])
	}

	return res
}

func checkRole(allow, role roleName) bool {
	for j := range roles {
		if roleTree := find(roles[j], allow); roleTree != nil {
			return containsRoles(roleTree, role)
		}
	}

	return false
}

func find(role *role, name roleName) *role {
	for curr := role; curr != nil; curr = curr.parent {
		if curr.name == name {
			return curr
		}
	}

	return nil
}

func containsString(xs []roleName, x roleName) bool {
	for i := range xs {
		if xs[i] == x {
			return true
		}
	}

	return false
}

func containsRoles(role *role, name roleName) bool {
	for curr := role; curr != nil; curr = curr.parent {
		if curr.name == name {
			return true
		}
	}

	return false
}