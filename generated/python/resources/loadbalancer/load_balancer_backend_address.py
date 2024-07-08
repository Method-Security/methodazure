# This file was auto-generated by Fern from our API Definition.

import datetime as dt
import typing

from ...core.datetime_utils import serialize_datetime
from ...core.pydantic_utilities import deep_union_pydantic_dicts, pydantic_v1
from .load_balancer_backend_address_admin_state import LoadBalancerBackendAddressAdminState
from .nat_rule_port_mapping import NatRulePortMapping
from .sub_resource import SubResource


class LoadBalancerBackendAddress(pydantic_v1.BaseModel):
    name: str
    admin_state: typing.Optional[LoadBalancerBackendAddressAdminState] = pydantic_v1.Field(
        alias="adminState", default=None
    )
    ip_address: typing.Optional[str] = pydantic_v1.Field(alias="ipAddress", default=None)
    load_balancer_frontend_ip_configuration: typing.Optional[SubResource] = pydantic_v1.Field(
        alias="loadBalancerFrontendIPConfiguration", default=None
    )
    subnet: typing.Optional[SubResource] = None
    virtual_network: typing.Optional[SubResource] = pydantic_v1.Field(alias="virtualNetwork", default=None)
    inbound_nat_rules_port_mapping: typing.Optional[typing.List[NatRulePortMapping]] = pydantic_v1.Field(
        alias="inboundNatRulesPortMapping", default=None
    )
    network_interface_ip_configurations: typing.Optional[SubResource] = pydantic_v1.Field(
        alias="NetworkInterfaceIpConfigurations", default=None
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