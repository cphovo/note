package model

func Models() []any {
	return []any{
		&Article{},
		&Tag{},
		&FileType{},
	}
}
