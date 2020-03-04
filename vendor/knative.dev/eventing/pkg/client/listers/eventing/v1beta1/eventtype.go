/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "knative.dev/eventing/pkg/apis/eventing/v1beta1"
)

// EventTypeLister helps list EventTypes.
type EventTypeLister interface {
	// List lists all EventTypes in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.EventType, err error)
	// EventTypes returns an object that can list and get EventTypes.
	EventTypes(namespace string) EventTypeNamespaceLister
	EventTypeListerExpansion
}

// eventTypeLister implements the EventTypeLister interface.
type eventTypeLister struct {
	indexer cache.Indexer
}

// NewEventTypeLister returns a new EventTypeLister.
func NewEventTypeLister(indexer cache.Indexer) EventTypeLister {
	return &eventTypeLister{indexer: indexer}
}

// List lists all EventTypes in the indexer.
func (s *eventTypeLister) List(selector labels.Selector) (ret []*v1beta1.EventType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.EventType))
	})
	return ret, err
}

// EventTypes returns an object that can list and get EventTypes.
func (s *eventTypeLister) EventTypes(namespace string) EventTypeNamespaceLister {
	return eventTypeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventTypeNamespaceLister helps list and get EventTypes.
type EventTypeNamespaceLister interface {
	// List lists all EventTypes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.EventType, err error)
	// Get retrieves the EventType from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.EventType, error)
	EventTypeNamespaceListerExpansion
}

// eventTypeNamespaceLister implements the EventTypeNamespaceLister
// interface.
type eventTypeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventTypes in the indexer for a given namespace.
func (s eventTypeNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.EventType, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.EventType))
	})
	return ret, err
}

// Get retrieves the EventType from the indexer for a given namespace and name.
func (s eventTypeNamespaceLister) Get(name string) (*v1beta1.EventType, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("eventtype"), name)
	}
	return obj.(*v1beta1.EventType), nil
}
