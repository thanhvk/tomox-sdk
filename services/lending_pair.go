package services

import (
	"github.com/tomochain/tomox-sdk/interfaces"
	"github.com/tomochain/tomox-sdk/types"
)

// LendingPairService struct with daos required, responsible for communicating with daos.
// LendingPairService functions are responsible for interacting with daos and implements business logics.
type LendingPairService struct {
	lendingPairDao interfaces.LendingPairDao
}

// NewLendingPairService returns a new instance of balance service
func NewLendingPairService(
	lendingPairDao interfaces.LendingPairDao,
) *LendingPairService {
	return &LendingPairService{lendingPairDao}
}

// GetAll is reponsible for fetching all the pairs in the DB
func (s *LendingPairService) GetAll() ([]types.LendingPair, error) {
	return s.lendingPairDao.GetAll()
}
