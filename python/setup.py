from setuptools import setup, find_packages

setup(
    name="netograph",
    version="0.6",
    description="A Python interface to netograph.io.",
    url="https://netograph.io",
    author="Netograph Ltd",
    author_email="aldo@netograph.io",
    packages=find_packages(),
    include_package_data=True,
    install_requires=[
        "grpcio",
        "googleapis-common-protos"
    ],
    extras_require={
        'dev': [
            "grpcio-tools",
            "flake8",
        ],
    }
)
