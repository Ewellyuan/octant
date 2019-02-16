package resourceviewer

import (
	"net/url"
	"testing"

	"github.com/heptio/developer-dash/internal/overview/link"

	"github.com/heptio/developer-dash/internal/conversion"

	"github.com/heptio/developer-dash/internal/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

func Test_Collector(t *testing.T) {
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{APIVersion: "apps/v1", Kind: "Deployment"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "deployment",
			UID:  types.UID("deployment"),
		},
		Status: appsv1.DeploymentStatus{
			Replicas:          1,
			AvailableReplicas: 1,
		},
	}

	replicaSet1 := &extv1beta1.ReplicaSet{
		TypeMeta: metav1.TypeMeta{APIVersion: "extensions/v1beta1", Kind: "ReplicaSet"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "replicaSet1",
			UID:  types.UID("replicaSet1"),
		},
		Spec: extv1beta1.ReplicaSetSpec{
			Replicas: conversion.PtrInt32(1),
		},
		Status: extv1beta1.ReplicaSetStatus{
			Replicas:          1,
			AvailableReplicas: 1,
		},
	}

	replicaSet2 := &extv1beta1.ReplicaSet{
		TypeMeta: metav1.TypeMeta{APIVersion: "extensions/v1beta1", Kind: "ReplicaSet"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "replicaSet2",
			UID:  types.UID("replicaSet2"),
		},
	}

	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "Pod"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod",
			UID:  types.UID("pod"),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(replicaSet1,
					schema.FromAPIVersionAndKind(replicaSet1.APIVersion,
						replicaSet1.Kind)),
			},
		},
	}

	c := NewCollector()

	err := c.Process(deployment)
	require.NoError(t, err)

	err = c.Process(replicaSet1)
	require.NoError(t, err)

	err = c.Process(replicaSet2)
	require.NoError(t, err)

	err = c.Process(pod)
	require.NoError(t, err)

	err = c.AddChild(deployment, replicaSet1, replicaSet2)
	require.NoError(t, err)

	err = c.AddChild(replicaSet1, pod)
	require.NoError(t, err)

	got, err := c.ViewComponent("deployment")
	require.NoError(t, err)

	q := url.Values{}
	q.Set("view", "summary")

	expected := component.NewResourceViewer("Resource Viewer")
	expected.AddEdge("deployment", "replicaSet1", component.EdgeTypeExplicit)
	expected.AddEdge("replicaSet1", "pods-replicaSet1", component.EdgeTypeExplicit)
	expected.AddNode("deployment", component.Node{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Name:       "deployment",
		Status:     "ok",
		Details:    component.Title(component.NewText("Deployment is OK")),
		Path:       link.ForObjectWithQuery(deployment, deployment.Name, q),
	})

	expected.AddNode("replicaSet1", component.Node{
		APIVersion: "extensions/v1beta1",
		Kind:       "ReplicaSet",
		Name:       "replicaSet1",
		Status:     "ok",
		Details:    component.Title(component.NewText("Replica Set is OK")),
		Path:       link.ForObjectWithQuery(replicaSet1, replicaSet1.Name, q),
	})
	expected.AddNode("pods-replicaSet1", component.Node{
		APIVersion: "v1",
		Kind:       "Pod",
		Name:       "replicaSet1 pods",
		Status:     "ok",
		Details:    component.Title(component.NewText("Pod count: 1")),
	})
	expected.Select("deployment")

	assert.Equal(t, expected, got)
}