# This file was auto-generated by Fern from our API Definition.

from ...core.pydantic_utilities import UniversalBaseModel
from .load_balancer_sku_name import LoadBalancerSkuName
from .load_balancer_sku_tier import LoadBalancerSkuTier
from ...core.pydantic_utilities import IS_PYDANTIC_V2
import typing
import pydantic


class LoadBalancerSku(UniversalBaseModel):
    name: LoadBalancerSkuName
    tier: LoadBalancerSkuTier

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(
            extra="allow"
        )  # type: ignore # Pydantic v2
    else:

        class Config:
            extra = pydantic.Extra.allow
