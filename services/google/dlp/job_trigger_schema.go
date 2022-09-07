// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dlp

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLJobTriggerSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Dlp/JobTrigger",
			Description: "The Dlp JobTrigger resource",
			StructName:  "JobTrigger",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a JobTrigger",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "jobTrigger",
						Required:    true,
						Description: "A full instance of a JobTrigger",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a JobTrigger",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "jobTrigger",
						Required:    true,
						Description: "A full instance of a JobTrigger",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a JobTrigger",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "jobTrigger",
						Required:    true,
						Description: "A full instance of a JobTrigger",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all JobTrigger",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many JobTrigger",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"JobTrigger": &dcl.Component{
					Title: "JobTrigger",
					ID:    "{{parent}}/jobTriggers/{{name}}",
					Locations: []string{
						"region",
					},
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"inspectJob",
							"triggers",
							"status",
							"parent",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The creation timestamp of a triggeredJob.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "User provided description (max 256 chars)",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Display name (max 100 chars)",
							},
							"errors": &dcl.Property{
								Type:        "array",
								GoName:      "Errors",
								ReadOnly:    true,
								Description: "Output only. A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "JobTriggerErrors",
									Properties: map[string]*dcl.Property{
										"details": &dcl.Property{
											Type:        "object",
											GoName:      "Details",
											GoType:      "JobTriggerErrorsDetails",
											Description: "Detailed error codes and messages.",
											Properties: map[string]*dcl.Property{
												"code": &dcl.Property{
													Type:        "integer",
													Format:      "int64",
													GoName:      "Code",
													Description: "The status code, which should be an enum value of google.rpc.Code.",
												},
												"details": &dcl.Property{
													Type:        "array",
													GoName:      "Details",
													Description: "A list of messages that carry the error details. There is a common set of message types for APIs to use.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "JobTriggerErrorsDetailsDetails",
														Properties: map[string]*dcl.Property{
															"typeUrl": &dcl.Property{
																Type:        "string",
																GoName:      "TypeUrl",
																Description: "A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one \"/\" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading \".\" is not accepted). In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows: * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a google.protobuf.Type value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the URL, or have them precompiled into a binary to avoid any lookup. Therefore, binary compatibility needs to be preserved on changes to types. (Use versioned type names to manage breaking changes.) Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com. Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics.",
															},
															"value": &dcl.Property{
																Type:        "string",
																GoName:      "Value",
																Description: "Must be a valid serialized protocol buffer of the above specified type.",
															},
														},
													},
												},
												"message": &dcl.Property{
													Type:        "string",
													GoName:      "Message",
													Description: "A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.",
												},
											},
										},
										"timestamps": &dcl.Property{
											Type:        "array",
											GoName:      "Timestamps",
											Description: "The times the error occurred.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												Format: "date-time",
												GoType: "string",
											},
										},
									},
								},
							},
							"inspectJob": &dcl.Property{
								Type:        "object",
								GoName:      "InspectJob",
								GoType:      "JobTriggerInspectJob",
								Description: "For inspect jobs, a snapshot of the configuration.",
								Required: []string{
									"storageConfig",
								},
								Properties: map[string]*dcl.Property{
									"actions": &dcl.Property{
										Type:        "array",
										GoName:      "Actions",
										Description: "Actions to execute at the completion of the job.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "JobTriggerInspectJobActions",
											Properties: map[string]*dcl.Property{
												"jobNotificationEmails": &dcl.Property{
													Type:        "object",
													GoName:      "JobNotificationEmails",
													GoType:      "JobTriggerInspectJobActionsJobNotificationEmails",
													Description: "Enable email notification for project owners and editors on job's completion/failure.",
													Conflicts: []string{
														"saveFindings",
														"pubSub",
														"publishSummaryToCscc",
														"publishFindingsToCloudDataCatalog",
														"publishToStackdriver",
													},
													SendEmpty:  true,
													Properties: map[string]*dcl.Property{},
												},
												"pubSub": &dcl.Property{
													Type:        "object",
													GoName:      "PubSub",
													GoType:      "JobTriggerInspectJobActionsPubSub",
													Description: "Publish a notification to a pubsub topic.",
													Conflicts: []string{
														"saveFindings",
														"publishSummaryToCscc",
														"publishFindingsToCloudDataCatalog",
														"jobNotificationEmails",
														"publishToStackdriver",
													},
													Properties: map[string]*dcl.Property{
														"topic": &dcl.Property{
															Type:        "string",
															GoName:      "Topic",
															Description: "Cloud Pub/Sub topic to send notifications to. The topic must have given publishing access rights to the DLP API service account executing the long running DlpJob sending the notifications. Format is projects/{project}/topics/{topic}.",
															ResourceReferences: []*dcl.PropertyResourceReference{
																&dcl.PropertyResourceReference{
																	Resource: "Pubsub/Topic",
																	Field:    "name",
																},
															},
														},
													},
												},
												"publishFindingsToCloudDataCatalog": &dcl.Property{
													Type:        "object",
													GoName:      "PublishFindingsToCloudDataCatalog",
													GoType:      "JobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog",
													Description: "Publish findings to Cloud Datahub.",
													Conflicts: []string{
														"saveFindings",
														"pubSub",
														"publishSummaryToCscc",
														"jobNotificationEmails",
														"publishToStackdriver",
													},
													SendEmpty:  true,
													Properties: map[string]*dcl.Property{},
												},
												"publishSummaryToCscc": &dcl.Property{
													Type:        "object",
													GoName:      "PublishSummaryToCscc",
													GoType:      "JobTriggerInspectJobActionsPublishSummaryToCscc",
													Description: "Publish summary to Cloud Security Command Center (Alpha).",
													Conflicts: []string{
														"saveFindings",
														"pubSub",
														"publishFindingsToCloudDataCatalog",
														"jobNotificationEmails",
														"publishToStackdriver",
													},
													SendEmpty:  true,
													Properties: map[string]*dcl.Property{},
												},
												"publishToStackdriver": &dcl.Property{
													Type:        "object",
													GoName:      "PublishToStackdriver",
													GoType:      "JobTriggerInspectJobActionsPublishToStackdriver",
													Description: "Enable Stackdriver metric dlp.googleapis.com/finding_count.",
													Conflicts: []string{
														"saveFindings",
														"pubSub",
														"publishSummaryToCscc",
														"publishFindingsToCloudDataCatalog",
														"jobNotificationEmails",
													},
													SendEmpty:  true,
													Properties: map[string]*dcl.Property{},
												},
												"saveFindings": &dcl.Property{
													Type:        "object",
													GoName:      "SaveFindings",
													GoType:      "JobTriggerInspectJobActionsSaveFindings",
													Description: "Save resulting findings in a provided location.",
													Conflicts: []string{
														"pubSub",
														"publishSummaryToCscc",
														"publishFindingsToCloudDataCatalog",
														"jobNotificationEmails",
														"publishToStackdriver",
													},
													Properties: map[string]*dcl.Property{
														"outputConfig": &dcl.Property{
															Type:        "object",
															GoName:      "OutputConfig",
															GoType:      "JobTriggerInspectJobActionsSaveFindingsOutputConfig",
															Description: "Location to store findings outside of DLP.",
															Properties: map[string]*dcl.Property{
																"dlpStorage": &dcl.Property{
																	Type:        "object",
																	GoName:      "DlpStorage",
																	GoType:      "JobTriggerInspectJobActionsSaveFindingsOutputConfigDlpStorage",
																	Description: "Store findings directly to DLP. If neither this or bigquery is chosen only summary stats of total infotype count will be stored. Quotes will not be stored to dlp findings. If quotes are needed, store to BigQuery. Currently only for inspect jobs.",
																	Conflicts: []string{
																		"table",
																	},
																	SendEmpty:  true,
																	Properties: map[string]*dcl.Property{},
																},
																"outputSchema": &dcl.Property{
																	Type:        "string",
																	GoName:      "OutputSchema",
																	GoType:      "JobTriggerInspectJobActionsSaveFindingsOutputConfigOutputSchemaEnum",
																	Description: "Schema used for writing the findings for Inspect jobs. This field is only used for Inspect and must be unspecified for Risk jobs. Columns are derived from the `Finding` object. If appending to an existing table, any columns from the predefined schema that are missing will be added. No columns in the existing table will be deleted. If unspecified, then all available columns will be used for a new table or an (existing) table with no schema, and no changes will be made to an existing table that has a schema. Only for use with external storage. Possible values: OUTPUT_SCHEMA_UNSPECIFIED, BASIC_COLUMNS, GCS_COLUMNS, DATASTORE_COLUMNS, BIG_QUERY_COLUMNS, ALL_COLUMNS",
																	Enum: []string{
																		"OUTPUT_SCHEMA_UNSPECIFIED",
																		"BASIC_COLUMNS",
																		"GCS_COLUMNS",
																		"DATASTORE_COLUMNS",
																		"BIG_QUERY_COLUMNS",
																		"ALL_COLUMNS",
																	},
																},
																"table": &dcl.Property{
																	Type:        "object",
																	GoName:      "Table",
																	GoType:      "JobTriggerInspectJobActionsSaveFindingsOutputConfigTable",
																	Description: "Store findings in an existing table or a new table in an existing dataset. If table_id is not set a new one will be generated for you with the following format: dlp_googleapis_yyyy_mm_dd_[dlp_job_id]. Pacific timezone will be used for generating the date details. For Inspect, each column in an existing output table must have the same name, type, and mode of a field in the `Finding` object. For Risk, an existing output table should be the output of a previous Risk analysis job run on the same source table, with the same privacy metric and quasi-identifiers. Risk jobs that analyze the same table but compute a different privacy metric, or use different sets of quasi-identifiers, cannot store their results in the same table.",
																	Conflicts: []string{
																		"dlpStorage",
																	},
																	Properties: map[string]*dcl.Property{
																		"datasetId": &dcl.Property{
																			Type:        "string",
																			GoName:      "DatasetId",
																			Description: "Dataset ID of the table.",
																			ResourceReferences: []*dcl.PropertyResourceReference{
																				&dcl.PropertyResourceReference{
																					Resource: "Bigquery/Dataset",
																					Field:    "name",
																				},
																			},
																		},
																		"projectId": &dcl.Property{
																			Type:        "string",
																			GoName:      "ProjectId",
																			Description: "The Google Cloud Platform project ID of the project containing the table. If omitted, project ID is inferred from the API call.",
																			ResourceReferences: []*dcl.PropertyResourceReference{
																				&dcl.PropertyResourceReference{
																					Resource: "Cloudresourcemanager/Project",
																					Field:    "name",
																				},
																			},
																		},
																		"tableId": &dcl.Property{
																			Type:        "string",
																			GoName:      "TableId",
																			Description: "Name of the table.",
																			ResourceReferences: []*dcl.PropertyResourceReference{
																				&dcl.PropertyResourceReference{
																					Resource: "Bigquery/Table",
																					Field:    "name",
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"inspectConfig": &dcl.Property{
										Type:        "object",
										GoName:      "InspectConfig",
										GoType:      "JobTriggerInspectJobInspectConfig",
										Description: "How and what to scan for.",
										Properties: map[string]*dcl.Property{
											"customInfoTypes": &dcl.Property{
												Type:        "array",
												GoName:      "CustomInfoTypes",
												Description: "CustomInfoTypes provided by the user. See https://cloud.google.com/dlp/docs/creating-custom-infotypes to learn more.",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "JobTriggerInspectJobInspectConfigCustomInfoTypes",
													Properties: map[string]*dcl.Property{
														"detectionRules": &dcl.Property{
															Type:        "array",
															GoName:      "DetectionRules",
															Description: "Set of detection rules to apply to all findings of this CustomInfoType. Rules are applied in order that they are specified. Not supported for the `surrogate_type` CustomInfoType.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "object",
																GoType: "JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRules",
																Properties: map[string]*dcl.Property{
																	"hotwordRule": &dcl.Property{
																		Type:        "object",
																		GoName:      "HotwordRule",
																		GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRule",
																		Description: "Hotword-based detection rule.",
																		Properties: map[string]*dcl.Property{
																			"hotwordRegex": &dcl.Property{
																				Type:        "object",
																				GoName:      "HotwordRegex",
																				GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleHotwordRegex",
																				Description: "Regular expression pattern defining what qualifies as a hotword.",
																				Properties: map[string]*dcl.Property{
																					"groupIndexes": &dcl.Property{
																						Type:        "array",
																						GoName:      "GroupIndexes",
																						Description: "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
																						SendEmpty:   true,
																						ListType:    "list",
																						Items: &dcl.Property{
																							Type:   "integer",
																							Format: "int64",
																							GoType: "int64",
																						},
																					},
																					"pattern": &dcl.Property{
																						Type:        "string",
																						GoName:      "Pattern",
																						Description: "Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
																					},
																				},
																			},
																			"likelihoodAdjustment": &dcl.Property{
																				Type:        "object",
																				GoName:      "LikelihoodAdjustment",
																				GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustment",
																				Description: "Likelihood adjustment to apply to all matching findings.",
																				Properties: map[string]*dcl.Property{
																					"fixedLikelihood": &dcl.Property{
																						Type:        "string",
																						GoName:      "FixedLikelihood",
																						GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum",
																						Description: "Set the likelihood of a finding to a fixed value. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY",
																						Conflicts: []string{
																							"relativeLikelihood",
																						},
																						Enum: []string{
																							"LIKELIHOOD_UNSPECIFIED",
																							"VERY_UNLIKELY",
																							"UNLIKELY",
																							"POSSIBLE",
																							"LIKELY",
																							"VERY_LIKELY",
																						},
																					},
																					"relativeLikelihood": &dcl.Property{
																						Type:        "integer",
																						Format:      "int64",
																						GoName:      "RelativeLikelihood",
																						Description: "Increase or decrease the likelihood by the specified number of levels. For example, if a finding would be `POSSIBLE` without the detection rule and `relative_likelihood` is 1, then it is upgraded to `LIKELY`, while a value of -1 would downgrade it to `UNLIKELY`. Likelihood may never drop below `VERY_UNLIKELY` or exceed `VERY_LIKELY`, so applying an adjustment of 1 followed by an adjustment of -1 when base likelihood is `VERY_LIKELY` will result in a final likelihood of `LIKELY`.",
																						Conflicts: []string{
																							"fixedLikelihood",
																						},
																					},
																				},
																			},
																			"proximity": &dcl.Property{
																				Type:        "object",
																				GoName:      "Proximity",
																				GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesDetectionRulesHotwordRuleProximity",
																				Description: "Proximity of the finding within which the entire hotword must reside. The total length of the window cannot exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be used to match substrings of the finding itself. For example, the certainty of a phone number regex \"(d{3}) d{3}-d{4}\" could be adjusted upwards if the area code is known to be the local area code of a company office using the hotword regex \"(xxx)\", where \"xxx\" is the area code in question.",
																				Properties: map[string]*dcl.Property{
																					"windowAfter": &dcl.Property{
																						Type:        "integer",
																						Format:      "int64",
																						GoName:      "WindowAfter",
																						Description: "Number of characters after the finding to consider.",
																					},
																					"windowBefore": &dcl.Property{
																						Type:        "integer",
																						Format:      "int64",
																						GoName:      "WindowBefore",
																						Description: "Number of characters before the finding to consider.",
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
														"dictionary": &dcl.Property{
															Type:        "object",
															GoName:      "Dictionary",
															GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesDictionary",
															Description: "A list of phrases to detect as a CustomInfoType.",
															Conflicts: []string{
																"regex",
																"surrogateType",
																"storedType",
															},
															Properties: map[string]*dcl.Property{
																"cloudStoragePath": &dcl.Property{
																	Type:        "object",
																	GoName:      "CloudStoragePath",
																	GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath",
																	Description: "Newline-delimited file of words in Cloud Storage. Only a single file is accepted.",
																	Conflicts: []string{
																		"wordList",
																	},
																	Properties: map[string]*dcl.Property{
																		"path": &dcl.Property{
																			Type:        "string",
																			GoName:      "Path",
																			Description: "A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt",
																		},
																	},
																},
																"wordList": &dcl.Property{
																	Type:        "object",
																	GoName:      "WordList",
																	GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryWordList",
																	Description: "List of words or phrases to search for.",
																	Conflicts: []string{
																		"cloudStoragePath",
																	},
																	Properties: map[string]*dcl.Property{
																		"words": &dcl.Property{
																			Type:        "array",
																			GoName:      "Words",
																			Description: "Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits. [required]",
																			SendEmpty:   true,
																			ListType:    "list",
																			Items: &dcl.Property{
																				Type:   "string",
																				GoType: "string",
																			},
																		},
																	},
																},
															},
														},
														"exclusionType": &dcl.Property{
															Type:        "string",
															GoName:      "ExclusionType",
															GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesExclusionTypeEnum",
															Description: "If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned. It still can be used for rules matching. Possible values: EXCLUSION_TYPE_UNSPECIFIED, EXCLUSION_TYPE_EXCLUDE",
															Enum: []string{
																"EXCLUSION_TYPE_UNSPECIFIED",
																"EXCLUSION_TYPE_EXCLUDE",
															},
														},
														"infoType": &dcl.Property{
															Type:        "object",
															GoName:      "InfoType",
															GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesInfoType",
															Description: "CustomInfoType can either be a new infoType, or an extension of built-in infoType, when the name matches one of existing infoTypes and that infoType is specified in `InspectContent.info_types` field. Specifying the latter adds findings to the one detected by the system. If built-in info type is not specified in `InspectContent.info_types` list then the name is treated as a custom info type.",
															Properties: map[string]*dcl.Property{
																"name": &dcl.Property{
																	Type:        "string",
																	GoName:      "Name",
																	Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
																},
																"version": &dcl.Property{
																	Type:        "string",
																	GoName:      "Version",
																	Description: "Optional version name for this InfoType.",
																},
															},
														},
														"likelihood": &dcl.Property{
															Type:        "string",
															GoName:      "Likelihood",
															GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesLikelihoodEnum",
															Description: "Likelihood to return for this CustomInfoType. This base value can be altered by a detection rule if the finding meets the criteria specified by the rule. Defaults to `VERY_LIKELY` if not specified. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY",
															Enum: []string{
																"LIKELIHOOD_UNSPECIFIED",
																"VERY_UNLIKELY",
																"UNLIKELY",
																"POSSIBLE",
																"LIKELY",
																"VERY_LIKELY",
															},
														},
														"regex": &dcl.Property{
															Type:        "object",
															GoName:      "Regex",
															GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesRegex",
															Description: "Regular expression based CustomInfoType.",
															Conflicts: []string{
																"dictionary",
																"surrogateType",
																"storedType",
															},
															Properties: map[string]*dcl.Property{
																"groupIndexes": &dcl.Property{
																	Type:        "array",
																	GoName:      "GroupIndexes",
																	Description: "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
																	SendEmpty:   true,
																	ListType:    "list",
																	Items: &dcl.Property{
																		Type:   "integer",
																		Format: "int64",
																		GoType: "int64",
																	},
																},
																"pattern": &dcl.Property{
																	Type:        "string",
																	GoName:      "Pattern",
																	Description: "Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
																},
															},
														},
														"storedType": &dcl.Property{
															Type:        "object",
															GoName:      "StoredType",
															GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesStoredType",
															Description: "Load an existing `StoredInfoType` resource for use in `InspectDataSource`. Not currently supported in `InspectContent`.",
															Conflicts: []string{
																"dictionary",
																"regex",
																"surrogateType",
															},
															Properties: map[string]*dcl.Property{
																"createTime": &dcl.Property{
																	Type:        "string",
																	Format:      "date-time",
																	GoName:      "CreateTime",
																	ReadOnly:    true,
																	Description: "Timestamp indicating when the version of the `StoredInfoType` used for inspection was created. Output-only field, populated by the system.",
																},
																"name": &dcl.Property{
																	Type:        "string",
																	GoName:      "Name",
																	Description: "Resource name of the requested `StoredInfoType`, for example `organizations/433245324/storedInfoTypes/432452342` or `projects/project-id/storedInfoTypes/432452342`.",
																	ResourceReferences: []*dcl.PropertyResourceReference{
																		&dcl.PropertyResourceReference{
																			Resource: "Dlp/StoredInfoType",
																			Field:    "name",
																			Parent:   true,
																		},
																	},
																},
															},
														},
														"surrogateType": &dcl.Property{
															Type:        "object",
															GoName:      "SurrogateType",
															GoType:      "JobTriggerInspectJobInspectConfigCustomInfoTypesSurrogateType",
															Description: "Message for detecting output from deidentification transformations that support reversing.",
															Conflicts: []string{
																"dictionary",
																"regex",
																"storedType",
															},
															SendEmpty:  true,
															Properties: map[string]*dcl.Property{},
														},
													},
												},
											},
											"excludeInfoTypes": &dcl.Property{
												Type:        "boolean",
												GoName:      "ExcludeInfoTypes",
												Description: "When true, excludes type information of the findings. This is not used for data profiling.",
											},
											"includeQuote": &dcl.Property{
												Type:        "boolean",
												GoName:      "IncludeQuote",
												Description: "When true, a contextual quote from the data that triggered a finding is included in the response; see Finding.quote. This is not used for data profiling.",
											},
											"infoTypes": &dcl.Property{
												Type:        "array",
												GoName:      "InfoTypes",
												Description: "Restricts what info_types to look for. The values must correspond to InfoType values returned by ListInfoTypes or listed at https://cloud.google.com/dlp/docs/infotypes-reference. When no InfoTypes or CustomInfoTypes are specified in a request, the system may automatically choose what detectors to run. By default this may be all types, but may change over time as detectors are updated. If you need precise control and predictability as to what detectors are run you should specify specific InfoTypes listed in the reference, otherwise a default list will be used, which may change over time.",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "JobTriggerInspectJobInspectConfigInfoTypes",
													Properties: map[string]*dcl.Property{
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
														},
													},
												},
											},
											"limits": &dcl.Property{
												Type:        "object",
												GoName:      "Limits",
												GoType:      "JobTriggerInspectJobInspectConfigLimits",
												Description: "Configuration to control the number of findings returned. This is not used for data profiling.",
												Properties: map[string]*dcl.Property{
													"maxFindingsPerInfoType": &dcl.Property{
														Type:        "array",
														GoName:      "MaxFindingsPerInfoType",
														Description: "Configuration of findings limit given for specified infoTypes.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "object",
															GoType: "JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType",
															Properties: map[string]*dcl.Property{
																"infoType": &dcl.Property{
																	Type:        "object",
																	GoName:      "InfoType",
																	GoType:      "JobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType",
																	Description: "Type of information the findings limit applies to. Only one limit per info_type should be provided. If InfoTypeLimit does not have an info_type, the DLP API applies the limit against all info_types that are found but not specified in another InfoTypeLimit.",
																	Properties: map[string]*dcl.Property{
																		"name": &dcl.Property{
																			Type:        "string",
																			GoName:      "Name",
																			Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
																		},
																		"version": &dcl.Property{
																			Type:        "string",
																			GoName:      "Version",
																			Description: "Optional version name for this InfoType.",
																		},
																	},
																},
																"maxFindings": &dcl.Property{
																	Type:        "integer",
																	Format:      "int64",
																	GoName:      "MaxFindings",
																	Description: "Max findings limit for the given infoType.",
																},
															},
														},
													},
													"maxFindingsPerItem": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "MaxFindingsPerItem",
														Description: "Max number of findings that will be returned for each item scanned. When set within `InspectJobConfig`, the maximum returned is 2000 regardless if this is set higher. When set within `InspectContentRequest`, this field is ignored.",
													},
													"maxFindingsPerRequest": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "MaxFindingsPerRequest",
														Description: "Max number of findings that will be returned per request/job. When set within `InspectContentRequest`, the maximum returned is 2000 regardless if this is set higher.",
													},
												},
											},
											"minLikelihood": &dcl.Property{
												Type:        "string",
												GoName:      "MinLikelihood",
												GoType:      "JobTriggerInspectJobInspectConfigMinLikelihoodEnum",
												Description: "Only returns findings equal or above this threshold. The default is POSSIBLE. See https://cloud.google.com/dlp/docs/likelihood to learn more. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY",
												Enum: []string{
													"LIKELIHOOD_UNSPECIFIED",
													"VERY_UNLIKELY",
													"UNLIKELY",
													"POSSIBLE",
													"LIKELY",
													"VERY_LIKELY",
												},
											},
											"ruleSet": &dcl.Property{
												Type:        "array",
												GoName:      "RuleSet",
												Description: "Set of rules to apply to the findings for this InspectConfig. Exclusion rules, contained in the set are executed in the end, other rules are executed in the order they are specified for each info type.",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "JobTriggerInspectJobInspectConfigRuleSet",
													Properties: map[string]*dcl.Property{
														"infoTypes": &dcl.Property{
															Type:        "array",
															GoName:      "InfoTypes",
															Description: "List of infoTypes this rule set is applied to.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "object",
																GoType: "JobTriggerInspectJobInspectConfigRuleSetInfoTypes",
																Properties: map[string]*dcl.Property{
																	"name": &dcl.Property{
																		Type:        "string",
																		GoName:      "Name",
																		Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
																	},
																	"version": &dcl.Property{
																		Type:        "string",
																		GoName:      "Version",
																		Description: "Optional version name for this InfoType.",
																	},
																},
															},
														},
														"rules": &dcl.Property{
															Type:        "array",
															GoName:      "Rules",
															Description: "Set of rules to be applied to infoTypes. The rules are applied in order.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "object",
																GoType: "JobTriggerInspectJobInspectConfigRuleSetRules",
																Properties: map[string]*dcl.Property{
																	"exclusionRule": &dcl.Property{
																		Type:        "object",
																		GoName:      "ExclusionRule",
																		GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule",
																		Description: "Exclusion rule.",
																		Conflicts: []string{
																			"hotwordRule",
																		},
																		Properties: map[string]*dcl.Property{
																			"dictionary": &dcl.Property{
																				Type:        "object",
																				GoName:      "Dictionary",
																				GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary",
																				Description: "Dictionary which defines the rule.",
																				Conflicts: []string{
																					"regex",
																					"excludeInfoTypes",
																				},
																				Properties: map[string]*dcl.Property{
																					"cloudStoragePath": &dcl.Property{
																						Type:        "object",
																						GoName:      "CloudStoragePath",
																						GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath",
																						Description: "Newline-delimited file of words in Cloud Storage. Only a single file is accepted.",
																						Conflicts: []string{
																							"wordList",
																						},
																						Properties: map[string]*dcl.Property{
																							"path": &dcl.Property{
																								Type:        "string",
																								GoName:      "Path",
																								Description: "A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt",
																							},
																						},
																					},
																					"wordList": &dcl.Property{
																						Type:        "object",
																						GoName:      "WordList",
																						GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordList",
																						Description: "List of words or phrases to search for.",
																						Conflicts: []string{
																							"cloudStoragePath",
																						},
																						Properties: map[string]*dcl.Property{
																							"words": &dcl.Property{
																								Type:        "array",
																								GoName:      "Words",
																								Description: "Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits. [required]",
																								SendEmpty:   true,
																								ListType:    "list",
																								Items: &dcl.Property{
																									Type:   "string",
																									GoType: "string",
																								},
																							},
																						},
																					},
																				},
																			},
																			"excludeInfoTypes": &dcl.Property{
																				Type:        "object",
																				GoName:      "ExcludeInfoTypes",
																				GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes",
																				Description: "Set of infoTypes for which findings would affect this rule.",
																				Conflicts: []string{
																					"dictionary",
																					"regex",
																				},
																				Properties: map[string]*dcl.Property{
																					"infoTypes": &dcl.Property{
																						Type:        "array",
																						GoName:      "InfoTypes",
																						Description: "InfoType list in ExclusionRule rule drops a finding when it overlaps or contained within with a finding of an infoType from this list. For example, for `InspectionRuleSet.info_types` containing \"PHONE_NUMBER\"` and `exclusion_rule` containing `exclude_info_types.info_types` with \"EMAIL_ADDRESS\" the phone number findings are dropped if they overlap with EMAIL_ADDRESS finding. That leads to \"555-222-2222@example.org\" to generate only a single finding, namely email address.",
																						SendEmpty:   true,
																						ListType:    "list",
																						Items: &dcl.Property{
																							Type:   "object",
																							GoType: "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes",
																							Properties: map[string]*dcl.Property{
																								"name": &dcl.Property{
																									Type:        "string",
																									GoName:      "Name",
																									Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
																								},
																								"version": &dcl.Property{
																									Type:        "string",
																									GoName:      "Version",
																									Description: "Optional version name for this InfoType.",
																								},
																							},
																						},
																					},
																				},
																			},
																			"matchingType": &dcl.Property{
																				Type:        "string",
																				GoName:      "MatchingType",
																				GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum",
																				Description: "How the rule is applied, see MatchingType documentation for details. Possible values: MATCHING_TYPE_UNSPECIFIED, MATCHING_TYPE_FULL_MATCH, MATCHING_TYPE_PARTIAL_MATCH, MATCHING_TYPE_INVERSE_MATCH",
																				Enum: []string{
																					"MATCHING_TYPE_UNSPECIFIED",
																					"MATCHING_TYPE_FULL_MATCH",
																					"MATCHING_TYPE_PARTIAL_MATCH",
																					"MATCHING_TYPE_INVERSE_MATCH",
																				},
																			},
																			"regex": &dcl.Property{
																				Type:        "object",
																				GoName:      "Regex",
																				GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex",
																				Description: "Regular expression which defines the rule.",
																				Conflicts: []string{
																					"dictionary",
																					"excludeInfoTypes",
																				},
																				Properties: map[string]*dcl.Property{
																					"groupIndexes": &dcl.Property{
																						Type:        "array",
																						GoName:      "GroupIndexes",
																						Description: "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
																						SendEmpty:   true,
																						ListType:    "list",
																						Items: &dcl.Property{
																							Type:   "integer",
																							Format: "int64",
																							GoType: "int64",
																						},
																					},
																					"pattern": &dcl.Property{
																						Type:        "string",
																						GoName:      "Pattern",
																						Description: "Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
																					},
																				},
																			},
																		},
																	},
																	"hotwordRule": &dcl.Property{
																		Type:   "object",
																		GoName: "HotwordRule",
																		GoType: "JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule",
																		Conflicts: []string{
																			"exclusionRule",
																		},
																		Properties: map[string]*dcl.Property{
																			"hotwordRegex": &dcl.Property{
																				Type:        "object",
																				GoName:      "HotwordRegex",
																				GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex",
																				Description: "Regular expression pattern defining what qualifies as a hotword.",
																				Properties: map[string]*dcl.Property{
																					"groupIndexes": &dcl.Property{
																						Type:        "array",
																						GoName:      "GroupIndexes",
																						Description: "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
																						SendEmpty:   true,
																						ListType:    "list",
																						Items: &dcl.Property{
																							Type:   "integer",
																							Format: "int64",
																							GoType: "int64",
																						},
																					},
																					"pattern": &dcl.Property{
																						Type:        "string",
																						GoName:      "Pattern",
																						Description: "Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
																					},
																				},
																			},
																			"likelihoodAdjustment": &dcl.Property{
																				Type:        "object",
																				GoName:      "LikelihoodAdjustment",
																				GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment",
																				Description: "Likelihood adjustment to apply to all matching findings.",
																				Properties: map[string]*dcl.Property{
																					"fixedLikelihood": &dcl.Property{
																						Type:        "string",
																						GoName:      "FixedLikelihood",
																						GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum",
																						Description: "Set the likelihood of a finding to a fixed value. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY",
																						Conflicts: []string{
																							"relativeLikelihood",
																						},
																						Enum: []string{
																							"LIKELIHOOD_UNSPECIFIED",
																							"VERY_UNLIKELY",
																							"UNLIKELY",
																							"POSSIBLE",
																							"LIKELY",
																							"VERY_LIKELY",
																						},
																					},
																					"relativeLikelihood": &dcl.Property{
																						Type:        "integer",
																						Format:      "int64",
																						GoName:      "RelativeLikelihood",
																						Description: "Increase or decrease the likelihood by the specified number of levels. For example, if a finding would be `POSSIBLE` without the detection rule and `relative_likelihood` is 1, then it is upgraded to `LIKELY`, while a value of -1 would downgrade it to `UNLIKELY`. Likelihood may never drop below `VERY_UNLIKELY` or exceed `VERY_LIKELY`, so applying an adjustment of 1 followed by an adjustment of -1 when base likelihood is `VERY_LIKELY` will result in a final likelihood of `LIKELY`.",
																						Conflicts: []string{
																							"fixedLikelihood",
																						},
																					},
																				},
																			},
																			"proximity": &dcl.Property{
																				Type:        "object",
																				GoName:      "Proximity",
																				GoType:      "JobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity",
																				Description: "Proximity of the finding within which the entire hotword must reside. The total length of the window cannot exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be used to match substrings of the finding itself. For example, the certainty of a phone number regex \"(d{3}) d{3}-d{4}\" could be adjusted upwards if the area code is known to be the local area code of a company office using the hotword regex \"(xxx)\", where \"xxx\" is the area code in question.",
																				Properties: map[string]*dcl.Property{
																					"windowAfter": &dcl.Property{
																						Type:        "integer",
																						Format:      "int64",
																						GoName:      "WindowAfter",
																						Description: "Number of characters after the finding to consider.",
																					},
																					"windowBefore": &dcl.Property{
																						Type:        "integer",
																						Format:      "int64",
																						GoName:      "WindowBefore",
																						Description: "Number of characters before the finding to consider.",
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"inspectTemplateName": &dcl.Property{
										Type:        "string",
										GoName:      "InspectTemplateName",
										Description: "If provided, will be used as the default for all values in InspectConfig. `inspect_config` will be merged into the values persisted as part of the template.",
									},
									"storageConfig": &dcl.Property{
										Type:        "object",
										GoName:      "StorageConfig",
										GoType:      "JobTriggerInspectJobStorageConfig",
										Description: "The data to scan.",
										Properties: map[string]*dcl.Property{
											"bigQueryOptions": &dcl.Property{
												Type:        "object",
												GoName:      "BigQueryOptions",
												GoType:      "JobTriggerInspectJobStorageConfigBigQueryOptions",
												Description: "BigQuery options.",
												Conflicts: []string{
													"datastoreOptions",
													"cloudStorageOptions",
													"hybridOptions",
												},
												Required: []string{
													"tableReference",
												},
												Properties: map[string]*dcl.Property{
													"excludedFields": &dcl.Property{
														Type:        "array",
														GoName:      "ExcludedFields",
														Description: "References to fields excluded from scanning. This allows you to skip inspection of entire columns which you know have no findings.",
														Conflicts: []string{
															"includedFields",
														},
														SendEmpty: true,
														ListType:  "list",
														Items: &dcl.Property{
															Type:   "object",
															GoType: "JobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFields",
															Properties: map[string]*dcl.Property{
																"name": &dcl.Property{
																	Type:        "string",
																	GoName:      "Name",
																	Description: "Name describing the field.",
																},
															},
														},
													},
													"identifyingFields": &dcl.Property{
														Type:        "array",
														GoName:      "IdentifyingFields",
														Description: "Table fields that may uniquely identify a row within the table. When `actions.saveFindings.outputConfig.table` is specified, the values of columns specified here are available in the output table under `location.content_locations.record_location.record_key.id_values`. Nested fields such as `person.birthdate.year` are allowed.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "object",
															GoType: "JobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFields",
															Properties: map[string]*dcl.Property{
																"name": &dcl.Property{
																	Type:        "string",
																	GoName:      "Name",
																	Description: "Name describing the field.",
																},
															},
														},
													},
													"includedFields": &dcl.Property{
														Type:        "array",
														GoName:      "IncludedFields",
														Description: "Limit scanning only to these fields.",
														Conflicts: []string{
															"excludedFields",
														},
														SendEmpty: true,
														ListType:  "list",
														Items: &dcl.Property{
															Type:   "object",
															GoType: "JobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFields",
															Properties: map[string]*dcl.Property{
																"name": &dcl.Property{
																	Type:        "string",
																	GoName:      "Name",
																	Description: "Name describing the field.",
																},
															},
														},
													},
													"rowsLimit": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "RowsLimit",
														Description: "Max number of rows to scan. If the table has more rows than this value, the rest of the rows are omitted. If not set, or if set to 0, all rows will be scanned. Only one of rows_limit and rows_limit_percent can be specified. Cannot be used in conjunction with TimespanConfig.",
														Conflicts: []string{
															"rowsLimitPercent",
														},
													},
													"rowsLimitPercent": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "RowsLimitPercent",
														Description: "Max percentage of rows to scan. The rest are omitted. The number of rows scanned is rounded down. Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0. Only one of rows_limit and rows_limit_percent can be specified. Cannot be used in conjunction with TimespanConfig.",
														Conflicts: []string{
															"rowsLimit",
														},
													},
													"sampleMethod": &dcl.Property{
														Type:        "string",
														GoName:      "SampleMethod",
														GoType:      "JobTriggerInspectJobStorageConfigBigQueryOptionsSampleMethodEnum",
														Description: " Possible values: SAMPLE_METHOD_UNSPECIFIED, TOP, RANDOM_START",
														Enum: []string{
															"SAMPLE_METHOD_UNSPECIFIED",
															"TOP",
															"RANDOM_START",
														},
													},
													"tableReference": &dcl.Property{
														Type:        "object",
														GoName:      "TableReference",
														GoType:      "JobTriggerInspectJobStorageConfigBigQueryOptionsTableReference",
														Description: "Complete BigQuery table reference.",
														Properties: map[string]*dcl.Property{
															"datasetId": &dcl.Property{
																Type:        "string",
																GoName:      "DatasetId",
																Description: "Dataset ID of the table.",
																ResourceReferences: []*dcl.PropertyResourceReference{
																	&dcl.PropertyResourceReference{
																		Resource: "Bigquery/Dataset",
																		Field:    "name",
																	},
																},
															},
															"projectId": &dcl.Property{
																Type:        "string",
																GoName:      "ProjectId",
																Description: "The Google Cloud Platform project ID of the project containing the table. If omitted, project ID is inferred from the API call.",
																ResourceReferences: []*dcl.PropertyResourceReference{
																	&dcl.PropertyResourceReference{
																		Resource: "Cloudresourcemanager/Project",
																		Field:    "name",
																	},
																},
															},
															"tableId": &dcl.Property{
																Type:        "string",
																GoName:      "TableId",
																Description: "Name of the table.",
																ResourceReferences: []*dcl.PropertyResourceReference{
																	&dcl.PropertyResourceReference{
																		Resource: "Bigquery/Table",
																		Field:    "name",
																	},
																},
															},
														},
													},
												},
											},
											"cloudStorageOptions": &dcl.Property{
												Type:        "object",
												GoName:      "CloudStorageOptions",
												GoType:      "JobTriggerInspectJobStorageConfigCloudStorageOptions",
												Description: "Google Cloud Storage options.",
												Conflicts: []string{
													"datastoreOptions",
													"bigQueryOptions",
													"hybridOptions",
												},
												Properties: map[string]*dcl.Property{
													"bytesLimitPerFile": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "BytesLimitPerFile",
														Description: "Max number of bytes to scan from a file. If a scanned file's size is bigger than this value then the rest of the bytes are omitted. Only one of bytes_limit_per_file and bytes_limit_per_file_percent can be specified. Cannot be set if de-identification is requested.",
														Conflicts: []string{
															"bytesLimitPerFilePercent",
														},
													},
													"bytesLimitPerFilePercent": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "BytesLimitPerFilePercent",
														Description: "Max percentage of bytes to scan from a file. The rest are omitted. The number of bytes scanned is rounded down. Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0. Only one of bytes_limit_per_file and bytes_limit_per_file_percent can be specified. Cannot be set if de-identification is requested.",
														Conflicts: []string{
															"bytesLimitPerFile",
														},
													},
													"fileSet": &dcl.Property{
														Type:        "object",
														GoName:      "FileSet",
														GoType:      "JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet",
														Description: "The set of one or more files to scan.",
														Properties: map[string]*dcl.Property{
															"regexFileSet": &dcl.Property{
																Type:        "object",
																GoName:      "RegexFileSet",
																GoType:      "JobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet",
																Description: "The regex-filtered set of files to scan. Exactly one of `url` or `regex_file_set` must be set.",
																Conflicts: []string{
																	"url",
																},
																Required: []string{
																	"bucketName",
																},
																Properties: map[string]*dcl.Property{
																	"bucketName": &dcl.Property{
																		Type:        "string",
																		GoName:      "BucketName",
																		Description: "The name of a Cloud Storage bucket. Required.",
																		ResourceReferences: []*dcl.PropertyResourceReference{
																			&dcl.PropertyResourceReference{
																				Resource: "Storage/Bucket",
																				Field:    "name",
																			},
																		},
																	},
																	"excludeRegex": &dcl.Property{
																		Type:        "array",
																		GoName:      "ExcludeRegex",
																		Description: "A list of regular expressions matching file paths to exclude. All files in the bucket that match at least one of these regular expressions will be excluded from the scan. Regular expressions use RE2 [syntax](https://github.com/google/re2/wiki/Syntax); a guide can be found under the google/re2 repository on GitHub.",
																		SendEmpty:   true,
																		ListType:    "list",
																		Items: &dcl.Property{
																			Type:   "string",
																			GoType: "string",
																		},
																	},
																	"includeRegex": &dcl.Property{
																		Type:        "array",
																		GoName:      "IncludeRegex",
																		Description: "A list of regular expressions matching file paths to include. All files in the bucket that match at least one of these regular expressions will be included in the set of files, except for those that also match an item in `exclude_regex`. Leaving this field empty will match all files by default (this is equivalent to including `.*` in the list). Regular expressions use RE2 [syntax](https://github.com/google/re2/wiki/Syntax); a guide can be found under the google/re2 repository on GitHub.",
																		SendEmpty:   true,
																		ListType:    "list",
																		Items: &dcl.Property{
																			Type:   "string",
																			GoType: "string",
																		},
																	},
																},
															},
															"url": &dcl.Property{
																Type:        "string",
																GoName:      "Url",
																Description: "The Cloud Storage url of the file(s) to scan, in the format `gs:///`. Trailing wildcard in the path is allowed. If the url ends in a trailing slash, the bucket or directory represented by the url will be scanned non-recursively (content in sub-directories will not be scanned). This means that `gs://mybucket/` is equivalent to `gs://mybucket/*`, and `gs://mybucket/directory/` is equivalent to `gs://mybucket/directory/*`. Exactly one of `url` or `regex_file_set` must be set.",
																Conflicts: []string{
																	"regexFileSet",
																},
															},
														},
													},
													"fileTypes": &dcl.Property{
														Type:        "array",
														GoName:      "FileTypes",
														Description: "List of file type groups to include in the scan. If empty, all files are scanned and available data format processors are applied. In addition, the binary content of the selected files is always scanned as well. Images are scanned only as binary if the specified region does not support image inspection and no file_types were specified. Image inspection is restricted to 'global', 'us', 'asia', and 'europe'.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "JobTriggerInspectJobStorageConfigCloudStorageOptionsFileTypesEnum",
															Enum: []string{
																"FILE_TYPE_UNSPECIFIED",
																"BINARY_FILE",
																"TEXT_FILE",
																"IMAGE",
																"WORD",
																"PDF",
																"AVRO",
																"CSV",
																"TSV",
															},
														},
													},
													"filesLimitPercent": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "FilesLimitPercent",
														Description: "Limits the number of files to scan to this percentage of the input FileSet. Number of files scanned is rounded down. Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0.",
													},
													"sampleMethod": &dcl.Property{
														Type:        "string",
														GoName:      "SampleMethod",
														GoType:      "JobTriggerInspectJobStorageConfigCloudStorageOptionsSampleMethodEnum",
														Description: " Possible values: SAMPLE_METHOD_UNSPECIFIED, TOP, RANDOM_START",
														Enum: []string{
															"SAMPLE_METHOD_UNSPECIFIED",
															"TOP",
															"RANDOM_START",
														},
													},
												},
											},
											"datastoreOptions": &dcl.Property{
												Type:        "object",
												GoName:      "DatastoreOptions",
												GoType:      "JobTriggerInspectJobStorageConfigDatastoreOptions",
												Description: "Google Cloud Datastore options.",
												Conflicts: []string{
													"cloudStorageOptions",
													"bigQueryOptions",
													"hybridOptions",
												},
												Properties: map[string]*dcl.Property{
													"kind": &dcl.Property{
														Type:        "object",
														GoName:      "Kind",
														GoType:      "JobTriggerInspectJobStorageConfigDatastoreOptionsKind",
														Description: "The kind to process.",
														Properties: map[string]*dcl.Property{
															"name": &dcl.Property{
																Type:        "string",
																GoName:      "Name",
																Description: "The name of the kind.",
															},
														},
													},
													"partitionId": &dcl.Property{
														Type:        "object",
														GoName:      "PartitionId",
														GoType:      "JobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId",
														Description: "A partition ID identifies a grouping of entities. The grouping is always by project namespace ID may be empty.",
														Properties: map[string]*dcl.Property{
															"namespaceId": &dcl.Property{
																Type:        "string",
																GoName:      "NamespaceId",
																Description: "If not empty, the ID of the namespace to which the entities belong.",
															},
															"projectId": &dcl.Property{
																Type:        "string",
																GoName:      "ProjectId",
																Description: "The ID of the project to which the entities belong.",
																ResourceReferences: []*dcl.PropertyResourceReference{
																	&dcl.PropertyResourceReference{
																		Resource: "Cloudresourcemanager/Project",
																		Field:    "name",
																	},
																},
															},
														},
													},
												},
											},
											"hybridOptions": &dcl.Property{
												Type:        "object",
												GoName:      "HybridOptions",
												GoType:      "JobTriggerInspectJobStorageConfigHybridOptions",
												Description: "Hybrid inspection options.",
												Conflicts: []string{
													"datastoreOptions",
													"cloudStorageOptions",
													"bigQueryOptions",
												},
												Properties: map[string]*dcl.Property{
													"description": &dcl.Property{
														Type:        "string",
														GoName:      "Description",
														Description: "A short description of where the data is coming from. Will be stored once in the job. 256 max length.",
													},
													"labels": &dcl.Property{
														Type: "object",
														AdditionalProperties: &dcl.Property{
															Type: "string",
														},
														GoName:      "Labels",
														Description: "To organize findings, these labels will be added to each finding. Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`. Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`. No more than 10 labels can be associated with a given finding. Examples: * `\"environment\" : \"production\"` * `\"pipeline\" : \"etl\"`",
													},
													"requiredFindingLabelKeys": &dcl.Property{
														Type:        "array",
														GoName:      "RequiredFindingLabelKeys",
														Description: "These are labels that each inspection request must include within their 'finding_labels' map. Request may contain others, but any missing one of these will be rejected. Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`. No more than 10 keys can be required.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "string",
														},
													},
													"tableOptions": &dcl.Property{
														Type:        "object",
														GoName:      "TableOptions",
														GoType:      "JobTriggerInspectJobStorageConfigHybridOptionsTableOptions",
														Description: "If the container is a table, additional information to make findings meaningful such as the columns that are primary keys.",
														Properties: map[string]*dcl.Property{
															"identifyingFields": &dcl.Property{
																Type:        "array",
																GoName:      "IdentifyingFields",
																Description: "The columns that are the primary keys for table objects included in ContentItem. A copy of this cell's value will stored alongside alongside each finding so that the finding can be traced to the specific row it came from. No more than 3 may be provided.",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "object",
																	GoType: "JobTriggerInspectJobStorageConfigHybridOptionsTableOptionsIdentifyingFields",
																	Properties: map[string]*dcl.Property{
																		"name": &dcl.Property{
																			Type:        "string",
																			GoName:      "Name",
																			Description: "Name describing the field.",
																		},
																	},
																},
															},
														},
													},
												},
											},
											"timespanConfig": &dcl.Property{
												Type:   "object",
												GoName: "TimespanConfig",
												GoType: "JobTriggerInspectJobStorageConfigTimespanConfig",
												Properties: map[string]*dcl.Property{
													"enableAutoPopulationOfTimespanConfig": &dcl.Property{
														Type:        "boolean",
														GoName:      "EnableAutoPopulationOfTimespanConfig",
														Description: "When the job is started by a JobTrigger we will automatically figure out a valid start_time to avoid scanning files that have not been modified since the last time the JobTrigger executed. This will be based on the time of the execution of the last run of the JobTrigger.",
													},
													"endTime": &dcl.Property{
														Type:        "string",
														Format:      "date-time",
														GoName:      "EndTime",
														Description: "Exclude files, tables, or rows newer than this value. If not set, no upper time limit is applied.",
													},
													"startTime": &dcl.Property{
														Type:        "string",
														Format:      "date-time",
														GoName:      "StartTime",
														Description: "Exclude files, tables, or rows older than this value. If not set, no lower time limit is applied.",
													},
													"timestampField": &dcl.Property{
														Type:        "object",
														GoName:      "TimestampField",
														GoType:      "JobTriggerInspectJobStorageConfigTimespanConfigTimestampField",
														Description: "Specification of the field containing the timestamp of scanned items. Used for data sources like Datastore and BigQuery. For BigQuery: If this value is not specified and the table was modified between the given start and end times, the entire table will be scanned. If this value is specified, then rows are filtered based on the given start and end times. Rows with a `NULL` value in the provided BigQuery column are skipped. Valid data types of the provided BigQuery column are: `INTEGER`, `DATE`, `TIMESTAMP`, and `DATETIME`. For Datastore: If this value is specified, then entities are filtered based on the given start and end times. If an entity does not contain the provided timestamp property or contains empty or invalid values, then it is included. Valid data types of the provided timestamp property are: `TIMESTAMP`.",
														Properties: map[string]*dcl.Property{
															"name": &dcl.Property{
																Type:        "string",
																GoName:      "Name",
																Description: "Name describing the field.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
							"lastRunTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "LastRunTime",
								ReadOnly:    true,
								Description: "Output only. The timestamp of the last time this trigger executed.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location of the resource",
								Immutable:   true,
							},
							"locationId": &dcl.Property{
								Type:        "string",
								GoName:      "LocationId",
								ReadOnly:    true,
								Description: "Output only. The geographic location where this resource is stored.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"parent": &dcl.Property{
								Type:                "string",
								GoName:              "Parent",
								Description:         "The parent of the resource",
								Immutable:           true,
								ForwardSlashAllowed: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"status": &dcl.Property{
								Type:        "string",
								GoName:      "Status",
								GoType:      "JobTriggerStatusEnum",
								Description: "Required. A status for this trigger. Possible values: STATUS_UNSPECIFIED, HEALTHY, PAUSED, CANCELLED",
								Immutable:   true,
								Enum: []string{
									"STATUS_UNSPECIFIED",
									"HEALTHY",
									"PAUSED",
									"CANCELLED",
								},
							},
							"triggers": &dcl.Property{
								Type:        "array",
								GoName:      "Triggers",
								Description: "A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "JobTriggerTriggers",
									Properties: map[string]*dcl.Property{
										"manual": &dcl.Property{
											Type:        "object",
											GoName:      "Manual",
											GoType:      "JobTriggerTriggersManual",
											Description: "For use with hybrid jobs. Jobs must be manually created and finished.",
											Conflicts: []string{
												"schedule",
											},
											SendEmpty:  true,
											Properties: map[string]*dcl.Property{},
										},
										"schedule": &dcl.Property{
											Type:        "object",
											GoName:      "Schedule",
											GoType:      "JobTriggerTriggersSchedule",
											Description: "Create a job on a repeating basis based on the elapse of time.",
											Conflicts: []string{
												"manual",
											},
											Properties: map[string]*dcl.Property{
												"recurrencePeriodDuration": &dcl.Property{
													Type:        "string",
													GoName:      "RecurrencePeriodDuration",
													Description: "With this option a job is started a regular periodic basis. For example: every day (86400 seconds). A scheduled start time will be skipped if the previous execution has not ended when its scheduled time occurs. This value must be set to a time duration greater than or equal to 1 day and can be no longer than 60 days.",
												},
											},
										},
									},
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The last update timestamp of a triggeredJob.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
