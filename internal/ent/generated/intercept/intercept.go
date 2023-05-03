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

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/load-balancer-api/internal/ent/generated"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancer"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancerannotation"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancerstatus"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/origin"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/pool"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/port"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/provider"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next generated.Querier) generated.Querier {
	return generated.QuerierFunc(func(ctx context.Context, q generated.Query) (generated.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q generated.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The LoadBalancerFunc type is an adapter to allow the use of ordinary function as a Querier.
type LoadBalancerFunc func(context.Context, *generated.LoadBalancerQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f LoadBalancerFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.LoadBalancerQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.LoadBalancerQuery", q)
}

// The TraverseLoadBalancer type is an adapter to allow the use of ordinary function as Traverser.
type TraverseLoadBalancer func(context.Context, *generated.LoadBalancerQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseLoadBalancer) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseLoadBalancer) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.LoadBalancerQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.LoadBalancerQuery", q)
}

// The LoadBalancerAnnotationFunc type is an adapter to allow the use of ordinary function as a Querier.
type LoadBalancerAnnotationFunc func(context.Context, *generated.LoadBalancerAnnotationQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f LoadBalancerAnnotationFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.LoadBalancerAnnotationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.LoadBalancerAnnotationQuery", q)
}

// The TraverseLoadBalancerAnnotation type is an adapter to allow the use of ordinary function as Traverser.
type TraverseLoadBalancerAnnotation func(context.Context, *generated.LoadBalancerAnnotationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseLoadBalancerAnnotation) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseLoadBalancerAnnotation) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.LoadBalancerAnnotationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.LoadBalancerAnnotationQuery", q)
}

// The LoadBalancerStatusFunc type is an adapter to allow the use of ordinary function as a Querier.
type LoadBalancerStatusFunc func(context.Context, *generated.LoadBalancerStatusQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f LoadBalancerStatusFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.LoadBalancerStatusQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.LoadBalancerStatusQuery", q)
}

// The TraverseLoadBalancerStatus type is an adapter to allow the use of ordinary function as Traverser.
type TraverseLoadBalancerStatus func(context.Context, *generated.LoadBalancerStatusQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseLoadBalancerStatus) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseLoadBalancerStatus) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.LoadBalancerStatusQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.LoadBalancerStatusQuery", q)
}

// The OriginFunc type is an adapter to allow the use of ordinary function as a Querier.
type OriginFunc func(context.Context, *generated.OriginQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f OriginFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.OriginQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.OriginQuery", q)
}

// The TraverseOrigin type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOrigin func(context.Context, *generated.OriginQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOrigin) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOrigin) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OriginQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.OriginQuery", q)
}

// The PoolFunc type is an adapter to allow the use of ordinary function as a Querier.
type PoolFunc func(context.Context, *generated.PoolQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f PoolFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.PoolQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.PoolQuery", q)
}

// The TraversePool type is an adapter to allow the use of ordinary function as Traverser.
type TraversePool func(context.Context, *generated.PoolQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePool) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePool) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.PoolQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.PoolQuery", q)
}

// The PortFunc type is an adapter to allow the use of ordinary function as a Querier.
type PortFunc func(context.Context, *generated.PortQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f PortFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.PortQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.PortQuery", q)
}

// The TraversePort type is an adapter to allow the use of ordinary function as Traverser.
type TraversePort func(context.Context, *generated.PortQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePort) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePort) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.PortQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.PortQuery", q)
}

// The ProviderFunc type is an adapter to allow the use of ordinary function as a Querier.
type ProviderFunc func(context.Context, *generated.ProviderQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f ProviderFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.ProviderQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.ProviderQuery", q)
}

// The TraverseProvider type is an adapter to allow the use of ordinary function as Traverser.
type TraverseProvider func(context.Context, *generated.ProviderQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseProvider) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseProvider) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.ProviderQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.ProviderQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q generated.Query) (Query, error) {
	switch q := q.(type) {
	case *generated.LoadBalancerQuery:
		return &query[*generated.LoadBalancerQuery, predicate.LoadBalancer, loadbalancer.OrderOption]{typ: generated.TypeLoadBalancer, tq: q}, nil
	case *generated.LoadBalancerAnnotationQuery:
		return &query[*generated.LoadBalancerAnnotationQuery, predicate.LoadBalancerAnnotation, loadbalancerannotation.OrderOption]{typ: generated.TypeLoadBalancerAnnotation, tq: q}, nil
	case *generated.LoadBalancerStatusQuery:
		return &query[*generated.LoadBalancerStatusQuery, predicate.LoadBalancerStatus, loadbalancerstatus.OrderOption]{typ: generated.TypeLoadBalancerStatus, tq: q}, nil
	case *generated.OriginQuery:
		return &query[*generated.OriginQuery, predicate.Origin, origin.OrderOption]{typ: generated.TypeOrigin, tq: q}, nil
	case *generated.PoolQuery:
		return &query[*generated.PoolQuery, predicate.Pool, pool.OrderOption]{typ: generated.TypePool, tq: q}, nil
	case *generated.PortQuery:
		return &query[*generated.PortQuery, predicate.Port, port.OrderOption]{typ: generated.TypePort, tq: q}, nil
	case *generated.ProviderQuery:
		return &query[*generated.ProviderQuery, predicate.Provider, provider.OrderOption]{typ: generated.TypeProvider, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}