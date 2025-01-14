// Copyright (C) 2021 ScyllaDB

package identity

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	"github.com/scylladb/scylla-operator/pkg/controllerhelpers"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func TestMember_GetSeeds(t *testing.T) {
	t.Parallel()

	createPodAndSvc := func(name, ip string, creationTimestamp time.Time) (*corev1.Pod, *corev1.Service) {
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: "namespace",
				Labels: map[string]string{
					"scylla/cluster":               "my-cluster",
					"app":                          "scylla",
					"app.kubernetes.io/name":       "scylla",
					"app.kubernetes.io/managed-by": "scylla-operator",
					"scylla/rack-ordinal":          "0",
				},
				CreationTimestamp: metav1.NewTime(creationTimestamp),
			},
		}
		svc := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: "namespace",
			},
			Spec: corev1.ServiceSpec{
				ClusterIP: ip,
			},
		}
		return pod, svc
	}

	now := time.Now()
	firstPod, firstService := createPodAndSvc("pod-0", "1.1.1.1", now)
	secondPod, secondService := createPodAndSvc("pod-1", "2.2.2.2", now.Add(time.Second))
	thirdPod, thirdService := createPodAndSvc("pod-2", "3.3.3.3", now.Add(2*time.Second))

	ts := []struct {
		name                       string
		memberService              *corev1.Service
		memberPod                  *corev1.Pod
		memberClientsBroadcastType scyllav1.BroadcastAddressType
		memberNodesBroadcastType   scyllav1.BroadcastAddressType
		externalSeeds              []string
		objects                    []runtime.Object
		expectSeeds                []string
		expectError                error
	}{
		{
			name:                       "error when no pods are found",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{},
			expectError:                fmt.Errorf("internal error: can't find any pod for this cluster, including itself"),
		},
		{
			name:                       "bootstrap with external seeds only when cluster is empty and external seeds are provided",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService},
			externalSeeds:              []string{"10.0.1.1", "10.0.1.2"},
			expectSeeds:                []string{"10.0.1.1", "10.0.1.2"},
		},
		{
			name:                       "bootstrap with external seeds and first created UN node when external seeds are provided",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService, markPodReady(secondPod), secondService, markPodReady(thirdPod), thirdService},
			externalSeeds:              []string{"10.0.1.1", "10.0.1.2"},
			expectSeeds:                []string{"10.0.1.1", "10.0.1.2", secondService.Spec.ClusterIP},
		},
		{
			name:                       "bootstrap with external seeds and first created Pod when all Pods from DC are down and external seeds are provided",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService, secondPod, secondService, thirdPod, thirdService},
			externalSeeds:              []string{"10.0.1.1", "10.0.1.2"},
			expectSeeds:                []string{"10.0.1.1", "10.0.1.2", secondService.Spec.ClusterIP},
		},
		{
			name:                       "bootstraps with itself using node-to-node identifier of ClusterIP type when cluster is empty",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService},
			expectSeeds:                []string{firstService.Spec.ClusterIP},
		},
		{
			name:                       "bootstraps with itself using node-to-node identifier of ClusterIP type when cluster is empty",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService},
			expectSeeds:                []string{firstService.Spec.ClusterIP},
		},
		{
			name: "bootstraps with itself using node-to-node identifier of PodIP type when cluster is empty",
			memberPod: func() *corev1.Pod {
				pod := firstPod.DeepCopy()
				pod.Status.PodIP = "1.2.3.4"
				return pod
			}(),
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypePodIP,
			objects:                    []runtime.Object{firstPod, firstService},
			expectSeeds:                []string{"1.2.3.4"},
		},
		{
			name:      "bootstraps with itself using node-to-node identifier of LoadBalancer External IP type when cluster is empty",
			memberPod: firstPod,
			memberService: func() *corev1.Service {
				svc := firstService.DeepCopy()
				svc.Spec.Type = corev1.ServiceTypeLoadBalancer
				svc.Status.LoadBalancer = corev1.LoadBalancerStatus{
					Ingress: []corev1.LoadBalancerIngress{
						{
							IP: "4.3.2.1",
						},
					},
				}
				return svc
			}(),
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceLoadBalancerIngress,
			objects:                    []runtime.Object{firstPod, firstService},
			expectSeeds:                []string{"4.3.2.1"},
		},
		{
			name:      "bootstraps with itself using node-to-node identifier of LoadBalancer hostname when cluster is empty",
			memberPod: firstPod,
			memberService: func() *corev1.Service {
				svc := firstService.DeepCopy()
				svc.Spec.Type = corev1.ServiceTypeLoadBalancer
				svc.Status.LoadBalancer = corev1.LoadBalancerStatus{
					Ingress: []corev1.LoadBalancerIngress{
						{
							Hostname: "node-1-hostname.scylladb.com",
						},
					},
				}
				return svc
			}(),
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceLoadBalancerIngress,
			objects:                    []runtime.Object{firstPod, firstService},
			expectSeeds:                []string{"node-1-hostname.scylladb.com"},
		},
		{
			name:                       "bootstrap with first created UN node",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService, markPodReady(secondPod), secondService, markPodReady(thirdPod), thirdService},
			expectSeeds:                []string{secondService.Spec.ClusterIP},
		},
		{
			name:                       "bootstrap only with UN node",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService, secondPod, secondService, markPodReady(thirdPod), thirdService},
			expectSeeds:                []string{thirdService.Spec.ClusterIP},
		},
		{
			name:                       "bootstrap with first created Pod when all are down",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService, secondPod, secondService, thirdPod, thirdService},
			expectSeeds:                []string{secondService.Spec.ClusterIP},
		},
		{
			name:                       "use PodIP from status when node broadcast address type is PodIP",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypePodIP,
			objects: []runtime.Object{
				firstPod,
				firstService,
				func() runtime.Object {
					pod := secondPod.DeepCopy()
					pod.Status.PodIP = "1.2.3.4"
					pod = markPodReady(pod)
					return pod
				}(),
				func() runtime.Object {
					svc := secondService.DeepCopy()
					svc.Spec.ClusterIP = corev1.ClusterIPNone
					return svc
				}(),
				thirdPod,
				thirdService,
			},
			expectSeeds: []string{"1.2.3.4"},
		},
		{
			name:                       "use ClusterIP from Service when node broadcast address type is ClusterIP",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects: []runtime.Object{
				firstPod,
				firstService,
				secondPod,
				func() runtime.Object {
					svc := secondService.DeepCopy()
					svc.Spec.ClusterIP = "1.2.3.4"
					return svc
				}(),
				thirdPod,
				thirdService,
			},
			expectSeeds: []string{"1.2.3.4"},
		},
		{
			name:                       "use preferred IP address from first Service ingress status when node broadcast address type is LoadBalancer Ingress",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceLoadBalancerIngress,
			objects: []runtime.Object{
				firstPod,
				firstService,
				secondPod,
				func() runtime.Object {
					svc := secondService.DeepCopy()
					svc.Status.LoadBalancer = corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{
								IP:       "1.2.3.4",
								Hostname: "second.service.scylladb.com",
							},
						},
					}
					return svc
				}(),
				thirdPod,
				thirdService,
			},
			expectSeeds: []string{"1.2.3.4"},
		},
		{
			name:                       "use hostname from first Service ingress status when node broadcast address type is LoadBalancer Ingress and IP is not available",
			memberPod:                  firstPod,
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceLoadBalancerIngress,
			objects: []runtime.Object{
				firstPod,
				firstService,
				secondPod,
				func() runtime.Object {
					svc := secondService.DeepCopy()
					svc.Status.LoadBalancer = corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{
								Hostname: "second.service.scylladb.com",
							},
						},
					}
					return svc
				}(),
				thirdPod,
				thirdService,
			},
			expectSeeds: []string{"second.service.scylladb.com"},
		},
		{
			name:                       "error when node is not first in rack, but there are no other pods",
			memberPod:                  secondPod,
			memberService:              secondService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{secondPod, secondService},
			expectSeeds:                nil,
			expectError:                fmt.Errorf("pod is not first in the cluster, but there are no other pods"),
		},
		{
			name: "error when node is first in rack of ordinal > 0, but there are no other pods",
			memberPod: func() *corev1.Pod {
				pod := firstPod.DeepCopy()
				pod.Labels["scylla/rack-ordinal"] = "1"
				return pod
			}(),
			memberService:              firstService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService},
			expectSeeds:                nil,
			expectError:                fmt.Errorf("pod is not first in the cluster, but there are no other pods"),
		},
		{
			name: "bootstrap with other pod when node is first in rack of ordinal > 0 and there are other pods",
			memberPod: func() *corev1.Pod {
				pod := secondPod.DeepCopy()
				pod.Labels["scylla/rack-ordinal"] = "1"
				return pod
			}(),
			memberService:              secondService,
			memberClientsBroadcastType: scyllav1.BroadcastAddressTypeServiceClusterIP,
			memberNodesBroadcastType:   scyllav1.BroadcastAddressTypeServiceClusterIP,
			objects:                    []runtime.Object{firstPod, firstService, secondPod, secondService},
			expectSeeds:                []string{firstService.Spec.ClusterIP},
		},
	}

	for i := range ts {
		test := ts[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			member, err := NewMember(test.memberService, test.memberPod, test.memberNodesBroadcastType, test.memberClientsBroadcastType)
			if err != nil {
				t.Fatal(err)
			}

			fakeClient := fake.NewSimpleClientset(test.objects...)
			seeds, err := member.GetSeeds(ctx, fakeClient.CoreV1(), test.externalSeeds)
			if !reflect.DeepEqual(err, test.expectError) {
				t.Errorf("expected error %v, got %v", test.expectError, err)
			}
			if !reflect.DeepEqual(seeds, test.expectSeeds) {
				t.Errorf("expected seeds %v, got %v", test.expectSeeds, seeds)
			}
		})
	}
}

func markPodReady(pod *corev1.Pod) *corev1.Pod {
	p := pod.DeepCopy()
	cond := controllerhelpers.GetPodCondition(p.Status.Conditions, corev1.PodReady)
	if cond != nil {
		cond.Status = corev1.ConditionTrue
		return p
	}

	p.Status.Conditions = append(p.Status.Conditions, corev1.PodCondition{
		Type:   corev1.PodReady,
		Status: corev1.ConditionTrue,
	})

	return p
}
