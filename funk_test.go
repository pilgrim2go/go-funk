package funk

import "database/sql"

// Bar is
type Bar struct {
	Name string
	Bar  *Bar
	Bars []*Bar
}

func (b Bar) TableName() string {
	return "bar"
}

// Foo is
type Foo struct {
	ID         int
	FirstName  string `tag_name:"tag 1"`
	LastName   string `tag_name:"tag 2"`
	Age        int    `tag_name:"tag 3"`
	Bar        *Bar
	Bars       []*Bar
	EmptyValue sql.NullInt64
}

func (f Foo) TableName() string {
	return "foo"
}

var bar *Bar = &Bar{
	Name: "Test",
	Bars: []*Bar{
		&Bar{
			Name: "Level1-1",
			Bar: &Bar{
				Name: "Level2-1",
			},
		},
		&Bar{
			Name: "Level1-2",
			Bar: &Bar{
				Name: "Level2-2",
			},
		},
	},
}

var foo *Foo = &Foo{
	ID:        1,
	FirstName: "Dark",
	LastName:  "Vador",
	Age:       30,
	Bar:       bar,
	EmptyValue: sql.NullInt64{
		Valid: true,
		Int64: 10,
	},
	Bars: []*Bar{
		bar,
		bar,
	},
}

var foo2 *Foo = &Foo{
	ID:        1,
	FirstName: "Dark",
	LastName:  "Vador",
	Age:       30,
}
