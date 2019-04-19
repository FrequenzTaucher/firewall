package firewall

func FilterCountry(iso string) bool {
	list := make(map[string]struct{})
	list["SK"] = struct{}{}

	_, found := list[iso]

	return found
}
