// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/superedge/superedge/pkg/application-grid-controller/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DeploymentGrids returns a DeploymentGridInformer.
	DeploymentGrids() DeploymentGridInformer
	// ServiceGrids returns a ServiceGridInformer.
	ServiceGrids() ServiceGridInformer
	// StatefulSetGrids returns a StatefulSetGridInformer.
	StatefulSetGrids() StatefulSetGridInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// DeploymentGrids returns a DeploymentGridInformer.
func (v *version) DeploymentGrids() DeploymentGridInformer {
	return &deploymentGridInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceGrids returns a ServiceGridInformer.
func (v *version) ServiceGrids() ServiceGridInformer {
	return &serviceGridInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StatefulSetGrids returns a StatefulSetGridInformer.
func (v *version) StatefulSetGrids() StatefulSetGridInformer {
	return &statefulSetGridInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
