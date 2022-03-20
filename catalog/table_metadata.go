// this code is from https://github.com/brunocalza/go-bustub
// there is license and copyright notice in licenses/go-bustub dir

package catalog

import (
	"github.com/ryogrid/SamehadaDB/storage/access"
	"github.com/ryogrid/SamehadaDB/storage/index"
	"github.com/ryogrid/SamehadaDB/storage/table/schema"
)

type TableMetadata struct {
	schema *schema.Schema
	name   string
	table  *access.TableHeap
	// index data class obj of each column
	// if column has no index, respond element is nil
	indexes []*index.Index
	oid     uint32
}

func NewTableMetadata(schema *schema.Schema, name string, table *access.TableHeap, oid uint32) *TableMetadata {
	ret := new(TableMetadata)
	ret.schema = schema
	ret.name = name
	ret.table = table
	ret.oid = oid

	// TODO: (SDB) need initialize indexes member according to schema at NewTableMetaData
	return ret
}

func (t *TableMetadata) Schema() *schema.Schema {
	return t.schema
}

func (t *TableMetadata) OID() uint32 {
	return t.oid
}

func (t *TableMetadata) Table() *access.TableHeap {
	return t.table
}
