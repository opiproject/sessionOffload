# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [openoffload.proto](#openoffload-proto)
    - [ActionParameters](#openoffload-v2-ActionParameters)
    - [AddSessionResponse](#openoffload-v2-AddSessionResponse)
    - [SessionId](#openoffload-v2-SessionId)
    - [SessionRequest](#openoffload-v2-SessionRequest)
    - [SessionRequestArgs](#openoffload-v2-SessionRequestArgs)
    - [SessionResponse](#openoffload-v2-SessionResponse)
    - [SessionResponseError](#openoffload-v2-SessionResponseError)
    - [SessionResponses](#openoffload-v2-SessionResponses)
    - [Uuid](#openoffload-v2-Uuid)
  
    - [ActionType](#openoffload-v2-ActionType)
    - [AddSessionStatus](#openoffload-v2-AddSessionStatus)
    - [IpVersion](#openoffload-v2-IpVersion)
    - [ProtocolId](#openoffload-v2-ProtocolId)
    - [RequestStatus](#openoffload-v2-RequestStatus)
    - [SessionCloseCode](#openoffload-v2-SessionCloseCode)
    - [SessionState](#openoffload-v2-SessionState)
  
    - [SessionTable](#openoffload-v2-SessionTable)
  
- [tunneloffload.proto](#tunneloffload-proto)
    - [CapabilityRequest](#tunneloffload-v2-CapabilityRequest)
    - [CapabilityResponse](#tunneloffload-v2-CapabilityResponse)
    - [CapabilityResponse.GeneveCapabilities](#tunneloffload-v2-CapabilityResponse-GeneveCapabilities)
    - [CapabilityResponse.IPSecCapabilities](#tunneloffload-v2-CapabilityResponse-IPSecCapabilities)
    - [CapabilityResponse.MatchCapabilities](#tunneloffload-v2-CapabilityResponse-MatchCapabilities)
    - [Counters](#tunneloffload-v2-Counters)
    - [CreateIpTunnelResponse](#tunneloffload-v2-CreateIpTunnelResponse)
    - [CreateIpTunnelResponses](#tunneloffload-v2-CreateIpTunnelResponses)
    - [Error](#tunneloffload-v2-Error)
    - [Geneve](#tunneloffload-v2-Geneve)
    - [GeneveDecap](#tunneloffload-v2-GeneveDecap)
    - [GeneveEncap](#tunneloffload-v2-GeneveEncap)
    - [GeneveOption](#tunneloffload-v2-GeneveOption)
    - [IPSecDec](#tunneloffload-v2-IPSecDec)
    - [IPSecEnc](#tunneloffload-v2-IPSecEnc)
    - [IPSecSAParams](#tunneloffload-v2-IPSecSAParams)
    - [IPSecTunnel](#tunneloffload-v2-IPSecTunnel)
    - [IPV4Match](#tunneloffload-v2-IPV4Match)
    - [IPV4Pair](#tunneloffload-v2-IPV4Pair)
    - [IPV6Match](#tunneloffload-v2-IPV6Match)
    - [IPV6Pair](#tunneloffload-v2-IPV6Pair)
    - [IpTunnelRequest](#tunneloffload-v2-IpTunnelRequest)
    - [IpTunnelResponse](#tunneloffload-v2-IpTunnelResponse)
    - [IpTunnelResponses](#tunneloffload-v2-IpTunnelResponses)
    - [IpTunnelStatsResponse](#tunneloffload-v2-IpTunnelStatsResponse)
    - [IpTunnelStatsResponses](#tunneloffload-v2-IpTunnelStatsResponses)
    - [MacPair](#tunneloffload-v2-MacPair)
    - [MatchCriteria](#tunneloffload-v2-MatchCriteria)
    - [MatchCriteria.GeneveMatch](#tunneloffload-v2-MatchCriteria-GeneveMatch)
    - [MatchCriteria.IPSecMatch](#tunneloffload-v2-MatchCriteria-IPSecMatch)
    - [MatchCriteria.VXLanMatch](#tunneloffload-v2-MatchCriteria-VXLanMatch)
    - [Nat](#tunneloffload-v2-Nat)
    - [TunnelAdditionError](#tunneloffload-v2-TunnelAdditionError)
    - [TunnelId](#tunneloffload-v2-TunnelId)
    - [TunnelRequestArgs](#tunneloffload-v2-TunnelRequestArgs)
  
    - [Action](#tunneloffload-v2-Action)
    - [AddTunnelStatus](#tunneloffload-v2-AddTunnelStatus)
    - [EncType](#tunneloffload-v2-EncType)
    - [GeneveError](#tunneloffload-v2-GeneveError)
    - [IPSecError](#tunneloffload-v2-IPSecError)
    - [IPSecTunnelType](#tunneloffload-v2-IPSecTunnelType)
    - [MatchError](#tunneloffload-v2-MatchError)
    - [Operation](#tunneloffload-v2-Operation)
    - [TunnelError](#tunneloffload-v2-TunnelError)
  
    - [IpTunnelService](#tunneloffload-v2-IpTunnelService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="openoffload-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## openoffload.proto



<a name="openoffload-v2-ActionParameters"></a>

### ActionParameters
MIRROR and SNOOP require an actionNextHop
DROP and FORWARD do not have an actionNextHop
The IPV4 nextHop definition maps to the V4 struct returned by inet_pton which is a uint32_t.
The IPV6 nextHop definition maps to the V6 struct returned by inet_ptoN which is a uint8_t s6_addr[16]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action_type | [ActionType](#openoffload-v2-ActionType) |  |  |
| action_next_hop | [uint32](#uint32) |  |  |
| action_next_hop_v6 | [bytes](#bytes) |  |  |






<a name="openoffload-v2-AddSessionResponse"></a>

### AddSessionResponse
In v1apha4 the errorstatus was added to act as a bitmask
of errors for each of the sesssions sent in a stream (max 64).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_status | [AddSessionStatus](#openoffload-v2-AddSessionStatus) |  |  |
| error_status | [uint64](#uint64) |  |  |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| response_error | [SessionResponseError](#openoffload-v2-SessionResponseError) | repeated |  |






<a name="openoffload-v2-SessionId"></a>

### SessionId
should the Application assign the sessionID on AddSession and avoid conflicts
or have the applications have a mechanism to avoid duplicate sessionIDs across 
applications since there will be many applications instances to 1 switch


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [uint64](#uint64) |  |  |






<a name="openoffload-v2-SessionRequest"></a>

### SessionRequest
SessionId is generated by client and passed in via gRPC call
The IPV4 definition maps to the V4 struct returned by inet_pton which is a uint32_t.
The IPV6 definition maps to the V6 struct returned by inet_ptoN which is a uint8_t s6_addr[16]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [uint64](#uint64) |  |  |
| in_lif | [int32](#int32) |  |  |
| out_lif | [int32](#int32) |  |  |
| ip_version | [IpVersion](#openoffload-v2-IpVersion) |  |  |
| source_ip | [uint32](#uint32) |  |  |
| source_ipv6 | [bytes](#bytes) |  |  |
| source_port | [uint32](#uint32) |  |  |
| destination_ip | [uint32](#uint32) |  |  |
| destination_ipv6 | [bytes](#bytes) |  |  |
| destination_port | [uint32](#uint32) |  |  |
| protocol_id | [ProtocolId](#openoffload-v2-ProtocolId) |  |  |
| action | [ActionParameters](#openoffload-v2-ActionParameters) |  |  |
| cache_timeout | [uint32](#uint32) |  |  |






<a name="openoffload-v2-SessionRequestArgs"></a>

### SessionRequestArgs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_size | [uint32](#uint32) |  | pageSize = 0 will turn off paging does paging make sense for a stream ? the client should read/process each event on the stream anyway. |
| page | [uint32](#uint32) |  |  |
| start_session | [uint64](#uint64) |  | what other arguments make sense for retrieving or filtering streams |






<a name="openoffload-v2-SessionResponse"></a>

### SessionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [uint64](#uint64) |  |  |
| in_packets | [uint64](#uint64) |  |  |
| out_packets | [uint64](#uint64) |  |  |
| in_bytes | [uint64](#uint64) |  |  |
| out_bytes | [uint64](#uint64) |  |  |
| session_state | [SessionState](#openoffload-v2-SessionState) |  |  |
| session_close_code | [SessionCloseCode](#openoffload-v2-SessionCloseCode) |  |  |
| request_status | [RequestStatus](#openoffload-v2-RequestStatus) |  |  |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="openoffload-v2-SessionResponseError"></a>

### SessionResponseError



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [uint64](#uint64) |  |  |
| error_status | [int32](#int32) |  |  |






<a name="openoffload-v2-SessionResponses"></a>

### SessionResponses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_info | [SessionResponse](#openoffload-v2-SessionResponse) | repeated |  |
| next_key | [uint64](#uint64) |  |  |






<a name="openoffload-v2-Uuid"></a>

### Uuid
Uuid for Session IDs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |





 


<a name="openoffload-v2-ActionType"></a>

### ActionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| _DROP | 0 |  |
| _FORWARD | 1 |  |
| _MIRROR | 2 |  |
| _SNOOP | 3 |  |



<a name="openoffload-v2-AddSessionStatus"></a>

### AddSessionStatus
Errors for adding a session
If all sessions are successful inserted return _ACCEPTED

If check of session capacity in offload device is insufficient to add all sessions 
do not insert any sessions and return  _REJECTED_SESSION_TABLE_FULL. It is the 
responsibility of the client to re-try

If the server is unavailable for some other reason then return _REJECTED_SESSION_TABLE_UNAVAILABLE.
It is the  responsibility of the client to re-try

All other errors will return _REJECTED with a buit mask of the failed sessions and it is the responsibility
of the client to address the issues

AddSessionStatus Codes Description

_SESSION_ACCEPTED: Session is accepted by the server and the client performs normal operation
_SESSION_REJECTED: Session is rejected by the server as the message 
   is invalid, the client needs to correct the error.
_SESSION_TABLE_FULL: Session is rejected by the server as its session table is full, 
   the client needs to backoff until more space is available
_SESSION_TABLE_UNAVAILABLE: Session is rejected by the server due to an internal error 
   in the server, the client needs to back off until error is corrected.
_SESSION_ALREADY_EXISTS: Session is rejected by the the server as it already exists 
   in the server session table, the client will take corrective action to ensure state is consistent.

| Name | Number | Description |
| ---- | ------ | ----------- |
| _SESSION_ACCEPTED | 0 |  |
| _SESSION_REJECTED | 1 |  |
| _SESSION_TABLE_FULL | 2 |  |
| _SESSION_TABLE_UNAVAILABLE | 3 |  |
| _SESSION_ALREADY_EXISTS | 4 |  |



<a name="openoffload-v2-IpVersion"></a>

### IpVersion


| Name | Number | Description |
| ---- | ------ | ----------- |
| _IPV4 | 0 |  |
| _IPV6 | 1 |  |



<a name="openoffload-v2-ProtocolId"></a>

### ProtocolId


| Name | Number | Description |
| ---- | ------ | ----------- |
| _HOPOPT | 0 |  |
| _TCP | 6 |  |
| _UDP | 17 |  |



<a name="openoffload-v2-RequestStatus"></a>

### RequestStatus
RequestStatus Codes Description

_ACCEPTED: Normal operation
_REJECTED: Unknown error in the format of the REQUEST message
_REJECTED_SESSION_NONEXISTENT: In getSession or deleteSession the server does not have the session
   in its session table. The client needs to reconcile the system state.
_REJECTED_SESSION_TABLE_FULL: This should never happen as getClosedSessions, getSession, deleteSession never add sessions.
_REJECTED_SESSION_ALREADY_EXISTS: This should never happen as getClosedSessions, getSession, deleteSession never add sessions.
_NO_CLOSED_SESSIONS: When getClosedSessions returns with no closed sessions it will return 0 sessions. There should be no
   message attached so not sure if this is valid.
_REJECTED_INTERNAL_ERROR: The server has an internal error and cannot serivce the request.
   The client must log the error and optionally retry or skip the request.

| Name | Number | Description |
| ---- | ------ | ----------- |
| _ACCEPTED | 0 |  |
| _REJECTED | 1 |  |
| _REJECTED_SESSION_NONEXISTENT | 2 |  |
| _REJECTED_SESSION_TABLE_FULL | 3 |  |
| _REJECTED_SESSION_ALREADY_EXISTS | 4 |  |
| _NO_CLOSED_SESSIONS | 5 |  |
| _REJECTED_INTERNAL_ERROR | 6 |  |



<a name="openoffload-v2-SessionCloseCode"></a>

### SessionCloseCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| _NOT_CLOSED | 0 |  |
| _FINACK | 1 |  |
| _RST | 2 |  |
| _TIMEOUT | 3 |  |
| _UNKNOWN_CLOSE_CODE | 4 |  |



<a name="openoffload-v2-SessionState"></a>

### SessionState


| Name | Number | Description |
| ---- | ------ | ----------- |
| _ESTABLISHED | 0 |  |
| _CLOSING_1 | 1 |  |
| _CLOSING_2 | 2 |  |
| _CLOSED | 3 |  |
| _UNKNOWN_STATE | 4 |  |


 

 


<a name="openoffload-v2-SessionTable"></a>

### SessionTable
The session table was combined with the statistices service
in v1alpha4 to simplfy the code.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddSession | [SessionRequest](#openoffload-v2-SessionRequest) stream | [AddSessionResponse](#openoffload-v2-AddSessionResponse) | Adds a session This was changed in v1alpha4 to be a streaming API, for performance reasons. |
| GetSession | [SessionId](#openoffload-v2-SessionId) | [SessionResponse](#openoffload-v2-SessionResponse) | Obtains the session |
| DeleteSession | [SessionId](#openoffload-v2-SessionId) | [SessionResponse](#openoffload-v2-SessionResponse) | Delete a session |
| GetAllSessions | [SessionRequestArgs](#openoffload-v2-SessionRequestArgs) | [SessionResponses](#openoffload-v2-SessionResponses) | Stream back a specific session or all current sessions To stream a single session, pass SessionId as zero |
| GetClosedSessions | [SessionRequestArgs](#openoffload-v2-SessionRequestArgs) | [SessionResponse](#openoffload-v2-SessionResponse) stream | statistics as a outgoing session from the WB to Applications ? grpc seems to need a request input streamId is a placeholder |

 



<a name="tunneloffload-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tunneloffload.proto



<a name="tunneloffload-v2-CapabilityRequest"></a>

### CapabilityRequest
Capabilty request is empty since no paramteres are supplied to it,
all capabilities will be provided at response






<a name="tunneloffload-v2-CapabilityResponse"></a>

### CapabilityResponse
We&#39;ll have capability for matching, and for every tunnel


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| match_capabilities | [CapabilityResponse.MatchCapabilities](#tunneloffload-v2-CapabilityResponse-MatchCapabilities) |  |  |
| ipsec_capabilities | [CapabilityResponse.IPSecCapabilities](#tunneloffload-v2-CapabilityResponse-IPSecCapabilities) |  |  |
| geneve_capabilities | [CapabilityResponse.GeneveCapabilities](#tunneloffload-v2-CapabilityResponse-GeneveCapabilities) |  |  |






<a name="tunneloffload-v2-CapabilityResponse-GeneveCapabilities"></a>

### CapabilityResponse.GeneveCapabilities



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| number_geneve_options_supported | [uint32](#uint32) |  | Number of options geneve is supporting in encap |






<a name="tunneloffload-v2-CapabilityResponse-IPSecCapabilities"></a>

### CapabilityResponse.IPSecCapabilities



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnel_type_supported | [IPSecTunnelType](#tunneloffload-v2-IPSecTunnelType) | repeated |  |
| encryption_supported | [EncType](#tunneloffload-v2-EncType) | repeated |  |






<a name="tunneloffload-v2-CapabilityResponse-MatchCapabilities"></a>

### CapabilityResponse.MatchCapabilities



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ingress_interface_matching | [bool](#bool) |  | Is interface can be matched for encapsulation / decapsulation |
| vxlan_matching | [bool](#bool) |  | Match with VXLAN VNI |
| geneve_matching | [bool](#bool) |  | Match with geneve can happen |
| tunnel_matching | [bool](#bool) |  | Matching on tunnel ID |
| spi_matching | [bool](#bool) |  | Can match on IPSec |






<a name="tunneloffload-v2-Counters"></a>

### Counters



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| in_packets | [uint64](#uint64) |  |  |
| out_packets | [uint64](#uint64) |  |  |
| in_bytes | [uint64](#uint64) |  |  |
| out_bytes | [uint64](#uint64) |  |  |
| in_packets_drops | [uint64](#uint64) |  |  |
| out_packets_drops | [uint64](#uint64) |  |  |
| in_bytes_drops | [uint64](#uint64) |  |  |
| out_bytes_drops | [uint64](#uint64) |  |  |






<a name="tunneloffload-v2-CreateIpTunnelResponse"></a>

### CreateIpTunnelResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnel_id | [uint64](#uint64) |  | Tunnel ID assigned to this tunnel |
| error | [Error](#tunneloffload-v2-Error) |  | Message appears only if there&#39;s error in the response |






<a name="tunneloffload-v2-CreateIpTunnelResponses"></a>

### CreateIpTunnelResponses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_status | [AddTunnelStatus](#tunneloffload-v2-AddTunnelStatus) |  |  |
| error_status | [uint64](#uint64) |  | bitmask of errors for each of the sesssions sent in a stream (max 64). |
| responses | [CreateIpTunnelResponse](#tunneloffload-v2-CreateIpTunnelResponse) | repeated |  |






<a name="tunneloffload-v2-Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_message | [TunnelAdditionError](#tunneloffload-v2-TunnelAdditionError) |  | Error code describing the error with the request |
| error_string | [string](#string) |  | Error string indicating the error |






<a name="tunneloffload-v2-Geneve"></a>

### Geneve



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| geneve_encap | [GeneveEncap](#tunneloffload-v2-GeneveEncap) |  |  |
| geneve_decap | [GeneveDecap](#tunneloffload-v2-GeneveDecap) |  |  |






<a name="tunneloffload-v2-GeneveDecap"></a>

### GeneveDecap
GeneveDecap can only be used if Geneve
is on the match of the tunnel






<a name="tunneloffload-v2-GeneveEncap"></a>

### GeneveEncap
Defining the Geneve Header at encpasulation
Fields names are identical to the fields as described in the RFC:
https://datatracker.ietf.org/doc/html/rfc8926
For details of each fields, please refer to the RFC

Notes:
- Version field is not present since always 0


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| outer_ipv4_pair | [IPV4Pair](#tunneloffload-v2-IPV4Pair) |  |  |
| outer_ipv6_pair | [IPV6Pair](#tunneloffload-v2-IPV6Pair) |  |  |
| inner_mac_pair | [MacPair](#tunneloffload-v2-MacPair) |  | Source &amp; Dest mac of inner Geneve packet |
| option_length | [uint32](#uint32) |  | 6 bits - Multiply of 4 bytes |
| control_packet | [bool](#bool) |  | O bit at rfc. True is &#39;1&#39;, False is &#39;0&#39;. Default is False. |
| critical_option_present | [bool](#bool) |  | C bit at rfc. True is &#39;1&#39;, False is &#39;0&#39;. Default is False. |
| vni | [uint32](#uint32) |  | Virtual Network Identifier - 24 bits |
| protocol_type | [uint32](#uint32) |  | Currently only &#34;Trans Ether Bridging&#34; is supported (0x6558) |
| geneve_option | [GeneveOption](#tunneloffload-v2-GeneveOption) | repeated |  |






<a name="tunneloffload-v2-GeneveOption"></a>

### GeneveOption



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| option_class | [uint32](#uint32) |  | 16 bits |
| type | [uint32](#uint32) |  | 8 bits |
| length | [uint32](#uint32) |  | Only 5 bits used |
| data | [bytes](#bytes) |  | Length is multiple of 4 bytes (see https://datatracker.ietf.org/doc/html/rfc8926#section-3.5)

Only 4-128 bytes are acceptable, |






<a name="tunneloffload-v2-IPSecDec"></a>

### IPSecDec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnel_type | [IPSecTunnelType](#tunneloffload-v2-IPSecTunnelType) |  | Transport / Tunnel... |
| encryption_type | [EncType](#tunneloffload-v2-EncType) |  | AES-256GCM |
| ipsec_sas | [IPSecSAParams](#tunneloffload-v2-IPSecSAParams) | repeated |  |






<a name="tunneloffload-v2-IPSecEnc"></a>

### IPSecEnc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnel_type | [IPSecTunnelType](#tunneloffload-v2-IPSecTunnelType) |  |  |
| encryption_type | [EncType](#tunneloffload-v2-EncType) |  |  |
| ipsec_sa | [IPSecSAParams](#tunneloffload-v2-IPSecSAParams) |  |  |
| ipv4_tunnel | [IPV4Pair](#tunneloffload-v2-IPV4Pair) |  |  |
| ipv6_tunnel | [IPV6Pair](#tunneloffload-v2-IPV6Pair) |  |  |






<a name="tunneloffload-v2-IPSecSAParams"></a>

### IPSecSAParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spi | [uint32](#uint32) |  |  |
| encryption_key | [bytes](#bytes) |  |  |
| operation | [Operation](#tunneloffload-v2-Operation) |  | Indicates if removing / updating / creating SA |






<a name="tunneloffload-v2-IPSecTunnel"></a>

### IPSecTunnel



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ipsec_enc | [IPSecEnc](#tunneloffload-v2-IPSecEnc) |  |  |
| ipsec_dec | [IPSecDec](#tunneloffload-v2-IPSecDec) |  |  |






<a name="tunneloffload-v2-IPV4Match"></a>

### IPV4Match



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_ip | [fixed32](#fixed32) |  |  |
| source_ip_prefix | [uint32](#uint32) |  |  |
| destination_ip | [fixed32](#fixed32) |  |  |
| destination_ip_prefix | [uint32](#uint32) |  |  |






<a name="tunneloffload-v2-IPV4Pair"></a>

### IPV4Pair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_ip | [fixed32](#fixed32) |  |  |
| destination_ip | [fixed32](#fixed32) |  |  |






<a name="tunneloffload-v2-IPV6Match"></a>

### IPV6Match



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_ip | [bytes](#bytes) |  |  |
| source_ip_prefix | [uint32](#uint32) |  |  |
| destination_ip | [bytes](#bytes) |  |  |
| destination_ip_prefix | [uint32](#uint32) |  |  |






<a name="tunneloffload-v2-IPV6Pair"></a>

### IPV6Pair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_ip | [bytes](#bytes) |  |  |
| destination_ip | [bytes](#bytes) |  |  |






<a name="tunneloffload-v2-IpTunnelRequest"></a>

### IpTunnelRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnel_id | [uint64](#uint64) |  |  |
| operation | [Operation](#tunneloffload-v2-Operation) |  |  |
| match_criteria | [MatchCriteria](#tunneloffload-v2-MatchCriteria) |  | When hitting this match, |
| next_action | [Action](#tunneloffload-v2-Action) |  | What we&#39;ll do after matching the packet, should we |
| ipsec_tunnel | [IPSecTunnel](#tunneloffload-v2-IPSecTunnel) |  | Tunnel that will be used for encapsulation, can be both |
| geneve | [Geneve](#tunneloffload-v2-Geneve) |  |  |
| nat | [Nat](#tunneloffload-v2-Nat) |  |  |






<a name="tunneloffload-v2-IpTunnelResponse"></a>

### IpTunnelResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnel_id | [uint64](#uint64) |  | Tunnel ID assigned to this tunnel |
| ip_tunnel | [IpTunnelRequest](#tunneloffload-v2-IpTunnelRequest) |  | Information regards the ipTunnel (including match, tunnel information) |
| tunnel_counters | [Counters](#tunneloffload-v2-Counters) |  | Counters of the session |
| error | [Error](#tunneloffload-v2-Error) |  | Message that appears only if there&#39;s a problem in the request |






<a name="tunneloffload-v2-IpTunnelResponses"></a>

### IpTunnelResponses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| responses | [IpTunnelResponse](#tunneloffload-v2-IpTunnelResponse) | repeated |  |






<a name="tunneloffload-v2-IpTunnelStatsResponse"></a>

### IpTunnelStatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnel_id | [uint64](#uint64) |  | Tunnel ID assigned to this tunnel |
| tunnel_counters | [Counters](#tunneloffload-v2-Counters) |  | Counters of the session |
| error | [Error](#tunneloffload-v2-Error) |  | Message that appears only if there&#39;s a problem in the request |






<a name="tunneloffload-v2-IpTunnelStatsResponses"></a>

### IpTunnelStatsResponses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| responses | [IpTunnelStatsResponse](#tunneloffload-v2-IpTunnelStatsResponse) | repeated |  |






<a name="tunneloffload-v2-MacPair"></a>

### MacPair



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destination_mac | [bytes](#bytes) |  |  |
| source_mac | [bytes](#bytes) |  |  |






<a name="tunneloffload-v2-MatchCriteria"></a>

### MatchCriteria



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ingress_interface | [string](#string) |  | In case it&#39;s not present, untagged traffic will be matched |
| mac_match | [MacPair](#tunneloffload-v2-MacPair) |  | MAC of the packet itself |
| ipv4_match | [IPV4Match](#tunneloffload-v2-IPV4Match) |  |  |
| ipv6_match | [IPV6Match](#tunneloffload-v2-IPV6Match) |  |  |
| tunnel_id | [uint64](#uint64) |  | Match on specific tunnel |
| ipsec_match | [MatchCriteria.IPSecMatch](#tunneloffload-v2-MatchCriteria-IPSecMatch) |  |  |
| geneve_match | [MatchCriteria.GeneveMatch](#tunneloffload-v2-MatchCriteria-GeneveMatch) |  |  |
| vxlan_match | [MatchCriteria.VXLanMatch](#tunneloffload-v2-MatchCriteria-VXLanMatch) |  |  |






<a name="tunneloffload-v2-MatchCriteria-GeneveMatch"></a>

### MatchCriteria.GeneveMatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vni | [uint32](#uint32) |  |  |
| mac_match | [MacPair](#tunneloffload-v2-MacPair) |  | Inner Match of Geneve Packet |
| protocol_type | [uint32](#uint32) |  | Currently only &#34;Trans Ether Bridging&#34; is supported (0x6558) |
| ipv4_match | [IPV4Match](#tunneloffload-v2-IPV4Match) |  |  |
| ipv6_match | [IPV6Match](#tunneloffload-v2-IPV6Match) |  |  |






<a name="tunneloffload-v2-MatchCriteria-IPSecMatch"></a>

### MatchCriteria.IPSecMatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spi | [uint32](#uint32) |  |  |
| sn | [uint32](#uint32) |  |  |






<a name="tunneloffload-v2-MatchCriteria-VXLanMatch"></a>

### MatchCriteria.VXLanMatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vni | [uint32](#uint32) |  |  |
| mac_match | [MacPair](#tunneloffload-v2-MacPair) |  |  |
| ipv4_match | [IPV4Match](#tunneloffload-v2-IPV4Match) |  |  |
| ipv6_match | [IPV6Match](#tunneloffload-v2-IPV6Match) |  |  |






<a name="tunneloffload-v2-Nat"></a>

### Nat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_ip | [uint32](#uint32) |  |  |






<a name="tunneloffload-v2-TunnelAdditionError"></a>

### TunnelAdditionError



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| match_error | [MatchError](#tunneloffload-v2-MatchError) |  | Message will only present if there&#39;s error in tunnel

Error in match of tunnel |
| tunnel_error | [TunnelError](#tunneloffload-v2-TunnelError) |  | General problem in tunnel definition |
| ipsec_error | [IPSecError](#tunneloffload-v2-IPSecError) |  | IPSec Error |
| geneve_error | [GeneveError](#tunneloffload-v2-GeneveError) |  | Error in geneve |






<a name="tunneloffload-v2-TunnelId"></a>

### TunnelId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnel_id | [uint64](#uint64) |  |  |






<a name="tunneloffload-v2-TunnelRequestArgs"></a>

### TunnelRequestArgs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tunnels_per_request | [uint32](#uint32) |  | How many tunnels will be returned per request |





 


<a name="tunneloffload-v2-Action"></a>

### Action


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| FORWARD | 1 | In this action packet will be forwarded right away |
| RECIRCULATE | 2 | In this action packet will continue processing, to other tunnel |



<a name="tunneloffload-v2-AddTunnelStatus"></a>

### AddTunnelStatus
Errors for adding a tunnel
If all tunnels are successful inserted return _ACCEPTED

If check of tunnel capacity in offload device is insufficient to add all tunnels
do not insert any tunnels and return  _REJECTED_TUNNEL_TABLE_FULL. It is the
responsibility of the client to re-try

If the server is unavailable for some other reason then return _REJECTED_TUNNEL_TABLE_UNAVAILABLE.
It is the  responsibility of the client to re-try

All other errors will return _REJECTED with a buit mask of the failed sessions and it is the responsibility
of the client to address the issues

AddTunnelStatus Codes Description

_TUNNEL_ACCEPTED: Tunnel is accepted by the server and the client performs normal operation
_TUNNEL_REJECTED: Tunnel is rejected by the server as the message
   is invalid, the client needs to correct the error.
_TUNNEL_TABLE_FULL: Tunnel is rejected by the server as its session table is full,
   the client needs to backoff until more space is available
_TUNNEL_TABLE_UNAVAILABLE: Tunnel is rejected by the server due to an internal error
   in the server, the client needs to back off until error is corrected.
_TUNNEL_ALREADY_EXISTS: Tunnel is rejected by the the server as it already exists
   in the server session table, the client will take corrective action to ensure state is consistent.

| Name | Number | Description |
| ---- | ------ | ----------- |
| _TUNNEL_ACCEPTED | 0 |  |
| _TUNNEL_REJECTED | 1 |  |
| _TUNNEL_TABLE_FULL | 2 |  |
| _TUNNEL_TABLE_UNAVAILABLE | 3 |  |
| _TUNNEL_ALREADY_EXISTS | 4 |  |



<a name="tunneloffload-v2-EncType"></a>

### EncType


| Name | Number | Description |
| ---- | ------ | ----------- |
| _AES256GCM64 | 0 |  |
| _AES256GCM96 | 1 |  |
| _AES256GCM128 | 2 |  |
| _AES128GCM64 | 3 |  |
| _AES128GCM96 | 4 |  |
| _AES128GCM128 | 5 |  |
| _AES256CCM64 | 6 |  |
| _AES256CCM96 | 7 |  |
| _AES256CCM128 | 8 |  |
| _AES128CCM64 | 9 |  |
| _AES128CCM96 | 10 |  |
| _AES128CCM128 | 11 |  |



<a name="tunneloffload-v2-GeneveError"></a>

### GeneveError


| Name | Number | Description |
| ---- | ------ | ----------- |
| INVALID_OPTION | 0 | One of the options supported isn&#39;t valid |
| TOO_MANY_OPTIONS | 1 | Too many options provided with the geneve tunnel |
| INVALID_GENEVE_FIELD | 2 | One of the fields provided in GENVE isn&#39;t valid (e.g. too big VNI) |



<a name="tunneloffload-v2-IPSecError"></a>

### IPSecError


| Name | Number | Description |
| ---- | ------ | ----------- |
| INVALID_KEY | 0 | Key got into IPSec is not matching the requested size |
| NON_SUPPORTED_ENCRYPTION | 1 | Encrypted type requested from IPSec is not supported |
| NON_SUPPORTED_TUNNEL_TYPE | 2 | Tunnel type requested by IPSec is not valid (TUNNEL mode requested but not valid) |
| IPSEC_MISSING_FIELDS | 3 | Some missing fields in IPSec, e.g. TUNNEL MODE without tunnel IP&#39;s provided |



<a name="tunneloffload-v2-IPSecTunnelType"></a>

### IPSecTunnelType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSPORT | 0 |  |
| TUNNEL | 1 |  |
| TRANSPORT_NAT_TRAVERSAL | 2 | Nat Traversal is a mechanism to overcome Nat happens between the two IPSec endpoints, by adding a UDP Header after IPSec This mode can happen both in TRANSPORT &amp; TUNNEL Mode Fore more information please refer to the following RFC&#39;s: https://datatracker.ietf.org/doc/html/rfc3947 https://datatracker.ietf.org/doc/html/rfc3715 |
| TUNNEL_NAT_TRAVERSAL | 3 |  |



<a name="tunneloffload-v2-MatchError"></a>

### MatchError


| Name | Number | Description |
| ---- | ------ | ----------- |
| MISING_FIELDS | 0 | Some missing fields are misisng in the match statment |
| INVALID_TUNNEL_ID | 1 | Tunnel ID match isn&#39;t valid or doesn&#39;t exist |
| INVALID_CAPABILITIES | 2 | The match isn&#39;t suitable with tunnel capabilities |
| INVALID_FIELD | 3 | Invalid field found in the match (e.g. invalid mac) |



<a name="tunneloffload-v2-Operation"></a>

### Operation


| Name | Number | Description |
| ---- | ------ | ----------- |
| _NONE | 0 |  |
| _CREATE | 1 |  |
| _UPDATE | 2 |  |
| _DELETE | 3 |  |



<a name="tunneloffload-v2-TunnelError"></a>

### TunnelError


| Name | Number | Description |
| ---- | ------ | ----------- |
| NOT_SUPPORTED_TUNNEL | 0 | Tunnel offload requested to non supported tunnel |


 

 


<a name="tunneloffload-v2-IpTunnelService"></a>

### IpTunnelService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Capabilities | [CapabilityRequest](#tunneloffload-v2-CapabilityRequest) | [CapabilityResponse](#tunneloffload-v2-CapabilityResponse) | Get which capabilities are available while using the |
| CreateIpTunnel | [IpTunnelRequest](#tunneloffload-v2-IpTunnelRequest) stream | [CreateIpTunnelResponses](#tunneloffload-v2-CreateIpTunnelResponses) | Creation of IP Tunnel This API should be generic and allow creations of many IP tunnels |
| GetIpTunnel | [TunnelId](#tunneloffload-v2-TunnelId) | [IpTunnelResponse](#tunneloffload-v2-IpTunnelResponse) | Getting a tunnel by it&#39;s ID |
| GetIpTunnelStats | [TunnelId](#tunneloffload-v2-TunnelId) | [IpTunnelStatsResponse](#tunneloffload-v2-IpTunnelStatsResponse) | Getting a tunnel by it&#39;s ID |
| GetAllIpTunnels | [TunnelRequestArgs](#tunneloffload-v2-TunnelRequestArgs) | [IpTunnelResponses](#tunneloffload-v2-IpTunnelResponses) stream | Getting all the ipTunnels currently configured |
| GetAllIpTunnelsStats | [TunnelRequestArgs](#tunneloffload-v2-TunnelRequestArgs) | [IpTunnelStatsResponses](#tunneloffload-v2-IpTunnelStatsResponses) stream | Get all the iptunnels stats responses |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

