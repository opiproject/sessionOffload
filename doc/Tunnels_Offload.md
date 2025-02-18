# IP Tunnel Offload

-------- DRAFT -----------

With openOffload it will be possible to also offload IP Tunnels into the underlying hardware.

A service called `ipTunnelTable` will be introduced, for CRUD operations on the offloaded sessions.

Through this service, it will be possible to offload several kinds of IP Tunnels, where the common configuration between them is the `match` criteria, which indicates which packets will be matched and go to a tunnel.

The API will also allow the definition of switching between tunnels with decapsulating and encapsulating appropriate to the tunnel type for the incoming and outgoing traffic.

## Offloading Rules

With openOffload, the user can request to match the packet encapsulate it on a tunnel, this will be done by providing matching criteria, tunnel properties and the next action to perform.

Each ipTunnelRule inserted will include *match* determines rule matching. Upon match tunnel provided by the *tunnel*
field will be applied to the packet.

`nextActinon` field will determine what will happen to the packet after applying the tunnel, with two possibilities: *recirculate* indicates the packet will overgo again the tunneling subsystem, and another tunnel rule may be applied. *forward* indicates that the packet will be forwarded directly from the device, without the process of any other tunnel.

To summarize the flow, see the following diagram:

![Matching](images/tunnelOffload/deviceDiagram.png)

### Tunnel ID

As a response to `ipTunnelRule` creation, the offloading device will send back a unique 32-bit identifier to the tunnel - **tunnelId**.

The tunnelId will be used for any tunnel action: getting tunnel operational data / deleting tunnel / etc

## Packet Matching

Matching of the packet can be based on several matching criteria, interface, IP, VRF, etc.

The matching criteria are comprised of three parts:

1. IP packet fields  - MAC, IP's, etc
2. Specific header matching: GENVE / VXLan / IPSec, etc.
3. Tunnel Matching: Matching packets according to their *last* match. More information on it in *Tunnel Chaining* section.

This matching structure is intended to provide maximum flexibility to the user, that, for example - can match on specific VXLan VNI & Inner Mac, and only then perform IPSec encryption (or any other action).

Field absence will count as wildcard matching. Fields not containing any values will be matched on any value received. e.g. Not providing ipv4/ipv6 match, will make the tunnel matched on any IP received by the device.

Let's take a look at the following packet:

![Matching](images/tunnelOffload/vxlanPacket.png)

Note: *All of the matches below will match this packet:*

```bash
Source IP: 1.2.3.0/24
VXLAN Packet {
   VNI: 466
}
```

Note: *Source IP & VNI Match, Destination IP is not provided - matching on any destination IP*

```bash
VXLAN Packet {
}
```

Note: *Just specifying that packet is VXLan packet is enough, done by matching VXLan without any field in message*

```bash
Source IP: 1.2.3.0/24
Destination IP: 5.5.5.0/24
VXLAN Packet {
   Source MAC: 0x102030405060
}
```

Note: *Source & Dest outer IP match, inner source MAC match*

## Packet Encapsulation / Decapsulation

After the packet is matched, it should be encapsulated / decapsulated according to the offloaded tunnel.

Tunnel configuration can be unidirectional or bidirectional, depending on the tunnel characteristics.

In the case of a bi-directional tunnel, the tunnel encapsulation will determine the match on the reverse side (e.g. NAT).

In the case of a uni-directional tunnel, tunnel definition will be for only encapsulation / decapsulation of tunnel.

GENEVE is a classic example of a bi-directional user - where the offloaded device can device to just encapsulate / decapsulate GENEVE.

```bash
message GENEVE {
  oneof {
    GENEVEEncap geneveEncap = 1;
    GENEVEDecap geneveDecap = 2;
  }
```

*Upon a match, the user decides to perform encapsulation / decapsulation, GENEVE is comprised of two distinct messages,
choosing one of them decide about the tunnel operation*

NAT is an example where the device is both encapsulating / decapsulating the tunnels and this is a bidirectional tunnel.

```bash
message NAT {
  SourceIP sourceIP = 1;
}
```

Note: *In NAT example, ipTunnelRule will be used for both encapsulation / decapsulation*

In a NAT example (of bidirectional tunnel), the match will be used for egress traffic only, and a packet that will be matched on it will perform NAT.

After maching on egress, the NAT will write the rule for ingress matching, and this will yield on matching of ingress tunnel (see example below).

![Matching](images/tunnelOffload/uni_bi_directional_tunnel.png)

**Match** is the only indicator for going into tunnel encapsulation / decapsulation.
While wanting to perform decapsulation of the packet in a uni-directional tunnel (e.g. the GENEVE Decapsulation in the example above), a GENEVE match **MUST** be on MATCH criteria (the same applies for IPSec, VXLan, etc)

### Tunnel Chaining

Tunnel chaining is possible with tunnelOffload, making a packet overgo several tunnels encapsulation / decapsulations.

According to the following chart:

![Matching](images/tunnelOffload/deviceDiagram.png)

Upon receipt of the packet into the offloading device, it will be matched and a tunnel will be applied.

If the `nextAction` equals `RECIRCULATE`, the packet will be processed via the matching logic again, and if a match will be found - another tunnel will be applied to the packet.

This iterative process will occur until one; the packet isn't matched via the matching logic. 2; the next action equals  `FORWARD`. That will yield immediate forwarding of the data.

Consider the following example:

![Matching](images/tunnelOffload/tunnel_chain_ip.png)

While the first match of packet **MUST** be based on IP's, there's a possibility to use "TunnelID" as the match criteria.

Rules having "TunnelID" in their match criteria can only be part of tunnel chaining and will be applied on the packet from the second tunnel and afterward.

Consider the following example:

![Matching](images/tunnelOffload/tunnel_chain_tunnel_id.png)

The advantage of using the "TunnelID" as a match, is the ability to know for sure that after some tunneling, the second will happen for sure.

## Capabilities

Capabilities are needed so the user can detect which features are available with tunnel offload,
user can detect which features are available with tunnel offload, both the tunnel capabilities & matching capabilities of the device.

Please see the capabilities rpc; in the tunnels offload proto for more detailed information about it.
