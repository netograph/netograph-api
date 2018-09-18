"""
    This example uses Netograph's social dataset to discover domains associated
    with a parent domain. It begins by enumerating all subdomains for a given
    parent domain. For each subdomain, it retrieves all IP addresses that we
    have seen serve up data for that domain. Finally, we do a reverse query to
    find all domains associated with each of these IP addresses. This gives us a
    list of all domains that were served from the same IPs as any of the parent
    domains.

    In many cases, this simply reveals CDNs. However, this often also turns up
    domains that are deeply associated with a parent domain through shared
    infrastructure.
"""
import os

import netograph
from netograph.dsetapi import dset_pb2

DSET = "netograph:social"
# Limit response size - CDNs mean that very large numbers of domains can be
# associated with a single IP.
LIMIT = 50

seen = set([])
domains = set([])

conn = netograph.connect_dset(os.environ["NGC_TOKEN"])
satq = dset_pb2.DomainSearchRequest(
    query="kaspersky.com",
    dataset=DSET,
    limit=LIMIT
)
for root in conn.DomainSearch(satq):
    print(f"{root.domain}")
    ipq = dset_pb2.IPsForDomainRequest(
        query=root.domain,
        dataset=DSET,
        limit=LIMIT
    )
    for ip in conn.IPsForDomain(ipq):
        print(f"\t{ip.ip}")
        if ip.ip in seen:
            print("\t\talready seen - skipping")
        else:
            sharedq = dset_pb2.DomainsForIPRequest(
                ip=ip.ip,
                dataset=DSET,
                limit=LIMIT
            )
            for shared in conn.DomainsForIP(sharedq):
                print(f"\t\t{shared.domain}")
                domains.add(shared.domain)
            seen.add(ip.ip)

print("\nAssociated domains:")
for d in sorted(domains):
    print("\t", d)

