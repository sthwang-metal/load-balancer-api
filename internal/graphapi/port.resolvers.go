package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"strings"

	"go.infratographer.com/load-balancer-api/internal/ent/generated"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/gidx"
)

// LoadBalancerPortCreate is the resolver for the loadBalancerPortCreate field.
func (r *mutationResolver) LoadBalancerPortCreate(ctx context.Context, input generated.CreateLoadBalancerPortInput) (*LoadBalancerPortCreatePayload, error) {
	if err := permissions.CheckAccess(ctx, input.LoadBalancerID, actionLoadBalancerUpdate); err != nil {
		return nil, err
	}

	lb, err := r.client.LoadBalancer.Get(ctx, input.LoadBalancerID)
	if err != nil {
		return nil, err
	}

	for _, poolId := range input.PoolIDs {
		if err := permissions.CheckAccess(ctx, poolId, actionLoadBalancerPoolGet); err != nil {
			return nil, err
		}

		pool, err := r.client.Pool.Get(ctx, poolId)
		if err != nil {
			return nil, err
		}

		if lb.OwnerID != pool.OwnerID {
			return nil, ErrOwnerConflict
		}
	}

	p, err := r.client.Port.Create().SetInput(input).Save(ctx)
	if err != nil {
		switch {
		case generated.IsConstraintError(err) && strings.Contains(err.Error(), "number"):
			return nil, ErrPortNumberInUse
		default:
			return nil, err
		}
	}

	return &LoadBalancerPortCreatePayload{LoadBalancerPort: p}, nil
}

// LoadBalancerPortUpdate is the resolver for the loadBalancerPortUpdate field.
func (r *mutationResolver) LoadBalancerPortUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateLoadBalancerPortInput) (*LoadBalancerPortUpdatePayload, error) {
	p, err := r.client.Port.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := permissions.CheckAccess(ctx, p.LoadBalancerID, actionLoadBalancerUpdate); err != nil {
		return nil, err
	}

	lb, err := r.client.LoadBalancer.Get(ctx, p.LoadBalancerID)
	if err != nil {
		return nil, err
	}

	for _, poolId := range input.AddPoolIDs {
		if err := permissions.CheckAccess(ctx, poolId, actionLoadBalancerPoolGet); err != nil {
			return nil, err
		}

		pool, err := r.client.Pool.Get(ctx, poolId)
		if err != nil {
			return nil, err
		}

		if lb.OwnerID != pool.OwnerID {
			return nil, ErrOwnerConflict
		}
	}

	p, err = p.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsConstraintError(err) && strings.Contains(err.Error(), "number") {
			return nil, ErrPortNumberInUse
		} else {
			return nil, err
		}
	}

	return &LoadBalancerPortUpdatePayload{LoadBalancerPort: p}, nil
}

// LoadBalancerPortDelete is the resolver for the loadBalancerPortDelete field.
func (r *mutationResolver) LoadBalancerPortDelete(ctx context.Context, id gidx.PrefixedID) (*LoadBalancerPortDeletePayload, error) {
	p, err := r.client.Port.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := permissions.CheckAccess(ctx, p.LoadBalancerID, actionLoadBalancerUpdate); err != nil {
		return nil, err
	}

	if err := r.client.Port.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}

	return &LoadBalancerPortDeletePayload{DeletedID: id}, nil
}
