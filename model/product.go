package model

import (
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"products"`
	Id            uint64 `bun:",pk,autoincrement"`
	Name          string `bun:",notnull"`
	Description   string `bun:",type:varchar(1000)"`
	Price         int64  `bun:",notnull"`
	Notes         string `bun:",type:varchar(1000)"`
	Dimensions    string `bun:",type:varchar(1000)"`
	Features      string `bun:",type:varchar(1000)"`
}
