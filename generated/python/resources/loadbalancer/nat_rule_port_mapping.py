# This file was auto-generated by Fern from our API Definition.

from ...core.pydantic_utilities import UniversalBaseModel
import pydantic
from ...core.pydantic_utilities import IS_PYDANTIC_V2
import typing


class NatRulePortMapping(UniversalBaseModel):
    backend_port: int = pydantic.Field(alias="backendPort")
    frontend_port: int = pydantic.Field(alias="frontendPort")
    inbound_nat_rule_name: str = pydantic.Field(alias="inboundNatRuleName")

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(
            extra="allow"
        )  # type: ignore # Pydantic v2
    else:

        class Config:
            extra = pydantic.Extra.allow
