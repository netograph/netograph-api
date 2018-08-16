# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/ngapi/ngapi.proto](#proto/ngapi/ngapi.proto)
    - [CapSummary](#.CapSummary)
    - [CapSummary.Plan](#.CapSummary.Plan)
    - [CapSummary.Root](#.CapSummary.Root)
    - [CapSummary.Stats](#.CapSummary.Stats)
    - [CaptureInfoRequest](#.CaptureInfoRequest)
    - [CaptureInfoResult](#.CaptureInfoResult)
    - [CaptureLogRequest](#.CaptureLogRequest)
    - [CaptureLogResult](#.CaptureLogResult)
    - [Dataset](#.Dataset)
    - [DatasetsRequest](#.DatasetsRequest)
    - [DomainHistoryRequest](#.DomainHistoryRequest)
    - [DomainHistoryResult](#.DomainHistoryResult)
    - [DomainSearchRequest](#.DomainSearchRequest)
    - [DomainSearchResult](#.DomainSearchResult)
    - [DomainsForIPRequest](#.DomainsForIPRequest)
    - [DomainsForIPResult](#.DomainsForIPResult)
    - [IPHistoryRequest](#.IPHistoryRequest)
    - [IPHistoryResult](#.IPHistoryResult)
    - [IPLogSearchRequest](#.IPLogSearchRequest)
    - [IPLogSearchResult](#.IPLogSearchResult)
    - [IPSearchRequest](#.IPSearchRequest)
    - [IPSearchResult](#.IPSearchResult)
    - [IPsForDomainRequest](#.IPsForDomainRequest)
    - [IPsForDomainResult](#.IPsForDomainResult)
    - [MetaForCaptureRequest](#.MetaForCaptureRequest)
    - [MetaForCaptureResult](#.MetaForCaptureResult)
    - [MetaSearchRequest](#.MetaSearchRequest)
    - [MetaSearchResult](#.MetaSearchResult)
    - [Metadata](#.Metadata)
    - [RedirsByDestinationRequest](#.RedirsByDestinationRequest)
    - [RedirsByDestinationResponse](#.RedirsByDestinationResponse)
    - [RedirsBySourceRequest](#.RedirsBySourceRequest)
    - [RedirsBySourceResponse](#.RedirsBySourceResponse)
    - [RootLogSearchRequest](#.RootLogSearchRequest)
    - [RootLogSearchResult](#.RootLogSearchResult)
    - [RootsForSatelliteRequest](#.RootsForSatelliteRequest)
    - [RootsForSatelliteResult](#.RootsForSatelliteResult)
    - [SatelliteLogSearchRequest](#.SatelliteLogSearchRequest)
    - [SatelliteLogSearchResult](#.SatelliteLogSearchResult)
    - [SatellitesForRootRequest](#.SatellitesForRootRequest)
    - [SatellitesForRootResult](#.SatellitesForRootResult)
    - [SubmitCaptureRequest](#.SubmitCaptureRequest)
    - [SubmitCaptureResult](#.SubmitCaptureResult)
    - [TempCaptureRequest](#.TempCaptureRequest)
    - [TempCaptureResult](#.TempCaptureResult)
    - [URLLogSearchRequest](#.URLLogSearchRequest)
    - [URLLogSearchResult](#.URLLogSearchResult)
  
  
  
    - [Netograph](#.Netograph)
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/ngapi/ngapi.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/ngapi/ngapi.proto



<a name=".CapSummary"></a>

### CapSummary



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| assets | [string](#string) |  |  |
| roots | [CapSummary.Root](#CapSummary.Root) | repeated |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| plan | [CapSummary.Plan](#CapSummary.Plan) |  |  |
| stats | [CapSummary.Stats](#CapSummary.Stats) |  |  |






<a name=".CapSummary.Plan"></a>

### CapSummary.Plan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| urls | [string](#string) | repeated |  |






<a name=".CapSummary.Root"></a>

### CapSummary.Root



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| urls | [string](#string) | repeated |  |






<a name=".CapSummary.Stats"></a>

### CapSummary.Stats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flows | [int32](#int32) |  |  |
| websockets | [int32](#int32) |  |  |
| hosts | [int32](#int32) |  |  |






<a name=".CaptureInfoRequest"></a>

### CaptureInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name=".CaptureInfoResult"></a>

### CaptureInfoResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#CapSummary) |  |  |






<a name=".CaptureLogRequest"></a>

### CaptureLogRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the start of the range is the most recent time. If start is zero, it&#39;s taken to be the largest possible time value. |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the end of the range is the least recent time. If end is zero, it&#39;s taken to be the smallest possible time value. |






<a name=".CaptureLogResult"></a>

### CaptureLogResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#CapSummary) |  |  |
| resume | [string](#string) |  |  |






<a name=".Dataset"></a>

### Dataset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| urlbase | [string](#string) |  |  |
| deleted | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| readonly | [bool](#bool) |  |  |






<a name=".DatasetsRequest"></a>

### DatasetsRequest







<a name=".DomainHistoryRequest"></a>

### DomainHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| domain | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |






<a name=".DomainHistoryResult"></a>

### DomainHistoryResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#CapSummary) |  |  |






<a name=".DomainSearchRequest"></a>

### DomainSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".DomainSearchResult"></a>

### DomainSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| latestcapture | [CapSummary](#CapSummary) |  |  |






<a name=".DomainsForIPRequest"></a>

### DomainsForIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".DomainsForIPResult"></a>

### DomainsForIPResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| latestcapture | [CapSummary](#CapSummary) |  |  |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  |  |






<a name=".IPHistoryRequest"></a>

### IPHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |






<a name=".IPHistoryResult"></a>

### IPHistoryResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#CapSummary) |  |  |






<a name=".IPLogSearchRequest"></a>

### IPLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".IPLogSearchResult"></a>

### IPLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#CapSummary) |  |  |
| ip | [string](#string) |  |  |
| resume | [string](#string) |  |  |






<a name=".IPSearchRequest"></a>

### IPSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| mask | [int32](#int32) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".IPSearchResult"></a>

### IPSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| latestcapture | [CapSummary](#CapSummary) |  |  |






<a name=".IPsForDomainRequest"></a>

### IPsForDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".IPsForDomainResult"></a>

### IPsForDomainResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| latestcapture | [CapSummary](#CapSummary) |  |  |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  |  |






<a name=".MetaForCaptureRequest"></a>

### MetaForCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| id | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".MetaForCaptureResult"></a>

### MetaForCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [Metadata](#Metadata) |  |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| resume | [string](#string) |  |  |






<a name=".MetaSearchRequest"></a>

### MetaSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".MetaSearchResult"></a>

### MetaSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [Metadata](#Metadata) |  |  |
| resume | [string](#string) |  |  |
| capsummary | [CapSummary](#CapSummary) |  |  |






<a name=".Metadata"></a>

### Metadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name=".RedirsByDestinationRequest"></a>

### RedirsByDestinationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".RedirsByDestinationResponse"></a>

### RedirsByDestinationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [string](#string) |  |  |
| destination | [string](#string) |  |  |
| latestcapture | [CapSummary](#CapSummary) |  |  |
| resume | [string](#string) |  |  |






<a name=".RedirsBySourceRequest"></a>

### RedirsBySourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".RedirsBySourceResponse"></a>

### RedirsBySourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [string](#string) |  |  |
| destination | [string](#string) |  |  |
| latestcapture | [CapSummary](#CapSummary) |  |  |
| resume | [string](#string) |  |  |






<a name=".RootLogSearchRequest"></a>

### RootLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".RootLogSearchResult"></a>

### RootLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| capsummary | [CapSummary](#CapSummary) |  |  |






<a name=".RootsForSatelliteRequest"></a>

### RootsForSatelliteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".RootsForSatelliteResult"></a>

### RootsForSatelliteResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| satellite | [string](#string) |  |  |
| capsummary | [CapSummary](#CapSummary) |  |  |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  |  |






<a name=".SatelliteLogSearchRequest"></a>

### SatelliteLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".SatelliteLogSearchResult"></a>

### SatelliteLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| satellite | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| capsummary | [CapSummary](#CapSummary) |  |  |






<a name=".SatellitesForRootRequest"></a>

### SatellitesForRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".SatellitesForRootResult"></a>

### SatellitesForRootResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| satellite | [string](#string) |  |  |
| capsummary | [CapSummary](#CapSummary) |  |  |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  |  |






<a name=".SubmitCaptureRequest"></a>

### SubmitCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| notification | [string](#string) |  |  |
| urls | [string](#string) | repeated |  |
| meta | [Metadata](#Metadata) | repeated |  |
| skiprecent | [int64](#int64) |  |  |






<a name=".SubmitCaptureResult"></a>

### SubmitCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| assets | [string](#string) |  |  |
| skipped | [bool](#bool) |  |  |
| id | [string](#string) |  |  |






<a name=".TempCaptureRequest"></a>

### TempCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notification | [string](#string) |  |  |
| urls | [string](#string) | repeated |  |
| meta | [Metadata](#Metadata) | repeated |  |






<a name=".TempCaptureResult"></a>

### TempCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| assets | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name=".URLLogSearchRequest"></a>

### URLLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  |  |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  |  |
| resume | [string](#string) |  |  |






<a name=".URLLogSearchResult"></a>

### URLLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| resume | [string](#string) |  |  |
| capsummary | [CapSummary](#CapSummary) |  |  |





 

 

 


<a name=".Netograph"></a>

### Netograph


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| TempCapture | [.TempCaptureRequest](#TempCaptureRequest) | [.TempCaptureResult](#TempCaptureResult) | Request a temporary capture. Temporary captures are not stored in a dataset, and the capture assets will be available for download for 24 hours before being deleted. |
| Datasets | [.DatasetsRequest](#DatasetsRequest) | [.Dataset](#Dataset) | List all datasets to which the authorizing account has access. This includes public datasets, which will be marked readonly. |
| SubmitCapture | [.SubmitCaptureRequest](#SubmitCaptureRequest) | [.SubmitCaptureResult](#SubmitCaptureResult) | Submit a capture request to a dataset. |
| CaptureInfo | [.CaptureInfoRequest](#CaptureInfoRequest) | [.CaptureInfoResult](#CaptureInfoResult) | Retrieve info for a specified capture by ID within a dataset. |
| CaptureLog | [.CaptureLogRequest](#CaptureLogRequest) | [.CaptureLogResult](#CaptureLogResult) | Retrieve the capture log for a dataset, in reverse chronological order. |
| DomainHistory | [.DomainHistoryRequest](#DomainHistoryRequest) | [.DomainHistoryResult](#DomainHistoryResult) | Retrieve the capture history for a specified domain in a dataset. The length of this history is capped at ~100. |
| DomainSearch | [.DomainSearchRequest](#DomainSearchRequest) | [.DomainSearchResult](#DomainSearchResult) | Retrieve the capture log for a dataset, in reverse chronological order. |
| DomainsForIP | [.DomainsForIPRequest](#DomainsForIPRequest) | [.DomainsForIPResult](#DomainsForIPResult) | Find all domains in the dataset associated with a given IP address. |
| IPHistory | [.IPHistoryRequest](#IPHistoryRequest) | [.IPHistoryResult](#IPHistoryResult) | Retrieve the capture history for a specified IP in a dataset. The length of this history is capped at ~100. |
| IPLogSearch | [.IPLogSearchRequest](#IPLogSearchRequest) | [.IPLogSearchResult](#IPLogSearchResult) | Search the dataset log for captures that contain a given IP. |
| IPSearch | [.IPSearchRequest](#IPSearchRequest) | [.IPSearchResult](#IPSearchResult) | Find all IPs in the dataset that match an address and integer netmask. |
| IPsForDomain | [.IPsForDomainRequest](#IPsForDomainRequest) | [.IPsForDomainResult](#IPsForDomainResult) | Find all IPs in a dataset associated with a given domain. |
| MetaForCapture | [.MetaForCaptureRequest](#MetaForCaptureRequest) | [.MetaForCaptureResult](#MetaForCaptureResult) | Get metadata associated with a specified capture within a dataset. |
| MetaSearch | [.MetaSearchRequest](#MetaSearchRequest) | [.MetaSearchResult](#MetaSearchResult) | Search the dataset log for captures matching a metadata query. |
| RedirsByDestination | [.RedirsByDestinationRequest](#RedirsByDestinationRequest) | [.RedirsByDestinationResponse](#RedirsByDestinationResponse) | Find all redirections in the dataset for a given destination domain query. |
| RedirsBySource | [.RedirsBySourceRequest](#RedirsBySourceRequest) | [.RedirsBySourceResponse](#RedirsBySourceResponse) | Find all redirections in the dataset for a given source domain query. |
| RootLogSearch | [.RootLogSearchRequest](#RootLogSearchRequest) | [.RootLogSearchResult](#RootLogSearchResult) | Search the dataset log for captures where any root domain matches a given query. |
| RootsForSatellite | [.RootsForSatelliteRequest](#RootsForSatelliteRequest) | [.RootsForSatelliteResult](#RootsForSatelliteResult) | Find all roots in a dataset that are associated with a given satellite query. |
| SatelliteLogSearch | [.SatelliteLogSearchRequest](#SatelliteLogSearchRequest) | [.SatelliteLogSearchResult](#SatelliteLogSearchResult) | Search the dataset log for captures where any satellite domain matches a given query. |
| SatellitesForRoot | [.SatellitesForRootRequest](#SatellitesForRootRequest) | [.SatellitesForRootResult](#SatellitesForRootResult) | Find all satellites in the dataset that are associated with a given root query. |
| URLLogSearch | [.URLLogSearchRequest](#URLLogSearchRequest) | [.URLLogSearchResult](#URLLogSearchResult) | Search the dataset log for captures where any root URL matches a given URL query. |

 



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

