# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/ngapi/ngapi.proto](#proto/ngapi/ngapi.proto)
    - [CapSummary](#io.netograph.CapSummary)
    - [CapSummary.Plan](#io.netograph.CapSummary.Plan)
    - [CapSummary.Root](#io.netograph.CapSummary.Root)
    - [CapSummary.Stats](#io.netograph.CapSummary.Stats)
    - [CaptureInfoRequest](#io.netograph.CaptureInfoRequest)
    - [CaptureInfoResult](#io.netograph.CaptureInfoResult)
    - [CaptureLogRequest](#io.netograph.CaptureLogRequest)
    - [CaptureLogResult](#io.netograph.CaptureLogResult)
    - [Dataset](#io.netograph.Dataset)
    - [DatasetsRequest](#io.netograph.DatasetsRequest)
    - [DomainHistoryRequest](#io.netograph.DomainHistoryRequest)
    - [DomainHistoryResult](#io.netograph.DomainHistoryResult)
    - [DomainSearchRequest](#io.netograph.DomainSearchRequest)
    - [DomainSearchResult](#io.netograph.DomainSearchResult)
    - [DomainsForIPRequest](#io.netograph.DomainsForIPRequest)
    - [DomainsForIPResult](#io.netograph.DomainsForIPResult)
    - [IPHistoryRequest](#io.netograph.IPHistoryRequest)
    - [IPHistoryResult](#io.netograph.IPHistoryResult)
    - [IPLogSearchRequest](#io.netograph.IPLogSearchRequest)
    - [IPLogSearchResult](#io.netograph.IPLogSearchResult)
    - [IPSearchRequest](#io.netograph.IPSearchRequest)
    - [IPSearchResult](#io.netograph.IPSearchResult)
    - [IPsForDomainRequest](#io.netograph.IPsForDomainRequest)
    - [IPsForDomainResult](#io.netograph.IPsForDomainResult)
    - [MetaForCaptureRequest](#io.netograph.MetaForCaptureRequest)
    - [MetaForCaptureResult](#io.netograph.MetaForCaptureResult)
    - [MetaSearchRequest](#io.netograph.MetaSearchRequest)
    - [MetaSearchResult](#io.netograph.MetaSearchResult)
    - [Metadata](#io.netograph.Metadata)
    - [RedirsByDestinationRequest](#io.netograph.RedirsByDestinationRequest)
    - [RedirsByDestinationResponse](#io.netograph.RedirsByDestinationResponse)
    - [RedirsBySourceRequest](#io.netograph.RedirsBySourceRequest)
    - [RedirsBySourceResponse](#io.netograph.RedirsBySourceResponse)
    - [RootLogSearchRequest](#io.netograph.RootLogSearchRequest)
    - [RootLogSearchResult](#io.netograph.RootLogSearchResult)
    - [RootsForSatelliteRequest](#io.netograph.RootsForSatelliteRequest)
    - [RootsForSatelliteResult](#io.netograph.RootsForSatelliteResult)
    - [SatelliteLogSearchRequest](#io.netograph.SatelliteLogSearchRequest)
    - [SatelliteLogSearchResult](#io.netograph.SatelliteLogSearchResult)
    - [SatellitesForRootRequest](#io.netograph.SatellitesForRootRequest)
    - [SatellitesForRootResult](#io.netograph.SatellitesForRootResult)
    - [SubmitCaptureRequest](#io.netograph.SubmitCaptureRequest)
    - [SubmitCaptureResult](#io.netograph.SubmitCaptureResult)
    - [TempCaptureRequest](#io.netograph.TempCaptureRequest)
    - [TempCaptureResult](#io.netograph.TempCaptureResult)
    - [URLLogSearchRequest](#io.netograph.URLLogSearchRequest)
    - [URLLogSearchResult](#io.netograph.URLLogSearchResult)
  
  
  
    - [Netograph](#io.netograph.Netograph)
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/ngapi/ngapi.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/ngapi/ngapi.proto



<a name="io.netograph.CapSummary"></a>

### CapSummary



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| assets | [string](#string) |  |  |
| roots | [CapSummary.Root](#io.netograph.CapSummary.Root) | repeated |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| plan | [CapSummary.Plan](#io.netograph.CapSummary.Plan) |  |  |
| stats | [CapSummary.Stats](#io.netograph.CapSummary.Stats) |  |  |






<a name="io.netograph.CapSummary.Plan"></a>

### CapSummary.Plan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| urls | [string](#string) | repeated |  |






<a name="io.netograph.CapSummary.Root"></a>

### CapSummary.Root



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| urls | [string](#string) | repeated |  |






<a name="io.netograph.CapSummary.Stats"></a>

### CapSummary.Stats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flows | [int32](#int32) |  |  |
| websockets | [int32](#int32) |  |  |
| hosts | [int32](#int32) |  |  |






<a name="io.netograph.CaptureInfoRequest"></a>

### CaptureInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="io.netograph.CaptureInfoResult"></a>

### CaptureInfoResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |






<a name="io.netograph.CaptureLogRequest"></a>

### CaptureLogRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the start of the range is the most recent time. If start is zero, it&#39;s taken to be the largest possible time value. |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the end of the range is the least recent time. If end is zero, it&#39;s taken to be the smallest possible time value. |






<a name="io.netograph.CaptureLogResult"></a>

### CaptureLogResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.Dataset"></a>

### Dataset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| urlbase | [string](#string) |  |  |
| deleted | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| readonly | [bool](#bool) |  |  |






<a name="io.netograph.DatasetsRequest"></a>

### DatasetsRequest







<a name="io.netograph.DomainHistoryRequest"></a>

### DomainHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| domain | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |






<a name="io.netograph.DomainHistoryResult"></a>

### DomainHistoryResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |






<a name="io.netograph.DomainSearchRequest"></a>

### DomainSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.DomainSearchResult"></a>

### DomainSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.CapSummary) |  |  |






<a name="io.netograph.DomainsForIPRequest"></a>

### DomainsForIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.DomainsForIPResult"></a>

### DomainsForIPResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.CapSummary) |  |  |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.IPHistoryRequest"></a>

### IPHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |






<a name="io.netograph.IPHistoryResult"></a>

### IPHistoryResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |






<a name="io.netograph.IPLogSearchRequest"></a>

### IPLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.IPLogSearchResult"></a>

### IPLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |
| ip | [string](#string) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.IPSearchRequest"></a>

### IPSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| mask | [int32](#int32) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.IPSearchResult"></a>

### IPSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.CapSummary) |  |  |






<a name="io.netograph.IPsForDomainRequest"></a>

### IPsForDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.IPsForDomainResult"></a>

### IPsForDomainResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.CapSummary) |  |  |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.MetaForCaptureRequest"></a>

### MetaForCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| id | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.MetaForCaptureResult"></a>

### MetaForCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [Metadata](#io.netograph.Metadata) |  |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.MetaSearchRequest"></a>

### MetaSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.MetaSearchResult"></a>

### MetaSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [Metadata](#io.netograph.Metadata) |  |  |
| resume | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |






<a name="io.netograph.Metadata"></a>

### Metadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="io.netograph.RedirsByDestinationRequest"></a>

### RedirsByDestinationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.RedirsByDestinationResponse"></a>

### RedirsByDestinationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [string](#string) |  |  |
| destination | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.CapSummary) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.RedirsBySourceRequest"></a>

### RedirsBySourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.RedirsBySourceResponse"></a>

### RedirsBySourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [string](#string) |  |  |
| destination | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.CapSummary) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.RootLogSearchRequest"></a>

### RootLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.RootLogSearchResult"></a>

### RootLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |






<a name="io.netograph.RootsForSatelliteRequest"></a>

### RootsForSatelliteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.RootsForSatelliteResult"></a>

### RootsForSatelliteResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| satellite | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.SatelliteLogSearchRequest"></a>

### SatelliteLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.SatelliteLogSearchResult"></a>

### SatelliteLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| satellite | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |






<a name="io.netograph.SatellitesForRootRequest"></a>

### SatellitesForRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.SatellitesForRootResult"></a>

### SatellitesForRootResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| satellite | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.SubmitCaptureRequest"></a>

### SubmitCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| notification | [string](#string) |  |  |
| urls | [string](#string) | repeated |  |
| meta | [Metadata](#io.netograph.Metadata) | repeated |  |
| skiprecent | [int64](#int64) |  |  |






<a name="io.netograph.SubmitCaptureResult"></a>

### SubmitCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| assets | [string](#string) |  |  |
| skipped | [bool](#bool) |  |  |
| id | [string](#string) |  |  |






<a name="io.netograph.TempCaptureRequest"></a>

### TempCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notification | [string](#string) |  |  |
| urls | [string](#string) | repeated |  |
| meta | [Metadata](#io.netograph.Metadata) | repeated |  |






<a name="io.netograph.TempCaptureResult"></a>

### TempCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| assets | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="io.netograph.URLLogSearchRequest"></a>

### URLLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.URLLogSearchResult"></a>

### URLLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.CapSummary) |  |  |





 

 

 


<a name="io.netograph.Netograph"></a>

### Netograph


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| TempCapture | [TempCaptureRequest](#io.netograph.TempCaptureRequest) | [TempCaptureResult](#io.netograph.TempCaptureResult) | Request a temporary capture. Temporary captures are not stored in a dataset, and the capture assets will be available for download for 24 hours before being deleted. |
| Datasets | [DatasetsRequest](#io.netograph.DatasetsRequest) | [Dataset](#io.netograph.Dataset) | List all datasets to which the authorizing account has access. This includes public datasets, which will be marked readonly. |
| SubmitCapture | [SubmitCaptureRequest](#io.netograph.SubmitCaptureRequest) | [SubmitCaptureResult](#io.netograph.SubmitCaptureResult) | Submit a capture request to a dataset. |
| CaptureInfo | [CaptureInfoRequest](#io.netograph.CaptureInfoRequest) | [CaptureInfoResult](#io.netograph.CaptureInfoResult) | Retrieve info for a specified capture by ID within a dataset. |
| CaptureLog | [CaptureLogRequest](#io.netograph.CaptureLogRequest) | [CaptureLogResult](#io.netograph.CaptureLogResult) | Retrieve the capture log for a dataset, in reverse chronological order. |
| DomainHistory | [DomainHistoryRequest](#io.netograph.DomainHistoryRequest) | [DomainHistoryResult](#io.netograph.DomainHistoryResult) | Retrieve the capture history for a specified domain in a dataset. The length of this history is capped at ~100. |
| DomainSearch | [DomainSearchRequest](#io.netograph.DomainSearchRequest) | [DomainSearchResult](#io.netograph.DomainSearchResult) | Retrieve the capture log for a dataset, in reverse chronological order. |
| DomainsForIP | [DomainsForIPRequest](#io.netograph.DomainsForIPRequest) | [DomainsForIPResult](#io.netograph.DomainsForIPResult) | Find all domains in the dataset associated with a given IP address. |
| IPHistory | [IPHistoryRequest](#io.netograph.IPHistoryRequest) | [IPHistoryResult](#io.netograph.IPHistoryResult) | Retrieve the capture history for a specified IP in a dataset. The length of this history is capped at ~100. |
| IPLogSearch | [IPLogSearchRequest](#io.netograph.IPLogSearchRequest) | [IPLogSearchResult](#io.netograph.IPLogSearchResult) | Search the dataset log for captures that contain a given IP. |
| IPSearch | [IPSearchRequest](#io.netograph.IPSearchRequest) | [IPSearchResult](#io.netograph.IPSearchResult) | Find all IPs in the dataset that match an address and integer netmask. |
| IPsForDomain | [IPsForDomainRequest](#io.netograph.IPsForDomainRequest) | [IPsForDomainResult](#io.netograph.IPsForDomainResult) | Find all IPs in a dataset associated with a given domain. |
| MetaForCapture | [MetaForCaptureRequest](#io.netograph.MetaForCaptureRequest) | [MetaForCaptureResult](#io.netograph.MetaForCaptureResult) | Get metadata associated with a specified capture within a dataset. |
| MetaSearch | [MetaSearchRequest](#io.netograph.MetaSearchRequest) | [MetaSearchResult](#io.netograph.MetaSearchResult) | Search the dataset log for captures matching a metadata query. |
| RedirsByDestination | [RedirsByDestinationRequest](#io.netograph.RedirsByDestinationRequest) | [RedirsByDestinationResponse](#io.netograph.RedirsByDestinationResponse) | Find all redirections in the dataset for a given destination domain query. |
| RedirsBySource | [RedirsBySourceRequest](#io.netograph.RedirsBySourceRequest) | [RedirsBySourceResponse](#io.netograph.RedirsBySourceResponse) | Find all redirections in the dataset for a given source domain query. |
| RootLogSearch | [RootLogSearchRequest](#io.netograph.RootLogSearchRequest) | [RootLogSearchResult](#io.netograph.RootLogSearchResult) | Search the dataset log for captures where any root domain matches a given query. |
| RootsForSatellite | [RootsForSatelliteRequest](#io.netograph.RootsForSatelliteRequest) | [RootsForSatelliteResult](#io.netograph.RootsForSatelliteResult) | Find all roots in a dataset that are associated with a given satellite query. |
| SatelliteLogSearch | [SatelliteLogSearchRequest](#io.netograph.SatelliteLogSearchRequest) | [SatelliteLogSearchResult](#io.netograph.SatelliteLogSearchResult) | Search the dataset log for captures where any satellite domain matches a given query. |
| SatellitesForRoot | [SatellitesForRootRequest](#io.netograph.SatellitesForRootRequest) | [SatellitesForRootResult](#io.netograph.SatellitesForRootResult) | Find all satellites in the dataset that are associated with a given root query. |
| URLLogSearch | [URLLogSearchRequest](#io.netograph.URLLogSearchRequest) | [URLLogSearchResult](#io.netograph.URLLogSearchResult) | Search the dataset log for captures where any root URL matches a given URL query. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

