// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AWSManagementClusterParams a w s management cluster params
// swagger:model AWSManagementClusterParams
type AWSManagementClusterParams struct {

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// aws account params
	AwsAccountParams *AWSAccountParams `json:"awsAccountParams,omitempty"`

	// bastion host enabled
	BastionHostEnabled bool `json:"bastionHostEnabled,omitempty"`

	// ceip opt in
	CeipOptIn *bool `json:"ceipOptIn,omitempty"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// control plane flavor
	ControlPlaneFlavor string `json:"controlPlaneFlavor,omitempty"`

	// control plane node type
	ControlPlaneNodeType string `json:"controlPlaneNodeType,omitempty"`

	// create cloud formation stack
	CreateCloudFormationStack bool `json:"createCloudFormationStack,omitempty"`

	// enable audit logging
	EnableAuditLogging bool `json:"enableAuditLogging,omitempty"`

	// identity management
	IdentityManagement *IdentityManagementConfig `json:"identityManagement,omitempty"`

	// kubernetes version
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// loadbalancer scheme internal
	LoadbalancerSchemeInternal bool `json:"loadbalancerSchemeInternal,omitempty"`

	// machine health check enabled
	MachineHealthCheckEnabled bool `json:"machineHealthCheckEnabled,omitempty"`

	// networking
	Networking *TKGNetwork `json:"networking,omitempty"`

	// num of worker node
	NumOfWorkerNode int64 `json:"numOfWorkerNode,omitempty"`

	// os
	Os *AWSVirtualMachine `json:"os,omitempty"`

	// ssh key name
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// vpc
	Vpc *AWSVpc `json:"vpc,omitempty"`

	// worker node type
	WorkerNodeType string `json:"workerNodeType,omitempty"`
}

// Validate validates this a w s management cluster params
func (m *AWSManagementClusterParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsAccountParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityManagement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AWSManagementClusterParams) validateAwsAccountParams(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsAccountParams) { // not required
		return nil
	}

	if m.AwsAccountParams != nil {
		if err := m.AwsAccountParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsAccountParams")
			}
			return err
		}
	}

	return nil
}

func (m *AWSManagementClusterParams) validateIdentityManagement(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentityManagement) { // not required
		return nil
	}

	if m.IdentityManagement != nil {
		if err := m.IdentityManagement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identityManagement")
			}
			return err
		}
	}

	return nil
}

func (m *AWSManagementClusterParams) validateNetworking(formats strfmt.Registry) error {

	if swag.IsZero(m.Networking) { // not required
		return nil
	}

	if m.Networking != nil {
		if err := m.Networking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networking")
			}
			return err
		}
	}

	return nil
}

func (m *AWSManagementClusterParams) validateOs(formats strfmt.Registry) error {

	if swag.IsZero(m.Os) { // not required
		return nil
	}

	if m.Os != nil {
		if err := m.Os.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os")
			}
			return err
		}
	}

	return nil
}

func (m *AWSManagementClusterParams) validateVpc(formats strfmt.Registry) error {

	if swag.IsZero(m.Vpc) { // not required
		return nil
	}

	if m.Vpc != nil {
		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AWSManagementClusterParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSManagementClusterParams) UnmarshalBinary(b []byte) error {
	var res AWSManagementClusterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}