// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCoreDNS.Inputs
{

    /// <summary>
    /// Specification defining the Helm chart repository to use.
    /// </summary>
    public sealed class RepositoryOptsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Repository's CA File
        /// </summary>
        [Input("caFile")]
        public Input<string>? CaFile { get; set; }

        /// <summary>
        /// The repository's cert file
        /// </summary>
        [Input("certFile")]
        public Input<string>? CertFile { get; set; }

        /// <summary>
        /// The repository's cert key file
        /// </summary>
        [Input("keyFile")]
        public Input<string>? KeyFile { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for HTTP basic authentication
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Repository where to locate the requested chart. If is a URL the chart is installed without installing the repository.
        /// </summary>
        [Input("repo")]
        public Input<string>? Repo { get; set; }

        /// <summary>
        /// Username for HTTP basic authentication
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public RepositoryOptsArgs()
        {
        }
    }
}
