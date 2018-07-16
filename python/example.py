import os

import netograph
from netograph import ngapi_pb2

DSET = "netograph:social"
LIMIT = 50

conn = netograph.channel(os.environ["NGC_TOKEN"])

seen = set([])
domains = set([])

satq = ngapi_pb2.DomainSearchRequest(
    query="twitter.com",
    dataset=DSET,
    limit=LIMIT
)
for root in conn.DomainSearch(satq):
    print(f"{root.domain}")
    ipq = ngapi_pb2.IPsForDomainRequest(
        query=root.domain,
        dataset=DSET,
        limit=LIMIT
    )
    for ip in conn.IPsForDomain(ipq):
        print(f"\t{ip.ip}")
        if ip.ip in seen:
            print("\t\talready seen - skipping")
        else:
            sharedq = ngapi_pb2.DomainsForIPRequest(
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

