// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

import * as pulumiKubernetes from "@pulumi/kubernetes";

export interface CoreDNSAutoscalerArgs {
    /**
     * Number of cores in the cluster per coredns replica.
     */
    coresPerReplica?: pulumi.Input<number>;
    /**
     * Enabled the cluster-proportional-autoscaler.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The image to pull from for the autoscaler.
     */
    image?: pulumi.Input<inputs.CoreDNSImageArgs>;
    /**
     * Whether to include unschedulable nodes in the nodes/cores calculations - this requires version 1.8.0+ of the autoscaler.
     */
    includeUnschedulableNodes?: pulumi.Input<boolean>;
    /**
     * Max size of replicaCount
     */
    max?: pulumi.Input<number>;
    /**
     * Min size of replicaCount
     */
    min?: pulumi.Input<number>;
    /**
     * Number of nodes in the cluster per coredns replica.
     */
    nodesPerReplica?: pulumi.Input<number>;
    /**
     * If true does not allow single points of failure to form.
     */
    preventSinglePointFailure?: pulumi.Input<boolean>;
}

export interface CoreDNSDeploymentArgs {
    /**
     * Optionally disable the main deployment and its respective resources.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Name of the deployment if deployment.enabled is true. Otherwise the name of an existing deployment for the autoscaler or HPA to target.
     */
    name?: pulumi.Input<string>;
}

export interface CoreDNSHPAArgs {
    enabled?: pulumi.Input<boolean>;
    maxReplicas?: pulumi.Input<number>;
    metrics?: pulumi.Input<pulumiKubernetes.types.input.autoscaling.v2beta2.MetricSpec>;
    minReplicas?: pulumi.Input<number>;
}

export interface CoreDNSImageArgs {
    /**
     * Image pull policy.
     */
    pullPolicy?: pulumi.Input<string>;
    /**
     * Specify container image pull secrets.
     */
    pullSecrets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The image repository to pull from.
     */
    repository?: pulumi.Input<string>;
    /**
     * The image tag to pull from.
     */
    tag?: pulumi.Input<string>;
}

export interface CoreDNSPrometheusArgs {
    monitor?: pulumi.Input<inputs.CoreDNSPrometheusMonitorArgs>;
    service?: pulumi.Input<inputs.CoreDNSPrometheusServiceArgs>;
}

export interface CoreDNSPrometheusMonitorArgs {
    /**
     * Additional labels that can be used so ServiceMonitor will be discovered by Prometheus.
     */
    additionalLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set this to true to create ServiceMonitor for Prometheus operator.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Selector to select which namespaces the Endpoints objects are discovered from.
     */
    namespace?: pulumi.Input<string>;
}

export interface CoreDNSPrometheusServiceArgs {
    /**
     * Annotations to add to the metrics Service.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set this to true to create Service for Prometheus metrics.
     */
    enabled?: pulumi.Input<boolean>;
}

export interface CoreDNSRBACArgs {
    /**
     * If true, create & use RBAC resources
     */
    create?: pulumi.Input<boolean>;
    /**
     * The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template.
     */
    name?: pulumi.Input<string>;
    /**
     * If true, create and use PodSecurityPolicy
     */
    pspEnable?: pulumi.Input<boolean>;
}

export interface CoreDNSServerArgs {
    /**
     * the plugins to use for this server block.
     */
    plugins?: pulumi.Input<pulumi.Input<inputs.CoreDNSServerPluginArgs>[]>;
    /**
     * optional, defaults to "" (which equals 53 in CoreDNS).
     */
    port?: pulumi.Input<number>;
    /**
     * the `zones` block can be left out entirely, defaults to "."
     */
    zones?: pulumi.Input<pulumi.Input<inputs.CoreDNSServerZoneArgs>[]>;
}

export interface CoreDNSServerPluginArgs {
    /**
     * if the plugin supports extra block style config, supply it here
     */
    configBlock?: pulumi.Input<string>;
    /**
     * name of plugin, if used multiple times ensure that the plugin supports it!
     */
    name?: pulumi.Input<string>;
    /**
     * list of parameters after the plugin
     */
    parameters?: pulumi.Input<string>;
}

export interface CoreDNSServerZoneArgs {
    /**
     * optional, defaults to "" (which equals "dns://" in CoreDNS)
     */
    scheme?: pulumi.Input<string>;
    /**
     * set this parameter to optionally expose the port on tcp as well as udp for the DNS protocol. Note that this will not work if you are also exposing tls or grpc on the same server.
     */
    use_tcp?: pulumi.Input<boolean>;
    /**
     * optional, defaults to "."
     */
    zone?: pulumi.Input<string>;
}

export interface CoreDNSServiceArgs {
    /**
     * Annotations to add to service.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * IP address to assign to service.
     */
    clusterIP?: pulumi.Input<string>;
    /**
     * External IP addresses.
     */
    externalIPs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable client source IP preservation.
     */
    externalTrafficPolicy?: pulumi.Input<string>;
    /**
     * IP address to assign to load balancer (if supported).
     */
    loadBalancerIP?: pulumi.Input<string>;
    /**
     * The name of the Service. If not set, a name is generated using the fullname template.
     */
    name?: pulumi.Input<string>;
}

export interface CoreDNSServiceAccountArgs {
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If true, create & use serviceAccount.
     */
    create?: pulumi.Input<boolean>;
    /**
     * The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template
     */
    name?: pulumi.Input<string>;
}

export interface CoreDNSZoneFileArgs {
    contents?: pulumi.Input<string>;
    domain?: pulumi.Input<string>;
    string?: pulumi.Input<string>;
}

/**
 * A Release is an instance of a chart running in a Kubernetes cluster.
 * A Chart is a Helm package. It contains all of the resource definitions necessary to run an application, tool, or service inside of a Kubernetes cluster.
 * Note - Helm Release is currently in BETA and may change. Use in production environment is discouraged.
 */
export interface ReleaseArgs {
    /**
     * If set, installation process purges chart on fail. `skipAwait` will be disabled automatically if atomic is used.
     */
    atomic?: pulumi.Input<boolean>;
    /**
     * Chart name to be installed. A path may be used.
     */
    chart?: pulumi.Input<string>;
    /**
     * Allow deletion of new resources created in this upgrade when upgrade fails.
     */
    cleanupOnFail?: pulumi.Input<boolean>;
    /**
     * Create the namespace if it does not exist.
     */
    createNamespace?: pulumi.Input<boolean>;
    /**
     * Run helm dependency update before installing the chart.
     */
    dependencyUpdate?: pulumi.Input<boolean>;
    /**
     * Add a custom description
     */
    description?: pulumi.Input<string>;
    /**
     * Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
     */
    devel?: pulumi.Input<boolean>;
    /**
     * Prevent CRD hooks from, running, but run other hooks.  See helm install --no-crd-hook
     */
    disableCRDHooks?: pulumi.Input<boolean>;
    /**
     * If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
     */
    disableOpenapiValidation?: pulumi.Input<boolean>;
    /**
     * Prevent hooks from running.
     */
    disableWebhooks?: pulumi.Input<boolean>;
    /**
     * Force resource update through delete/recreate if needed.
     */
    forceUpdate?: pulumi.Input<boolean>;
    /**
     * Location of public keys used for verification. Used only if `verify` is true
     */
    keyring?: pulumi.Input<string>;
    /**
     * Run helm lint when planning.
     */
    lint?: pulumi.Input<boolean>;
    /**
     * The rendered manifests as JSON. Not yet supported.
     */
    manifest?: pulumi.Input<{[key: string]: any}>;
    /**
     * Limit the maximum number of revisions saved per release. Use 0 for no limit.
     */
    maxHistory?: pulumi.Input<number>;
    /**
     * Release name.
     */
    name?: pulumi.Input<string>;
    /**
     * Namespace to install the release into.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Postrender command to run.
     */
    postrender?: pulumi.Input<string>;
    /**
     * Perform pods restart during upgrade/rollback.
     */
    recreatePods?: pulumi.Input<boolean>;
    /**
     * If set, render subchart notes along with the parent.
     */
    renderSubchartNotes?: pulumi.Input<boolean>;
    /**
     * Re-use the given name, even if that name is already used. This is unsafe in production
     */
    replace?: pulumi.Input<boolean>;
    /**
     * Specification defining the Helm chart repository to use.
     */
    repositoryOpts?: pulumi.Input<inputs.RepositoryOptsArgs>;
    /**
     * When upgrading, reset the values to the ones built into the chart.
     */
    resetValues?: pulumi.Input<boolean>;
    /**
     * Names of resources created by the release grouped by "kind/version".
     */
    resourceNames?: pulumi.Input<{[key: string]: pulumi.Input<pulumi.Input<string>[]>}>;
    /**
     * When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
     */
    reuseValues?: pulumi.Input<boolean>;
    /**
     * By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
     */
    skipAwait?: pulumi.Input<boolean>;
    /**
     * If set, no CRDs will be installed. By default, CRDs are installed if not already present.
     */
    skipCrds?: pulumi.Input<boolean>;
    /**
     * Time in seconds to wait for any individual kubernetes operation.
     */
    timeout?: pulumi.Input<number>;
    /**
     * List of assets (raw yaml files). Content is read and merged with values. Not yet supported.
     */
    valueYamlFiles?: pulumi.Input<pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>[]>;
    /**
     * Custom values set for the release.
     */
    values?: pulumi.Input<{[key: string]: any}>;
    /**
     * Verify the package before installing it.
     */
    verify?: pulumi.Input<boolean>;
    /**
     * Specify the exact chart version to install. If this is not specified, the latest version is installed.
     */
    version?: pulumi.Input<string>;
    /**
     * Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipAwait` is enabled.
     */
    waitForJobs?: pulumi.Input<boolean>;
}

/**
 * Specification defining the Helm chart repository to use.
 */
export interface RepositoryOptsArgs {
    /**
     * The Repository's CA File
     */
    caFile?: pulumi.Input<string>;
    /**
     * The repository's cert file
     */
    certFile?: pulumi.Input<string>;
    /**
     * The repository's cert key file
     */
    keyFile?: pulumi.Input<string>;
    /**
     * Password for HTTP basic authentication
     */
    password?: pulumi.Input<string>;
    /**
     * Repository where to locate the requested chart. If is a URL the chart is installed without installing the repository.
     */
    repo?: pulumi.Input<string>;
    /**
     * Username for HTTP basic authentication
     */
    username?: pulumi.Input<string>;
}
