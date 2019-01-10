# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/ngapi/dsetapi/dset.proto](#proto/ngapi/dsetapi/dset.proto)
    - [CapSummary](#io.netograph.dset.CapSummary)
    - [CapSummary.Plan](#io.netograph.dset.CapSummary.Plan)
    - [CapSummary.Root](#io.netograph.dset.CapSummary.Root)
    - [CapSummary.Stats](#io.netograph.dset.CapSummary.Stats)
    - [CaptureInfoRequest](#io.netograph.dset.CaptureInfoRequest)
    - [CaptureInfoResult](#io.netograph.dset.CaptureInfoResult)
    - [CaptureLogRequest](#io.netograph.dset.CaptureLogRequest)
    - [CaptureLogResult](#io.netograph.dset.CaptureLogResult)
    - [DomainHistoryRequest](#io.netograph.dset.DomainHistoryRequest)
    - [DomainHistoryResult](#io.netograph.dset.DomainHistoryResult)
    - [DomainSearchRequest](#io.netograph.dset.DomainSearchRequest)
    - [DomainSearchResult](#io.netograph.dset.DomainSearchResult)
    - [DomainsForIPRequest](#io.netograph.dset.DomainsForIPRequest)
    - [DomainsForIPResult](#io.netograph.dset.DomainsForIPResult)
    - [IPHistoryRequest](#io.netograph.dset.IPHistoryRequest)
    - [IPHistoryResult](#io.netograph.dset.IPHistoryResult)
    - [IPLogSearchRequest](#io.netograph.dset.IPLogSearchRequest)
    - [IPLogSearchResult](#io.netograph.dset.IPLogSearchResult)
    - [IPSearchRequest](#io.netograph.dset.IPSearchRequest)
    - [IPSearchResult](#io.netograph.dset.IPSearchResult)
    - [IPsForDomainRequest](#io.netograph.dset.IPsForDomainRequest)
    - [IPsForDomainResult](#io.netograph.dset.IPsForDomainResult)
    - [MetaForCaptureRequest](#io.netograph.dset.MetaForCaptureRequest)
    - [MetaForCaptureResult](#io.netograph.dset.MetaForCaptureResult)
    - [MetaSearchRequest](#io.netograph.dset.MetaSearchRequest)
    - [MetaSearchResult](#io.netograph.dset.MetaSearchResult)
    - [Metadata](#io.netograph.dset.Metadata)
    - [PoliciesForDomainRequest](#io.netograph.dset.PoliciesForDomainRequest)
    - [PoliciesForDomainResult](#io.netograph.dset.PoliciesForDomainResult)
    - [RedirsByDestinationRequest](#io.netograph.dset.RedirsByDestinationRequest)
    - [RedirsByDestinationResult](#io.netograph.dset.RedirsByDestinationResult)
    - [RedirsBySourceRequest](#io.netograph.dset.RedirsBySourceRequest)
    - [RedirsBySourceResult](#io.netograph.dset.RedirsBySourceResult)
    - [RootLogSearchRequest](#io.netograph.dset.RootLogSearchRequest)
    - [RootLogSearchResult](#io.netograph.dset.RootLogSearchResult)
    - [RootsForSatelliteRequest](#io.netograph.dset.RootsForSatelliteRequest)
    - [RootsForSatelliteResult](#io.netograph.dset.RootsForSatelliteResult)
    - [SatelliteLogSearchRequest](#io.netograph.dset.SatelliteLogSearchRequest)
    - [SatelliteLogSearchResult](#io.netograph.dset.SatelliteLogSearchResult)
    - [SatellitesForRootRequest](#io.netograph.dset.SatellitesForRootRequest)
    - [SatellitesForRootResult](#io.netograph.dset.SatellitesForRootResult)
    - [SubmitCaptureRequest](#io.netograph.dset.SubmitCaptureRequest)
    - [SubmitCaptureResult](#io.netograph.dset.SubmitCaptureResult)
    - [URLLogSearchRequest](#io.netograph.dset.URLLogSearchRequest)
    - [URLLogSearchResult](#io.netograph.dset.URLLogSearchResult)
  
  
  
    - [Dset](#io.netograph.dset.Dset)
  

- [proto/ngapi/userapi/user.proto](#proto/ngapi/userapi/user.proto)
    - [Dataset](#io.netograph.user.Dataset)
    - [DatasetsRequest](#io.netograph.user.DatasetsRequest)
    - [Metadata](#io.netograph.user.Metadata)
    - [TempCaptureRequest](#io.netograph.user.TempCaptureRequest)
    - [TempCaptureResult](#io.netograph.user.TempCaptureResult)
  
  
  
    - [User](#io.netograph.user.User)
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/ngapi/dsetapi/dset.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/ngapi/dsetapi/dset.proto



<a name="io.netograph.dset.CapSummary"></a>

### CapSummary
A capture summary - this is the data that you will recieve for any query that
returns captures.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | The globally unique ID for this capture. |
| assets | [string](#string) |  | The URL under which this capture&#39;s static assets can be found. |
| roots | [CapSummary.Root](#io.netograph.dset.CapSummary.Root) | repeated | The roots for each URL in the capture plan. |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The time of capture. |
| plan | [CapSummary.Plan](#io.netograph.dset.CapSummary.Plan) |  | The capture plan. |
| stats | [CapSummary.Stats](#io.netograph.dset.CapSummary.Stats) |  | Some basic statistics on the resulting capture. |






<a name="io.netograph.dset.CapSummary.Plan"></a>

### CapSummary.Plan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| urls | [string](#string) | repeated | URLs submitted for capture. |






<a name="io.netograph.dset.CapSummary.Root"></a>

### CapSummary.Root
Roots that resulted from loading a URL in the the submitted plan. You can
think of this as the trajectory of redirections followed to reach the
final resting URL.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| urls | [string](#string) | repeated |  |






<a name="io.netograph.dset.CapSummary.Stats"></a>

### CapSummary.Stats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flows | [int32](#int32) |  | The number of flows. |
| websockets | [int32](#int32) |  | The number of websocket connections. |
| hosts | [int32](#int32) |  | The number of distinct hosts. |






<a name="io.netograph.dset.CaptureInfoRequest"></a>

### CaptureInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| id | [string](#string) |  | The unique ID of the capture. |






<a name="io.netograph.dset.CaptureInfoResult"></a>

### CaptureInfoResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |






<a name="io.netograph.dset.CaptureLogRequest"></a>

### CaptureLogRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the start of the range is the most recent time. If start is zero, it&#39;s taken to be the largest possible time value. |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the end of the range is the least recent time. If end is zero, it&#39;s taken to be the smallest possible time value. |






<a name="io.netograph.dset.CaptureLogResult"></a>

### CaptureLogResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.DomainHistoryRequest"></a>

### DomainHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| domain | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |






<a name="io.netograph.dset.DomainHistoryResult"></a>

### DomainHistoryResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |






<a name="io.netograph.dset.DomainSearchRequest"></a>

### DomainSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  | A domain query. This is a domain prefix, which will return results for all relevant subdomains. To search for an exact domain, prefix with &#34;$&#34; - e.g. &#34;$rt.com&#34;. |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.DomainSearchResult"></a>

### DomainSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| latestcapture | [CapSummary](#io.netograph.dset.CapSummary) |  | The latest capture relevant to this result. |






<a name="io.netograph.dset.DomainsForIPRequest"></a>

### DomainsForIPRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.DomainsForIPResult"></a>

### DomainsForIPResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.dset.CapSummary) |  | The latest capture relevant to this result. |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.IPHistoryRequest"></a>

### IPHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |






<a name="io.netograph.dset.IPHistoryResult"></a>

### IPHistoryResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |






<a name="io.netograph.dset.IPLogSearchRequest"></a>

### IPLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| ip | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.IPLogSearchResult"></a>

### IPLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |
| ip | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.IPSearchRequest"></a>

### IPSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| ip | [string](#string) |  |  |
| mask | [int32](#int32) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.IPSearchResult"></a>

### IPSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| latestcapture | [CapSummary](#io.netograph.dset.CapSummary) |  | The latest capture relevant to this result. |






<a name="io.netograph.dset.IPsForDomainRequest"></a>

### IPsForDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.IPsForDomainResult"></a>

### IPsForDomainResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.dset.CapSummary) |  | The latest capture relevant to this result. |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.MetaForCaptureRequest"></a>

### MetaForCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| id | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.MetaForCaptureResult"></a>

### MetaForCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [Metadata](#io.netograph.dset.Metadata) |  |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.MetaSearchRequest"></a>

### MetaSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.MetaSearchResult"></a>

### MetaSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meta | [Metadata](#io.netograph.dset.Metadata) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |






<a name="io.netograph.dset.Metadata"></a>

### Metadata
Metadata is arbitrary information associated with a capture.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The key for this metadata pair. Keys do not have to be unique. |
| value | [string](#string) |  | The value for this metadata pair. |






<a name="io.netograph.dset.PoliciesForDomainRequest"></a>

### PoliciesForDomainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.PoliciesForDomainResult"></a>

### PoliciesForDomainResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| url | [string](#string) |  |  |
| root | [string](#string) |  |  |
| policyid | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | The most recent capture on which we observed this policy |
| policycapsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | The most recent capture for the policy itself |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.RedirsByDestinationRequest"></a>

### RedirsByDestinationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.RedirsByDestinationResult"></a>

### RedirsByDestinationResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [string](#string) |  |  |
| destination | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.dset.CapSummary) |  | The latest capture relevant to this result. |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.RedirsBySourceRequest"></a>

### RedirsBySourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.RedirsBySourceResult"></a>

### RedirsBySourceResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [string](#string) |  |  |
| destination | [string](#string) |  |  |
| latestcapture | [CapSummary](#io.netograph.dset.CapSummary) |  | The latest capture relevant to this result. |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.RootLogSearchRequest"></a>

### RootLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.RootLogSearchResult"></a>

### RootLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |






<a name="io.netograph.dset.RootsForSatelliteRequest"></a>

### RootsForSatelliteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.RootsForSatelliteResult"></a>

### RootsForSatelliteResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| satellite | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.SatelliteLogSearchRequest"></a>

### SatelliteLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.SatelliteLogSearchResult"></a>

### SatelliteLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| satellite | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |






<a name="io.netograph.dset.SatellitesForRootRequest"></a>

### SatellitesForRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.SatellitesForRootResult"></a>

### SatellitesForRootResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [string](#string) |  |  |
| satellite | [string](#string) |  |  |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |
| associations | [string](#string) | repeated |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.SubmitCaptureRequest"></a>

### SubmitCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| notification | [string](#string) |  | A URL to which a notification will be posted when capture completes |
| urls | [string](#string) | repeated | A sequence of URLs to visit in order |
| meta | [Metadata](#io.netograph.dset.Metadata) | repeated | Metadata to associate with the apture |
| skiprecent | [int64](#int64) |  | Skip capture if we&#39;ve seen this exact URL within a specified number of seconds |
| zone | [string](#string) |  | Capture zone - &#34;us&#34; or &#34;eu&#34;. If unspecified, we choose based on availability. |






<a name="io.netograph.dset.SubmitCaptureResult"></a>

### SubmitCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| assets | [string](#string) |  | A URL under which capture assets may be found |
| skipped | [bool](#bool) |  | Was capture skipped due to a skiprecent specification? |
| id | [string](#string) |  | The capture ID |






<a name="io.netograph.dset.URLLogSearchRequest"></a>

### URLLogSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.URLLogSearchResult"></a>

### URLLogSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |





 

 

 


<a name="io.netograph.dset.Dset"></a>

### Dset
Methods that operate on an individual dataset, either public or private.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SubmitCapture | [SubmitCaptureRequest](#io.netograph.dset.SubmitCaptureRequest) | [SubmitCaptureResult](#io.netograph.dset.SubmitCaptureResult) | Submit a capture request to a dataset. |
| CaptureInfo | [CaptureInfoRequest](#io.netograph.dset.CaptureInfoRequest) | [CaptureInfoResult](#io.netograph.dset.CaptureInfoResult) | Retrieve info for a specified capture by ID within a dataset. |
| CaptureLog | [CaptureLogRequest](#io.netograph.dset.CaptureLogRequest) | [CaptureLogResult](#io.netograph.dset.CaptureLogResult) | Retrieve the capture log for a dataset, in reverse chronological order. |
| DomainHistory | [DomainHistoryRequest](#io.netograph.dset.DomainHistoryRequest) | [DomainHistoryResult](#io.netograph.dset.DomainHistoryResult) | Retrieve the capture history for a specified domain in a dataset. The length of this history is capped at ~100. |
| DomainSearch | [DomainSearchRequest](#io.netograph.dset.DomainSearchRequest) | [DomainSearchResult](#io.netograph.dset.DomainSearchResult) | Retrieve the capture log for a specified domain in a dataset. |
| DomainsForIP | [DomainsForIPRequest](#io.netograph.dset.DomainsForIPRequest) | [DomainsForIPResult](#io.netograph.dset.DomainsForIPResult) | Find all domains in the dataset associated with a given IP address. |
| IPHistory | [IPHistoryRequest](#io.netograph.dset.IPHistoryRequest) | [IPHistoryResult](#io.netograph.dset.IPHistoryResult) | Retrieve the capture history for a specified IP in a dataset. The length of this history is capped at ~100. |
| IPLogSearch | [IPLogSearchRequest](#io.netograph.dset.IPLogSearchRequest) | [IPLogSearchResult](#io.netograph.dset.IPLogSearchResult) | Search the dataset log for captures that contain a given IP. |
| IPSearch | [IPSearchRequest](#io.netograph.dset.IPSearchRequest) | [IPSearchResult](#io.netograph.dset.IPSearchResult) | Find all IPs in the dataset that match an address and integer netmask. |
| IPsForDomain | [IPsForDomainRequest](#io.netograph.dset.IPsForDomainRequest) | [IPsForDomainResult](#io.netograph.dset.IPsForDomainResult) | Find all IPs in a dataset associated with a given domain. |
| MetaForCapture | [MetaForCaptureRequest](#io.netograph.dset.MetaForCaptureRequest) | [MetaForCaptureResult](#io.netograph.dset.MetaForCaptureResult) | Get metadata associated with a specified capture within a dataset. |
| MetaSearch | [MetaSearchRequest](#io.netograph.dset.MetaSearchRequest) | [MetaSearchResult](#io.netograph.dset.MetaSearchResult) | Search the dataset log for captures matching a metadata query. |
| PoliciesForDomain | [PoliciesForDomainRequest](#io.netograph.dset.PoliciesForDomainRequest) | [PoliciesForDomainResult](#io.netograph.dset.PoliciesForDomainResult) | Find all policies for a specified domain query. |
| RedirsByDestination | [RedirsByDestinationRequest](#io.netograph.dset.RedirsByDestinationRequest) | [RedirsByDestinationResult](#io.netograph.dset.RedirsByDestinationResult) | Find all redirections in the dataset for a given destination domain query. |
| RedirsBySource | [RedirsBySourceRequest](#io.netograph.dset.RedirsBySourceRequest) | [RedirsBySourceResult](#io.netograph.dset.RedirsBySourceResult) | Find all redirections in the dataset for a given source domain query. |
| RootLogSearch | [RootLogSearchRequest](#io.netograph.dset.RootLogSearchRequest) | [RootLogSearchResult](#io.netograph.dset.RootLogSearchResult) | Search the dataset log for captures where any root domain matches a given query. |
| RootsForSatellite | [RootsForSatelliteRequest](#io.netograph.dset.RootsForSatelliteRequest) | [RootsForSatelliteResult](#io.netograph.dset.RootsForSatelliteResult) | Find all roots in a dataset that are associated with a given satellite query. |
| SatelliteLogSearch | [SatelliteLogSearchRequest](#io.netograph.dset.SatelliteLogSearchRequest) | [SatelliteLogSearchResult](#io.netograph.dset.SatelliteLogSearchResult) | Search the dataset log for captures where any satellite domain matches a given query. |
| SatellitesForRoot | [SatellitesForRootRequest](#io.netograph.dset.SatellitesForRootRequest) | [SatellitesForRootResult](#io.netograph.dset.SatellitesForRootResult) | Find all satellites in the dataset that are associated with a given root query. |
| URLLogSearch | [URLLogSearchRequest](#io.netograph.dset.URLLogSearchRequest) | [URLLogSearchResult](#io.netograph.dset.URLLogSearchResult) | Search the dataset log for captures where any root URL matches a given URL query. |

 



<a name="proto/ngapi/userapi/user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/ngapi/userapi/user.proto



<a name="io.netograph.user.Dataset"></a>

### Dataset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the dataset. |
| description | [string](#string) |  | The dataset text description. |
| urlbase | [string](#string) |  | The base URL under which this dataset&#39;s assets are exposed. |
| deleted | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Non-zero if the dataset has been deleted. |
| readonly | [bool](#bool) |  | Is the current user restricted to readonly access? If so, capture submission and other write operations will be denied. |






<a name="io.netograph.user.DatasetsRequest"></a>

### DatasetsRequest







<a name="io.netograph.user.Metadata"></a>

### Metadata
Metadata is arbitrary information associated with a capture.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | The key for this metadata pair. Keys do not have to be unique. |
| value | [string](#string) |  | The value for this metadata pair. |






<a name="io.netograph.user.TempCaptureRequest"></a>

### TempCaptureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notification | [string](#string) |  |  |
| urls | [string](#string) | repeated |  |
| meta | [Metadata](#io.netograph.user.Metadata) | repeated |  |
| zone | [string](#string) |  |  |






<a name="io.netograph.user.TempCaptureResult"></a>

### TempCaptureResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| assets | [string](#string) |  |  |
| id | [string](#string) |  |  |





 

 

 


<a name="io.netograph.user.User"></a>

### User
Methods that operate at the level of the user account.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| TempCapture | [TempCaptureRequest](#io.netograph.user.TempCaptureRequest) | [TempCaptureResult](#io.netograph.user.TempCaptureResult) | Request a temporary capture. Temporary captures are not stored in a dataset, and the capture assets will be available for download for 24 hours before being deleted. |
| Datasets | [DatasetsRequest](#io.netograph.user.DatasetsRequest) | [Dataset](#io.netograph.user.Dataset) | List all datasets to which the authorizing account has access. This includes public datasets, which will be marked readonly. |

 



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

