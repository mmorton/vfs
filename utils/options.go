package utils

import "github.com/c2fo/vfs/v6"

func BindFileOptions[Of any](opts []vfs.FileOption, fn func(Of)) {
	for _, opt := range opts {
		switch opt := opt.(type) {
		case vfs.FileOptionGenerator:
			BindFileOptions[Of](opt.Options(), fn)
		case Of:
			fn(opt)
		}
	}
}
