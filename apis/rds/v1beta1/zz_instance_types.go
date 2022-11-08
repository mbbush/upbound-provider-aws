/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceObservation struct {

	// The hostname of the RDS instance. See also endpoint and port.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The ARN of the RDS instance.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The connection endpoint in address:port format.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The running version of the database.
	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	// The canonical hosted zone ID of the DB instance (to be used
	// in a Route 53 Alias record).
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// The RDS instance ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The latest time, in UTC RFC3339 format, to which a database can be restored with point-in-time restore.
	LatestRestorableTime *string `json:"latestRestorableTime,omitempty" tf:"latest_restorable_time,omitempty"`

	Replicas []*string `json:"replicas,omitempty" tf:"replicas,omitempty"`

	// The RDS Resource ID of this instance.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// The RDS instance status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type InstanceParameters struct {

	// The allocated storage in gibibytes. If max_allocated_storage is configured, this argument represents the initial storage allocation and differences from the configuration will be ignored automatically when Storage Autoscaling occurs. If replicate_source_db is set, the value is ignored during the creation of the instance.
	// +kubebuilder:validation:Optional
	AllocatedStorage *float64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// Indicates that major version
	// upgrades are allowed. Changing this parameter does not result in an outage and
	// the change is asynchronously applied as soon as possible.
	// +kubebuilder:validation:Optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// false. See Amazon RDS Documentation for more
	// information.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// Indicates that minor engine upgrades
	// will be applied automatically to the DB instance during the maintenance window.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// The AZ for the RDS instance.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The days to retain backups for. Must be
	// between 0 and 35. Must be greater than 0 if the database is used as a source for a Read Replica. See Read Replica.
	// +kubebuilder:validation:Optional
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// The daily time range (in UTC) during which
	// automated backups are created if they are enabled. Example: "09:46-10:16". Must
	// not overlap with maintenance_window.
	// +kubebuilder:validation:Optional
	BackupWindow *string `json:"backupWindow,omitempty" tf:"backup_window,omitempty"`

	// The identifier of the CA certificate for the DB instance.
	// +kubebuilder:validation:Optional
	CACertIdentifier *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier,omitempty"`

	// The character set name to use for DB
	// encoding in Oracle and Microsoft SQL instances (collation). This can't be changed. See Oracle Character Sets
	// Supported in Amazon RDS
	// or Server-Level Collation for Microsoft SQL Server for more information.
	// +kubebuilder:validation:Optional
	CharacterSetName *string `json:"characterSetName,omitempty" tf:"character_set_name,omitempty"`

	// –  Copy all Instance tags to snapshots. Default is false.
	// +kubebuilder:validation:Optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// Indicates whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance. See CoIP for RDS on Outposts for more information.
	// +kubebuilder:validation:Optional
	CustomerOwnedIPEnabled *bool `json:"customerOwnedIpEnabled,omitempty" tf:"customer_owned_ip_enabled,omitempty"`

	// The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance. Note that this does not apply for Oracle or SQL Server engines. See the AWS documentation for more details on what applies for those engines. If you are providing an Oracle db name, it needs to be in all upper case. Cannot be specified for a replica.
	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Name of DB subnet group. DB instance will
	// be created in the VPC associated with the DB subnet group. If unspecified, will
	// be created in the default VPC, or in EC2 Classic, if available. When working
	// with read replicas, it should be specified only if the source database
	// specifies an instance in another AWS Region. See DBSubnetGroupName in API
	// action CreateDBInstanceReadReplica
	// for additional read replica contraints.
	// +crossplane:generate:reference:type=SubnetGroup
	// +kubebuilder:validation:Optional
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name,omitempty"`

	// Reference to a SubnetGroup to populate dbSubnetGroupName.
	// +kubebuilder:validation:Optional
	DBSubnetGroupNameRef *v1.Reference `json:"dbSubnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup to populate dbSubnetGroupName.
	// +kubebuilder:validation:Optional
	DBSubnetGroupNameSelector *v1.Selector `json:"dbSubnetGroupNameSelector,omitempty" tf:"-"`

	// Specifies whether to remove automated backups immediately after the DB instance is deleted. Default is true.
	// +kubebuilder:validation:Optional
	DeleteAutomatedBackups *bool `json:"deleteAutomatedBackups,omitempty" tf:"delete_automated_backups,omitempty"`

	// If the DB instance should have deletion protection enabled. The database can't be deleted when this value is set to true. The default is false.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The ID of the Directory Service Active Directory domain to create the instance in.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The name of the IAM role to be used when making API calls to the Directory Service.
	// +kubebuilder:validation:Optional
	DomainIAMRoleName *string `json:"domainIamRoleName,omitempty" tf:"domain_iam_role_name,omitempty"`

	// Set of log types to enable for exporting to CloudWatch logs. If omitted, no logs will be exported. Valid values (depending on engine). MySQL and MariaDB: audit, error, general, slowquery. PostgreSQL: postgresql, upgrade. MSSQL: agent , error. Oracle: alert, audit, listener, trace.
	// +kubebuilder:validation:Optional
	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports,omitempty"`

	// The database engine to use.  For supported values, see the Engine parameter in API action CreateDBInstance. Cannot be specified for a replica.
	// Note that for Amazon Aurora instances the engine must match the DB cluster's engine'.
	// For information on the difference between the available Aurora MySQL engines
	// see Comparison between Aurora MySQL 1 and Aurora MySQL 2
	// in the Amazon RDS User Guide.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The engine version to use. If auto_minor_version_upgrade
	// is enabled, you can provide a prefix of the version such as 5.7 (for 5.7.10).
	// The actual engine version used is returned in the attribute engine_version_actual, see Attributes Reference below.
	// For supported values, see the EngineVersion parameter in API action CreateDBInstance.
	// Note that for Amazon Aurora instances the engine version must match the DB cluster's engine version'. Cannot be specified for a replica.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The name of your final DB snapshot
	// when this DB instance is deleted. Must be provided if skip_final_snapshot is
	// set to false. The value must begin with a letter, only contain alphanumeric characters and hyphens, and not end with a hyphen or contain two consecutive hyphens. Must not be provided when deleting a read replica.
	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// Specifies whether or
	// mappings of AWS Identity and Access Management (IAM) accounts to database
	// accounts is enabled.
	// +kubebuilder:validation:Optional
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`

	// The instance type of the RDS instance.
	// +kubebuilder:validation:Required
	InstanceClass *string `json:"instanceClass" tf:"instance_class,omitempty"`

	// The amount of provisioned IOPS. Setting this implies a
	// storage_type of "io1".
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The ARN for the KMS encryption key. If creating an
	// encrypted replica, set this to the destination KMS ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// License model information for this DB instance.
	// +kubebuilder:validation:Optional
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`

	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00". See RDS
	// Maintenance Window
	// docs
	// for more information.
	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// When configured, the upper limit to which Amazon RDS can automatically scale the storage of the DB instance. Configuring this will automatically ignore differences to allocated_storage. Must be greater than or equal to allocated_storage or 0 to disable Storage Autoscaling.
	// +kubebuilder:validation:Optional
	MaxAllocatedStorage *float64 `json:"maxAllocatedStorage,omitempty" tf:"max_allocated_storage,omitempty"`

	// The interval, in seconds, between points
	// when Enhanced Monitoring metrics are collected for the DB instance. To disable
	// collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid
	// Values: 0, 1, 5, 10, 15, 30, 60.
	// +kubebuilder:validation:Optional
	MonitoringInterval *float64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval,omitempty"`

	// The ARN for the IAM role that permits RDS
	// to send enhanced monitoring metrics to CloudWatch Logs. You can find more
	// information on the AWS
	// Documentation
	// what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
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

	// Specifies if the RDS instance is multi-AZ
	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance. Note that this does not apply for Oracle or SQL Server engines. See the AWS documentation for more details on what applies for those engines. If you are providing an Oracle db name, it needs to be in all upper case. Cannot be specified for a replica.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The national character set is used in the NCHAR, NVARCHAR2, and NCLOB data types for Oracle instances. This can't be changed. See Oracle Character Sets
	// Supported in Amazon RDS.
	// +kubebuilder:validation:Optional
	NcharCharacterSetName *string `json:"ncharCharacterSetName,omitempty" tf:"nchar_character_set_name,omitempty"`

	// Name of the DB option group to associate.
	// +kubebuilder:validation:Optional
	OptionGroupName *string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`

	// Name of the DB parameter group to
	// associate.
	// +kubebuilder:validation:Optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// Password for the master DB user. Note that this may show up in
	// logs, and it will be stored in the state file.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Specifies whether Performance Insights are enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled,omitempty"`

	// The ARN for the KMS key to encrypt Performance Insights data. When specifying performance_insights_kms_key_id, performance_insights_enabled needs to be set to true. Once KMS key is set, it can never be changed.
	// +kubebuilder:validation:Optional
	PerformanceInsightsKMSKeyID *string `json:"performanceInsightsKmsKeyId,omitempty" tf:"performance_insights_kms_key_id,omitempty"`

	// The amount of time in days to retain Performance Insights data. Either 7 (7 days) or 731 (2 years). When specifying performance_insights_retention_period, performance_insights_enabled needs to be set to true. Defaults to '7'.
	// +kubebuilder:validation:Optional
	PerformanceInsightsRetentionPeriod *float64 `json:"performanceInsightsRetentionPeriod,omitempty" tf:"performance_insights_retention_period,omitempty"`

	// The port on which the DB accepts connections.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Bool to control if instance is publicly
	// accessible. Default is false.
	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies whether the replica is in either mounted or open-read-only mode. This attribute
	// is only supported by Oracle instances. Oracle replicas operate in open-read-only mode unless otherwise specified. See Working with Oracle Read Replicas for more information.
	// +kubebuilder:validation:Optional
	ReplicaMode *string `json:"replicaMode,omitempty" tf:"replica_mode,omitempty"`

	// Specifies that this resource is a Replicate
	// database, and to use this value as the source database. This correlates to the
	// identifier of another Amazon RDS Database to replicate (if replicating within
	// a single region) or ARN of the Amazon RDS Database to replicate (if replicating
	// cross-region). Note that if you are
	// creating a cross-region replica of an encrypted database you will also need to
	// specify a kms_key_id. See DB Instance Replication and Working with
	// PostgreSQL and MySQL Read Replicas
	// for more information on using Replication.
	// +kubebuilder:validation:Optional
	ReplicateSourceDB *string `json:"replicateSourceDb,omitempty" tf:"replicate_source_db,omitempty"`

	// A configuration block for restoring a DB instance to an arbitrary point in time. Requires the identifier argument to be set with the name of the new DB instance to be created. See Restore To Point In Time below for details.
	// +kubebuilder:validation:Optional
	RestoreToPointInTime []RestoreToPointInTimeParameters `json:"restoreToPointInTime,omitempty" tf:"restore_to_point_in_time,omitempty"`

	// Restore from a Percona Xtrabackup in S3.  See Importing Data into an Amazon RDS MySQL DB Instance
	// +kubebuilder:validation:Optional
	S3Import []S3ImportParameters `json:"s3Import,omitempty" tf:"s3_import,omitempty"`

	// List of DB Security Groups to
	// associate. Only used for DB Instances on the .
	// +kubebuilder:validation:Optional
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// Determines whether a final DB snapshot is
	// created before the DB instance is deleted. If true is specified, no DBSnapshot
	// is created. If false is specified, a DB snapshot is created before the DB
	// instance is deleted, using the value from final_snapshot_identifier. Default
	// is false.
	// +kubebuilder:validation:Optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// Specifies whether or not to create this
	// database from a snapshot. This correlates to the snapshot ID you'd find in the
	// RDS console, e.g: rds:production-2015-06-26-06-05.
	// +kubebuilder:validation:Optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Specifies whether the DB instance is
	// encrypted. Note that if you are creating a cross-region read replica this field
	// is ignored and you should instead declare kms_key_id with a valid ARN. The
	// default is false if not specified.
	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// One of "standard" (magnetic), "gp2" (general
	// purpose SSD), or "io1" (provisioned IOPS SSD). The default is "io1" if iops is
	// specified, "gp2" if not.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Time zone of the DB instance. timezone is currently
	// only supported by Microsoft SQL Server. The timezone can only be set on
	// creation. See MSSQL User
	// Guide
	// for more information.
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// Username for the master DB user. Cannot be specified for a replica.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// List of VPC security groups to
	// associate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type RestoreToPointInTimeObservation struct {
}

type RestoreToPointInTimeParameters struct {

	// The date and time to restore from. Value must be a time in Universal Coordinated Time (UTC) format and must be before the latest restorable time for the DB instance. Cannot be specified with use_latest_restorable_time.
	// +kubebuilder:validation:Optional
	RestoreTime *string `json:"restoreTime,omitempty" tf:"restore_time,omitempty"`

	// The ARN of the automated backup from which to restore. Required if source_db_instance_identifier or source_dbi_resource_id is not specified.
	// +kubebuilder:validation:Optional
	SourceDBInstanceAutomatedBackupsArn *string `json:"sourceDbInstanceAutomatedBackupsArn,omitempty" tf:"source_db_instance_automated_backups_arn,omitempty"`

	// The identifier of the source DB instance from which to restore. Must match the identifier of an existing DB instance. Required if source_db_instance_automated_backups_arn or source_dbi_resource_id is not specified.
	// +kubebuilder:validation:Optional
	SourceDBInstanceIdentifier *string `json:"sourceDbInstanceIdentifier,omitempty" tf:"source_db_instance_identifier,omitempty"`

	// The resource ID of the source DB instance from which to restore. Required if source_db_instance_identifier or source_db_instance_automated_backups_arn is not specified.
	// +kubebuilder:validation:Optional
	SourceDbiResourceID *string `json:"sourceDbiResourceId,omitempty" tf:"source_dbi_resource_id,omitempty"`

	// A boolean value that indicates whether the DB instance is restored from the latest backup time. Defaults to false. Cannot be specified with restore_time.
	// +kubebuilder:validation:Optional
	UseLatestRestorableTime *bool `json:"useLatestRestorableTime,omitempty" tf:"use_latest_restorable_time,omitempty"`
}

type S3ImportObservation struct {
}

type S3ImportParameters struct {

	// The bucket name where your backup is stored
	// +kubebuilder:validation:Required
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// Can be blank, but is the path to your backup
	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// Role applied to load the data.
	// +kubebuilder:validation:Required
	IngestionRole *string `json:"ingestionRole" tf:"ingestion_role,omitempty"`

	// Source engine for the backup
	// +kubebuilder:validation:Required
	SourceEngine *string `json:"sourceEngine" tf:"source_engine,omitempty"`

	// Version of the source engine used to make the backup
	// +kubebuilder:validation:Required
	SourceEngineVersion *string `json:"sourceEngineVersion" tf:"source_engine_version,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Provides an RDS instance resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}