package storage

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/bdjuno/v4/types"
)

// UpdateParams gets the updated params and stores them inside the database
func (m *Module) UpdateParams(height int64) error {
	log.Debug().Str("module", "storage").Int64("height", height).
		Msg("updating params")

	params, err := m.source.Params(height)
	if err != nil {
		return fmt.Errorf("error while getting params: %s", err)
	}

	return m.db.SaveStorageParams(types.NewStorageParams(params, height))

}

// UpdateProviders gets the updated providers list and stores it inside the database
func (m *Module) UpdateProviders(height int64) error {
	log.Debug().Str("module", "storage").Int64("height", height).
		Msg("updating providers")

	providers, err := m.source.Providers(height)
	if err != nil {
		return fmt.Errorf("error while getting providers list: %s", err)
	}

	return m.db.SaveStorageProviders(providers, height)

}