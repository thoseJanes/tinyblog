package id

import(
	shortid "github.com/0x5487/go-short-id"
)

func GenShortId() string {
	opt := shortid.Options{
		Number: 5,
		StartWithYear: true,
		EndWithHost: false,
	}

	return shortid.Generate(opt)
}