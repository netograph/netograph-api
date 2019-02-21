import grpc

from netograph.dsetapi import dset_pb2_grpc
from netograph.userapi import user_pb2_grpc


common_options = [
    ('grpc.keepalive_time_ms', 300000),
    ('grpc.ssl_target_name_override', "grpc.netograph.io"),
    ('grpc.keepalive_permit_without_calls', 1),
]


def connect_dset(token):
    channel = grpc.secure_channel(
        'grpc.netograph.io:443',
        grpc.composite_channel_credentials(
            grpc.ssl_channel_credentials(),
            grpc.access_token_call_credentials(token),
        ),
        options=common_options,
    )
    return dset_pb2_grpc.DsetStub(channel)


def connect_user(token):
    channel = grpc.secure_channel(
        'grpc.netograph.io:443',
        grpc.composite_channel_credentials(
            grpc.ssl_channel_credentials(),
            grpc.access_token_call_credentials(token),
        ),
        options=common_options,
    )
    return user_pb2_grpc.UserStub(channel)