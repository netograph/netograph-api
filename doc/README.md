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
    - [CaptureTextRequest](#io.netograph.dset.CaptureTextRequest)
    - [CaptureTextResult](#io.netograph.dset.CaptureTextResult)
    - [CaptureTextResult.Document](#io.netograph.dset.CaptureTextResult.Document)
    - [CaptureTextResult.Group](#io.netograph.dset.CaptureTextResult.Group)
    - [CaptureTextResult.Page](#io.netograph.dset.CaptureTextResult.Page)
    - [CaptureTextResult.Rectangle](#io.netograph.dset.CaptureTextResult.Rectangle)
    - [CaptureTextResult.TextNode](#io.netograph.dset.CaptureTextResult.TextNode)
    - [Cert](#io.netograph.dset.Cert)
    - [Cert.BasicConstraints](#io.netograph.dset.Cert.BasicConstraints)
    - [Cert.DistinguishedName](#io.netograph.dset.Cert.DistinguishedName)
    - [Cert.Extension](#io.netograph.dset.Cert.Extension)
    - [Cert.FingerprintsEntry](#io.netograph.dset.Cert.FingerprintsEntry)
    - [Cert.PublickeyEntry](#io.netograph.dset.Cert.PublickeyEntry)
    - [CertDomainSearchRequest](#io.netograph.dset.CertDomainSearchRequest)
    - [CertDomainSearchResult](#io.netograph.dset.CertDomainSearchResult)
    - [CertDomainStatsRequest](#io.netograph.dset.CertDomainStatsRequest)
    - [CertDomainStatsResult](#io.netograph.dset.CertDomainStatsResult)
    - [CertIPSearchRequest](#io.netograph.dset.CertIPSearchRequest)
    - [CertIPSearchResult](#io.netograph.dset.CertIPSearchResult)
    - [CertSearchRequest](#io.netograph.dset.CertSearchRequest)
    - [CertSearchResult](#io.netograph.dset.CertSearchResult)
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
    - [PoliciesForRootRequest](#io.netograph.dset.PoliciesForRootRequest)
    - [PoliciesForRootResult](#io.netograph.dset.PoliciesForRootResult)
    - [PolicyDomainCapturesRequest](#io.netograph.dset.PolicyDomainCapturesRequest)
    - [PolicyDomainCapturesResult](#io.netograph.dset.PolicyDomainCapturesResult)
    - [PolicyDomainStatsRequest](#io.netograph.dset.PolicyDomainStatsRequest)
    - [PolicyDomainStatsResult](#io.netograph.dset.PolicyDomainStatsResult)
    - [PolicyURLCapturesRequest](#io.netograph.dset.PolicyURLCapturesRequest)
    - [PolicyURLCapturesResult](#io.netograph.dset.PolicyURLCapturesResult)
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
A capture summary - this is the data that you will receive for any query that
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
| extended | [bool](#bool) |  |  |






<a name="io.netograph.dset.CapSummary.Root"></a>

### CapSummary.Root
Roots that resulted from loading a URL in the submitted plan. You can
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






<a name="io.netograph.dset.CaptureTextRequest"></a>

### CaptureTextRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| id | [string](#string) |  | The unique ID of the capture. |






<a name="io.netograph.dset.CaptureTextResult"></a>

### CaptureTextResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pages | [CaptureTextResult.Page](#io.netograph.dset.CaptureTextResult.Page) | repeated |  |
| groups | [CaptureTextResult.Group](#io.netograph.dset.CaptureTextResult.Group) | repeated |  |






<a name="io.netograph.dset.CaptureTextResult.Document"></a>

### CaptureTextResult.Document



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | The root URL for this document. |
| text | [CaptureTextResult.TextNode](#io.netograph.dset.CaptureTextResult.TextNode) | repeated | Amalgamated text nodes in this document. |






<a name="io.netograph.dset.CaptureTextResult.Group"></a>

### CaptureTextResult.Group



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of this group. |
| page | [int64](#int64) |  | The offset of the page this group applies to |
| document | [int64](#int64) |  | The offset of the document within the page. |
| ids | [int64](#int64) | repeated | A list of node IDs within the page included in this group. |






<a name="io.netograph.dset.CaptureTextResult.Page"></a>

### CaptureTextResult.Page



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| documents | [CaptureTextResult.Document](#io.netograph.dset.CaptureTextResult.Document) | repeated | Documents within a page include the root frame and any iframes. |






<a name="io.netograph.dset.CaptureTextResult.Rectangle"></a>

### CaptureTextResult.Rectangle



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| x | [float](#float) |  |  |
| y | [float](#float) |  |  |
| with | [float](#float) |  |  |
| height | [float](#float) |  |  |






<a name="io.netograph.dset.CaptureTextResult.TextNode"></a>

### CaptureTextResult.TextNode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | A unique ID for this text node. |
| nodes | [int64](#int64) | repeated | The underlying content nodes included in this text node. |
| bounds | [CaptureTextResult.Rectangle](#io.netograph.dset.CaptureTextResult.Rectangle) |  | The pixel bounds for this text node. |
| text | [string](#string) |  | The assembled text. |






<a name="io.netograph.dset.Cert"></a>

### Cert



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authoritykeyid | [string](#string) |  |  |
| basicconstraints | [Cert.BasicConstraints](#io.netograph.dset.Cert.BasicConstraints) |  | Basic Constraints extension |
| crldistributionpoints | [string](#string) | repeated | CRL Distribution Points extension |
| dnsnames | [string](#string) | repeated | Subject Alternate Names extension |
| emailaddresses | [string](#string) | repeated |  |
| extendedkeyusage | [string](#string) | repeated | ExtendedKeyUsage extension |
| extensions | [Cert.Extension](#io.netograph.dset.Cert.Extension) | repeated |  |
| fingerprints | [Cert.FingerprintsEntry](#io.netograph.dset.Cert.FingerprintsEntry) | repeated |  |
| ipaddresses | [string](#string) | repeated | Subject Alternate Names extension |
| issuer | [Cert.DistinguishedName](#io.netograph.dset.Cert.DistinguishedName) |  |  |
| issuingcerturl | [string](#string) | repeated | RFC 5280, 4.2.2.1 (Authority Information Access |
| keyusage | [string](#string) | repeated | Key Usage extension |
| notafter | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| notbefore | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| ocspserver | [string](#string) | repeated | RFC 5280, 4.2.2.1 (Authority Information Access |
| permitteddnsdomains | [string](#string) | repeated |  |
| permitteddnsdomainscritical | [bool](#bool) |  |  |
| publickey | [Cert.PublickeyEntry](#io.netograph.dset.Cert.PublickeyEntry) | repeated |  |
| publickeyalgorithm | [string](#string) |  |  |
| serialnumber | [string](#string) |  |  |
| signaturealgorithm | [string](#string) |  |  |
| subject | [Cert.DistinguishedName](#io.netograph.dset.Cert.DistinguishedName) |  |  |
| subjectkeyid | [string](#string) |  |  |
| version | [int64](#int64) |  |  |






<a name="io.netograph.dset.Cert.BasicConstraints"></a>

### Cert.BasicConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isca | [bool](#bool) |  |  |
| maxpathlen | [int64](#int64) |  |  |






<a name="io.netograph.dset.Cert.DistinguishedName"></a>

### Cert.DistinguishedName



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| commonname | [string](#string) |  |  |
| country | [string](#string) | repeated |  |
| locality | [string](#string) | repeated |  |
| organization | [string](#string) | repeated |  |
| organizationalunit | [string](#string) | repeated |  |
| postalcode | [string](#string) | repeated |  |
| province | [string](#string) | repeated |  |
| serialnumber | [string](#string) |  |  |
| streetaddress | [string](#string) | repeated |  |






<a name="io.netograph.dset.Cert.Extension"></a>

### Cert.Extension



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| critical | [bool](#bool) |  |  |
| oid | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="io.netograph.dset.Cert.FingerprintsEntry"></a>

### Cert.FingerprintsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="io.netograph.dset.Cert.PublickeyEntry"></a>

### Cert.PublickeyEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="io.netograph.dset.CertDomainSearchRequest"></a>

### CertDomainSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  | A domain query. This is a domain prefix, which will return results for all relevant subdomains. To search for an exact domain, prefix with &#34;$&#34; - e.g. &#34;$rt.com&#34;. |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.CertDomainSearchResult"></a>

### CertDomainSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| cert | [Cert](#io.netograph.dset.Cert) |  | A matching certificate. |
| chain | [string](#string) | repeated | Fingerprints for certificates in the chain. |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="io.netograph.dset.CertDomainStatsRequest"></a>

### CertDomainStatsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  | A domain query. This is a domain prefix, which will return results for all relevant subdomains. To search for an exact domain, prefix with &#34;$&#34; - e.g. &#34;$rt.com&#34;. |






<a name="io.netograph.dset.CertDomainStatsResult"></a>

### CertDomainStatsResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domains | [uint64](#uint64) |  | Number of matching subdomains |
| certs | [uint64](#uint64) |  | Number of certificates at all matching subdomains |
| certshere | [uint64](#uint64) |  | Number of certificates at this exact domain |






<a name="io.netograph.dset.CertIPSearchRequest"></a>

### CertIPSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| ip | [string](#string) |  |  |
| mask | [int32](#int32) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.CertIPSearchResult"></a>

### CertIPSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [string](#string) |  |  |
| domain | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| cert | [Cert](#io.netograph.dset.Cert) |  | The latest capture relevant to this result. |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="io.netograph.dset.CertSearchRequest"></a>

### CertSearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| field | [string](#string) |  |  |
| text | [string](#string) |  |  |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.CertSearchResult"></a>

### CertSearchResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| cert | [Cert](#io.netograph.dset.Cert) |  | The latest capture relevant to this result. |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






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
| id | [string](#string) |  | The ID of the capture |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |
| prefix | [string](#string) |  | An optional key prefix |






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






<a name="io.netograph.dset.PoliciesForRootRequest"></a>

### PoliciesForRootRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  |  |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the start of the range is the most recent time. If start is zero, it&#39;s taken to be the largest possible time value. |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the end of the range is the least recent time. If end is zero, it&#39;s taken to be the smallest possible time value. |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.PoliciesForRootResult"></a>

### PoliciesForRootResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rootdomain | [string](#string) |  |  |
| url | [string](#string) |  |  |
| time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| lastreference | [CapSummary](#io.netograph.dset.CapSummary) |  |  |
| lastcapture | [CapSummary](#io.netograph.dset.CapSummary) |  |  |
| resume | [string](#string) |  |  |






<a name="io.netograph.dset.PolicyDomainCapturesRequest"></a>

### PolicyDomainCapturesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  | A domain query. This is a domain prefix, which will return results for all relevant subdomains. To search for an exact domain, prefix with &#34;$&#34; - e.g. &#34;$rt.com&#34;. |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the start of the range is the most recent time. If start is zero, it&#39;s taken to be the largest possible time value. If a start or end time is specified, the query must be for an exact domain. |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The capture log is in reverse chronological order, so the end of the range is the least recent time. If end is zero, it&#39;s taken to be the smallest possible time value. If a start or end time is specified, the query must be for an exact domain. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.PolicyDomainCapturesResult"></a>

### PolicyDomainCapturesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  | The matching domain. |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |






<a name="io.netograph.dset.PolicyDomainStatsRequest"></a>

### PolicyDomainStatsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  | A domain query. This is a domain prefix, which will return results for all relevant subdomains. To search for an exact domain, prefix with &#34;$&#34; - e.g. &#34;$rt.com&#34;. |






<a name="io.netograph.dset.PolicyDomainStatsResult"></a>

### PolicyDomainStatsResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policydomains | [uint64](#uint64) |  | Number of matching domains with policies |
| subdomainpolicies | [uint64](#uint64) |  | Number of policies at all matching subdomains |
| policies | [uint64](#uint64) |  | Number of policies at the exact matching domain |
| subdomaincaptures | [uint64](#uint64) |  | Total number of policy captures at all matching subdomains |
| captures | [uint64](#uint64) |  | Number of policy captures at the exact matching domain |
| capturedomains | [uint64](#uint64) |  | Total number of subdomains with policy captures |






<a name="io.netograph.dset.PolicyURLCapturesRequest"></a>

### PolicyURLCapturesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataset | [string](#string) |  | The name of the dataset to query. |
| query | [string](#string) |  | A URL query |
| limit | [int64](#int64) |  | Limit the number of records that will be returned. |
| resume | [string](#string) |  | A resumption token, previously returned by an identical query. |






<a name="io.netograph.dset.PolicyURLCapturesResult"></a>

### PolicyURLCapturesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| resume | [string](#string) |  | A resumption token that can be passed to an identical query to resume results. |
| capsummary | [CapSummary](#io.netograph.dset.CapSummary) |  | A capture summary for this result. |






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
| meta | [Metadata](#io.netograph.dset.Metadata) | repeated | Metadata to associate with the capture |
| skiprecent | [int64](#int64) |  | Skip capture if we&#39;ve seen this exact URL within a specified number of seconds |
| zone | [string](#string) |  | Capture zone - &#34;us&#34; or &#34;eu&#34;. If unspecified, we choose based on availability. |
| extended | [bool](#bool) |  | Extended capture includes full-page screenshot and page content formats |






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
| CaptureLog | [CaptureLogRequest](#io.netograph.dset.CaptureLogRequest) | [CaptureLogResult](#io.netograph.dset.CaptureLogResult) stream | Retrieve the capture log for a dataset, in reverse chronological order. |
| CertDomainSearch | [CertDomainSearchRequest](#io.netograph.dset.CertDomainSearchRequest) | [CertDomainSearchResult](#io.netograph.dset.CertDomainSearchResult) stream | Retrieve certificates for a specified domain query. |
| CertDomainStats | [CertDomainStatsRequest](#io.netograph.dset.CertDomainStatsRequest) | [CertDomainStatsResult](#io.netograph.dset.CertDomainStatsResult) | Retrieve certificate statistics for a specified domain query. |
| CertIPSearch | [CertIPSearchRequest](#io.netograph.dset.CertIPSearchRequest) | [CertIPSearchResult](#io.netograph.dset.CertIPSearchResult) stream | Retrieve certificates for a specified IP query. |
| CertSearch | [CertSearchRequest](#io.netograph.dset.CertSearchRequest) | [CertSearchResult](#io.netograph.dset.CertSearchResult) stream | Retrieve certificates based on a field query. |
| DomainHistory | [DomainHistoryRequest](#io.netograph.dset.DomainHistoryRequest) | [DomainHistoryResult](#io.netograph.dset.DomainHistoryResult) stream | Retrieve the capture history for a specified domain. The length of this history is capped at ~100. |
| DomainSearch | [DomainSearchRequest](#io.netograph.dset.DomainSearchRequest) | [DomainSearchResult](#io.netograph.dset.DomainSearchResult) stream | Retrieve the capture log for a specified domain in a dataset. |
| DomainsForIP | [DomainsForIPRequest](#io.netograph.dset.DomainsForIPRequest) | [DomainsForIPResult](#io.netograph.dset.DomainsForIPResult) stream | Find all domains in the dataset associated with a given IP address. |
| IPHistory | [IPHistoryRequest](#io.netograph.dset.IPHistoryRequest) | [IPHistoryResult](#io.netograph.dset.IPHistoryResult) stream | Retrieve the capture history for a specified IP in a dataset. The length of this history is capped at ~100. |
| IPLogSearch | [IPLogSearchRequest](#io.netograph.dset.IPLogSearchRequest) | [IPLogSearchResult](#io.netograph.dset.IPLogSearchResult) stream | Search the dataset log for captures that contain a given IP. |
| IPSearch | [IPSearchRequest](#io.netograph.dset.IPSearchRequest) | [IPSearchResult](#io.netograph.dset.IPSearchResult) stream | Find all IPs in the dataset that match an address and integer netmask. |
| IPsForDomain | [IPsForDomainRequest](#io.netograph.dset.IPsForDomainRequest) | [IPsForDomainResult](#io.netograph.dset.IPsForDomainResult) stream | Find all IPs in a dataset associated with a given domain. |
| MetaForCapture | [MetaForCaptureRequest](#io.netograph.dset.MetaForCaptureRequest) | [MetaForCaptureResult](#io.netograph.dset.MetaForCaptureResult) stream | Get metadata associated with a specified capture within a dataset. |
| MetaSearch | [MetaSearchRequest](#io.netograph.dset.MetaSearchRequest) | [MetaSearchResult](#io.netograph.dset.MetaSearchResult) stream | Search the dataset log for captures matching a metadata query. |
| PoliciesForRoot | [PoliciesForRootRequest](#io.netograph.dset.PoliciesForRootRequest) | [PoliciesForRootResult](#io.netograph.dset.PoliciesForRootResult) stream | Find all policies for a specified domain query. |
| PolicyDomainCaptures | [PolicyDomainCapturesRequest](#io.netograph.dset.PolicyDomainCapturesRequest) | [PolicyDomainCapturesResult](#io.netograph.dset.PolicyDomainCapturesResult) stream | Retrieve the policy capture log for a domain query, in reverse chronological order. |
| PolicyDomainStats | [PolicyDomainStatsRequest](#io.netograph.dset.PolicyDomainStatsRequest) | [PolicyDomainStatsResult](#io.netograph.dset.PolicyDomainStatsResult) | Retrieve statistics for a policy domain query. |
| PolicyURLCaptures | [PolicyURLCapturesRequest](#io.netograph.dset.PolicyURLCapturesRequest) | [PolicyURLCapturesResult](#io.netograph.dset.PolicyURLCapturesResult) stream | Retrieve the policy capture log for specific policy URL, in reverse chronological order. |
| RedirsByDestination | [RedirsByDestinationRequest](#io.netograph.dset.RedirsByDestinationRequest) | [RedirsByDestinationResult](#io.netograph.dset.RedirsByDestinationResult) stream | Find all redirections in the dataset for a given destination domain query. |
| RedirsBySource | [RedirsBySourceRequest](#io.netograph.dset.RedirsBySourceRequest) | [RedirsBySourceResult](#io.netograph.dset.RedirsBySourceResult) stream | Find all redirections in the dataset for a given source domain query. |
| RootLogSearch | [RootLogSearchRequest](#io.netograph.dset.RootLogSearchRequest) | [RootLogSearchResult](#io.netograph.dset.RootLogSearchResult) stream | Search the dataset log for captures where any root domain matches a given query. |
| RootsForSatellite | [RootsForSatelliteRequest](#io.netograph.dset.RootsForSatelliteRequest) | [RootsForSatelliteResult](#io.netograph.dset.RootsForSatelliteResult) stream | Find all roots in a dataset that are associated with a given satellite query. |
| SatelliteLogSearch | [SatelliteLogSearchRequest](#io.netograph.dset.SatelliteLogSearchRequest) | [SatelliteLogSearchResult](#io.netograph.dset.SatelliteLogSearchResult) stream | Search the dataset log for captures where any satellite domain matches a given query. |
| SatellitesForRoot | [SatellitesForRootRequest](#io.netograph.dset.SatellitesForRootRequest) | [SatellitesForRootResult](#io.netograph.dset.SatellitesForRootResult) stream | Find all satellites in the dataset that are associated with a given root query. |
| URLLogSearch | [URLLogSearchRequest](#io.netograph.dset.URLLogSearchRequest) | [URLLogSearchResult](#io.netograph.dset.URLLogSearchResult) stream | Search the dataset log for captures where any root URL matches a given URL query. |

 



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
| extended | [bool](#bool) |  | Extended capture includes full-page screenshot and page content formats |






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
| Datasets | [DatasetsRequest](#io.netograph.user.DatasetsRequest) | [Dataset](#io.netograph.user.Dataset) stream | List all datasets to which the authorizing account has access. This includes public datasets, which will be marked readonly. |

 



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

