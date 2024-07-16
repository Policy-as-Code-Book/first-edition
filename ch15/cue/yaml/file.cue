#Policy: {
	// Name: =~"^\\p{Lu}{1}\\p{Ll}\\-*$"
	// Tag:  =~"^\\p{Ll}{4,7}$"
	Name: =~"^[A-Z]{1}[a-z\\-]*$"
	Tag:  =~"^[a-z]{4,7}$"
}
policies: [...#Policy]