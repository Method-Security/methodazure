# This file was auto-generated by Fern from our API Definition.

import datetime as dt
import typing

from ...core.datetime_utils import serialize_datetime
from ...core.pydantic_utilities import deep_union_pydantic_dicts, pydantic_v1
from ..interface.interface_ip_configuration import InterfaceIpConfiguration
from .load_balancer_backend_address import LoadBalancerBackendAddress
from .sub_resource import SubResource
from .sync_mode import SyncMode


class BackendAddressPool(pydantic_v1.BaseModel):
    """
    Collection of backend address pools used by the load balancer:
    https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5#BackendAddressPool
    """

    id: str
    name: str
    type: str
    load_balancer_backend_addresses: typing.Optional[typing.List[LoadBalancerBackendAddress]] = pydantic_v1.Field(
        alias="loadBalancerBackendAddresses", default=None
    )
    location: typing.Optional[str] = None
    sync_mode: typing.Optional[SyncMode] = pydantic_v1.Field(alias="syncMode", default=None)
    virtual_network: typing.Optional[SubResource] = pydantic_v1.Field(alias="virtualNetwork", default=None)
    backend_ip_configurations: typing.Optional[typing.List[InterfaceIpConfiguration]] = pydantic_v1.Field(
        alias="backendIpConfigurations", default=None
    )

    def json(self, **kwargs: typing.Any) -> str:
        kwargs_with_defaults: typing.Any = {"by_alias": True, "exclude_unset": True, **kwargs}
        return super().json(**kwargs_with_defaults)

    def dict(self, **kwargs: typing.Any) -> typing.Dict[str, typing.Any]:
        kwargs_with_defaults_exclude_unset: typing.Any = {"by_alias": True, "exclude_unset": True, **kwargs}
        kwargs_with_defaults_exclude_none: typing.Any = {"by_alias": True, "exclude_none": True, **kwargs}

        return deep_union_pydantic_dicts(
            super().dict(**kwargs_with_defaults_exclude_unset), super().dict(**kwargs_with_defaults_exclude_none)
        )

    class Config:
        allow_population_by_field_name = True
        populate_by_name = True
        extra = pydantic_v1.Extra.allow
        json_encoders = {dt.datetime: serialize_datetime}
