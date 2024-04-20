// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterInstanceInitParameters struct {

	// Specifies whether any database modifications are applied immediately, or during the next maintenance window. Default isfalse.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default true.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// EC2 Availability Zone that the DB instance is created in. See docs about the details.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Identifier of the CA certificate for the DB instance.
	CACertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier,omitempty"`

	// Identifier of the aws_rds_cluster in which to launch this instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in rds to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in rds to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// defined tags from the DB instance to snapshots of the DB instance. Default false.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// Instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
	CustomIAMInstanceProfile *string `json:"customIamInstanceProfile,omitempty" tf:"custom_iam_instance_profile,omitempty"`

	// Name of the DB parameter group to associate with this instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.ParameterGroup
	DBParameterGroupName *string `json:"dbParameterGroupName,omitempty" tf:"db_parameter_group_name,omitempty"`

	// Reference to a ParameterGroup in rds to populate dbParameterGroupName.
	// +kubebuilder:validation:Optional
	DBParameterGroupNameRef *v1.Reference `json:"dbParameterGroupNameRef,omitempty" tf:"-"`

	// Selector for a ParameterGroup in rds to populate dbParameterGroupName.
	// +kubebuilder:validation:Optional
	DBParameterGroupNameSelector *v1.Selector `json:"dbParameterGroupNameSelector,omitempty" tf:"-"`

	// DB subnet group to associate with this DB instance. NOTE: This must match the db_subnet_group_name of the attached aws_rds_cluster.
	// +crossplane:generate:reference:type=SubnetGroup
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// Reference to a SubnetGroup to populate dbSubnetGroupName.
	// +kubebuilder:validation:Optional
	DBSubnetGroupNameRef *v1.Reference `json:"dbSubnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup to populate dbSubnetGroupName.
	// +kubebuilder:validation:Optional
	DBSubnetGroupNameSelector *v1.Selector `json:"dbSubnetGroupNameSelector,omitempty" tf:"-"`

	// Name of the database engine to be used for the RDS cluster instance.
	// Valid Values: aurora-mysql, aurora-postgresql, mysql, postgres.(Note that mysql and postgres are Multi-AZ RDS clusters).
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Database engine version. Please note that to upgrade the engine_version of the instance, it must be done on the aws_rds_cluster engine_version. Trying to upgrade in aws_cluster_instance will not update the engine_version.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Instance class to use. For details on CPU and memory, see Scaling Aurora DB Instances. Aurora uses db.* instance classes/types. Please see AWS Documentation for currently available instance classes and complete details.
	InstanceClass *string `json:"instanceClass,omitempty" tf:"instance_class,omitempty"`

	// Interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval *float64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`

	// ARN for the IAM role that permits RDS to send enhanced monitoring metrics to CloudWatch Logs. You can find more information on the AWS Documentation what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	MonitoringRoleArn *string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn,omitempty"`

	// Reference to a Role in iam to populate monitoringRoleArn.
	// +kubebuilder:validation:Optional
	MonitoringRoleArnRef *v1.Reference `json:"monitoringRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate monitoringRoleArn.
	// +kubebuilder:validation:Optional
	MonitoringRoleArnSelector *v1.Selector `json:"monitoringRoleArnSelector,omitempty" tf:"-"`

	// Specifies whether Performance Insights is enabled or not.
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`

	// ARN for the KMS key to encrypt Performance Insights data. When specifying performance_insights_kms_key_id, performance_insights_enabled needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	PerformanceInsightsKMSKeyID *string `json:"performanceInsightsKmsKeyId,omitempty" tf:"performance_insights_kms_key_id,omitempty"`

	// Reference to a Key in kms to populate performanceInsightsKmsKeyId.
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyIDRef *v1.Reference `json:"performanceInsightsKmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate performanceInsightsKmsKeyId.
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyIDSelector *v1.Selector `json:"performanceInsightsKmsKeyIdSelector,omitempty" tf:"-"`

	// Amount of time in days to retain Performance Insights data. Valid values are 7, 731 (2 years) or a multiple of 31. When specifying performance_insights_retention_period, performance_insights_enabled needs to be set to true. Defaults to '7'.
	PerformanceInsightsRetentionPeriod *float64 `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period,omitempty"`

	// Daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00". NOTE: If preferred_backup_window is set at the cluster level, this argument must be omitted.
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// Window to perform maintenance in. Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	PromotionTier *float64 `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`

	// Bool to control if instance is publicly accessible. Default false. See the documentation on Creating DB Instances for more details on controlling this property.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterInstanceObservation struct {

	// Specifies whether any database modifications are applied immediately, or during the next maintenance window. Default isfalse.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Amazon Resource Name (ARN) of cluster instance
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default true.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// EC2 Availability Zone that the DB instance is created in. See docs about the details.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Identifier of the CA certificate for the DB instance.
	CACertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier,omitempty"`

	// Identifier of the aws_rds_cluster in which to launch this instance.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// defined tags from the DB instance to snapshots of the DB instance. Default false.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// Instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
	CustomIAMInstanceProfile *string `json:"customIamInstanceProfile,omitempty" tf:"custom_iam_instance_profile,omitempty"`

	// Name of the DB parameter group to associate with this instance.
	DBParameterGroupName *string `json:"dbParameterGroupName,omitempty" tf:"db_parameter_group_name,omitempty"`

	// DB subnet group to associate with this DB instance. NOTE: This must match the db_subnet_group_name of the attached aws_rds_cluster.
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// Region-unique, immutable identifier for the DB instance.
	DbiResourceID *string `json:"dbiResourceId,omitempty" tf:"dbi_resource_id,omitempty"`

	// DNS address for this instance. May not be writable
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Name of the database engine to be used for the RDS cluster instance.
	// Valid Values: aurora-mysql, aurora-postgresql, mysql, postgres.(Note that mysql and postgres are Multi-AZ RDS clusters).
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Database engine version. Please note that to upgrade the engine_version of the instance, it must be done on the aws_rds_cluster engine_version. Trying to upgrade in aws_cluster_instance will not update the engine_version.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Database engine version
	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	// Instance identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance class to use. For details on CPU and memory, see Scaling Aurora DB Instances. Aurora uses db.* instance classes/types. Please see AWS Documentation for currently available instance classes and complete details.
	InstanceClass *string `json:"instanceClass,omitempty" tf:"instance_class,omitempty"`

	// ARN for the KMS encryption key if one is set to the cluster.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	MonitoringInterval *float64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`

	// ARN for the IAM role that permits RDS to send enhanced monitoring metrics to CloudWatch Logs. You can find more information on the AWS Documentation what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	MonitoringRoleArn *string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn,omitempty"`

	// Network type of the DB instance.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// Specifies whether Performance Insights is enabled or not.
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`

	// ARN for the KMS key to encrypt Performance Insights data. When specifying performance_insights_kms_key_id, performance_insights_enabled needs to be set to true.
	PerformanceInsightsKMSKeyID *string `json:"performanceInsightsKmsKeyId,omitempty" tf:"performance_insights_kms_key_id,omitempty"`

	// Amount of time in days to retain Performance Insights data. Valid values are 7, 731 (2 years) or a multiple of 31. When specifying performance_insights_retention_period, performance_insights_enabled needs to be set to true. Defaults to '7'.
	PerformanceInsightsRetentionPeriod *float64 `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period,omitempty"`

	// Database port
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00". NOTE: If preferred_backup_window is set at the cluster level, this argument must be omitted.
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// Window to perform maintenance in. Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	PromotionTier *float64 `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`

	// Bool to control if instance is publicly accessible. Default false. See the documentation on Creating DB Instances for more details on controlling this property.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// – Boolean indicating if this instance is writable. False indicates this instance is a read replica.
	Writer *bool `json:"writer,omitempty" tf:"writer,omitempty"`
}

type ClusterInstanceParameters struct {

	// Specifies whether any database modifications are applied immediately, or during the next maintenance window. Default isfalse.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default true.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// EC2 Availability Zone that the DB instance is created in. See docs about the details.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Identifier of the CA certificate for the DB instance.
	// +kubebuilder:validation:Optional
	CACertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier,omitempty"`

	// Identifier of the aws_rds_cluster in which to launch this instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in rds to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in rds to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// defined tags from the DB instance to snapshots of the DB instance. Default false.
	// +kubebuilder:validation:Optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// Instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
	// +kubebuilder:validation:Optional
	CustomIAMInstanceProfile *string `json:"customIamInstanceProfile,omitempty" tf:"custom_iam_instance_profile,omitempty"`

	// Name of the DB parameter group to associate with this instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.ParameterGroup
	// +kubebuilder:validation:Optional
	DBParameterGroupName *string `json:"dbParameterGroupName,omitempty" tf:"db_parameter_group_name,omitempty"`

	// Reference to a ParameterGroup in rds to populate dbParameterGroupName.
	// +kubebuilder:validation:Optional
	DBParameterGroupNameRef *v1.Reference `json:"dbParameterGroupNameRef,omitempty" tf:"-"`

	// Selector for a ParameterGroup in rds to populate dbParameterGroupName.
	// +kubebuilder:validation:Optional
	DBParameterGroupNameSelector *v1.Selector `json:"dbParameterGroupNameSelector,omitempty" tf:"-"`

	// DB subnet group to associate with this DB instance. NOTE: This must match the db_subnet_group_name of the attached aws_rds_cluster.
	// +crossplane:generate:reference:type=SubnetGroup
	// +kubebuilder:validation:Optional
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// Reference to a SubnetGroup to populate dbSubnetGroupName.
	// +kubebuilder:validation:Optional
	DBSubnetGroupNameRef *v1.Reference `json:"dbSubnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup to populate dbSubnetGroupName.
	// +kubebuilder:validation:Optional
	DBSubnetGroupNameSelector *v1.Selector `json:"dbSubnetGroupNameSelector,omitempty" tf:"-"`

	// Name of the database engine to be used for the RDS cluster instance.
	// Valid Values: aurora-mysql, aurora-postgresql, mysql, postgres.(Note that mysql and postgres are Multi-AZ RDS clusters).
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Database engine version. Please note that to upgrade the engine_version of the instance, it must be done on the aws_rds_cluster engine_version. Trying to upgrade in aws_cluster_instance will not update the engine_version.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Instance class to use. For details on CPU and memory, see Scaling Aurora DB Instances. Aurora uses db.* instance classes/types. Please see AWS Documentation for currently available instance classes and complete details.
	// +kubebuilder:validation:Optional
	InstanceClass *string `json:"instanceClass,omitempty" tf:"instance_class,omitempty"`

	// Interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
	// +kubebuilder:validation:Optional
	MonitoringInterval *float64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`

	// ARN for the IAM role that permits RDS to send enhanced monitoring metrics to CloudWatch Logs. You can find more information on the AWS Documentation what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	MonitoringRoleArn *string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn,omitempty"`

	// Reference to a Role in iam to populate monitoringRoleArn.
	// +kubebuilder:validation:Optional
	MonitoringRoleArnRef *v1.Reference `json:"monitoringRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate monitoringRoleArn.
	// +kubebuilder:validation:Optional
	MonitoringRoleArnSelector *v1.Selector `json:"monitoringRoleArnSelector,omitempty" tf:"-"`

	// Specifies whether Performance Insights is enabled or not.
	// +kubebuilder:validation:Optional
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`

	// ARN for the KMS key to encrypt Performance Insights data. When specifying performance_insights_kms_key_id, performance_insights_enabled needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyID *string `json:"performanceInsightsKmsKeyId,omitempty" tf:"performance_insights_kms_key_id,omitempty"`

	// Reference to a Key in kms to populate performanceInsightsKmsKeyId.
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyIDRef *v1.Reference `json:"performanceInsightsKmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate performanceInsightsKmsKeyId.
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyIDSelector *v1.Selector `json:"performanceInsightsKmsKeyIdSelector,omitempty" tf:"-"`

	// Amount of time in days to retain Performance Insights data. Valid values are 7, 731 (2 years) or a multiple of 31. When specifying performance_insights_retention_period, performance_insights_enabled needs to be set to true. Defaults to '7'.
	// +kubebuilder:validation:Optional
	PerformanceInsightsRetentionPeriod *float64 `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period,omitempty"`

	// Daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00". NOTE: If preferred_backup_window is set at the cluster level, this argument must be omitted.
	// +kubebuilder:validation:Optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// Window to perform maintenance in. Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
	// +kubebuilder:validation:Optional
	PromotionTier *float64 `json:"promotionTier,omitempty" tf:"promotion_tier,omitempty"`

	// Bool to control if instance is publicly accessible. Default false. See the documentation on Creating DB Instances for more details on controlling this property.
	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ClusterInstanceSpec defines the desired state of ClusterInstance
type ClusterInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterInstanceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClusterInstanceInitParameters `json:"initProvider,omitempty"`
}

// ClusterInstanceStatus defines the observed state of ClusterInstance.
type ClusterInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClusterInstance is the Schema for the ClusterInstances API. Provides an RDS Cluster Resource Instance
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ClusterInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || (has(self.initProvider) && has(self.initProvider.engine))",message="spec.forProvider.engine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceClass) || (has(self.initProvider) && has(self.initProvider.instanceClass))",message="spec.forProvider.instanceClass is a required parameter"
	Spec   ClusterInstanceSpec   `json:"spec"`
	Status ClusterInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterInstanceList contains a list of ClusterInstances
type ClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInstance `json:"items"`
}

// Repository type metadata.
var (
	ClusterInstance_Kind             = "ClusterInstance"
	ClusterInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterInstance_Kind}.String()
	ClusterInstance_KindAPIVersion   = ClusterInstance_Kind + "." + CRDGroupVersion.String()
	ClusterInstance_GroupVersionKind = CRDGroupVersion.WithKind(ClusterInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterInstance{}, &ClusterInstanceList{})
}
