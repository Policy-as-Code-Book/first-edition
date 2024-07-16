#Policy: {
	Name: =~"^[A-Z]{1}[a-z\\-]*$"
	Tag:  "deny" | "violate" | "warn"
}
policies: [...#Policy]