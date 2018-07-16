import os
from codecs import open

import re
from setuptools import setup, find_packages

setup(
    name="netograph",
    version="0.1",
    description="A Python interface to netograph.io.",
    url="http://mitmproxy.org",
    author="Netograph Ltd",
    author_email="aldo@netograph.io",
    packages=find_packages(include=["netograph"]),
    include_package_data=True,
    install_requires=[
        "grpcio>=1.13",
        "googleapis-common-protos"
    ],
    extras_require={
        'dev': [
            "grpcio-tools"
        ],
    }
)
