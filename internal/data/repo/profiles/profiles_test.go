package profiles

import (
	"context"
	"path/filepath"
	"time"

	"github.com/London57/profiles/pkg/testhelpers"
	"github.com/stretchr/testify/suite"
)

var (
	TableCreateScriptPath = filepath.Join("internal", "data", "preferences_table.sql")
)

type ProfilesRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repo ProfilesRepo
}


func (suite *ProfilesRepoTestSuite) SetUpSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	container, err := testhelpers.RunPostgresContainer(
		ctx,
		[]string{TableCreateScriptPath},
		""
	)
}