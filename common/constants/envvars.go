package constants

/*
 Copyright (c) 2019 Dell Inc, or its subsidiaries.

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

const (
	// CustomPluginName is the name of the provider. It is used for multi array support
	CustomPluginName = "X_CSI_ISILON_PLUGIN_NAME"

	// EnvCSIEndpoint is the name of the unix domain socket that the csi driver is listening on
	EnvCSIEndpoint = "CSI_ENDPOINT"

	// EnvEndpoint is the name of the enviroment variable used to set the
	// HTTPS endpoint of the Isilon OneFS API server
	EnvEndpoint = "X_CSI_ISI_ENDPOINT"

	// EnvPort is the name of the enviroment variable used to set the
	// HTTPS port number of the Isilon OneFS API server
	EnvPort = "X_CSI_ISI_PORT"

	// EnvUser is the name of the enviroment variable used to set the
	// username when authenticating to the Isilon OneFS API server
	EnvUser = "X_CSI_ISI_USER"

	// EnvPassword is the name of the enviroment variable used to set the
	// user's password when authenticating to the Isilon OneFS API server
	EnvPassword = "X_CSI_ISI_PASSWORD"

	// EnvInsecure is the name of the enviroment variable used to specify
	// that the Isilon OneFS API server's certificate chain and host name should not
	// be verified
	EnvInsecure = "X_CSI_ISI_INSECURE"

	// EnvPath is the root path under which all the volumes (directories) will be provisioned, e.g. /ifs/engineering
	EnvPath = "X_CSI_ISI_PATH"

	// EnvAutoProbe is the name of the environment variable used to specify
	// that the controller service should automatically probe itself if it
	// receives incoming requests before having been probed, in direct
	// violation of the CSI spec
	EnvAutoProbe = "X_CSI_ISI_AUTOPROBE"

	// EnvDebug indicates whether the driver is in debug mode
	EnvDebug = "X_CSI_DEBUG"

	// EnvVerbose indicates whether the driver should log OneFS REST API response body content
	EnvVerbose = "X_CSI_VERBOSE"

	// EnvQuotaEnabled is the boolean flag that indicates whether the provisioner should attempt to set (later unset) quota on a newly provisioned volume
	EnvQuotaEnabled = "X_CSI_ISI_QUOTA_ENABLED"

	// EnvAccessZone is the name of the access zone a volume can be created in, e.g. "System"
	EnvAccessZone = "X_CSI_ISI_ACCESS_ZONE"

	// EnvNoProbeOnStart indicates whether a probe should be attempted upon start
	EnvNoProbeOnStart = "X_CSI_ISILON_NO_PROBE_ON_START"

	// EnvNfsV3 indicates whether to add "-o ver=3" option to the mount command when mounting an NFS export
	EnvNfsV3 = "X_CSI_ISILON_NFS_V3"

	// EnvNodeName is the name of a k8s node
	EnvNodeName = "X_CSI_NODE_NAME"

	// EnvNodeIP is the ip address of a k8s node
	EnvNodeIP = "X_CSI_NODE_IP"
)
