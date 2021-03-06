/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package overview

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/vmware/octant/internal/describer"
	"github.com/vmware/octant/pkg/icon"
	"github.com/vmware/octant/pkg/store"
)

var (
	workloadsCronJobs = describer.NewResource(describer.ResourceOptions{
		Path:           "/workloads/cron-jobs",
		ObjectStoreKey: store.Key{APIVersion: "batch/v1beta1", Kind: "CronJob"},
		ListType:       &batchv1beta1.CronJobList{},
		ObjectType:     &batchv1beta1.CronJob{},
		Titles:         describer.ResourceTitle{List: "Workloads / Cron Jobs", Object: "Cron Job"},
		IconName:       icon.OverviewCronJob,
	})

	workloadsDaemonSets = describer.NewResource(describer.ResourceOptions{
		Path:           "/workloads/daemon-sets",
		ObjectStoreKey: store.Key{APIVersion: "apps/v1", Kind: "DaemonSet"},
		ListType:       &appsv1.DaemonSetList{},
		ObjectType:     &appsv1.DaemonSet{},
		Titles:         describer.ResourceTitle{List: "Workloads / Daemon Sets", Object: "Daemon Set"},
		IconName:       icon.OverviewDaemonSet,
	})

	workloadsDeployments = describer.NewResource(describer.ResourceOptions{
		Path:           "/workloads/deployments",
		ObjectStoreKey: store.Key{APIVersion: "apps/v1", Kind: "Deployment"},
		ListType:       &appsv1.DeploymentList{},
		ObjectType:     &appsv1.Deployment{},
		Titles:         describer.ResourceTitle{List: "Workloads / Deployments", Object: "Deployment"},
		IconName:       icon.OverviewDeployment,
	})

	workloadsJobs = describer.NewResource(describer.ResourceOptions{
		Path:           "/workloads/jobs",
		ObjectStoreKey: store.Key{APIVersion: "batch/v1", Kind: "Job"},
		ListType:       &batchv1.JobList{},
		ObjectType:     &batchv1.Job{},
		Titles:         describer.ResourceTitle{List: "Workloads / Jobs", Object: "Job"},
		IconName:       icon.OverviewJob,
	})

	workloadsPods = describer.NewResource(describer.ResourceOptions{
		Path:           "/workloads/pods",
		ObjectStoreKey: store.Key{APIVersion: "v1", Kind: "Pod"},
		ListType:       &corev1.PodList{},
		ObjectType:     &corev1.Pod{},
		Titles:         describer.ResourceTitle{List: "Workloads / Pods", Object: "Pod"},
		IconName:       icon.OverviewPod,
	})

	workloadsReplicaSets = describer.NewResource(describer.ResourceOptions{
		Path:           "/workloads/replica-sets",
		ObjectStoreKey: store.Key{APIVersion: "apps/v1", Kind: "ReplicaSet"},
		ListType:       &appsv1.ReplicaSetList{},
		ObjectType:     &appsv1.ReplicaSet{},
		Titles:         describer.ResourceTitle{List: "Workloads / Replica Sets", Object: "Replica Set"},
		IconName:       icon.OverviewReplicaSet,
	})

	workloadsReplicationControllers = describer.NewResource(describer.ResourceOptions{
		Path:           "/workloads/replication-controllers",
		ObjectStoreKey: store.Key{APIVersion: "v1", Kind: "ReplicationController"},
		ListType:       &corev1.ReplicationControllerList{},
		ObjectType:     &corev1.ReplicationController{},
		Titles:         describer.ResourceTitle{List: "Workloads / Replication Controllers", Object: "Replication Controller"},
		IconName:       icon.OverviewReplicationController,
	})
	workloadsStatefulSets = describer.NewResource(describer.ResourceOptions{
		Path:           "/workloads/stateful-sets",
		ObjectStoreKey: store.Key{APIVersion: "apps/v1", Kind: "StatefulSet"},
		ListType:       &appsv1.StatefulSetList{},
		ObjectType:     &appsv1.StatefulSet{},
		Titles:         describer.ResourceTitle{List: "Workloads / Stateful Sets", Object: "Stateful Set"},
		IconName:       icon.OverviewStatefulSet,
	})

	workloadsDescriber = describer.NewSection(
		"/workloads",
		"Workloads",
		workloadsCronJobs,
		workloadsDaemonSets,
		workloadsDeployments,
		workloadsJobs,
		workloadsPods,
		workloadsReplicaSets,
		workloadsReplicationControllers,
		workloadsStatefulSets,
	)

	dlbIngresses = describer.NewResource(describer.ResourceOptions{
		Path:           "/discovery-and-load-balancing/ingresses",
		ObjectStoreKey: store.Key{APIVersion: "extensions/v1beta1", Kind: "Ingress"},
		ListType:       &v1beta1.IngressList{},
		ObjectType:     &v1beta1.Ingress{},
		Titles:         describer.ResourceTitle{List: "Discovery & Load Balancing / Ingresses", Object: "Ingress"},
		IconName:       icon.OverviewIngress,
	})

	dlbServices = describer.NewResource(describer.ResourceOptions{
		Path:           "/discovery-and-load-balancing/services",
		ObjectStoreKey: store.Key{APIVersion: "v1", Kind: "Service"},
		ListType:       &corev1.ServiceList{},
		ObjectType:     &corev1.Service{},
		Titles:         describer.ResourceTitle{List: "Discovery & Load Balancing / Services", Object: "Service"},
		IconName:       icon.OverviewService,
	})

	discoveryAndLoadBalancingDescriber = describer.NewSection(
		"/discovery-and-load-balancing",
		"Discovery and Load Balancing",
		dlbIngresses,
		dlbServices,
	)

	csConfigMaps = describer.NewResource(describer.ResourceOptions{
		Path:           "/config-and-storage/config-maps",
		ObjectStoreKey: store.Key{APIVersion: "v1", Kind: "ConfigMap"},
		ListType:       &corev1.ConfigMapList{},
		ObjectType:     &corev1.ConfigMap{},
		Titles:         describer.ResourceTitle{List: "Config & Storage / Config Maps", Object: "Config Map"},
		IconName:       icon.OverviewConfigMap,
	})

	csPVCs = describer.NewResource(describer.ResourceOptions{
		Path:           "/config-and-storage/persistent-volume-claims",
		ObjectStoreKey: store.Key{APIVersion: "v1", Kind: "PersistentVolumeClaim"},
		ListType:       &corev1.PersistentVolumeClaimList{},
		ObjectType:     &corev1.PersistentVolumeClaim{},
		Titles:         describer.ResourceTitle{List: "Config & Storage / Persistent Volume Claims", Object: "Persistent Volume Claim"},
		IconName:       icon.OverviewPersistentVolumeClaim,
	})

	csSecrets = describer.NewResource(describer.ResourceOptions{
		Path:           "/config-and-storage/secrets",
		ObjectStoreKey: store.Key{APIVersion: "v1", Kind: "Secret"},
		ListType:       &corev1.SecretList{},
		ObjectType:     &corev1.Secret{},
		Titles:         describer.ResourceTitle{List: "Config & Storage / Secrets", Object: "Secret"},
		IconName:       icon.OverviewSecret,
	})

	csServiceAccounts = describer.NewResource(describer.ResourceOptions{
		Path:           "/config-and-storage/service-accounts",
		ObjectStoreKey: store.Key{APIVersion: "v1", Kind: "ServiceAccount"},
		ListType:       &corev1.ServiceAccountList{},
		ObjectType:     &corev1.ServiceAccount{},
		Titles:         describer.ResourceTitle{List: "Config & Storage / Service Accounts", Object: "Service Account"},
		IconName:       icon.OverviewServiceAccount,
	})

	configAndStorageDescriber = describer.NewSection(
		"/config-and-storage",
		"Config and Storage",
		csConfigMaps,
		csPVCs,
		csSecrets,
		csServiceAccounts,
	)

	customResourcesDescriber = describer.NewCRDSection(
		"/custom-resources",
		"Custom Resources",
	)

	rbacRoles = describer.NewResource(describer.ResourceOptions{
		Path:           "/rbac/roles",
		ObjectStoreKey: store.Key{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "Role"},
		ListType:       &rbacv1.RoleList{},
		ObjectType:     &rbacv1.Role{},
		Titles:         describer.ResourceTitle{List: "RBAC / Roles", Object: "Role"},
		IconName:       icon.OverviewRole,
	})

	rbacRoleBindings = describer.NewResource(describer.ResourceOptions{
		Path:           "/rbac/role-bindings",
		ObjectStoreKey: store.Key{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "RoleBinding"},
		ListType:       &rbacv1.RoleBindingList{},
		ObjectType:     &rbacv1.RoleBinding{},
		Titles:         describer.ResourceTitle{List: "RBAC / Role Bindings", Object: "Role Binding"},
		IconName:       icon.OverviewRoleBinding,
	})

	rbacDescriber = describer.NewSection(
		"/rbac",
		"RBAC",
		rbacRoles,
		rbacRoleBindings,
	)

	rootDescriber = describer.NewSection(
		"/",
		"Overview",
		workloadsDescriber,
		discoveryAndLoadBalancingDescriber,
		configAndStorageDescriber,
		customResourcesDescriber,
		rbacDescriber,
	)

	eventsDescriber = describer.NewResource(describer.ResourceOptions{
		Path:                  "/events",
		ObjectStoreKey:        store.Key{APIVersion: "v1", Kind: "Event"},
		ListType:              &corev1.EventList{},
		ObjectType:            &corev1.Event{},
		Titles:                describer.ResourceTitle{List: "Events", Object: "Event"},
		DisableResourceViewer: true,
	})
)
