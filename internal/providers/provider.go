package providers

import (
	"github.com/ronanwatkins/terraform-plugin-sdk/v2/internal/configs/configschema"
)

// Schema pairs a provider or resource schema with that schema's version.
// This is used to be able to upgrade the schema in UpgradeResourceState.
type Schema struct {
	Version int64
	Block   *configschema.Block
}
