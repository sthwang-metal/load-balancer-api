// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package origin

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/x/gidx"
)

const (
	// Label holds the string label denoting the origin type in the database.
	Label = "origin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTarget holds the string denoting the target field in the database.
	FieldTarget = "target"
	// FieldPortNumber holds the string denoting the port_number field in the database.
	FieldPortNumber = "port_number"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldPoolID holds the string denoting the pool_id field in the database.
	FieldPoolID = "pool_id"
	// EdgePool holds the string denoting the pool edge name in mutations.
	EdgePool = "pool"
	// Table holds the table name of the origin in the database.
	Table = "origins"
	// PoolTable is the table that holds the pool relation/edge.
	PoolTable = "origins"
	// PoolInverseTable is the table name for the Pool entity.
	// It exists in this package in order to avoid circular dependency with the "pool" package.
	PoolInverseTable = "pools"
	// PoolColumn is the table column denoting the pool relation/edge.
	PoolColumn = "pool_id"
)

// Columns holds all SQL columns for origin fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldTarget,
	FieldPortNumber,
	FieldActive,
	FieldPoolID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// TargetValidator is a validator for the "target" field. It is called by the builders before save.
	TargetValidator func(string) error
	// PortNumberValidator is a validator for the "port_number" field. It is called by the builders before save.
	PortNumberValidator func(int) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// PoolIDValidator is a validator for the "pool_id" field. It is called by the builders before save.
	PoolIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() gidx.PrefixedID
)

// OrderOption defines the ordering options for the Origin queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTarget orders the results by the target field.
func ByTarget(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTarget, opts...).ToFunc()
}

// ByPortNumber orders the results by the port_number field.
func ByPortNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPortNumber, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByPoolID orders the results by the pool_id field.
func ByPoolID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoolID, opts...).ToFunc()
}

// ByPoolField orders the results by pool field.
func ByPoolField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPoolStep(), sql.OrderByField(field, opts...))
	}
}
func newPoolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PoolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PoolTable, PoolColumn),
	)
}