package horizon

import (
	"github.com/fonero-project/fonero-golang/support/db"
	"github.com/fonero-project/fonero-horizon/db2/core"
	"github.com/fonero-project/fonero-horizon/db2/history"
	"github.com/fonero-project/fonero-horizon/log"
)

func initHorizonDb(app *App) {
	session, err := db.Open("postgres", app.config.DatabaseURL)

	if err != nil {
		log.Panic(err)
	}
	session.DB.SetMaxIdleConns(4)
	session.DB.SetMaxOpenConns(12)

	app.historyQ = &history.Q{session}
}

func initCoreDb(app *App) {
	session, err := db.Open("postgres", app.config.FoneroCoreDatabaseURL)

	if err != nil {
		log.Panic(err)
	}

	session.DB.SetMaxIdleConns(4)
	session.DB.SetMaxOpenConns(12)
	app.coreQ = &core.Q{session}
}

func init() {
	appInit.Add("horizon-db", initHorizonDb, "app-context", "log")
	appInit.Add("core-db", initCoreDb, "app-context", "log")
}
