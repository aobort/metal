//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockSpec) DeepCopyInto(out *BlockSpec) {
	*out = *in
	if in.PartitionTable != nil {
		in, out := &in.PartitionTable, &out.PartitionTable
		*out = new(PartitionTableSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockSpec.
func (in *BlockSpec) DeepCopy() *BlockSpec {
	if in == nil {
		return nil
	}
	out := new(BlockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUSpec) DeepCopyInto(out *CPUSpec) {
	*out = *in
	if in.LogicalIDs != nil {
		in, out := &in.LogicalIDs, &out.LogicalIDs
		*out = make([]uint64, len(*in))
		copy(*out, *in)
	}
	out.MHz = in.MHz.DeepCopy()
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VMXFlags != nil {
		in, out := &in.VMXFlags, &out.VMXFlags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Bugs != nil {
		in, out := &in.Bugs, &out.Bugs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.BogoMIPS = in.BogoMIPS.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUSpec.
func (in *CPUSpec) DeepCopy() *CPUSpec {
	if in == nil {
		return nil
	}
	out := new(CPUSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleProtocol) DeepCopyInto(out *ConsoleProtocol) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleProtocol.
func (in *ConsoleProtocol) DeepCopy() *ConsoleProtocol {
	if in == nil {
		return nil
	}
	out := new(ConsoleProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistroSpec) DeepCopyInto(out *DistroSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistroSpec.
func (in *DistroSpec) DeepCopy() *DistroSpec {
	if in == nil {
		return nil
	}
	out := new(DistroSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostSpec) DeepCopyInto(out *HostSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostSpec.
func (in *HostSpec) DeepCopy() *HostSpec {
	if in == nil {
		return nil
	}
	out := new(HostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPMISpec) DeepCopyInto(out *IPMISpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPMISpec.
func (in *IPMISpec) DeepCopy() *IPMISpec {
	if in == nil {
		return nil
	}
	out := new(IPMISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inventory) DeepCopyInto(out *Inventory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inventory.
func (in *Inventory) DeepCopy() *Inventory {
	if in == nil {
		return nil
	}
	out := new(Inventory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Inventory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InventoryList) DeepCopyInto(out *InventoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Inventory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InventoryList.
func (in *InventoryList) DeepCopy() *InventoryList {
	if in == nil {
		return nil
	}
	out := new(InventoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InventoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InventorySpec) DeepCopyInto(out *InventorySpec) {
	*out = *in
	if in.System != nil {
		in, out := &in.System, &out.System
		*out = new(SystemSpec)
		**out = **in
	}
	if in.IPMIs != nil {
		in, out := &in.IPMIs, &out.IPMIs
		*out = make([]IPMISpec, len(*in))
		copy(*out, *in)
	}
	if in.Blocks != nil {
		in, out := &in.Blocks, &out.Blocks
		*out = make([]BlockSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(MemorySpec)
		**out = **in
	}
	if in.CPUs != nil {
		in, out := &in.CPUs, &out.CPUs
		*out = make([]CPUSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NUMA != nil {
		in, out := &in.NUMA, &out.NUMA
		*out = make([]NumaSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PCIDevices != nil {
		in, out := &in.PCIDevices, &out.PCIDevices
		*out = make([]PCIDeviceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NICs != nil {
		in, out := &in.NICs, &out.NICs
		*out = make([]NICSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Virt != nil {
		in, out := &in.Virt, &out.Virt
		*out = new(VirtSpec)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(HostSpec)
		**out = **in
	}
	if in.Distro != nil {
		in, out := &in.Distro, &out.Distro
		*out = new(DistroSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InventorySpec.
func (in *InventorySpec) DeepCopy() *InventorySpec {
	if in == nil {
		return nil
	}
	out := new(InventorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InventoryStatus) DeepCopyInto(out *InventoryStatus) {
	*out = *in
	out.InventoryStatuses = in.InventoryStatuses
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InventoryStatus.
func (in *InventoryStatus) DeepCopy() *InventoryStatus {
	if in == nil {
		return nil
	}
	out := new(InventoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InventoryStatuses) DeepCopyInto(out *InventoryStatuses) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InventoryStatuses.
func (in *InventoryStatuses) DeepCopy() *InventoryStatuses {
	if in == nil {
		return nil
	}
	out := new(InventoryStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLDPSpec) DeepCopyInto(out *LLDPSpec) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]LLDPCapabilities, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLDPSpec.
func (in *LLDPSpec) DeepCopy() *LLDPSpec {
	if in == nil {
		return nil
	}
	out := new(LLDPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Machine) DeepCopyInto(out *Machine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Machine.
func (in *Machine) DeepCopy() *Machine {
	if in == nil {
		return nil
	}
	out := new(Machine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Machine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClaim) DeepCopyInto(out *MachineClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClaim.
func (in *MachineClaim) DeepCopy() *MachineClaim {
	if in == nil {
		return nil
	}
	out := new(MachineClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClaimList) DeepCopyInto(out *MachineClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClaimList.
func (in *MachineClaimList) DeepCopy() *MachineClaimList {
	if in == nil {
		return nil
	}
	out := new(MachineClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClaimNetworkInterface) DeepCopyInto(out *MachineClaimNetworkInterface) {
	*out = *in
	in.Prefix.DeepCopyInto(&out.Prefix)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClaimNetworkInterface.
func (in *MachineClaimNetworkInterface) DeepCopy() *MachineClaimNetworkInterface {
	if in == nil {
		return nil
	}
	out := new(MachineClaimNetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClaimSpec) DeepCopyInto(out *MachineClaimSpec) {
	*out = *in
	if in.MachineRef != nil {
		in, out := &in.MachineRef, &out.MachineRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.MachineSelector != nil {
		in, out := &in.MachineSelector, &out.MachineSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.IgnitionSecretRef != nil {
		in, out := &in.IgnitionSecretRef, &out.IgnitionSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.NetworkInterfaces != nil {
		in, out := &in.NetworkInterfaces, &out.NetworkInterfaces
		*out = make([]MachineClaimNetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClaimSpec.
func (in *MachineClaimSpec) DeepCopy() *MachineClaimSpec {
	if in == nil {
		return nil
	}
	out := new(MachineClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineClaimStatus) DeepCopyInto(out *MachineClaimStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineClaimStatus.
func (in *MachineClaimStatus) DeepCopy() *MachineClaimStatus {
	if in == nil {
		return nil
	}
	out := new(MachineClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineList) DeepCopyInto(out *MachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Machine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineList.
func (in *MachineList) DeepCopy() *MachineList {
	if in == nil {
		return nil
	}
	out := new(MachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineNetworkInterface) DeepCopyInto(out *MachineNetworkInterface) {
	*out = *in
	if in.IPRef != nil {
		in, out := &in.IPRef, &out.IPRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.SwitchRef != nil {
		in, out := &in.SwitchRef, &out.SwitchRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineNetworkInterface.
func (in *MachineNetworkInterface) DeepCopy() *MachineNetworkInterface {
	if in == nil {
		return nil
	}
	out := new(MachineNetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpec) DeepCopyInto(out *MachineSpec) {
	*out = *in
	out.OOBRef = in.OOBRef
	if in.InventoryRef != nil {
		in, out := &in.InventoryRef, &out.InventoryRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.MachineClaimRef != nil {
		in, out := &in.MachineClaimRef, &out.MachineClaimRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.LoopbackAddressRef != nil {
		in, out := &in.LoopbackAddressRef, &out.LoopbackAddressRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpec.
func (in *MachineSpec) DeepCopy() *MachineSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineStatus) DeepCopyInto(out *MachineStatus) {
	*out = *in
	if in.ShutdownDeadline != nil {
		in, out := &in.ShutdownDeadline, &out.ShutdownDeadline
		*out = (*in).DeepCopy()
	}
	if in.NetworkInterfaces != nil {
		in, out := &in.NetworkInterfaces, &out.NetworkInterfaces
		*out = make([]MachineNetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineStatus.
func (in *MachineStatus) DeepCopy() *MachineStatus {
	if in == nil {
		return nil
	}
	out := new(MachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemorySpec) DeepCopyInto(out *MemorySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemorySpec.
func (in *MemorySpec) DeepCopy() *MemorySpec {
	if in == nil {
		return nil
	}
	out := new(MemorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NDPSpec) DeepCopyInto(out *NDPSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NDPSpec.
func (in *NDPSpec) DeepCopy() *NDPSpec {
	if in == nil {
		return nil
	}
	out := new(NDPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NICSpec) DeepCopyInto(out *NICSpec) {
	*out = *in
	if in.LLDPs != nil {
		in, out := &in.LLDPs, &out.LLDPs
		*out = make([]LLDPSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NDPs != nil {
		in, out := &in.NDPs, &out.NDPs
		*out = make([]NDPSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NICSpec.
func (in *NICSpec) DeepCopy() *NICSpec {
	if in == nil {
		return nil
	}
	out := new(NICSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaSpec) DeepCopyInto(out *NumaSpec) {
	*out = *in
	if in.CPUs != nil {
		in, out := &in.CPUs, &out.CPUs
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Distances != nil {
		in, out := &in.Distances, &out.Distances
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(MemorySpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaSpec.
func (in *NumaSpec) DeepCopy() *NumaSpec {
	if in == nil {
		return nil
	}
	out := new(NumaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOB) DeepCopyInto(out *OOB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOB.
func (in *OOB) DeepCopy() *OOB {
	if in == nil {
		return nil
	}
	out := new(OOB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OOB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOBList) DeepCopyInto(out *OOBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OOB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOBList.
func (in *OOBList) DeepCopy() *OOBList {
	if in == nil {
		return nil
	}
	out := new(OOBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OOBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOBSecret) DeepCopyInto(out *OOBSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOBSecret.
func (in *OOBSecret) DeepCopy() *OOBSecret {
	if in == nil {
		return nil
	}
	out := new(OOBSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OOBSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOBSecretList) DeepCopyInto(out *OOBSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OOBSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOBSecretList.
func (in *OOBSecretList) DeepCopy() *OOBSecretList {
	if in == nil {
		return nil
	}
	out := new(OOBSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OOBSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOBSecretSpec) DeepCopyInto(out *OOBSecretSpec) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOBSecretSpec.
func (in *OOBSecretSpec) DeepCopy() *OOBSecretSpec {
	if in == nil {
		return nil
	}
	out := new(OOBSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOBSecretStatus) DeepCopyInto(out *OOBSecretStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOBSecretStatus.
func (in *OOBSecretStatus) DeepCopy() *OOBSecretStatus {
	if in == nil {
		return nil
	}
	out := new(OOBSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOBSpec) DeepCopyInto(out *OOBSpec) {
	*out = *in
	if in.EndpointRef != nil {
		in, out := &in.EndpointRef, &out.EndpointRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(Protocol)
		**out = **in
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConsoleProtocol != nil {
		in, out := &in.ConsoleProtocol, &out.ConsoleProtocol
		*out = new(ConsoleProtocol)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOBSpec.
func (in *OOBSpec) DeepCopy() *OOBSpec {
	if in == nil {
		return nil
	}
	out := new(OOBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOBStatus) DeepCopyInto(out *OOBStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOBStatus.
func (in *OOBStatus) DeepCopy() *OOBStatus {
	if in == nil {
		return nil
	}
	out := new(OOBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PCIDeviceDescriptionSpec) DeepCopyInto(out *PCIDeviceDescriptionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PCIDeviceDescriptionSpec.
func (in *PCIDeviceDescriptionSpec) DeepCopy() *PCIDeviceDescriptionSpec {
	if in == nil {
		return nil
	}
	out := new(PCIDeviceDescriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PCIDeviceSpec) DeepCopyInto(out *PCIDeviceSpec) {
	*out = *in
	if in.Vendor != nil {
		in, out := &in.Vendor, &out.Vendor
		*out = new(PCIDeviceDescriptionSpec)
		**out = **in
	}
	if in.Subvendor != nil {
		in, out := &in.Subvendor, &out.Subvendor
		*out = new(PCIDeviceDescriptionSpec)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(PCIDeviceDescriptionSpec)
		**out = **in
	}
	if in.Subtype != nil {
		in, out := &in.Subtype, &out.Subtype
		*out = new(PCIDeviceDescriptionSpec)
		**out = **in
	}
	if in.Class != nil {
		in, out := &in.Class, &out.Class
		*out = new(PCIDeviceDescriptionSpec)
		**out = **in
	}
	if in.Subclass != nil {
		in, out := &in.Subclass, &out.Subclass
		*out = new(PCIDeviceDescriptionSpec)
		**out = **in
	}
	if in.ProgrammingInterface != nil {
		in, out := &in.ProgrammingInterface, &out.ProgrammingInterface
		*out = new(PCIDeviceDescriptionSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PCIDeviceSpec.
func (in *PCIDeviceSpec) DeepCopy() *PCIDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(PCIDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionSpec) DeepCopyInto(out *PartitionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionSpec.
func (in *PartitionSpec) DeepCopy() *PartitionSpec {
	if in == nil {
		return nil
	}
	out := new(PartitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionTableSpec) DeepCopyInto(out *PartitionTableSpec) {
	*out = *in
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = make([]PartitionSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionTableSpec.
func (in *PartitionTableSpec) DeepCopy() *PartitionTableSpec {
	if in == nil {
		return nil
	}
	out := new(PartitionTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Protocol) DeepCopyInto(out *Protocol) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Protocol.
func (in *Protocol) DeepCopy() *Protocol {
	if in == nil {
		return nil
	}
	out := new(Protocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemSpec) DeepCopyInto(out *SystemSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemSpec.
func (in *SystemSpec) DeepCopy() *SystemSpec {
	if in == nil {
		return nil
	}
	out := new(SystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtSpec) DeepCopyInto(out *VirtSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtSpec.
func (in *VirtSpec) DeepCopy() *VirtSpec {
	if in == nil {
		return nil
	}
	out := new(VirtSpec)
	in.DeepCopyInto(out)
	return out
}
