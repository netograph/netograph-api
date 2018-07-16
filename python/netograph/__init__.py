import grpc
from netograph import ngapi_pb2_grpc


def channel(token):
    channel = grpc.secure_channel(
        'grpc.netograph.io:443',
        grpc.composite_channel_credentials(
            grpc.ssl_channel_credentials(),
            grpc.access_token_call_credentials(token),
        ),
        options=[
            ('grpc.ssl_target_name_override', "grpc.netograph.io"),
        ]
    )
    return ngapi_pb2_grpc.NetographStub(channel)