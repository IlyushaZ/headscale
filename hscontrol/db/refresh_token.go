package db

import (
	"github.com/juanfont/headscale/hscontrol/types"
	"gorm.io/gorm"
)

func (hsdb *HSDatabase) GetRefreshTokens(nodeIDs ...types.NodeID) (map[types.NodeID]*types.RefreshToken, error) {
	return Read(hsdb.DB, func(tx *gorm.DB) (map[types.NodeID]*types.RefreshToken, error) {
		return GetRefreshTokens(tx, nodeIDs...)
	})
}

func GetRefreshTokens(tx *gorm.DB, nodeIDs ...types.NodeID) (map[types.NodeID]*types.RefreshToken, error) {
	tokens := []*types.RefreshToken{}
	result := make(map[types.NodeID]*types.RefreshToken)

	if len(nodeIDs) > 0 {
		tx = tx.Where("node_id IN ?", nodeIDs)
	}

	if err := tx.Find(&tokens).Error; err != nil {
		return nil, err
	}

	for _, t := range tokens {
		result[t.NodeID] = t
	}

	return result, nil
}

func (hsdb *HSDatabase) SaveRefreshToken(token *types.RefreshToken) (*types.RefreshToken, error) {
	return Write(hsdb.DB, func(tx *gorm.DB) (*types.RefreshToken, error) {
		return SaveRefreshToken(tx, token)
	})
}

func SaveRefreshToken(tx *gorm.DB, token *types.RefreshToken) (*types.RefreshToken, error) {
	if err := tx.Save(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}
