// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/kubeflow/training-operator/pkg/apis/kubeflow.org/v1"
	v2alpha1 "github.com/kubeflow/training-operator/pkg/apis/kubeflow.org/v2alpha1"
	internal "github.com/kubeflow/training-operator/pkg/client/applyconfiguration/internal"
	kubefloworgv1 "github.com/kubeflow/training-operator/pkg/client/applyconfiguration/kubeflow.org/v1"
	kubefloworgv2alpha1 "github.com/kubeflow/training-operator/pkg/client/applyconfiguration/kubeflow.org/v2alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=kubeflow.org, Version=v1
	case v1.SchemeGroupVersion.WithKind("ElasticPolicy"):
		return &kubefloworgv1.ElasticPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JAXJob"):
		return &kubefloworgv1.JAXJobApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JAXJobSpec"):
		return &kubefloworgv1.JAXJobSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JobCondition"):
		return &kubefloworgv1.JobConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JobStatus"):
		return &kubefloworgv1.JobStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MPIJob"):
		return &kubefloworgv1.MPIJobApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MPIJobSpec"):
		return &kubefloworgv1.MPIJobSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PaddleElasticPolicy"):
		return &kubefloworgv1.PaddleElasticPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PaddleJob"):
		return &kubefloworgv1.PaddleJobApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PaddleJobSpec"):
		return &kubefloworgv1.PaddleJobSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PyTorchJob"):
		return &kubefloworgv1.PyTorchJobApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PyTorchJobSpec"):
		return &kubefloworgv1.PyTorchJobSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RDZVConf"):
		return &kubefloworgv1.RDZVConfApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ReplicaSpec"):
		return &kubefloworgv1.ReplicaSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ReplicaStatus"):
		return &kubefloworgv1.ReplicaStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RunPolicy"):
		return &kubefloworgv1.RunPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SchedulingPolicy"):
		return &kubefloworgv1.SchedulingPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TFJob"):
		return &kubefloworgv1.TFJobApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TFJobSpec"):
		return &kubefloworgv1.TFJobSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("XGBoostJob"):
		return &kubefloworgv1.XGBoostJobApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("XGBoostJobSpec"):
		return &kubefloworgv1.XGBoostJobSpecApplyConfiguration{}

		// Group=kubeflow.org, Version=v2alpha1
	case v2alpha1.SchemeGroupVersion.WithKind("ClusterTrainingRuntime"):
		return &kubefloworgv2alpha1.ClusterTrainingRuntimeApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("ContainerOverride"):
		return &kubefloworgv2alpha1.ContainerOverrideApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("CoschedulingPodGroupPolicySource"):
		return &kubefloworgv2alpha1.CoschedulingPodGroupPolicySourceApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("DatasetConfig"):
		return &kubefloworgv2alpha1.DatasetConfigApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("InputModel"):
		return &kubefloworgv2alpha1.InputModelApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("JobSetTemplateSpec"):
		return &kubefloworgv2alpha1.JobSetTemplateSpecApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("JobStatus"):
		return &kubefloworgv2alpha1.JobStatusApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("MLPolicy"):
		return &kubefloworgv2alpha1.MLPolicyApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("MLPolicySource"):
		return &kubefloworgv2alpha1.MLPolicySourceApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("ModelConfig"):
		return &kubefloworgv2alpha1.ModelConfigApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("MPIMLPolicySource"):
		return &kubefloworgv2alpha1.MPIMLPolicySourceApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("OutputModel"):
		return &kubefloworgv2alpha1.OutputModelApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("PodGroupPolicy"):
		return &kubefloworgv2alpha1.PodGroupPolicyApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("PodGroupPolicySource"):
		return &kubefloworgv2alpha1.PodGroupPolicySourceApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("PodSpecOverride"):
		return &kubefloworgv2alpha1.PodSpecOverrideApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("PodSpecOverrideTargetJob"):
		return &kubefloworgv2alpha1.PodSpecOverrideTargetJobApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("RuntimeRef"):
		return &kubefloworgv2alpha1.RuntimeRefApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("TorchElasticPolicy"):
		return &kubefloworgv2alpha1.TorchElasticPolicyApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("TorchMLPolicySource"):
		return &kubefloworgv2alpha1.TorchMLPolicySourceApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("Trainer"):
		return &kubefloworgv2alpha1.TrainerApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("TrainingRuntime"):
		return &kubefloworgv2alpha1.TrainingRuntimeApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("TrainingRuntimeSpec"):
		return &kubefloworgv2alpha1.TrainingRuntimeSpecApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("TrainJob"):
		return &kubefloworgv2alpha1.TrainJobApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("TrainJobSpec"):
		return &kubefloworgv2alpha1.TrainJobSpecApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("TrainJobStatus"):
		return &kubefloworgv2alpha1.TrainJobStatusApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}