package colums_c

import "reflect"

type Colum struct {
	Name    string
	Type    reflect.Type
	Content any
}
type colum_append interface {
	Append_content(any)
}

func (c *Colum)Append_content(sty any){
	
}