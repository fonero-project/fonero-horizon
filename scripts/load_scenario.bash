#! /usr/bin/env bash

set -e

SCENARIO=$1
CORE_SQL=./src/github.com/fonero-project/fonero-horizon/test/scenarios/$SCENARIO-core.sql
HORIZON_SQL=./src/github.com/fonero-project/fonero-horizon/test/scenarios/$SCENARIO-horizon.sql

echo "psql $FONERO_CORE_DATABASE_URL < $CORE_SQL" 
psql $FONERO_CORE_DATABASE_URL < $CORE_SQL 
echo "psql $DATABASE_URL < $HORIZON_SQL"
psql $DATABASE_URL < $HORIZON_SQL 
