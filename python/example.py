import os

import grpc
from netograph import ngapi_pb2_grpc
from netograph import ngapi_pb2

channel = grpc.secure_channel(
    'localhost:8081',
    grpc.composite_channel_credentials(
        grpc.ssl_channel_credentials(),
        grpc.access_token_call_credentials(os.environ["NGC_TOKEN"]),
    ),
    options=[
        ('grpc.ssl_target_name_override', "grpc.netograph.io"),
    ]
)
stub = ngapi_pb2_grpc.NetographStub(channel)
for i in stub.Datasets(ngapi_pb2.DatasetsRequest()):
    print(i)

