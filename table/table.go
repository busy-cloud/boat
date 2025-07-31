package table

import (
	"context"
	"errors"
	"fmt"
	"github.com/busy-cloud/boat/db"
	"github.com/rs/xid"
	"github.com/spf13/cast"
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

func (f *Field) Cast(v any) (ret any, err error) {
	switch f.Type {
	case "int", "int8", "int16", "int32", "int64":
		ret, err = cast.ToInt64E(v)
	case "uint", "uint8", "uint16", "uint32", "uint64":
		ret, err = cast.ToUint64E(v)
	case "float32", "float64", "float", "double":
		ret, err = cast.ToFloat64E(v)
	default: //string
		ret = v
	}
	return
}

func (f *Field) Condition(val string) (cond builder.Cond, err error) {
	var v any
	switch val[0] {
	case '>':
		if val[1] == '=' {
			v, err = f.Cast(val[2:])
			if err != nil {
				return
			}
			cond = builder.Gte{f.Name: v}
		} else {
			v, err = f.Cast(val[1:])
			if err != nil {
				return
			}
			cond = builder.Gt{f.Name: v}
		}
	case '<':
		if val[1] == '=' {
			v, err = f.Cast(val[2:])
			if err != nil {
				return
			}
			cond = builder.Lte{f.Name: v}
		} else {
			v, err = f.Cast(val[1:])
			if err != nil {
				return
			}
			cond = builder.Lt{f.Name: v}
		}
	case '=': //此处冗余了
		if val[1] == '=' {
			v, err = f.Cast(val[2:])
			if err != nil {
				return
			}
			cond = builder.Eq{f.Name: v}
		} else {
			v, err = f.Cast(val[1:])
			if err != nil {
				return
			}
			cond = builder.Eq{f.Name: v}
		}
	case '!', '~':
		if val[1] == '=' {
			v, err = f.Cast(val[2:])
			if err != nil {
				return
			}
			cond = builder.Neq{f.Name: v}
		} else {
			v, err = f.Cast(val[1:])
			if err != nil {
				return
			}
			cond = builder.Neq{f.Name: v}
		}
	case '%':
		v, err = f.Cast(val[2:])
		if err != nil {
			return
		}
		cond = builder.Like{f.Name, val}
	default:
		cond = builder.Eq{f.Name: v}
	}
	return
}

type Table struct {
	Name          string   `json:"name,omitempty"`
	Fields        []*Field `json:"fields,omitempty"`
	DisableInsert bool     `json:"disable_insert,omitempty"`
	DisableUpdate bool     `json:"disable_update,omitempty"`
	DisableDelete bool     `json:"disable_delete,omitempty"`

	//原生钩子
	Hook

	indexedFields map[string]*Field
}

func (t *Table) Init() error {
	t.indexedFields = make(map[string]*Field)
	for _, field := range t.Fields {
		t.indexedFields[field.Name] = field
	}

	return t.Hook.Compile()
}

func (t *Table) Field(name string) *Field {
	return t.indexedFields[name]
}

func (t *Table) PrimaryKeys() []*Field {
	var fields []*Field
	for _, field := range t.Fields {
		if field.PrimaryKey {
			fields = append(fields, field)
		}
	}
	return fields
}

func (t *Table) condId(id any) (builder.Cond, error) {
	//取id列
	field := t.Field("id")

	//取主键
	if field == nil {
		keys := t.PrimaryKeys()
		if len(keys) == 0 {
			return nil, errors.New("table has no primary key")
		}
		if len(keys) > 1 {
			return nil, errors.New("table has more than one primary key")
		}
		field = keys[0]
	}

	//转换（有需要的情况下，字符串转数值）
	val, err := field.Cast(id)
	if err != nil {
		return nil, err
	}

	return builder.Eq{field.Name: val}, nil
}

func (t *Table) condWhere(filter map[string]any) (conds []builder.Cond, err error) {
	for k, v := range filter {
		field := t.Field(k)
		if field == nil {
			return nil, fmt.Errorf("field %s not found", k)
		}

		switch val := v.(type) {
		case []string:
			for _, s := range val {
				cond, err := field.Condition(s)
				if err != nil {
					return nil, err
				}
				conds = append(conds, cond)
			}
		case string:
			cond, err := field.Condition(val)
			if err != nil {
				return nil, err
			}
			conds = append(conds, cond)
		default:
			conds = append(conds, builder.Eq{field.Name: val})
		}
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

	var increment bool

	for _, field := range t.Fields {
		//查询自增主键
		if field.PrimaryKey && field.Increment {
			increment = true
		}

		//主键，生成默认ID
		if field.PrimaryKey && field.Name == "id" && field.Type == "string" {
			if val, ok := values[field.Name]; ok {
				if v, ok := val.(string); ok && v == "" {
					id = xid.New().String()
					values[field.Name] = id
				}
			} else {
				id = xid.New().String()
				values[field.Name] = id
			}
		}

		if field.Created {
			values[field.Name] = time.Now()
		}
	}

	if t.BeforeInsert != nil {
		err = t.BeforeInsert(values)
		if err != nil {
			return
		}
	}

	var vs []interface{}
	for k, v := range values {
		vs = append(vs, builder.Eq{k: v})
	}
	bdr := builder.Insert(vs...).Into(t.Name)
	res, err := db.Engine().Exec(bdr)
	if err != nil {
		return id, err
	}

	//获取自增ID
	if increment {
		id, err = res.LastInsertId()
	}

	//_, err = db.Engine().Table(t.Name).Insert(values) 原始方式

	if t.AfterInsert != nil {
		err = t.AfterInsert(id, values)
	}

	return
}

func (t *Table) Update(filter map[string]any, values map[string]any) (rows int64, err error) {
	for _, field := range t.Fields {
		if field.Updated {
			values[field.Name] = time.Now()
		}
	}

	var updates []builder.Cond
	for k, v := range values {
		updates = append(updates, builder.Eq{k: v})
	}

	bdr := builder.Update(updates...).From(t.Name)

	cs, err := t.condWhere(filter)
	if err != nil {
		return 0, err
	}
	for _, c := range cs {
		bdr.Where(c)
	}

	res, err := db.Engine().Exec(bdr)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (t *Table) UpdateById(id any, values map[string]any) (rows int64, err error) {
	for _, field := range t.Fields {
		if field.Updated {
			values[field.Name] = time.Now()
		}
	}

	if t.BeforeUpdate != nil {
		err = t.BeforeUpdate(id, values)
		if err != nil {
			return
		}
	}

	var updates []builder.Cond
	for k, v := range values {
		updates = append(updates, builder.Eq{k: v})
	}
	bdr := builder.Update(updates...).From(t.Name)

	//bdr.Where(builder.Eq{"id": id})
	cond, err := t.condId(id)
	if err != nil {
		return 0, err
	}
	bdr.Where(cond)

	res, err := db.Engine().ID(id).Exec(bdr)
	if err != nil {
		return 0, err
	}

	if t.AfterUpdate != nil {
		err = t.AfterUpdate(id, values, values)
	}

	return res.RowsAffected()
}

func (t *Table) Delete(filter map[string]any) (rows int64, err error) {
	cs, err := t.condWhere(filter)
	if err != nil {
		return 0, err
	}

	bdr := builder.Delete(cs...).From(t.Name)

	res, err := db.Engine().Exec(bdr)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (t *Table) DeleteById(id any) (rows int64, err error) {
	bdr := builder.Delete().From(t.Name)

	if t.BeforeDelete != nil {
		err = t.BeforeDelete(id)
		if err != nil {
			return
		}
	}

	cond, err := t.condId(id)
	if err != nil {
		return 0, err
	}
	bdr.Where(cond)

	res, err := db.Engine().Exec(bdr)
	if err != nil {
		return 0, err
	}

	if t.AfterDelete != nil {
		err = t.AfterDelete(id, nil)
		if err != nil {
			return 0, err
		}
	}

	return res.RowsAffected()
}

func (t *Table) Find(filter map[string]any, fields []string, skip, limit int) (rows []map[string]any, err error) {
	bdr := builder.Select(fields...)

	bdr.From(t.Name)

	cs, err := t.condWhere(filter)
	if err != nil {
		return nil, err
	}
	for _, c := range cs {
		bdr.Where(c)
	}

	//bdr.OrderBy()

	bdr.Limit(skip, limit)

	return db.Engine().QueryInterface(bdr)
}

func (t *Table) Get(id any, fields []string) (Document, error) {
	bdr := builder.Select(fields...)

	bdr.From(t.Name)

	//bdr.Where(builder.Eq{"id": id})
	cond, err := t.condId(id)
	if err != nil {
		return nil, err
	}
	bdr.Where(cond)

	res, err := db.Engine().QueryInterface(bdr)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil //TODO 记录不存在
	}
	return res[0], nil
}

func (t *Table) Count(filter map[string]any) (cnt int64, err error) {
	bdr := builder.Select("count(*)").From(t.Name)

	cs, err := t.condWhere(filter)
	if err != nil {
		return 0, err
	}
	for _, c := range cs {
		bdr.Where(c)
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
