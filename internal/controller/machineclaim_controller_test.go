// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	. "sigs.k8s.io/controller-runtime/pkg/envtest/komega"

	metalv1alpha1 "github.com/ironcore-dev/metal/api/v1alpha1"
)

var _ = Describe("MachineClaim Controller", Serial, func() {
	var ns *v1.Namespace

	BeforeEach(func(ctx SpecContext) {
		ns = &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
			},
		}
		Expect(k8sClient.Create(ctx, ns)).To(Succeed())
		DeferCleanup(func(ctx SpecContext) {
			Expect(k8sClient.Delete(ctx, ns)).To(Succeed())
		})

		Eventually(ObjectList(&metalv1alpha1.MachineList{})).Should(HaveField("Items", HaveLen(0)))
		Eventually(ObjectList(&metalv1alpha1.MachineClaimList{}, &client.ListOptions{
			Namespace: ns.Name,
		})).Should(HaveField("Items", HaveLen(0)))

		DeferCleanup(func(ctx SpecContext) {
			Eventually(ctx, func(g Gomega, ctx SpecContext) {
				var machines metalv1alpha1.MachineList
				g.Expect(ObjectList(&machines)()).To(SatisfyAll())
				if len(machines.Items) > 0 {
					g.Expect(k8sClient.DeleteAllOf(ctx, &machines.Items[0])).To(Succeed())
				}
				var claims metalv1alpha1.MachineClaimList
				g.Expect(ObjectList(&claims)()).To(SatisfyAll())
				if len(claims.Items) > 0 {
					g.Expect(k8sClient.DeleteAllOf(ctx, &claims.Items[0], &client.DeleteAllOfOptions{
						ListOptions: client.ListOptions{
							Namespace: ns.Name,
						},
					})).To(Succeed())
				}

				g.Expect(ObjectList(&machines)()).To(HaveField("Items", BeEmpty()))
				g.Expect(ObjectList(&claims)()).To(HaveField("Items", BeEmpty()))
			}, time.Second*3).Should(Succeed())
		})
	})

	It("should claim a Machine by ref", func(ctx SpecContext) {
		By("Creating a Machine")
		machine := &metalv1alpha1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
			},
			Spec: metalv1alpha1.MachineSpec{
				UUID: uuid.NewString(),
				OOBRef: v1.LocalObjectReference{
					Name: "doesnotexist",
				},
			},
		}
		Expect(k8sClient.Create(ctx, machine)).To(Succeed())

		By("Patching Machine state to Ready")
		Eventually(UpdateStatus(machine, func() {
			machine.Status.State = metalv1alpha1.MachineStateReady
		})).Should(Succeed())

		By("Creating a MachineClaim referencing the Machine")
		claim := &metalv1alpha1.MachineClaim{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
				Namespace:    ns.Name,
			},
			Spec: metalv1alpha1.MachineClaimSpec{
				MachineRef: &v1.LocalObjectReference{
					Name: machine.Name,
				},
				Image: "test",
				Power: metalv1alpha1.PowerOn,
			},
		}
		Expect(k8sClient.Create(ctx, claim)).To(Succeed())

		By("Expecting finalizer and phase to be correct on the MachineClaim")
		Eventually(Object(claim)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Status.Phase", metalv1alpha1.MachineClaimPhaseBound),
		))

		By("Expecting finalizer and machineclaimref to be correct on the Machine")
		Eventually(Object(machine)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Spec.MachineClaimRef.Namespace", claim.Namespace),
			HaveField("Spec.MachineClaimRef.Name", claim.Name),
			HaveField("Spec.MachineClaimRef.UID", claim.UID),
		))

		By("Deleting the MachineClaim")
		Expect(k8sClient.Delete(ctx, claim)).To(Succeed())

		By("Expecting machineclaimref and finalizer to be removed from the Machine")
		Eventually(Object(machine)).Should(SatisfyAll(
			HaveField("Finalizers", Not(ContainElement(MachineClaimFinalizer))),
			HaveField("Spec.MachineClaimRef", BeNil()),
		))

		By("Expecting MachineClaim to be removed")
		Eventually(Get(claim)).Should(Satisfy(errors.IsNotFound))
	})

	It("should claim a Machine by selector", func(ctx SpecContext) {
		By("Creating a Machine")
		machine := &metalv1alpha1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
				Labels: map[string]string{
					"test": "test",
				},
			},
			Spec: metalv1alpha1.MachineSpec{
				UUID: uuid.NewString(),
				OOBRef: v1.LocalObjectReference{
					Name: "doesnotexist",
				},
			},
		}
		Expect(k8sClient.Create(ctx, machine)).To(Succeed())

		By("Patching Machine state to Ready")
		Eventually(UpdateStatus(machine, func() {
			machine.Status.State = metalv1alpha1.MachineStateReady
		})).Should(Succeed())

		By("Creating a MachineClaim with a matching selector")
		claim := &metalv1alpha1.MachineClaim{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
				Namespace:    ns.Name,
			},
			Spec: metalv1alpha1.MachineClaimSpec{
				MachineSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"test": "test",
					},
				},
				Image: "test",
				Power: metalv1alpha1.PowerOn,
			},
		}
		Expect(k8sClient.Create(ctx, claim)).To(Succeed())

		By("Expecting finalizer, machineref, and phase to be correct on the MachineClaim")
		Eventually(Object(claim)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Spec.MachineRef.Name", machine.Name),
			HaveField("Status.Phase", metalv1alpha1.MachineClaimPhaseBound),
		))

		By("Expecting finalizer and machineclaimref to be correct on the Machine")
		Eventually(Object(machine)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Spec.MachineClaimRef.Namespace", claim.Namespace),
			HaveField("Spec.MachineClaimRef.Name", claim.Name),
			HaveField("Spec.MachineClaimRef.UID", claim.UID),
		))

		By("Deleting the MachineClaim")
		Expect(k8sClient.Delete(ctx, claim)).To(Succeed())

		By("Expecting machineclaimref and finalizer to be removed from the Machine")
		Eventually(Object(machine)).Should(SatisfyAll(
			HaveField("Finalizers", Not(ContainElement(MachineClaimFinalizer))),
			HaveField("Spec.MachineClaimRef", BeNil()),
		))

		By("Expecting MachineClaim to be removed")
		Eventually(Get(claim)).Should(Satisfy(errors.IsNotFound))
	})

	It("should not claim a Machine with a wrong ref", func(ctx SpecContext) {
		By("Creating a MachineClaim referencing the Machine")
		claim := &metalv1alpha1.MachineClaim{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
				Namespace:    ns.Name,
			},
			Spec: metalv1alpha1.MachineClaimSpec{
				MachineRef: &v1.LocalObjectReference{
					Name: "doesnotexist",
				},
				Image: "test",
				Power: metalv1alpha1.PowerOn,
			},
		}
		Expect(k8sClient.Create(ctx, claim)).To(Succeed())

		By("Expecting finalizer and phase to be correct on the MachineClaim")
		Eventually(Object(claim)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Status.Phase", metalv1alpha1.MachineClaimPhaseUnbound),
		))
	})

	It("should not claim a Machine with no matching selector", func(ctx SpecContext) {
		By("Creating a Machine")
		machine := &metalv1alpha1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
				Labels: map[string]string{
					"test": "test",
				},
			},
			Spec: metalv1alpha1.MachineSpec{
				UUID: uuid.NewString(),
				OOBRef: v1.LocalObjectReference{
					Name: "doesnotexist",
				},
			},
		}
		Expect(k8sClient.Create(ctx, machine)).To(Succeed())

		By("Patching Machine state to Ready")
		Eventually(UpdateStatus(machine, func() {
			machine.Status.State = metalv1alpha1.MachineStateReady
		})).Should(Succeed())

		By("Creating a MachineClaim referencing the Machine")
		claim := &metalv1alpha1.MachineClaim{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
				Namespace:    ns.Name,
			},
			Spec: metalv1alpha1.MachineClaimSpec{
				MachineSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"doesnotexist": "doesnotexist",
					},
				},
				Image: "test",
				Power: metalv1alpha1.PowerOn,
			},
		}
		Expect(k8sClient.Create(ctx, claim)).To(Succeed())

		By("Expecting finalizer and phase to be correct on the MachineClaim")
		Eventually(Object(claim)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Status.Phase", metalv1alpha1.MachineClaimPhaseUnbound),
		))

		By("Expecting no finalizer or claimref on the Machine")
		Eventually(Object(machine)).Should(SatisfyAll(
			HaveField("Finalizers", Not(ContainElement(MachineClaimFinalizer))),
			HaveField("Spec.MachineClaimRef", BeNil()),
		))

		By("Deleting the MachineClaim")
		Expect(k8sClient.Delete(ctx, claim)).To(Succeed())

		By("Expecting MachineClaim to be removed")
		Eventually(Get(claim)).Should(Satisfy(errors.IsNotFound))
	})

	It("should claim a Machine by ref once the Machine becomes Ready", func(ctx SpecContext) {
		By("Creating a Machine")
		machine := &metalv1alpha1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
			},
			Spec: metalv1alpha1.MachineSpec{
				UUID: uuid.NewString(),
				OOBRef: v1.LocalObjectReference{
					Name: "doesnotexist",
				},
			},
		}
		Expect(k8sClient.Create(ctx, machine)).To(Succeed())

		By("Patching Machine state to Error")
		Eventually(UpdateStatus(machine, func() {
			machine.Status.State = metalv1alpha1.MachineStateError
		})).Should(Succeed())

		By("Creating a MachineClaim referencing the Machine")
		claim := &metalv1alpha1.MachineClaim{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-",
				Namespace:    ns.Name,
			},
			Spec: metalv1alpha1.MachineClaimSpec{
				MachineRef: &v1.LocalObjectReference{
					Name: machine.Name,
				},
				Image: "test",
				Power: metalv1alpha1.PowerOn,
			},
		}
		Expect(k8sClient.Create(ctx, claim)).To(Succeed())

		By("Expecting finalizer and phase to be correct on the MachineClaim")
		Eventually(Object(claim)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Status.Phase", metalv1alpha1.MachineClaimPhaseUnbound),
		))

		By("Expecting no finalizer or claimref on the Machine")
		Eventually(Object(machine)).Should(SatisfyAll(
			HaveField("Finalizers", Not(ContainElement(MachineClaimFinalizer))),
			HaveField("Spec.MachineClaimRef", BeNil()),
		))

		By("Patching Machine state to Ready")
		Eventually(UpdateStatus(machine, func() {
			machine.Status.State = metalv1alpha1.MachineStateReady
		})).Should(Succeed())

		By("Expecting finalizer and phase to be correct on the MachineClaim")
		Eventually(Object(claim)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Status.Phase", metalv1alpha1.MachineClaimPhaseBound),
		))

		By("Expecting finalizer and machineclaimref to be correct on the Machine")
		Eventually(Object(machine)).Should(SatisfyAll(
			HaveField("Finalizers", ContainElement(MachineClaimFinalizer)),
			HaveField("Spec.MachineClaimRef.Namespace", claim.Namespace),
			HaveField("Spec.MachineClaimRef.Name", claim.Name),
			HaveField("Spec.MachineClaimRef.UID", claim.UID),
		))
	})
})
