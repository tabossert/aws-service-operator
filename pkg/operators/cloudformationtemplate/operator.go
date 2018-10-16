// >>>>>>> DO NOT EDIT THIS FILE <<<<<<<<<<
// This file is autogenerated via `aws-operator generate`
// If you'd like the change anything about this file make edits to the .templ
// file in the pkg/codegen/assets directory.

package cloudformationtemplate

import (
	"github.com/awslabs/aws-service-operator/pkg/config"
	"github.com/awslabs/aws-service-operator/pkg/operator"
	"k8s.io/client-go/tools/cache"

	awsV1alpha1 "github.com/awslabs/aws-service-operator/pkg/apis/service-operator.aws/v1alpha1"
	"github.com/awslabs/aws-service-operator/pkg/customizations/cloudformationtemplate"
)

// Operator represents a controller object for object store custom resources
type Operator struct {
	config   *config.Config
	topicARN string
}

// NewOperator create controller for watching object store custom resources created
func NewOperator(config *config.Config) *Operator {
	return &Operator{
		config: config,
	}
}

// StartWatch watches for instances of Object Store custom resources and acts on them
func (c *Operator) StartWatch(namespace string, stopCh chan struct{}) error {
	resourceHandlers := cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	}

	oper := operator.New("cloudformationtemplates", namespace, resourceHandlers, c.config.AWSClientset.RESTClient())
	go oper.Watch(&awsV1alpha1.CloudFormationTemplate{}, stopCh)

	return nil
}

func (c *Operator) onAdd(obj interface{}) {
	s := obj.(*awsV1alpha1.CloudFormationTemplate).DeepCopy()
	cloudformationtemplate.OnAdd(c.config, s)
}

func (c *Operator) onUpdate(oldObj, newObj interface{}) {
	oo := oldObj.(*awsV1alpha1.CloudFormationTemplate).DeepCopy()
	no := newObj.(*awsV1alpha1.CloudFormationTemplate).DeepCopy()

	if no.Status.ResourceStatus == "DELETE_COMPLETE" {
		c.onAdd(no)
	}
	cloudformationtemplate.OnUpdate(c.config, oo, no)
}

func (c *Operator) onDelete(obj interface{}) {
	s := obj.(*awsV1alpha1.CloudFormationTemplate).DeepCopy()
	cloudformationtemplate.OnDelete(c.config, s)
}
