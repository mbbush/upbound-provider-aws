// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LogConfigurationInitParameters struct {

	// The Amazon Resource Name (ARN) that specifies the destination of the logs. The name of the Amazon CloudWatch Logs log group must begin with the /aws/fsx prefix. If you do not provide a destination, Amazon FSx will create and use a log stream in the CloudWatch Logs /aws/fsx/lustre log group.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Sets which data repository events are logged by Amazon FSx. Valid values are WARN_ONLY, FAILURE_ONLY, ERROR_ONLY, WARN_ERROR and DISABLED. Default value is DISABLED.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`
}

type LogConfigurationObservation struct {

	// The Amazon Resource Name (ARN) that specifies the destination of the logs. The name of the Amazon CloudWatch Logs log group must begin with the /aws/fsx prefix. If you do not provide a destination, Amazon FSx will create and use a log stream in the CloudWatch Logs /aws/fsx/lustre log group.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Sets which data repository events are logged by Amazon FSx. Valid values are WARN_ONLY, FAILURE_ONLY, ERROR_ONLY, WARN_ERROR and DISABLED. Default value is DISABLED.
	Level *string `json:"level,omitempty" tf:"level,omitempty"`
}

type LogConfigurationParameters struct {

	// The Amazon Resource Name (ARN) that specifies the destination of the logs. The name of the Amazon CloudWatch Logs log group must begin with the /aws/fsx prefix. If you do not provide a destination, Amazon FSx will create and use a log stream in the CloudWatch Logs /aws/fsx/lustre log group.
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Sets which data repository events are logged by Amazon FSx. Valid values are WARN_ONLY, FAILURE_ONLY, ERROR_ONLY, WARN_ERROR and DISABLED. Default value is DISABLED.
	// +kubebuilder:validation:Optional
	Level *string `json:"level,omitempty" tf:"level,omitempty"`
}

type LustreFileSystemInitParameters struct {

	// How Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. see Auto Import Data Repo for more details. Only supported on PERSISTENT_1 deployment types.
	AutoImportPolicy *string `json:"autoImportPolicy,omitempty" tf:"auto_import_policy,omitempty"`

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days. only valid for PERSISTENT_1 and PERSISTENT_2 deployment_type.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// The ID of the source backup to create the filesystem from.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// A boolean flag indicating whether tags for the file system should be copied to backups. Applicable for PERSISTENT_1 and PERSISTENT_2 deployment_type. The default value is false.
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. only valid for PERSISTENT_1 and PERSISTENT_2 deployment_type. Requires automatic_backup_retention_days to be set.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// Sets the data compression configuration for the file system. Valid values are LZ4 and NONE. Default value is NONE. Unsetting this value reverts the compression type back to NONE.
	DataCompressionType *string `json:"dataCompressionType,omitempty" tf:"data_compression_type,omitempty"`

	// - The filesystem deployment type. One of: SCRATCH_1, SCRATCH_2, PERSISTENT_1, PERSISTENT_2.
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// - The type of drive cache used by PERSISTENT_1 filesystems that are provisioned with HDD storage_type. Required for HDD storage_type, set to either READ or NONE.
	DriveCacheType *string `json:"driveCacheType,omitempty" tf:"drive_cache_type,omitempty"`

	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with import_path argument and the path must use the same Amazon S3 bucket as specified in import_path. Set equal to import_path to overwrite files on export. Defaults to s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}. Only supported on PERSISTENT_1 deployment types.
	ExportPath *string `json:"exportPath,omitempty" tf:"export_path,omitempty"`

	// Sets the Lustre version for the file system that you're creating. Valid values are 2.10 for SCRATCH_1, SCRATCH_2 and PERSISTENT_1 deployment types. Valid values for 2.12 include all deployment types.
	FileSystemTypeVersion *string `json:"fileSystemTypeVersion,omitempty" tf:"file_system_type_version,omitempty"`

	// A map of tags to apply to the file system's final backup.
	// +mapType=granular
	FinalBackupTags map[string]*string `json:"finalBackupTags,omitempty" tf:"final_backup_tags,omitempty"`

	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, s3://example-bucket/optional-prefix/. Only supported on PERSISTENT_1 deployment types.
	ImportPath *string `json:"importPath,omitempty" tf:"import_path,omitempty"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with import_path argument. Defaults to 1024. Minimum of 1 and maximum of 512000. Only supported on PERSISTENT_1 deployment types.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize,omitempty" tf:"imported_file_chunk_size,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, applicable for PERSISTENT_1 and PERSISTENT_2 deployment_type. Defaults to an AWS managed KMS Key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The Lustre logging configuration used when creating an Amazon FSx for Lustre file system. When logging is enabled, Lustre logs error and warning events for data repositories associated with your file system to Amazon CloudWatch Logs. See log_configuration Block for details.
	LogConfiguration *LogConfigurationInitParameters `json:"logConfiguration,omitempty" tf:"log_configuration,omitempty"`

	// The Lustre metadata configuration used when creating an Amazon FSx for Lustre file system. This can be used to specify a user provisioned metadata scale. This is only supported when deployment_type is set to PERSISTENT_2. See metadata_configuration Block for details.
	MetadataConfiguration *MetadataConfigurationInitParameters `json:"metadataConfiguration,omitempty" tf:"metadata_configuration,omitempty"`

	// - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the PERSISTENT_1 and PERSISTENT_2 deployment_type. Valid values for PERSISTENT_1 deployment_type and SSD storage_type are 50, 100, 200. Valid values for PERSISTENT_1 deployment_type and HDD storage_type are 12, 40. Valid values for PERSISTENT_2 deployment_type and  SSD storage_type are 125, 250, 500, 1000.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput,omitempty" tf:"per_unit_storage_throughput,omitempty"`

	// The Lustre root squash configuration used when creating an Amazon FSx for Lustre file system. When enabled, root squash restricts root-level access from clients that try to access your file system as a root user. See root_squash_configuration Block for details.
	RootSquashConfiguration *RootSquashConfigurationInitParameters `json:"rootSquashConfiguration,omitempty" tf:"root_squash_configuration,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to true.
	SkipFinalBackup *bool `json:"skipFinalBackup,omitempty" tf:"skip_final_backup,omitempty"`

	// The storage capacity (GiB) of the file system. Minimum of 1200. See more details at Allowed values for Fsx storage capacity. Update is allowed only for SCRATCH_2, PERSISTENT_1 and PERSISTENT_2 deployment types, See more details at Fsx Storage Capacity Update. Required when not creating filesystem for a backup.
	StorageCapacity *float64 `json:"storageCapacity,omitempty" tf:"storage_capacity,omitempty"`

	// - The filesystem storage type. Either SSD or HDD, defaults to SSD. HDD is only supported on PERSISTENT_1 deployment types.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The preferred start time (in d:HH:MM format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

type LustreFileSystemObservation struct {

	// Amazon Resource Name of the file system.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// How Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. see Auto Import Data Repo for more details. Only supported on PERSISTENT_1 deployment types.
	AutoImportPolicy *string `json:"autoImportPolicy,omitempty" tf:"auto_import_policy,omitempty"`

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days. only valid for PERSISTENT_1 and PERSISTENT_2 deployment_type.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// The ID of the source backup to create the filesystem from.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// A boolean flag indicating whether tags for the file system should be copied to backups. Applicable for PERSISTENT_1 and PERSISTENT_2 deployment_type. The default value is false.
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups,omitempty"`

	// DNS name for the file system, e.g., fs-12345678.fsx.us-west-2.amazonaws.com
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. only valid for PERSISTENT_1 and PERSISTENT_2 deployment_type. Requires automatic_backup_retention_days to be set.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// Sets the data compression configuration for the file system. Valid values are LZ4 and NONE. Default value is NONE. Unsetting this value reverts the compression type back to NONE.
	DataCompressionType *string `json:"dataCompressionType,omitempty" tf:"data_compression_type,omitempty"`

	// - The filesystem deployment type. One of: SCRATCH_1, SCRATCH_2, PERSISTENT_1, PERSISTENT_2.
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// - The type of drive cache used by PERSISTENT_1 filesystems that are provisioned with HDD storage_type. Required for HDD storage_type, set to either READ or NONE.
	DriveCacheType *string `json:"driveCacheType,omitempty" tf:"drive_cache_type,omitempty"`

	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with import_path argument and the path must use the same Amazon S3 bucket as specified in import_path. Set equal to import_path to overwrite files on export. Defaults to s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}. Only supported on PERSISTENT_1 deployment types.
	ExportPath *string `json:"exportPath,omitempty" tf:"export_path,omitempty"`

	// Sets the Lustre version for the file system that you're creating. Valid values are 2.10 for SCRATCH_1, SCRATCH_2 and PERSISTENT_1 deployment types. Valid values for 2.12 include all deployment types.
	FileSystemTypeVersion *string `json:"fileSystemTypeVersion,omitempty" tf:"file_system_type_version,omitempty"`

	// A map of tags to apply to the file system's final backup.
	// +mapType=granular
	FinalBackupTags map[string]*string `json:"finalBackupTags,omitempty" tf:"final_backup_tags,omitempty"`

	// Identifier of the file system, e.g., fs-12345678
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, s3://example-bucket/optional-prefix/. Only supported on PERSISTENT_1 deployment types.
	ImportPath *string `json:"importPath,omitempty" tf:"import_path,omitempty"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with import_path argument. Defaults to 1024. Minimum of 1 and maximum of 512000. Only supported on PERSISTENT_1 deployment types.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize,omitempty" tf:"imported_file_chunk_size,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, applicable for PERSISTENT_1 and PERSISTENT_2 deployment_type. Defaults to an AWS managed KMS Key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The Lustre logging configuration used when creating an Amazon FSx for Lustre file system. When logging is enabled, Lustre logs error and warning events for data repositories associated with your file system to Amazon CloudWatch Logs. See log_configuration Block for details.
	LogConfiguration *LogConfigurationObservation `json:"logConfiguration,omitempty" tf:"log_configuration,omitempty"`

	// The Lustre metadata configuration used when creating an Amazon FSx for Lustre file system. This can be used to specify a user provisioned metadata scale. This is only supported when deployment_type is set to PERSISTENT_2. See metadata_configuration Block for details.
	MetadataConfiguration *MetadataConfigurationObservation `json:"metadataConfiguration,omitempty" tf:"metadata_configuration,omitempty"`

	// The value to be used when mounting the filesystem.
	MountName *string `json:"mountName,omitempty" tf:"mount_name,omitempty"`

	// Set of Elastic Network Interface identifiers from which the file system is accessible. As explained in the documentation, the first network interface returned is the primary network interface.
	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	// AWS account identifier that created the file system.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the PERSISTENT_1 and PERSISTENT_2 deployment_type. Valid values for PERSISTENT_1 deployment_type and SSD storage_type are 50, 100, 200. Valid values for PERSISTENT_1 deployment_type and HDD storage_type are 12, 40. Valid values for PERSISTENT_2 deployment_type and  SSD storage_type are 125, 250, 500, 1000.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput,omitempty" tf:"per_unit_storage_throughput,omitempty"`

	// The Lustre root squash configuration used when creating an Amazon FSx for Lustre file system. When enabled, root squash restricts root-level access from clients that try to access your file system as a root user. See root_squash_configuration Block for details.
	RootSquashConfiguration *RootSquashConfigurationObservation `json:"rootSquashConfiguration,omitempty" tf:"root_squash_configuration,omitempty"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to true.
	SkipFinalBackup *bool `json:"skipFinalBackup,omitempty" tf:"skip_final_backup,omitempty"`

	// The storage capacity (GiB) of the file system. Minimum of 1200. See more details at Allowed values for Fsx storage capacity. Update is allowed only for SCRATCH_2, PERSISTENT_1 and PERSISTENT_2 deployment types, See more details at Fsx Storage Capacity Update. Required when not creating filesystem for a backup.
	StorageCapacity *float64 `json:"storageCapacity,omitempty" tf:"storage_capacity,omitempty"`

	// - The filesystem storage type. Either SSD or HDD, defaults to SSD. HDD is only supported on PERSISTENT_1 deployment types.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Identifier of the Virtual Private Cloud for the file system.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The preferred start time (in d:HH:MM format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

type LustreFileSystemParameters struct {

	// How Amazon FSx keeps your file and directory listings up to date as you add or modify objects in your linked S3 bucket. see Auto Import Data Repo for more details. Only supported on PERSISTENT_1 deployment types.
	// +kubebuilder:validation:Optional
	AutoImportPolicy *string `json:"autoImportPolicy,omitempty" tf:"auto_import_policy,omitempty"`

	// The number of days to retain automatic backups. Setting this to 0 disables automatic backups. You can retain automatic backups for a maximum of 90 days. only valid for PERSISTENT_1 and PERSISTENT_2 deployment_type.
	// +kubebuilder:validation:Optional
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// The ID of the source backup to create the filesystem from.
	// +kubebuilder:validation:Optional
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// A boolean flag indicating whether tags for the file system should be copied to backups. Applicable for PERSISTENT_1 and PERSISTENT_2 deployment_type. The default value is false.
	// +kubebuilder:validation:Optional
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups,omitempty"`

	// A recurring daily time, in the format HH:MM. HH is the zero-padded hour of the day (0-23), and MM is the zero-padded minute of the hour. For example, 05:00 specifies 5 AM daily. only valid for PERSISTENT_1 and PERSISTENT_2 deployment_type. Requires automatic_backup_retention_days to be set.
	// +kubebuilder:validation:Optional
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// Sets the data compression configuration for the file system. Valid values are LZ4 and NONE. Default value is NONE. Unsetting this value reverts the compression type back to NONE.
	// +kubebuilder:validation:Optional
	DataCompressionType *string `json:"dataCompressionType,omitempty" tf:"data_compression_type,omitempty"`

	// - The filesystem deployment type. One of: SCRATCH_1, SCRATCH_2, PERSISTENT_1, PERSISTENT_2.
	// +kubebuilder:validation:Optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// - The type of drive cache used by PERSISTENT_1 filesystems that are provisioned with HDD storage_type. Required for HDD storage_type, set to either READ or NONE.
	// +kubebuilder:validation:Optional
	DriveCacheType *string `json:"driveCacheType,omitempty" tf:"drive_cache_type,omitempty"`

	// S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with import_path argument and the path must use the same Amazon S3 bucket as specified in import_path. Set equal to import_path to overwrite files on export. Defaults to s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}. Only supported on PERSISTENT_1 deployment types.
	// +kubebuilder:validation:Optional
	ExportPath *string `json:"exportPath,omitempty" tf:"export_path,omitempty"`

	// Sets the Lustre version for the file system that you're creating. Valid values are 2.10 for SCRATCH_1, SCRATCH_2 and PERSISTENT_1 deployment types. Valid values for 2.12 include all deployment types.
	// +kubebuilder:validation:Optional
	FileSystemTypeVersion *string `json:"fileSystemTypeVersion,omitempty" tf:"file_system_type_version,omitempty"`

	// A map of tags to apply to the file system's final backup.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FinalBackupTags map[string]*string `json:"finalBackupTags,omitempty" tf:"final_backup_tags,omitempty"`

	// S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, s3://example-bucket/optional-prefix/. Only supported on PERSISTENT_1 deployment types.
	// +kubebuilder:validation:Optional
	ImportPath *string `json:"importPath,omitempty" tf:"import_path,omitempty"`

	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with import_path argument. Defaults to 1024. Minimum of 1 and maximum of 512000. Only supported on PERSISTENT_1 deployment types.
	// +kubebuilder:validation:Optional
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize,omitempty" tf:"imported_file_chunk_size,omitempty"`

	// ARN for the KMS Key to encrypt the file system at rest, applicable for PERSISTENT_1 and PERSISTENT_2 deployment_type. Defaults to an AWS managed KMS Key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The Lustre logging configuration used when creating an Amazon FSx for Lustre file system. When logging is enabled, Lustre logs error and warning events for data repositories associated with your file system to Amazon CloudWatch Logs. See log_configuration Block for details.
	// +kubebuilder:validation:Optional
	LogConfiguration *LogConfigurationParameters `json:"logConfiguration,omitempty" tf:"log_configuration,omitempty"`

	// The Lustre metadata configuration used when creating an Amazon FSx for Lustre file system. This can be used to specify a user provisioned metadata scale. This is only supported when deployment_type is set to PERSISTENT_2. See metadata_configuration Block for details.
	// +kubebuilder:validation:Optional
	MetadataConfiguration *MetadataConfigurationParameters `json:"metadataConfiguration,omitempty" tf:"metadata_configuration,omitempty"`

	// - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the PERSISTENT_1 and PERSISTENT_2 deployment_type. Valid values for PERSISTENT_1 deployment_type and SSD storage_type are 50, 100, 200. Valid values for PERSISTENT_1 deployment_type and HDD storage_type are 12, 40. Valid values for PERSISTENT_2 deployment_type and  SSD storage_type are 125, 250, 500, 1000.
	// +kubebuilder:validation:Optional
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput,omitempty" tf:"per_unit_storage_throughput,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Lustre root squash configuration used when creating an Amazon FSx for Lustre file system. When enabled, root squash restricts root-level access from clients that try to access your file system as a root user. See root_squash_configuration Block for details.
	// +kubebuilder:validation:Optional
	RootSquashConfiguration *RootSquashConfigurationParameters `json:"rootSquashConfiguration,omitempty" tf:"root_squash_configuration,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to true.
	// +kubebuilder:validation:Optional
	SkipFinalBackup *bool `json:"skipFinalBackup,omitempty" tf:"skip_final_backup,omitempty"`

	// The storage capacity (GiB) of the file system. Minimum of 1200. See more details at Allowed values for Fsx storage capacity. Update is allowed only for SCRATCH_2, PERSISTENT_1 and PERSISTENT_2 deployment types, See more details at Fsx Storage Capacity Update. Required when not creating filesystem for a backup.
	// +kubebuilder:validation:Optional
	StorageCapacity *float64 `json:"storageCapacity,omitempty" tf:"storage_capacity,omitempty"`

	// - The filesystem storage type. Either SSD or HDD, defaults to SSD. HDD is only supported on PERSISTENT_1 deployment types.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The preferred start time (in d:HH:MM format) to perform weekly maintenance, in the UTC time zone.
	// +kubebuilder:validation:Optional
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

type MetadataConfigurationInitParameters struct {

	// Amount of IOPS provisioned for metadata. This parameter should only be used when the mode is set to USER_PROVISIONED. Valid Values are 1500,3000,6000 and 12000 through 192000 in increments of 12000.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Mode for the metadata configuration of the file system. Valid values are AUTOMATIC, and USER_PROVISIONED.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type MetadataConfigurationObservation struct {

	// Amount of IOPS provisioned for metadata. This parameter should only be used when the mode is set to USER_PROVISIONED. Valid Values are 1500,3000,6000 and 12000 through 192000 in increments of 12000.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Mode for the metadata configuration of the file system. Valid values are AUTOMATIC, and USER_PROVISIONED.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type MetadataConfigurationParameters struct {

	// Amount of IOPS provisioned for metadata. This parameter should only be used when the mode is set to USER_PROVISIONED. Valid Values are 1500,3000,6000 and 12000 through 192000 in increments of 12000.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Mode for the metadata configuration of the file system. Valid values are AUTOMATIC, and USER_PROVISIONED.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type RootSquashConfigurationInitParameters struct {

	// When root squash is enabled, you can optionally specify an array of NIDs of clients for which root squash does not apply. A client NID is a Lustre Network Identifier used to uniquely identify a client. You can specify the NID as either a single address or a range of addresses: 1. A single address is described in standard Lustre NID format by specifying the client’s IP address followed by the Lustre network ID (for example, 10.0.1.6@tcp). 2. An address range is described using a dash to separate the range (for example, 10.0.[2-10].[1-255]@tcp).
	// +listType=set
	NoSquashNids []*string `json:"noSquashNids,omitempty" tf:"no_squash_nids,omitempty"`

	// You enable root squash by setting a user ID (UID) and group ID (GID) for the file system in the format UID:GID (for example, 365534:65534). The UID and GID values can range from 0 to 4294967294.
	RootSquash *string `json:"rootSquash,omitempty" tf:"root_squash,omitempty"`
}

type RootSquashConfigurationObservation struct {

	// When root squash is enabled, you can optionally specify an array of NIDs of clients for which root squash does not apply. A client NID is a Lustre Network Identifier used to uniquely identify a client. You can specify the NID as either a single address or a range of addresses: 1. A single address is described in standard Lustre NID format by specifying the client’s IP address followed by the Lustre network ID (for example, 10.0.1.6@tcp). 2. An address range is described using a dash to separate the range (for example, 10.0.[2-10].[1-255]@tcp).
	// +listType=set
	NoSquashNids []*string `json:"noSquashNids,omitempty" tf:"no_squash_nids,omitempty"`

	// You enable root squash by setting a user ID (UID) and group ID (GID) for the file system in the format UID:GID (for example, 365534:65534). The UID and GID values can range from 0 to 4294967294.
	RootSquash *string `json:"rootSquash,omitempty" tf:"root_squash,omitempty"`
}

type RootSquashConfigurationParameters struct {

	// When root squash is enabled, you can optionally specify an array of NIDs of clients for which root squash does not apply. A client NID is a Lustre Network Identifier used to uniquely identify a client. You can specify the NID as either a single address or a range of addresses: 1. A single address is described in standard Lustre NID format by specifying the client’s IP address followed by the Lustre network ID (for example, 10.0.1.6@tcp). 2. An address range is described using a dash to separate the range (for example, 10.0.[2-10].[1-255]@tcp).
	// +kubebuilder:validation:Optional
	// +listType=set
	NoSquashNids []*string `json:"noSquashNids,omitempty" tf:"no_squash_nids,omitempty"`

	// You enable root squash by setting a user ID (UID) and group ID (GID) for the file system in the format UID:GID (for example, 365534:65534). The UID and GID values can range from 0 to 4294967294.
	// +kubebuilder:validation:Optional
	RootSquash *string `json:"rootSquash,omitempty" tf:"root_squash,omitempty"`
}

// LustreFileSystemSpec defines the desired state of LustreFileSystem
type LustreFileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LustreFileSystemParameters `json:"forProvider"`
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
	InitProvider LustreFileSystemInitParameters `json:"initProvider,omitempty"`
}

// LustreFileSystemStatus defines the observed state of LustreFileSystem.
type LustreFileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LustreFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// LustreFileSystem is the Schema for the LustreFileSystems API. Manages a FSx Lustre File System.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LustreFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LustreFileSystemSpec   `json:"spec"`
	Status            LustreFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LustreFileSystemList contains a list of LustreFileSystems
type LustreFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LustreFileSystem `json:"items"`
}

// Repository type metadata.
var (
	LustreFileSystem_Kind             = "LustreFileSystem"
	LustreFileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LustreFileSystem_Kind}.String()
	LustreFileSystem_KindAPIVersion   = LustreFileSystem_Kind + "." + CRDGroupVersion.String()
	LustreFileSystem_GroupVersionKind = CRDGroupVersion.WithKind(LustreFileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&LustreFileSystem{}, &LustreFileSystemList{})
}
