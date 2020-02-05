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

// Code generated by injection-gen. DO NOT EDIT.

package samplesource

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	controller "knative.dev/pkg/controller"
	logging "knative.dev/pkg/logging"
	samplesScheme "knative.dev/sample-source/pkg/client/clientset/versioned/scheme"
	client "knative.dev/sample-source/pkg/client/injection/client"
	samplesource "knative.dev/sample-source/pkg/client/injection/informers/samples/v1alpha1/samplesource"
)

const (
	defaultControllerAgentName = "samplesource-controller"
	defaultFinalizerName       = "samplesource"
)

func NewImpl(ctx context.Context, r Interface) *controller.Impl {
	logger := logging.FromContext(ctx)

	samplesourceInformer := samplesource.Get(ctx)

	c := &reconcilerImpl{
		Client: client.Get(ctx),
		Lister: samplesourceInformer.Lister(),
		Recorder: record.NewBroadcaster().NewRecorder(
			scheme.Scheme, v1.EventSource{Component: defaultControllerAgentName}),
		FinalizerName: defaultFinalizerName,
		reconciler:    r,
	}
	impl := controller.NewImpl(c, logger, "samplesources")

	return impl
}

func init() {
	samplesScheme.AddToScheme(scheme.Scheme)
}
