// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ChartCertManager.Inputs
{

    public sealed class CoreDNSServerZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// optional, defaults to "" (which equals "dns://" in CoreDNS)
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// set this parameter to optionally expose the port on tcp as well as udp for the DNS protocol. Note that this will not work if you are also exposing tls or grpc on the same server.
        /// </summary>
        [Input("use_tcp")]
        public Input<bool>? Use_tcp { get; set; }

        /// <summary>
        /// optional, defaults to "."
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public CoreDNSServerZoneArgs()
        {
        }
    }
}
