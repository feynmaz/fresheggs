package eth

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
	"github.com/feynmaz/fresheggs/internal/types"
)

type EthClient struct {
	log logger.Logger
	cfg config.EthClient

	client *ethclient.Client
}

func NewClient(cfg config.EthClient, log logger.Logger) (*EthClient, error) {
	client, err := ethclient.Dial(cfg.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to dial eth client: %w", err)
	}

	return &EthClient{
		log:    log,
		cfg:    cfg,
		client: client,
	}, nil
}

func (e *EthClient) GetLastBlockNumber() (*types.BlockNumber, error) {
	block, err := e.client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get block by number: %w", err)
	}

	return &types.BlockNumber{
		ChainName: e.cfg.ChainName,
		Number:    block.Number().String(),
	}, nil
}
