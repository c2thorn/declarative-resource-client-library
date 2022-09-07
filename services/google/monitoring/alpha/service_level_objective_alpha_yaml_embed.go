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
// gen_go_data -package alpha -var YAML_service_level_objective blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/alpha/service_level_objective.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/alpha/service_level_objective.yaml
var YAML_service_level_objective = []byte("info:\n  title: Monitoring/ServiceLevelObjective\n  description: The Monitoring ServiceLevelObjective resource\n  x-dcl-struct-name: ServiceLevelObjective\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a ServiceLevelObjective\n    parameters:\n    - name: serviceLevelObjective\n      required: true\n      description: A full instance of a ServiceLevelObjective\n  apply:\n    description: The function used to apply information about a ServiceLevelObjective\n    parameters:\n    - name: serviceLevelObjective\n      required: true\n      description: A full instance of a ServiceLevelObjective\n  delete:\n    description: The function used to delete a ServiceLevelObjective\n    parameters:\n    - name: serviceLevelObjective\n      required: true\n      description: A full instance of a ServiceLevelObjective\n  deleteAll:\n    description: The function used to delete all ServiceLevelObjective\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: service\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ServiceLevelObjective\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: service\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ServiceLevelObjective:\n      title: ServiceLevelObjective\n      x-dcl-id: projects/{{project}}/services/{{service}}/serviceLevelObjectives/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: userLabels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - goal\n      - project\n      - service\n      properties:\n        calendarPeriod:\n          type: string\n          x-dcl-go-name: CalendarPeriod\n          x-dcl-go-type: ServiceLevelObjectiveCalendarPeriodEnum\n          description: 'A calendar period, semantically \"since the start of the current\n            ``\". At this time, only `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH` are supported.\n            Possible values: CALENDAR_PERIOD_UNSPECIFIED, DAY, WEEK, FORTNIGHT, MONTH,\n            QUARTER, HALF, YEAR'\n          x-dcl-conflicts:\n          - rollingPeriod\n          enum:\n          - CALENDAR_PERIOD_UNSPECIFIED\n          - DAY\n          - WEEK\n          - FORTNIGHT\n          - MONTH\n          - QUARTER\n          - HALF\n          - YEAR\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Time stamp of the `Create` or most recent `Update` command\n            on this `Slo`.\n          x-kubernetes-immutable: true\n        deleteTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: DeleteTime\n          readOnly: true\n          description: Time stamp of the `Update` or `Delete` command that made this\n            no longer a current `Slo`. This field is not populated in `ServiceLevelObjective`s\n            returned from calls to `GetServiceLevelObjective` and `ListServiceLevelObjectives`,\n            because it is always empty in the current version. It is populated in\n            `ServiceLevelObjective`s representing previous versions in the output\n            of `ListServiceLevelObjectiveVersions`. Because all old configuration\n            versions are stored, `Update` operations mark the obsoleted version as\n            deleted.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Name used for UI elements listing this SLO.\n        goal:\n          type: number\n          format: double\n          x-dcl-go-name: Goal\n          description: The fraction of service that must be good in order for this\n            objective to be met. `0 < goal <= 0.999`.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Resource name for this `ServiceLevelObjective`. The format\n            is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]'\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        rollingPeriod:\n          type: string\n          x-dcl-go-name: RollingPeriod\n          description: A rolling time period, semantically \"in the past ``\". Must\n            be an integer multiple of 1 day no larger than 30 days.\n          x-dcl-conflicts:\n          - calendarPeriod\n        service:\n          type: string\n          x-dcl-go-name: Service\n          description: The service for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Monitoring/Service\n            field: name\n            parent: true\n        serviceLevelIndicator:\n          type: object\n          x-dcl-go-name: ServiceLevelIndicator\n          x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicator\n          description: The definition of good service, used to measure and calculate\n            the quality of the `Service`'s performance with respect to a single aspect\n            of service quality.\n          properties:\n            basicSli:\n              type: object\n              x-dcl-go-name: BasicSli\n              x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorBasicSli\n              description: Basic SLI on a well-known service type.\n              x-dcl-conflicts:\n              - requestBased\n              - windowsBased\n              properties:\n                availability:\n                  type: object\n                  x-dcl-go-name: Availability\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability\n                  description: Good service is defined to be the count of requests\n                    made to this service that return successfully.\n                  x-dcl-conflicts:\n                  - latency\n                  - operationAvailability\n                  - operationLatency\n                latency:\n                  type: object\n                  x-dcl-go-name: Latency\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency\n                  description: Good service is defined to be the count of requests\n                    made to this service that are fast enough with respect to `latency.threshold`.\n                  x-dcl-conflicts:\n                  - availability\n                  - operationAvailability\n                  - operationLatency\n                  properties:\n                    experience:\n                      type: string\n                      x-dcl-go-name: Experience\n                      x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum\n                      description: 'A description of the experience associated with\n                        failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED,\n                        DELIGHTING, SATISFYING, ANNOYING'\n                      enum:\n                      - LATENCY_EXPERIENCE_UNSPECIFIED\n                      - DELIGHTING\n                      - SATISFYING\n                      - ANNOYING\n                    threshold:\n                      type: string\n                      x-dcl-go-name: Threshold\n                      description: Good service is defined to be the count of requests\n                        made to this service that return in no more than `threshold`.\n                location:\n                  type: array\n                  x-dcl-go-name: Location\n                  description: 'OPTIONAL: The set of locations to which this SLI is\n                    relevant. Telemetry from other locations will not be used to calculate\n                    performance for this SLI. If omitted, this SLI applies to all\n                    locations in which the Service has activity. For service types\n                    that don''t support breaking down by location, setting this field\n                    will result in an error.'\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n                method:\n                  type: array\n                  x-dcl-go-name: Method\n                  description: 'OPTIONAL: The set of RPCs to which this SLI is relevant.\n                    Telemetry from other methods will not be used to calculate performance\n                    for this SLI. If omitted, this SLI applies to all the Service''s\n                    methods. For service types that don''t support breaking down by\n                    method, setting this field will result in an error.'\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n                operationAvailability:\n                  type: object\n                  x-dcl-go-name: OperationAvailability\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability\n                  description: Good service is defined to be the count of operations\n                    performed by this service that return successfully\n                  x-dcl-conflicts:\n                  - availability\n                  - latency\n                  - operationLatency\n                operationLatency:\n                  type: object\n                  x-dcl-go-name: OperationLatency\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency\n                  description: Good service is defined to be the count of operations\n                    performed by this service that are fast enough with respect to\n                    `operation_latency.threshold`.\n                  x-dcl-conflicts:\n                  - availability\n                  - latency\n                  - operationAvailability\n                  properties:\n                    experience:\n                      type: string\n                      x-dcl-go-name: Experience\n                      x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum\n                      description: 'A description of the experience associated with\n                        failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED,\n                        DELIGHTING, SATISFYING, ANNOYING'\n                      enum:\n                      - LATENCY_EXPERIENCE_UNSPECIFIED\n                      - DELIGHTING\n                      - SATISFYING\n                      - ANNOYING\n                    threshold:\n                      type: string\n                      x-dcl-go-name: Threshold\n                      description: Good service is defined to be the count of operations\n                        that are completed in no more than `threshold`.\n                version:\n                  type: array\n                  x-dcl-go-name: Version\n                  description: 'OPTIONAL: The set of API versions to which this SLI\n                    is relevant. Telemetry from other API versions will not be used\n                    to calculate performance for this SLI. If omitted, this SLI applies\n                    to all API versions. For service types that don''t support breaking\n                    down by version, setting this field will result in an error.'\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n            requestBased:\n              type: object\n              x-dcl-go-name: RequestBased\n              x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorRequestBased\n              description: Request-based SLIs\n              x-dcl-conflicts:\n              - basicSli\n              - windowsBased\n              properties:\n                distributionCut:\n                  type: object\n                  x-dcl-go-name: DistributionCut\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut\n                  description: '`distribution_cut` is used when `good_service` is\n                    a count of values aggregated in a `Distribution` that fall into\n                    a good range. The `total_service` is the total count of all values\n                    aggregated in the `Distribution`.'\n                  x-dcl-conflicts:\n                  - goodTotalRatio\n                  properties:\n                    distributionFilter:\n                      type: string\n                      x-dcl-go-name: DistributionFilter\n                      description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                        specifying a `TimeSeries` aggregating values. Must have `ValueType\n                        = DISTRIBUTION` and `MetricKind = DELTA` or `MetricKind =\n                        CUMULATIVE`.\n                    range:\n                      type: object\n                      x-dcl-go-name: Range\n                      x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange\n                      description: Range of values considered \"good.\" For a one-sided\n                        range, set one bound to an infinite value.\n                      properties:\n                        max:\n                          type: number\n                          format: double\n                          x-dcl-go-name: Max\n                          description: Range maximum.\n                        min:\n                          type: number\n                          format: double\n                          x-dcl-go-name: Min\n                          description: Range minimum.\n                goodTotalRatio:\n                  type: object\n                  x-dcl-go-name: GoodTotalRatio\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio\n                  description: '`good_total_ratio` is used when the ratio of `good_service`\n                    to `total_service` is computed from two `TimeSeries`.'\n                  x-dcl-conflicts:\n                  - distributionCut\n                  properties:\n                    badServiceFilter:\n                      type: string\n                      x-dcl-go-name: BadServiceFilter\n                      description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                        specifying a `TimeSeries` quantifying bad service, either\n                        demanded service that was not provided or demanded service\n                        that was of inadequate quality. Must have `ValueType = DOUBLE`\n                        or `ValueType = INT64` and must have `MetricKind = DELTA`\n                        or `MetricKind = CUMULATIVE`.\n                    goodServiceFilter:\n                      type: string\n                      x-dcl-go-name: GoodServiceFilter\n                      description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                        specifying a `TimeSeries` quantifying good service provided.\n                        Must have `ValueType = DOUBLE` or `ValueType = INT64` and\n                        must have `MetricKind = DELTA` or `MetricKind = CUMULATIVE`.\n                    totalServiceFilter:\n                      type: string\n                      x-dcl-go-name: TotalServiceFilter\n                      description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                        specifying a `TimeSeries` quantifying total demanded service.\n                        Must have `ValueType = DOUBLE` or `ValueType = INT64` and\n                        must have `MetricKind = DELTA` or `MetricKind = CUMULATIVE`.\n            windowsBased:\n              type: object\n              x-dcl-go-name: WindowsBased\n              x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBased\n              description: Windows-based SLIs\n              x-dcl-conflicts:\n              - basicSli\n              - requestBased\n              properties:\n                goodBadMetricFilter:\n                  type: string\n                  x-dcl-go-name: GoodBadMetricFilter\n                  description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                    specifying a `TimeSeries` with `ValueType = BOOL`. The window\n                    is good if any `true` values appear in the window.\n                  x-dcl-conflicts:\n                  - goodTotalRatioThreshold\n                  - metricMeanInRange\n                  - metricSumInRange\n                goodTotalRatioThreshold:\n                  type: object\n                  x-dcl-go-name: GoodTotalRatioThreshold\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold\n                  description: A window is good if its `performance` is high enough.\n                  x-dcl-conflicts:\n                  - goodBadMetricFilter\n                  - metricMeanInRange\n                  - metricSumInRange\n                  properties:\n                    basicSliPerformance:\n                      type: object\n                      x-dcl-go-name: BasicSliPerformance\n                      x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance\n                      description: '`BasicSli` to evaluate to judge window quality.'\n                      x-dcl-conflicts:\n                      - performance\n                      properties:\n                        availability:\n                          type: object\n                          x-dcl-go-name: Availability\n                          x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability\n                          description: Good service is defined to be the count of\n                            requests made to this service that return successfully.\n                          x-dcl-conflicts:\n                          - latency\n                          - operationAvailability\n                          - operationLatency\n                        latency:\n                          type: object\n                          x-dcl-go-name: Latency\n                          x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency\n                          description: Good service is defined to be the count of\n                            requests made to this service that are fast enough with\n                            respect to `latency.threshold`.\n                          x-dcl-conflicts:\n                          - availability\n                          - operationAvailability\n                          - operationLatency\n                          properties:\n                            experience:\n                              type: string\n                              x-dcl-go-name: Experience\n                              x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum\n                              description: 'A description of the experience associated\n                                with failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED,\n                                DELIGHTING, SATISFYING, ANNOYING'\n                              enum:\n                              - LATENCY_EXPERIENCE_UNSPECIFIED\n                              - DELIGHTING\n                              - SATISFYING\n                              - ANNOYING\n                            threshold:\n                              type: string\n                              x-dcl-go-name: Threshold\n                              description: Good service is defined to be the count\n                                of requests made to this service that return in no\n                                more than `threshold`.\n                        location:\n                          type: array\n                          x-dcl-go-name: Location\n                          description: 'OPTIONAL: The set of locations to which this\n                            SLI is relevant. Telemetry from other locations will not\n                            be used to calculate performance for this SLI. If omitted,\n                            this SLI applies to all locations in which the Service\n                            has activity. For service types that don''t support breaking\n                            down by location, setting this field will result in an\n                            error.'\n                          x-dcl-send-empty: true\n                          x-dcl-list-type: list\n                          items:\n                            type: string\n                            x-dcl-go-type: string\n                        method:\n                          type: array\n                          x-dcl-go-name: Method\n                          description: 'OPTIONAL: The set of RPCs to which this SLI\n                            is relevant. Telemetry from other methods will not be\n                            used to calculate performance for this SLI. If omitted,\n                            this SLI applies to all the Service''s methods. For service\n                            types that don''t support breaking down by method, setting\n                            this field will result in an error.'\n                          x-dcl-send-empty: true\n                          x-dcl-list-type: list\n                          items:\n                            type: string\n                            x-dcl-go-type: string\n                        operationAvailability:\n                          type: object\n                          x-dcl-go-name: OperationAvailability\n                          x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability\n                          description: Good service is defined to be the count of\n                            operations performed by this service that return successfully\n                          x-dcl-conflicts:\n                          - availability\n                          - latency\n                          - operationLatency\n                        operationLatency:\n                          type: object\n                          x-dcl-go-name: OperationLatency\n                          x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency\n                          description: Good service is defined to be the count of\n                            operations performed by this service that are fast enough\n                            with respect to `operation_latency.threshold`.\n                          x-dcl-conflicts:\n                          - availability\n                          - latency\n                          - operationAvailability\n                          properties:\n                            experience:\n                              type: string\n                              x-dcl-go-name: Experience\n                              x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum\n                              description: 'A description of the experience associated\n                                with failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED,\n                                DELIGHTING, SATISFYING, ANNOYING'\n                              enum:\n                              - LATENCY_EXPERIENCE_UNSPECIFIED\n                              - DELIGHTING\n                              - SATISFYING\n                              - ANNOYING\n                            threshold:\n                              type: string\n                              x-dcl-go-name: Threshold\n                              description: Good service is defined to be the count\n                                of operations that are completed in no more than `threshold`.\n                        version:\n                          type: array\n                          x-dcl-go-name: Version\n                          description: 'OPTIONAL: The set of API versions to which\n                            this SLI is relevant. Telemetry from other API versions\n                            will not be used to calculate performance for this SLI.\n                            If omitted, this SLI applies to all API versions. For\n                            service types that don''t support breaking down by version,\n                            setting this field will result in an error.'\n                          x-dcl-send-empty: true\n                          x-dcl-list-type: list\n                          items:\n                            type: string\n                            x-dcl-go-type: string\n                    performance:\n                      type: object\n                      x-dcl-go-name: Performance\n                      x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance\n                      description: '`RequestBasedSli` to evaluate to judge window\n                        quality.'\n                      x-dcl-conflicts:\n                      - basicSliPerformance\n                      properties:\n                        distributionCut:\n                          type: object\n                          x-dcl-go-name: DistributionCut\n                          x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut\n                          description: '`distribution_cut` is used when `good_service`\n                            is a count of values aggregated in a `Distribution` that\n                            fall into a good range. The `total_service` is the total\n                            count of all values aggregated in the `Distribution`.'\n                          x-dcl-conflicts:\n                          - goodTotalRatio\n                          properties:\n                            distributionFilter:\n                              type: string\n                              x-dcl-go-name: DistributionFilter\n                              description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                                specifying a `TimeSeries` aggregating values. Must\n                                have `ValueType = DISTRIBUTION` and `MetricKind =\n                                DELTA` or `MetricKind = CUMULATIVE`.\n                            range:\n                              type: object\n                              x-dcl-go-name: Range\n                              x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange\n                              description: Range of values considered \"good.\" For\n                                a one-sided range, set one bound to an infinite value.\n                              properties:\n                                max:\n                                  type: number\n                                  format: double\n                                  x-dcl-go-name: Max\n                                  description: Range maximum.\n                                min:\n                                  type: number\n                                  format: double\n                                  x-dcl-go-name: Min\n                                  description: Range minimum.\n                        goodTotalRatio:\n                          type: object\n                          x-dcl-go-name: GoodTotalRatio\n                          x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio\n                          description: '`good_total_ratio` is used when the ratio\n                            of `good_service` to `total_service` is computed from\n                            two `TimeSeries`.'\n                          x-dcl-conflicts:\n                          - distributionCut\n                          properties:\n                            badServiceFilter:\n                              type: string\n                              x-dcl-go-name: BadServiceFilter\n                              description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                                specifying a `TimeSeries` quantifying bad service,\n                                either demanded service that was not provided or demanded\n                                service that was of inadequate quality. Must have\n                                `ValueType = DOUBLE` or `ValueType = INT64` and must\n                                have `MetricKind = DELTA` or `MetricKind = CUMULATIVE`.\n                            goodServiceFilter:\n                              type: string\n                              x-dcl-go-name: GoodServiceFilter\n                              description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                                specifying a `TimeSeries` quantifying good service\n                                provided. Must have `ValueType = DOUBLE` or `ValueType\n                                = INT64` and must have `MetricKind = DELTA` or `MetricKind\n                                = CUMULATIVE`.\n                            totalServiceFilter:\n                              type: string\n                              x-dcl-go-name: TotalServiceFilter\n                              description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                                specifying a `TimeSeries` quantifying total demanded\n                                service. Must have `ValueType = DOUBLE` or `ValueType\n                                = INT64` and must have `MetricKind = DELTA` or `MetricKind\n                                = CUMULATIVE`.\n                    threshold:\n                      type: number\n                      format: double\n                      x-dcl-go-name: Threshold\n                      description: If window `performance >= threshold`, the window\n                        is counted as good.\n                metricMeanInRange:\n                  type: object\n                  x-dcl-go-name: MetricMeanInRange\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange\n                  description: A window is good if the metric's value is in a good\n                    range, averaged across returned streams.\n                  x-dcl-conflicts:\n                  - goodBadMetricFilter\n                  - goodTotalRatioThreshold\n                  - metricSumInRange\n                  properties:\n                    range:\n                      type: object\n                      x-dcl-go-name: Range\n                      x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange\n                      description: Range of values considered \"good.\" For a one-sided\n                        range, set one bound to an infinite value.\n                      properties:\n                        max:\n                          type: number\n                          format: double\n                          x-dcl-go-name: Max\n                          description: Range maximum.\n                        min:\n                          type: number\n                          format: double\n                          x-dcl-go-name: Min\n                          description: Range minimum.\n                    timeSeries:\n                      type: string\n                      x-dcl-go-name: TimeSeries\n                      description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                        specifying the `TimeSeries` to use for evaluating window quality.\n                metricSumInRange:\n                  type: object\n                  x-dcl-go-name: MetricSumInRange\n                  x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange\n                  description: A window is good if the metric's value is in a good\n                    range, summed across returned streams.\n                  x-dcl-conflicts:\n                  - goodBadMetricFilter\n                  - goodTotalRatioThreshold\n                  - metricMeanInRange\n                  properties:\n                    range:\n                      type: object\n                      x-dcl-go-name: Range\n                      x-dcl-go-type: ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange\n                      description: Range of values considered \"good.\" For a one-sided\n                        range, set one bound to an infinite value.\n                      properties:\n                        max:\n                          type: number\n                          format: double\n                          x-dcl-go-name: Max\n                          description: Range maximum.\n                        min:\n                          type: number\n                          format: double\n                          x-dcl-go-name: Min\n                          description: Range minimum.\n                    timeSeries:\n                      type: string\n                      x-dcl-go-name: TimeSeries\n                      description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)\n                        specifying the `TimeSeries` to use for evaluating window quality.\n                windowPeriod:\n                  type: string\n                  x-dcl-go-name: WindowPeriod\n                  description: Duration over which window quality is evaluated. Must\n                    be an integer fraction of a day and at least `60s`.\n        serviceManagementOwned:\n          type: boolean\n          x-dcl-go-name: ServiceManagementOwned\n          readOnly: true\n          description: Output only. If set, this SLO is managed at the [Service Management](https://cloud.google.com/service-management/overview)\n            level. Therefore the service yaml file is the source of truth for this\n            SLO, and API `Update` and `Delete` operations are forbidden.\n          x-kubernetes-immutable: true\n        userLabels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: UserLabels\n          description: Labels which have been used to annotate the service-level objective.\n            Label keys must start with a letter. Label keys and values may contain\n            lowercase letters, numbers, underscores, and dashes. Label keys and values\n            have a maximum length of 63 characters, and must be less than 128 bytes\n            in size. Up to 64 label entries may be stored. For labels which do not\n            have a semantic value, the empty string may be supplied for the label\n            value.\n")

// 35986 bytes
// MD5: 5b8759abf4bffde4575272089d73685b
