package ctxutil

type CtxKey int

const (
	CtxUserID CtxKey = iota + 1 // userid
	CtxOffSet                   // offset
	CtxLimit                    // limit
	CtxSort                     // sort
)
