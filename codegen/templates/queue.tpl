package {{.Package}}

type {{.MyType}}Queue struct {
	q []{{.MyType}}
}

func New{{.MyType}}Queue() *{{.MyType}}Queue {
	return &{{.MyType}}Queue{
		q: []{{.MyType}}{},
	}
}

func (o *{{.MyType}}Queue) Insert(v {{.MyType}}) {
	o.q = append(o.q, v)
}

func (o *{{.MyType}}Queue) Remove() {{.MyType}} {
	if len(o.q) == 0 {
		panic("Oops.")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
