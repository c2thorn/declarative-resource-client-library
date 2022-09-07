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
// gen_go_data -package alpha -var YAML_zone blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dataplex/alpha/zone.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dataplex/alpha/zone.yaml
var YAML_zone = []byte("info:\n  title: Dataplex/Zone\n  description: The Dataplex Zone resource\n  x-dcl-struct-name: Zone\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Zone\n    parameters:\n    - name: zone\n      required: true\n      description: A full instance of a Zone\n  apply:\n    description: The function used to apply information about a Zone\n    parameters:\n    - name: zone\n      required: true\n      description: A full instance of a Zone\n  delete:\n    description: The function used to delete a Zone\n    parameters:\n    - name: zone\n      required: true\n      description: A full instance of a Zone\n  deleteAll:\n    description: The function used to delete all Zone\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: lake\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Zone\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: lake\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Zone:\n      title: Zone\n      x-dcl-id: projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - type\n      - discoverySpec\n      - resourceSpec\n      - project\n      - location\n      - lake\n      properties:\n        assetStatus:\n          type: object\n          x-dcl-go-name: AssetStatus\n          x-dcl-go-type: ZoneAssetStatus\n          readOnly: true\n          description: Output only. Aggregated status of the underlying assets of\n            the zone.\n          properties:\n            activeAssets:\n              type: integer\n              format: int64\n              x-dcl-go-name: ActiveAssets\n              description: Number of active assets.\n            securityPolicyApplyingAssets:\n              type: integer\n              format: int64\n              x-dcl-go-name: SecurityPolicyApplyingAssets\n              description: Number of assets that are in process of updating the security\n                policy on attached resources.\n            updateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: UpdateTime\n              description: Last update time of the status.\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when the zone was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Description of the zone.\n        discoverySpec:\n          type: object\n          x-dcl-go-name: DiscoverySpec\n          x-dcl-go-type: ZoneDiscoverySpec\n          description: Required. Specification of the discovery feature applied to\n            data in this zone.\n          required:\n          - enabled\n          properties:\n            csvOptions:\n              type: object\n              x-dcl-go-name: CsvOptions\n              x-dcl-go-type: ZoneDiscoverySpecCsvOptions\n              description: Optional. Configuration for CSV data.\n              x-dcl-server-default: true\n              properties:\n                delimiter:\n                  type: string\n                  x-dcl-go-name: Delimiter\n                  description: Optional. The delimiter being used to separate values.\n                    This defaults to ','.\n                disableTypeInference:\n                  type: boolean\n                  x-dcl-go-name: DisableTypeInference\n                  description: Optional. Whether to disable the inference of data\n                    type for CSV data. If true, all columns will be registered as\n                    strings.\n                encoding:\n                  type: string\n                  x-dcl-go-name: Encoding\n                  description: Optional. The character encoding of the data. The default\n                    is UTF-8.\n                headerRows:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: HeaderRows\n                  description: Optional. The number of rows to interpret as header\n                    rows that should be skipped when reading data rows.\n            enabled:\n              type: boolean\n              x-dcl-go-name: Enabled\n              description: Required. Whether discovery is enabled.\n            excludePatterns:\n              type: array\n              x-dcl-go-name: ExcludePatterns\n              description: Optional. The list of patterns to apply for selecting data\n                to exclude during discovery. For Cloud Storage bucket assets, these\n                are interpreted as glob patterns used to match object names. For BigQuery\n                dataset assets, these are interpreted as patterns to match table names.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            includePatterns:\n              type: array\n              x-dcl-go-name: IncludePatterns\n              description: Optional. The list of patterns to apply for selecting data\n                to include during discovery if only a subset of the data should considered.\n                For Cloud Storage bucket assets, these are interpreted as glob patterns\n                used to match object names. For BigQuery dataset assets, these are\n                interpreted as patterns to match table names.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            jsonOptions:\n              type: object\n              x-dcl-go-name: JsonOptions\n              x-dcl-go-type: ZoneDiscoverySpecJsonOptions\n              description: Optional. Configuration for Json data.\n              x-dcl-server-default: true\n              properties:\n                disableTypeInference:\n                  type: boolean\n                  x-dcl-go-name: DisableTypeInference\n                  description: Optional. Whether to disable the inference of data\n                    type for Json data. If true, all columns will be registered as\n                    their primitive types (strings, number or boolean).\n                encoding:\n                  type: string\n                  x-dcl-go-name: Encoding\n                  description: Optional. The character encoding of the data. The default\n                    is UTF-8.\n            schedule:\n              type: string\n              x-dcl-go-name: Schedule\n              description: 'Optional. Cron schedule (https://en.wikipedia.org/wiki/Cron)\n                for running discovery periodically. Successive discovery runs must\n                be scheduled at least 60 minutes apart. The default value is to run\n                discovery every 60 minutes. To explicitly set a timezone to the cron\n                tab, apply a prefix in the cron tab: \"CRON_TZ=${IANA_TIME_ZONE}\" or\n                TZ=${IANA_TIME_ZONE}\". The ${IANA_TIME_ZONE} may only be a valid string\n                from IANA time zone database. For example, \"CRON_TZ=America/New_York\n                1 * * * *\", or \"TZ=America/New_York 1 * * * *\".'\n              x-dcl-server-default: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Optional. User friendly display name.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. User defined labels for the zone.\n        lake:\n          type: string\n          x-dcl-go-name: Lake\n          description: The lake for the resource\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the zone.\n          x-dcl-references:\n          - resource: Dataplex/Zone\n            field: selfLink\n            parent: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        resourceSpec:\n          type: object\n          x-dcl-go-name: ResourceSpec\n          x-dcl-go-type: ZoneResourceSpec\n          description: Required. Immutable. Specification of the resources that are\n            referenced by the assets within this zone.\n          x-kubernetes-immutable: true\n          required:\n          - locationType\n          properties:\n            locationType:\n              type: string\n              x-dcl-go-name: LocationType\n              x-dcl-go-type: ZoneResourceSpecLocationTypeEnum\n              description: 'Required. Immutable. The location type of the resources\n                that are allowed to be attached to the assets within this zone. Possible\n                values: LOCATION_TYPE_UNSPECIFIED, SINGLE_REGION, MULTI_REGION'\n              x-kubernetes-immutable: true\n              enum:\n              - LOCATION_TYPE_UNSPECIFIED\n              - SINGLE_REGION\n              - MULTI_REGION\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: ZoneStateEnum\n          readOnly: true\n          description: 'Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED,\n            ACTIVE, CREATING, DELETING, ACTION_REQUIRED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - CREATING\n          - DELETING\n          - ACTION_REQUIRED\n        type:\n          type: string\n          x-dcl-go-name: Type\n          x-dcl-go-type: ZoneTypeEnum\n          description: 'Required. Immutable. The type of the zone. Possible values:\n            TYPE_UNSPECIFIED, RAW, CURATED'\n          x-kubernetes-immutable: true\n          enum:\n          - TYPE_UNSPECIFIED\n          - RAW\n          - CURATED\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. System generated globally unique ID for the zone.\n            This ID will be different if the zone is deleted and re-created with the\n            same name.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time when the zone was last updated.\n          x-kubernetes-immutable: true\n")

// 11301 bytes
// MD5: 540a1c50e6af4e87e302c5e773949e0c
