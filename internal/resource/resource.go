package resource

import (
	"context"
	"time"
)

type ResourceKind string

const (
	ResourceKindVM        ResourceKind = "vm"
	ResourceKindContainer ResourceKind = "container"
	ResourceKindLXC       ResourceKind = "lxc"
	ResourceKindNetwork   ResourceKind = "network"
	ResourceKindVolume    ResourceKind = "volume"
	ResourceKindSnapshot  ResourceKind = "snapshot"
	ResourceKindSecret    ResourceKind = "secret"
	ResourceKindNode      ResourceKind = "node"
)

type Resource[TSpec any, TStatus any] struct {
	ID       string       `json:"id"`
	Kind     ResourceKind `json:"kind"`
	Metadata Metadata     `json:"metadata"`
	Spec     TSpec        `json:"spec"`
	Status   TStatus      `json:"status"`
}

type Driver[TSpec any, TStatus any] interface {
	Create(ctx context.Context, res Resource[TSpec, TStatus]) error
	Start(ctx context.Context, name string) error
	Stop(ctx context.Context, name string) error
	Delete(ctx context.Context, name string) error
	Snapshot(ctx context.Context, name string) error
	Inspect(ctx context.Context, name string) (Resource[TSpec, TStatus], error)
	Reconcile(ctx context.Context, name string) (Resource[TSpec, TStatus], error)
}

type Metadata struct {
	Name        string            `db:"name" json:"name"`
	Namespace   string            `db:"namespace" json:"namespace,omitempty"`
	Labels      map[string]string `db:"labels,omitempty" json:"labels,omitempty"`
	Annotations map[string]string `db:"annotations,omitempty" json:"annotations,omitempty"`
	CreatedAt   time.Time         `db:"created_at" json:"created_at"`
	DependsOn   []ResourceRef     `db:"depends_on,omitempty" json:"depends_on,omitempty"`
	Owner       string            `db:"owner,omitempty" json:"owner,omitempty"`
}

type ResourceRef struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type Condition struct {
	Type               string    `db:"type" json:"type"`
	Status             string    `db:"status" json:"status"`
	LastTransitionTime time.Time `db:"last_transition_time" json:"last_transition_time"`
	Reason             string    `db:"reason" json:"reason"`
	Message            string    `db:"message" json:"message"`
}

type IResource interface {
	GetKind() ResourceKind
	GetMetadata() *Metadata
}
