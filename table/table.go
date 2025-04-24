package table

import (
	"context"
	"errors"
	"github.com/busy-cloud/boat/db"
	"github.com/spf13/cast"
	"strings"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm/schemas"
)

func TypeToSqlType(t string) (st schemas.SQLType) {
	switch t {
	case "int", "int8", "int16", "int32":
		st = schemas.SQLType{schemas.Int, 0, 0}
	case "uint", "uint8", "uint16", "uint32":
		st = schemas.SQLType{schemas.UnsignedInt, 0, 0}
	case "int64":
		st = schemas.SQLType{schemas.BigInt, 0, 0}
	case "uint64":
		st = schemas.SQLType{schemas.UnsignedBigInt, 0, 0}
	case "float32", "float":
		st = schemas.SQLType{schemas.Float, 0, 0}
	case "float64", "double":
		st = schemas.SQLType{schemas.Double, 0, 0}
	case "complex64", "complex128":
		st = schemas.SQLType{schemas.Varchar, 64, 0}
	case "blob":
		st = schemas.SQLType{schemas.Blob, 0, 0}
	case "text":
		st = schemas.SQLType{schemas.Text, 0, 0}
	case "bool", "boolean":
		st = schemas.SQLType{schemas.Bool, 0, 0}
	case "string":
		st = schemas.SQLType{schemas.Varchar, 255, 0}
	case "datetime":
		st = schemas.SQLType{schemas.DateTime, 0, 0}
	default:
		st = schemas.SQLType{schemas.Text, 0, 0}
	}
	return
}

type Field struct {
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
	Default    string `json:"default,omitempty"`
	NotNull    bool   `json:"not_null,omitempty"`
	Length     int64  `json:"length,omitempty"`
	Length2    int64  `json:"length2,omitempty"`
	PrimaryKey bool   `json:"primary_key,omitempty"`
	Increment  bool   `json:"increment,omitempty"`
	Indexed    bool   `json:"indexed,omitempty"`
	Created    bool   `json:"created,omitempty"`
	Updated    bool   `json:"updated,omitempty"`
}

type Table struct {
	Name          string   `json:"name,omitempty"`
	Fields        []*Field `json:"fields,omitempty"`
	DisableInsert bool     `json:"disable_insert,omitempty"`
	DisableUpdate bool     `json:"disable_update,omitempty"`
	DisableDelete bool     `json:"disable_delete,omitempty"`
}

func (t *Table) parseId(id string, bdr *builder.Builder) (err error) {
	ids := strings.Split(id, ",")
	var i = 0
	for _, field := range t.Fields {
		if field.PrimaryKey {
			if i < len(ids) {
				var val any
				switch field.Type {
				case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
					val, err = cast.ToInt64E(ids[i])
				case "float32", "float64", "float", "double":
					val, err = cast.ToFloat64E(ids[i])
				default:
					val = ids[i]
				}

				bdr.Where(builder.Eq{field.Name: val})
				i++
			} else {
				return errors.New("id length error")
			}
		}
	}
	if i == 0 {
		return errors.New("table has not primary key")
	}
	return
}

func (t *Table) Schema() *schemas.Table {
	//构建xorm schema
	var table schemas.Table
	table.Name = t.Name

	//转化列
	for _, field := range t.Fields {
		col := schemas.NewColumn(field.Name, "", TypeToSqlType(field.Type), field.Length, field.Length2, !field.NotNull)
		col.IsPrimaryKey = field.PrimaryKey
		col.IsAutoIncrement = field.Increment
		col.Default = field.Default
		col.IsCreated = field.Created
		col.IsUpdated = field.Updated
		if field.Indexed {
			col.Indexes[field.Name] = schemas.IndexType
		}
		table.AddColumn(col)
	}

	return &table
}

func (t *Table) Create() error {
	schema := t.Schema()

	//创建表
	sql, _, err := db.Engine().Dialect().CreateTableSQL(context.Background(), db.Engine().DB(), schema, t.Name)
	if err != nil {
		return err
	}
	_, err = db.Engine().Exec(sql)
	if err != nil {
		return err
	}

	//创建索引
	for _, index := range schema.Indexes {
		sql := db.Engine().Dialect().CreateIndexSQL(t.Name, index)
		_, err := db.Engine().Exec(sql)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Table) Drop() error {
	//第二个参数 checkIfExist 没有处理
	sql, _ := db.Engine().Dialect().DropTableSQL(t.Name)
	_, err := db.Engine().Exec(sql)
	return err
}

func (t *Table) Insert(values map[string]any) (id any, err error) {
	if len(values) == 0 {
		err = errors.New("no values to insert")
		return
	}

	for _, field := range t.Fields {
		if field.Created {
			values[field.Name] = time.Now()
		}
	}

	_, err = db.Engine().Table(t.Name).Insert(values)
	if err != nil {
		return id, err
	}
	return
}

func (t *Table) Update(cond map[string]any, values map[string]any) (rows int64, err error) {
	for _, field := range t.Fields {
		if field.Updated {
			values[field.Name] = time.Now()
		}
	}

	var updates []builder.Cond
	for k, v := range values {
		updates = append(updates, builder.Eq{k: v})
	}
	bdr := builder.Update(updates...)

	bdr.From(t.Name)

	for k, v := range cond {
		bdr.Where(builder.Eq{k: v})
	}

	res, err := db.Engine().Exec(bdr)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (t *Table) UpdateById(id string, values map[string]any) (rows int64, err error) {
	for _, field := range t.Fields {
		if field.Updated {
			values[field.Name] = time.Now()
		}
	}

	var updates []builder.Cond
	for k, v := range values {
		updates = append(updates, builder.Eq{k: v})
	}
	bdr := builder.Update(updates...)

	bdr.From(t.Name)

	err = t.parseId(id, bdr)
	if err != nil {
		return 0, err
	}
	//bdr.Where(builder.Eq{"id": id})

	res, err := db.Engine().Exec(bdr)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (t *Table) Delete(cond map[string]any) (rows int64, err error) {
	var conds []builder.Cond
	for k, v := range cond {
		conds = append(conds, builder.Eq{k: v})
	}

	bdr := builder.Delete(conds...).From(t.Name)
	res, err := db.Engine().Exec(bdr)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (t *Table) DeleteById(id string) (rows int64, err error) {
	bdr := builder.Delete()

	err = t.parseId(id, bdr)
	if err != nil {
		return 0, err
	}
	bdr.From(t.Name)

	res, err := db.Engine().Exec(bdr)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (t *Table) Find(cond map[string]any, fields []string, skip, limit int) (rows []map[string]any, err error) {
	bdr := builder.Select(fields...)

	bdr.From(t.Name)

	for k, v := range cond {
		bdr.Where(builder.Eq{k: v})
	}

	//bdr.OrderBy()

	bdr.Limit(skip, limit)

	return db.Engine().QueryInterface(bdr)
}

func (t *Table) Get(id string, fields []string) (Document, error) {
	bdr := builder.Select(fields...)

	bdr.From(t.Name)

	err := t.parseId(id, bdr)
	if err != nil {
		return nil, err
	}
	//bdr.Where(builder.Eq{"id": id})

	res, err := db.Engine().QueryInterface(bdr)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil //TODO 记录不存在
	}
	return res[0], nil
}

func (t *Table) Count(cond map[string]any) (cnt int64, err error) {
	bdr := builder.Select("count(*)").From(t.Name)

	for k, v := range cond {
		bdr.Where(builder.Eq{k: v})
	}

	res, err := db.Engine().QueryInterface(bdr)
	if err != nil {
		return 0, err
	}

	if len(res) == 0 {
		return 0, errors.New("no values to count")
	}

	for _, v := range res[0] {
		return cast.ToInt64(v), nil
	}

	return 0, errors.New("no values to count")
}
