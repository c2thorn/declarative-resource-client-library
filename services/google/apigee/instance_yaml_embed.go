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
// GENERATED BY gen_go_data.go
// gen_go_data -package apigee -var YAML_instance blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/instance.yaml

package apigee

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/instance.yaml
var YAML_instance = []byte("info:\n  title: Apigee/Instance\n  description: The Apigee Instance resource\n  x-dcl-struct-name: Instance\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Instance\n    parameters:\n    - name: instance\n      required: true\n      description: A full instance of a Instance\n  apply:\n    description: The function used to apply information about a Instance\n    parameters:\n    - name: instance\n      required: true\n      description: A full instance of a Instance\n  delete:\n    description: The function used to delete a Instance\n    parameters:\n    - name: instance\n      required: true\n      description: A full instance of a Instance\n  deleteAll:\n    description: The function used to delete all Instance\n    parameters:\n    - name: apigeeOrganization\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Instance\n    parameters:\n    - name: apigeeOrganization\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Instance:\n      title: Instance\n      x-dcl-id: organizations/{{apigee_organization}}/instances/{{name}}\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 4800\n      x-dcl-delete-timeout: 4800\n      type: object\n      required:\n      - name\n      - location\n      - apigeeOrganization\n      properties:\n        apigeeOrganization:\n          type: string\n          x-dcl-go-name: ApigeeOrganization\n          description: The apigee organization for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Apigee/Organization\n            field: name\n            parent: true\n        createdAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: CreatedAt\n          readOnly: true\n          description: Output only. Time the instance was created in milliseconds\n            since epoch.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Description of the instance.\n          x-kubernetes-immutable: true\n        diskEncryptionKeyName:\n          type: string\n          x-dcl-go-name: DiskEncryptionKeyName\n          description: 'Customer Managed Encryption Key (CMEK) used for disk and volume\n            encryption. Required for Apigee paid subscriptions only. Use the following\n            format: projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudkms/CryptoKey\n            field: name\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Optional. Display name for the instance.\n          x-kubernetes-immutable: true\n        host:\n          type: string\n          x-dcl-go-name: Host\n          readOnly: true\n          description: Output only. Internal hostname or IP address of the Apigee\n            endpoint used by clients to connect to the service.\n          x-kubernetes-immutable: true\n        lastModifiedAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: LastModifiedAt\n          readOnly: true\n          description: Output only. Time the instance was last modified in milliseconds\n            since epoch.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: Required. Compute Engine location where the instance resides.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Resource ID of the instance. Values must match the\n            regular expression ^[a-z][a-z\\-\\d]{0,30}[a-z\\d]$.\n          x-kubernetes-immutable: true\n        peeringCidrRange:\n          type: string\n          x-dcl-go-name: PeeringCidrRange\n          x-dcl-go-type: InstancePeeringCidrRangeEnum\n          description: 'Optional. Size of the CIDR block range that will be reserved\n            by the instance. PAID apigee_organizations support SLASH_16 to SLASH_20\n            and defaults to SLASH_16. Evaluation organizations support only SLASH_23.\n            Possible values: CIDR_RANGE_UNSPECIFIED, SLASH_16, SLASH_17, SLASH_18,\n            SLASH_19, SLASH_20, SLASH_23'\n          x-kubernetes-immutable: true\n          enum:\n          - CIDR_RANGE_UNSPECIFIED\n          - SLASH_16\n          - SLASH_17\n          - SLASH_18\n          - SLASH_19\n          - SLASH_20\n          - SLASH_23\n        port:\n          type: string\n          x-dcl-go-name: Port\n          readOnly: true\n          description: Output only. Port number of the exposed Apigee endpoint.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: InstanceStateEnum\n          readOnly: true\n          description: 'Output only. State of the instance. Values other than `ACTIVE`\n            means the resource is not ready to use. Possible values: SNAPSHOT_STATE_UNSPECIFIED,\n            MISSING, OK_DOCSTORE, OK_SUBMITTED, OK_EXTERNAL, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - SNAPSHOT_STATE_UNSPECIFIED\n          - MISSING\n          - OK_DOCSTORE\n          - OK_SUBMITTED\n          - OK_EXTERNAL\n          - DELETED\n")

// 5456 bytes
// MD5: 7dee735a9ee8b30b664d567a0b3cb520
