package graphapi_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	ent "go.infratographer.com/load-balancer-api/internal/ent/generated"
	"go.infratographer.com/load-balancer-api/internal/graphapi"
	"go.infratographer.com/load-balancer-api/internal/graphclient"
	"go.uber.org/zap"
)

func TestFullLoadBalancerPortLifecycle(t *testing.T) {
	ctx := context.Background()
	graphClient := graphclient.New(graphapi.NewResolver(EntClient, zap.NewNop().Sugar()))

	lb := (&LoadBalancerBuilder{}).MustNew(ctx)
	name := gofakeit.DomainName()

	createdPort, err := graphClient.LoadBalancerPortCreate(ent.CreateLoadBalancerPortInput{
		Name:           name,
		LoadBalancerID: lb.ID,
		Number:         22,
	})

	require.NoError(t, err)
	require.NotNil(t, createdPort)
	require.NotNil(t, createdPort.ID)
	require.Equal(t, name, createdPort.Name)
	require.Equal(t, 22, createdPort.Number)
	assert.Equal(t, "loadprt", createdPort.ID.Prefix())
	assert.Equal(t, lb.ID, createdPort.LoadBalancerID.ID)

	// Update the Port
	newPort := gofakeit.Number(1, 65535)
	updatedPort, err := graphClient.LoadBalancerPortUpdate(createdPort.ID, ent.UpdateLoadBalancerPortInput{Number: &newPort})

	require.NoError(t, err)
	require.NotNil(t, updatedPort)
	require.EqualValues(t, createdPort.ID, updatedPort.ID)
	require.Equal(t, newPort, updatedPort.Number)

	// Query the Port
	queryPort, err := graphClient.QueryLoadBalancerPortByID(lb.ID, createdPort.ID)
	require.NoError(t, err)
	require.NotNil(t, queryPort)
	require.Equal(t, newPort, queryPort.Number)

	// Delete the Port
	deletedID, err := graphClient.LoadBalancerPortDelete(createdPort.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedID)
	require.Equal(t, createdPort.ID, deletedID)

	// Query the Port
	queryPort, err = graphClient.QueryLoadBalancerPortByID(lb.ID, createdPort.ID)
	require.NoError(t, err)
	require.Nil(t, queryPort)
}