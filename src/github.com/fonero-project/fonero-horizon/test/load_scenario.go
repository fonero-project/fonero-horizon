package test

import (
	"github.com/fonero-project/fonero-horizon/test/scenarios"
)

func loadScenario(scenarioName string, includeHorizon bool) {
	foneroCorePath := scenarioName + "-core.sql"
	horizonPath := scenarioName + "-horizon.sql"

	if !includeHorizon {
		horizonPath = "blank-horizon.sql"
	}

	scenarios.Load(FoneroCoreDatabaseURL(), foneroCorePath)
	scenarios.Load(DatabaseURL(), horizonPath)
}
