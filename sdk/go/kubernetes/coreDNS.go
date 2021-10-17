// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	autoscalingv2beta2 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/autoscaling/v2beta2"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	policyv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/policy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enable fast and flexible in-cluster DNS.
type CoreDNS struct {
	pulumi.ResourceState

	// Detailed information about the status of the underlying Helm deployment.
	Status ReleaseStatusOutput `pulumi:"status"`
}

// NewCoreDNS registers a new resource with the given unique name, arguments, and options.
func NewCoreDNS(ctx *pulumi.Context,
	name string, args *CoreDNSArgs, opts ...pulumi.ResourceOption) (*CoreDNS, error) {
	if args == nil {
		args = &CoreDNSArgs{}
	}

	var resource CoreDNS
	err := ctx.RegisterRemoteComponentResource("kubernetes-coredns:index:CoreDNS", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type coreDNSArgs struct {
	// Affinity settings for pod assignment	.
	Affinity *corev1.Affinity `pulumi:"affinity"`
	// Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
	Autoscaler *CoreDNSAutoscaler `pulumi:"autoscaler"`
	// Create HorizontalPodAutoscaler object.
	Autoscaling *autoscalingv2beta2.HorizontalPodAutoscalerSpec `pulumi:"autoscaling"`
	// Configure the CoreDNS Deployment.
	Deployment *CoreDNSDeployment `pulumi:"deployment"`
	// Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
	ExtraSecrets []corev1.VolumeMount `pulumi:"extraSecrets"`
	// Optional array of mount points for extraVolumes.
	ExtraVolumeMounts []corev1.VolumeMount `pulumi:"extraVolumeMounts"`
	// Optional array of extra volumes to create.
	ExtraVolumes []corev1.Volume `pulumi:"extraVolumes"`
	// HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions *Release `pulumi:"helmOptions"`
	// Alternative configuration for HPA deployment if wanted.
	Hpa *CoreDNSHPA `pulumi:"hpa"`
	// The image to pull.
	Image *CoreDNSImage `pulumi:"image"`
	// Specifies whether chart should be deployed as cluster-service or normal k8s app.
	IsClusterService *bool `pulumi:"isClusterService"`
	// Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS' server config.
	LivenessProbe *corev1.Probe `pulumi:"livenessProbe"`
	// Node labels for pod assignment.
	NodeSelector map[string]string `pulumi:"nodeSelector"`
	// Optional Pod only Annotations.
	PodAnnotations map[string]string `pulumi:"podAnnotations"`
	// Optional PodDisruptionBudget.
	PodDisruptionBudget *policyv1.PodDisruptionBudgetSpec `pulumi:"podDisruptionBudget"`
	// Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.
	PreStopSleep *int `pulumi:"preStopSleep"`
	// Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
	PriorityClassName *string `pulumi:"priorityClassName"`
	// Configure Prometheus installation.
	Prometheus *CoreDNSPrometheus `pulumi:"prometheus"`
	// Configure CoreDNS RBAC resources.
	Rbac *CoreDNSRBAC `pulumi:"rbac"`
	// Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS' server config.
	ReadinessProbe *corev1.Probe `pulumi:"readinessProbe"`
	// Number of replicas.
	ReplicaCount *int `pulumi:"replicaCount"`
	// Container resource limits.
	Resources     *corev1.ResourceRequirements    `pulumi:"resources"`
	RollingUpdate *appsv1.RollingUpdateDeployment `pulumi:"rollingUpdate"`
	// Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
	Servers []CoreDNSServer `pulumi:"servers"`
	// Configure CoreDNS Service parameters.
	Service *CoreDNSService `pulumi:"service"`
	// Configure CoreDNS Service Account.
	ServiceAccount *CoreDNSServiceAccount `pulumi:"serviceAccount"`
	// Kubernetes Service type.
	ServiceType *string `pulumi:"serviceType"`
	// Optional duration in seconds the pod needs to terminate gracefully.
	TerminationGracePeriodSeconds *int `pulumi:"terminationGracePeriodSeconds"`
	// Tolerations for pod assignment.
	Tolerations []corev1.Toleration `pulumi:"tolerations"`
	// Configure custom Zone files.
	ZoneFiles []CoreDNSZoneFile `pulumi:"zoneFiles"`
}

// The set of arguments for constructing a CoreDNS resource.
type CoreDNSArgs struct {
	// Affinity settings for pod assignment	.
	Affinity corev1.AffinityPtrInput
	// Configue a cluster-proportional-autoscaler for coredns. See https://github.com/kubernetes-incubator/cluster-proportional-autoscaler.
	Autoscaler CoreDNSAutoscalerPtrInput
	// Create HorizontalPodAutoscaler object.
	Autoscaling autoscalingv2beta2.HorizontalPodAutoscalerSpecPtrInput
	// Configure the CoreDNS Deployment.
	Deployment CoreDNSDeploymentPtrInput
	// Optional array of secrets to mount inside coredns container. Possible usecase: need for secure connection with etcd backend. Optional array of mount points for extraVolumes.
	ExtraSecrets corev1.VolumeMountArrayInput
	// Optional array of mount points for extraVolumes.
	ExtraVolumeMounts corev1.VolumeMountArrayInput
	// Optional array of extra volumes to create.
	ExtraVolumes corev1.VolumeArrayInput
	// HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions ReleasePtrInput
	// Alternative configuration for HPA deployment if wanted.
	Hpa CoreDNSHPAPtrInput
	// The image to pull.
	Image CoreDNSImagePtrInput
	// Specifies whether chart should be deployed as cluster-service or normal k8s app.
	IsClusterService pulumi.BoolPtrInput
	// Configure the liveness probe. To use the livenessProbe, the health plugin needs to be enabled in CoreDNS' server config.
	LivenessProbe corev1.ProbePtrInput
	// Node labels for pod assignment.
	NodeSelector pulumi.StringMapInput
	// Optional Pod only Annotations.
	PodAnnotations pulumi.StringMapInput
	// Optional PodDisruptionBudget.
	PodDisruptionBudget policyv1.PodDisruptionBudgetSpecPtrInput
	// Under heavy load it takes more that standard time to remove Pod endpoint from a cluster. This will delay termination of our pod by `preStopSleep`. To make sure kube-proxy has enough time to catch up.
	PreStopSleep pulumi.IntPtrInput
	// Optional priority class to be used for the coredns pods. Used for autoscaler if autoscaler.priorityClassName not set.
	PriorityClassName pulumi.StringPtrInput
	// Configure Prometheus installation.
	Prometheus CoreDNSPrometheusPtrInput
	// Configure CoreDNS RBAC resources.
	Rbac CoreDNSRBACPtrInput
	// Configure the readiness probe. To use the readinessProbe, the health plugin needs to be enabled in CoreDNS' server config.
	ReadinessProbe corev1.ProbePtrInput
	// Number of replicas.
	ReplicaCount pulumi.IntPtrInput
	// Container resource limits.
	Resources     corev1.ResourceRequirementsPtrInput
	RollingUpdate appsv1.RollingUpdateDeploymentPtrInput
	// Configuration for CoreDNS and plugins. Default zone is what Kubernetes recommends: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#coredns-configmap-options
	Servers CoreDNSServerArrayInput
	// Configure CoreDNS Service parameters.
	Service CoreDNSServicePtrInput
	// Configure CoreDNS Service Account.
	ServiceAccount CoreDNSServiceAccountPtrInput
	// Kubernetes Service type.
	ServiceType pulumi.StringPtrInput
	// Optional duration in seconds the pod needs to terminate gracefully.
	TerminationGracePeriodSeconds pulumi.IntPtrInput
	// Tolerations for pod assignment.
	Tolerations corev1.TolerationArrayInput
	// Configure custom Zone files.
	ZoneFiles CoreDNSZoneFileArrayInput
}

func (CoreDNSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*coreDNSArgs)(nil)).Elem()
}

type CoreDNSInput interface {
	pulumi.Input

	ToCoreDNSOutput() CoreDNSOutput
	ToCoreDNSOutputWithContext(ctx context.Context) CoreDNSOutput
}

func (*CoreDNS) ElementType() reflect.Type {
	return reflect.TypeOf((*CoreDNS)(nil))
}

func (i *CoreDNS) ToCoreDNSOutput() CoreDNSOutput {
	return i.ToCoreDNSOutputWithContext(context.Background())
}

func (i *CoreDNS) ToCoreDNSOutputWithContext(ctx context.Context) CoreDNSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDNSOutput)
}

func (i *CoreDNS) ToCoreDNSPtrOutput() CoreDNSPtrOutput {
	return i.ToCoreDNSPtrOutputWithContext(context.Background())
}

func (i *CoreDNS) ToCoreDNSPtrOutputWithContext(ctx context.Context) CoreDNSPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDNSPtrOutput)
}

type CoreDNSPtrInput interface {
	pulumi.Input

	ToCoreDNSPtrOutput() CoreDNSPtrOutput
	ToCoreDNSPtrOutputWithContext(ctx context.Context) CoreDNSPtrOutput
}

type coreDNSPtrType CoreDNSArgs

func (*coreDNSPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreDNS)(nil))
}

func (i *coreDNSPtrType) ToCoreDNSPtrOutput() CoreDNSPtrOutput {
	return i.ToCoreDNSPtrOutputWithContext(context.Background())
}

func (i *coreDNSPtrType) ToCoreDNSPtrOutputWithContext(ctx context.Context) CoreDNSPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDNSPtrOutput)
}

// CoreDNSArrayInput is an input type that accepts CoreDNSArray and CoreDNSArrayOutput values.
// You can construct a concrete instance of `CoreDNSArrayInput` via:
//
//          CoreDNSArray{ CoreDNSArgs{...} }
type CoreDNSArrayInput interface {
	pulumi.Input

	ToCoreDNSArrayOutput() CoreDNSArrayOutput
	ToCoreDNSArrayOutputWithContext(context.Context) CoreDNSArrayOutput
}

type CoreDNSArray []CoreDNSInput

func (CoreDNSArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CoreDNS)(nil)).Elem()
}

func (i CoreDNSArray) ToCoreDNSArrayOutput() CoreDNSArrayOutput {
	return i.ToCoreDNSArrayOutputWithContext(context.Background())
}

func (i CoreDNSArray) ToCoreDNSArrayOutputWithContext(ctx context.Context) CoreDNSArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDNSArrayOutput)
}

// CoreDNSMapInput is an input type that accepts CoreDNSMap and CoreDNSMapOutput values.
// You can construct a concrete instance of `CoreDNSMapInput` via:
//
//          CoreDNSMap{ "key": CoreDNSArgs{...} }
type CoreDNSMapInput interface {
	pulumi.Input

	ToCoreDNSMapOutput() CoreDNSMapOutput
	ToCoreDNSMapOutputWithContext(context.Context) CoreDNSMapOutput
}

type CoreDNSMap map[string]CoreDNSInput

func (CoreDNSMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CoreDNS)(nil)).Elem()
}

func (i CoreDNSMap) ToCoreDNSMapOutput() CoreDNSMapOutput {
	return i.ToCoreDNSMapOutputWithContext(context.Background())
}

func (i CoreDNSMap) ToCoreDNSMapOutputWithContext(ctx context.Context) CoreDNSMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDNSMapOutput)
}

type CoreDNSOutput struct{ *pulumi.OutputState }

func (CoreDNSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CoreDNS)(nil))
}

func (o CoreDNSOutput) ToCoreDNSOutput() CoreDNSOutput {
	return o
}

func (o CoreDNSOutput) ToCoreDNSOutputWithContext(ctx context.Context) CoreDNSOutput {
	return o
}

func (o CoreDNSOutput) ToCoreDNSPtrOutput() CoreDNSPtrOutput {
	return o.ToCoreDNSPtrOutputWithContext(context.Background())
}

func (o CoreDNSOutput) ToCoreDNSPtrOutputWithContext(ctx context.Context) CoreDNSPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CoreDNS) *CoreDNS {
		return &v
	}).(CoreDNSPtrOutput)
}

type CoreDNSPtrOutput struct{ *pulumi.OutputState }

func (CoreDNSPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreDNS)(nil))
}

func (o CoreDNSPtrOutput) ToCoreDNSPtrOutput() CoreDNSPtrOutput {
	return o
}

func (o CoreDNSPtrOutput) ToCoreDNSPtrOutputWithContext(ctx context.Context) CoreDNSPtrOutput {
	return o
}

func (o CoreDNSPtrOutput) Elem() CoreDNSOutput {
	return o.ApplyT(func(v *CoreDNS) CoreDNS {
		if v != nil {
			return *v
		}
		var ret CoreDNS
		return ret
	}).(CoreDNSOutput)
}

type CoreDNSArrayOutput struct{ *pulumi.OutputState }

func (CoreDNSArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CoreDNS)(nil))
}

func (o CoreDNSArrayOutput) ToCoreDNSArrayOutput() CoreDNSArrayOutput {
	return o
}

func (o CoreDNSArrayOutput) ToCoreDNSArrayOutputWithContext(ctx context.Context) CoreDNSArrayOutput {
	return o
}

func (o CoreDNSArrayOutput) Index(i pulumi.IntInput) CoreDNSOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CoreDNS {
		return vs[0].([]CoreDNS)[vs[1].(int)]
	}).(CoreDNSOutput)
}

type CoreDNSMapOutput struct{ *pulumi.OutputState }

func (CoreDNSMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CoreDNS)(nil))
}

func (o CoreDNSMapOutput) ToCoreDNSMapOutput() CoreDNSMapOutput {
	return o
}

func (o CoreDNSMapOutput) ToCoreDNSMapOutputWithContext(ctx context.Context) CoreDNSMapOutput {
	return o
}

func (o CoreDNSMapOutput) MapIndex(k pulumi.StringInput) CoreDNSOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CoreDNS {
		return vs[0].(map[string]CoreDNS)[vs[1].(string)]
	}).(CoreDNSOutput)
}

func init() {
	pulumi.RegisterOutputType(CoreDNSOutput{})
	pulumi.RegisterOutputType(CoreDNSPtrOutput{})
	pulumi.RegisterOutputType(CoreDNSArrayOutput{})
	pulumi.RegisterOutputType(CoreDNSMapOutput{})
}
