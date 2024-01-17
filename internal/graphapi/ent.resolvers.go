package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"

	"entgo.io/contrib/entgql"
	"go.infratographer.com/load-balancer-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// LoadBalancerPools is the resolver for the loadBalancerPools field.
func (r *queryResolver) LoadBalancerPools(ctx context.Context, after *entgql.Cursor[gidx.PrefixedID], first *int, before *entgql.Cursor[gidx.PrefixedID], last *int, orderBy *generated.LoadBalancerPoolOrder, where *generated.LoadBalancerPoolWhereInput) (*generated.LoadBalancerPoolConnection, error) {
	return r.client.Pool.Query().Paginate(ctx, after, first, before, last, generated.WithLoadBalancerPoolOrder(orderBy), generated.WithLoadBalancerPoolFilter(where.Filter))
}

// LoadBalancer returns LoadBalancerResolver implementation.
func (r *Resolver) LoadBalancer() LoadBalancerResolver { return &loadBalancerResolver{r} }

// LoadBalancerPool returns LoadBalancerPoolResolver implementation.
func (r *Resolver) LoadBalancerPool() LoadBalancerPoolResolver { return &loadBalancerPoolResolver{r} }

// LoadBalancerProvider returns LoadBalancerProviderResolver implementation.
func (r *Resolver) LoadBalancerProvider() LoadBalancerProviderResolver {
	return &loadBalancerProviderResolver{r}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type loadBalancerResolver struct{ *Resolver }
type loadBalancerPoolResolver struct{ *Resolver }
type loadBalancerProviderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
