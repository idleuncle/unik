package aws

import (
	"github.com/emc-advanced-dev/unik/pkg/types"
	"github.com/layer-x/layerx-commons/lxlog"
	"github.com/emc-advanced-dev/unik/pkg/providers/common"
)

func (p *AwsProvider) GetVolume(logger lxlog.Logger, nameOrIdPrefix string) (*types.Volume, error) {
	return common.GetVolume(logger, p, nameOrIdPrefix)
}
