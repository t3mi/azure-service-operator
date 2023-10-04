// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Service_Spec_ARM struct {
	// Identity: Managed service identity of the Api Management service.
	Identity *ApiManagementServiceIdentity_ARM `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the API Management service.
	Properties *ApiManagementServiceProperties_ARM `json:"properties,omitempty"`

	// Sku: SKU properties of the API Management service.
	Sku *ApiManagementServiceSkuProperties_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Service_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (service Service_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (service *Service_Spec_ARM) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service"
func (service *Service_Spec_ARM) GetType() string {
	return "Microsoft.ApiManagement/service"
}

// Identity properties of the Api Management service resource.
type ApiManagementServiceIdentity_ARM struct {
	// Type: The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly
	// created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
	Type                   *ApiManagementServiceIdentity_Type         `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityDetails_ARM `json:"userAssignedIdentities,omitempty"`
}

// Properties of an API Management service resource description.
type ApiManagementServiceProperties_ARM struct {
	// AdditionalLocations: Additional datacenter locations of the API Management service.
	AdditionalLocations []AdditionalLocation_ARM `json:"additionalLocations,omitempty"`

	// ApiVersionConstraint: Control Plane Apis version constraint for the API Management service.
	ApiVersionConstraint *ApiVersionConstraint_ARM `json:"apiVersionConstraint,omitempty"`

	// Certificates: List of Certificates that need to be installed in the API Management service. Max supported certificates
	// that can be installed is 10.
	Certificates []CertificateConfiguration_ARM `json:"certificates,omitempty"`

	// CustomProperties: Custom properties of the API Management service.</br>Setting
	// `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher
	// TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting
	// `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1.</br>Setting
	// `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API
	// Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11` can be
	// used to disable just TLS 1.1 for communications with backends.</br>Setting
	// `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10` can be used to disable TLS 1.0 for
	// communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2` can be
	// used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH
	// operation will reset omitted properties' values to their defaults. For all the settings except Http2 the default value
	// is `True` if the service was created on or before April 1, 2018 and `False` otherwise. Http2 setting's default value is
	// `False`.</br></br>You can disable any of the following ciphers by using settings
	// `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]`: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	// TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	// TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256,
	// TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example,
	// `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256`:`false`. The default
	// value is `true` for them.</br> Note: The following ciphers can't be disabled since they are required by internal
	// platform components:
	// TLS_AES_256_GCM_SHA384,TLS_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
	// disable the gateway in master region.
	DisableGateway *bool `json:"disableGateway,omitempty"`

	// EnableClientCertificate: Property only meant to be used for Consumption SKU Service. This enforces a client certificate
	// to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the
	// policy on the gateway.
	EnableClientCertificate *bool `json:"enableClientCertificate,omitempty"`

	// HostnameConfigurations: Custom hostname configuration of the API Management service.
	HostnameConfigurations []HostnameConfiguration_ARM `json:"hostnameConfigurations,omitempty"`

	// NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.
	NatGatewayState *ApiManagementServiceProperties_NatGatewayState `json:"natGatewayState,omitempty"`

	// NotificationSenderEmail: Email address from which the notification will be sent.
	NotificationSenderEmail *string `json:"notificationSenderEmail,omitempty"`

	// PublicIpAddressId: Public Standard SKU IP V4 based IP address to be associated with Virtual Network deployed service in
	// the region. Supported only for Developer and Premium SKU being deployed in Virtual Network.
	PublicIpAddressId *string `json:"publicIpAddressId,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this API Management service.  Value is
	// optional but if passed in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access
	// method. Default value is 'Enabled'
	PublicNetworkAccess *ApiManagementServiceProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// PublisherEmail: Publisher email.
	PublisherEmail *string `json:"publisherEmail,omitempty"`

	// PublisherName: Publisher name.
	PublisherName *string `json:"publisherName,omitempty"`

	// Restore: Undelete Api Management Service if it was previously soft-deleted. If this flag is specified and set to True
	// all other properties will be ignored.
	Restore *bool `json:"restore,omitempty"`

	// VirtualNetworkConfiguration: Virtual network configuration of the API Management service.
	VirtualNetworkConfiguration *VirtualNetworkConfiguration_ARM `json:"virtualNetworkConfiguration,omitempty"`

	// VirtualNetworkType: The type of VPN in which API Management service needs to be configured in. None (Default Value)
	// means the API Management service is not part of any Virtual Network, External means the API Management deployment is set
	// up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is
	// setup inside a Virtual Network having an Intranet Facing Endpoint only.
	VirtualNetworkType *ApiManagementServiceProperties_VirtualNetworkType `json:"virtualNetworkType,omitempty"`
}

// API Management service resource SKU properties.
type ApiManagementServiceSkuProperties_ARM struct {
	// Capacity: Capacity of the SKU (number of deployed units of the SKU). For Consumption SKU capacity must be specified as 0.
	Capacity *int `json:"capacity,omitempty"`

	// Name: Name of the Sku.
	Name *ApiManagementServiceSkuProperties_Name `json:"name,omitempty"`
}

// Description of an additional API Management resource location.
type AdditionalLocation_ARM struct {
	// DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
	// disable the gateway in this additional location.
	DisableGateway *bool `json:"disableGateway,omitempty"`

	// Location: The location name of the additional region among Azure Data center regions.
	Location *string `json:"location,omitempty"`

	// NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.
	NatGatewayState   *AdditionalLocation_NatGatewayState `json:"natGatewayState,omitempty"`
	PublicIpAddressId *string                             `json:"publicIpAddressId,omitempty"`

	// Sku: SKU properties of the API Management service.
	Sku *ApiManagementServiceSkuProperties_ARM `json:"sku,omitempty"`

	// VirtualNetworkConfiguration: Virtual network configuration for the location.
	VirtualNetworkConfiguration *VirtualNetworkConfiguration_ARM `json:"virtualNetworkConfiguration,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type ApiManagementServiceIdentity_Type string

const (
	ApiManagementServiceIdentity_Type_None                       = ApiManagementServiceIdentity_Type("None")
	ApiManagementServiceIdentity_Type_SystemAssigned             = ApiManagementServiceIdentity_Type("SystemAssigned")
	ApiManagementServiceIdentity_Type_SystemAssignedUserAssigned = ApiManagementServiceIdentity_Type("SystemAssigned, UserAssigned")
	ApiManagementServiceIdentity_Type_UserAssigned               = ApiManagementServiceIdentity_Type("UserAssigned")
)

// +kubebuilder:validation:Enum={"Basic","Consumption","Developer","Isolated","Premium","Standard"}
type ApiManagementServiceSkuProperties_Name string

const (
	ApiManagementServiceSkuProperties_Name_Basic       = ApiManagementServiceSkuProperties_Name("Basic")
	ApiManagementServiceSkuProperties_Name_Consumption = ApiManagementServiceSkuProperties_Name("Consumption")
	ApiManagementServiceSkuProperties_Name_Developer   = ApiManagementServiceSkuProperties_Name("Developer")
	ApiManagementServiceSkuProperties_Name_Isolated    = ApiManagementServiceSkuProperties_Name("Isolated")
	ApiManagementServiceSkuProperties_Name_Premium     = ApiManagementServiceSkuProperties_Name("Premium")
	ApiManagementServiceSkuProperties_Name_Standard    = ApiManagementServiceSkuProperties_Name("Standard")
)

// Control Plane Apis version constraint for the API Management service.
type ApiVersionConstraint_ARM struct {
	// MinApiVersion: Limit control plane API calls to API Management service with version equal to or newer than this value.
	MinApiVersion *string `json:"minApiVersion,omitempty"`
}

// Certificate configuration which consist of non-trusted intermediates and root certificates.
type CertificateConfiguration_ARM struct {
	// Certificate: Certificate information.
	Certificate *CertificateInformation_ARM `json:"certificate,omitempty"`

	// CertificatePassword: Certificate Password.
	CertificatePassword *string `json:"certificatePassword,omitempty"`

	// EncodedCertificate: Base64 Encoded certificate.
	EncodedCertificate *string `json:"encodedCertificate,omitempty"`

	// StoreName: The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and
	// CertificateAuthority are valid locations.
	StoreName *CertificateConfiguration_StoreName `json:"storeName,omitempty"`
}

// Custom hostname configuration.
type HostnameConfiguration_ARM struct {
	// Certificate: Certificate information.
	Certificate *CertificateInformation_ARM `json:"certificate,omitempty"`

	// CertificatePassword: Certificate Password.
	CertificatePassword *string `json:"certificatePassword,omitempty"`

	// CertificateSource: Certificate Source.
	CertificateSource *HostnameConfiguration_CertificateSource `json:"certificateSource,omitempty"`

	// CertificateStatus: Certificate Status.
	CertificateStatus *HostnameConfiguration_CertificateStatus `json:"certificateStatus,omitempty"`

	// DefaultSslBinding: Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate.
	// If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is
	// useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The
	// setting only applied to gateway Hostname Type.
	DefaultSslBinding *bool `json:"defaultSslBinding,omitempty"`

	// EncodedCertificate: Base64 Encoded certificate.
	EncodedCertificate *string `json:"encodedCertificate,omitempty"`

	// HostName: Hostname to configure on the Api Management service.
	HostName *string `json:"hostName,omitempty"`

	// IdentityClientId: System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to
	// the keyVault containing the SSL certificate.
	IdentityClientId *string `json:"identityClientId,omitempty" optionalConfigMapPair:"IdentityClientId"`

	// KeyVaultId: Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided,
	// auto-update of ssl certificate will not work. This requires Api Management service to be configured with aka.ms/apimmsi.
	// The secret should be of type *application/x-pkcs12*
	KeyVaultId *string `json:"keyVaultId,omitempty"`

	// NegotiateClientCertificate: Specify true to always negotiate client certificate on the hostname. Default Value is false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty"`

	// Type: Hostname type.
	Type *HostnameConfiguration_Type `json:"type,omitempty"`
}

// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails_ARM struct {
}

// Configuration of a virtual network to which API Management service is deployed.
type VirtualNetworkConfiguration_ARM struct {
	SubnetResourceId *string `json:"subnetResourceId,omitempty"`
}

// SSL certificate information.
type CertificateInformation_ARM struct {
	// Expiry: Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as
	// specified by the ISO 8601 standard.
	Expiry *string `json:"expiry,omitempty" optionalConfigMapPair:"Expiry"`

	// Subject: Subject of the certificate.
	Subject *string `json:"subject,omitempty" optionalConfigMapPair:"Subject"`

	// Thumbprint: Thumbprint of the certificate.
	Thumbprint *string `json:"thumbprint,omitempty" optionalConfigMapPair:"Thumbprint"`
}
