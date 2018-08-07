package monitor

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// collect gets interested ethereum logs and sends them to the ethereum logs channel for further processing.
func (m *Monitor) collect() {
	header, err := m.ethConn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		m.errChan <- err
		return
	}

	from := new(big.Int).SetUint64(m.lastSeenBlock + 1)
	to := new(big.Int).SetUint64(header.Number.Uint64() - 1)

	addresses, err := m.getAddresses()
	if err != nil {
		fmt.Println(err)
		m.errChan <- err
		return
	}

	filterQueries := []ethereum.FilterQuery{
		{
			FromBlock: from,
			ToBlock:   to,
			Addresses: m.contracts,
			Topics:    [][]common.Hash{nil, addresses},
		},
		{
			FromBlock: from,
			ToBlock:   to,
			Addresses: m.contracts,
			Topics:    [][]common.Hash{nil, nil, addresses},
		},
	}
	for _, q := range filterQueries {
		logs, _ := m.ethConn.FilterLogs(context.Background(), q)
		for _, log := range logs {
			if log.BlockNumber > m.lastSeenBlock {
				m.lastSeenBlock = log.BlockNumber
			}
			go func(log ethtypes.Log) {
				m.ethlogChan <- &log
			}(log)
		}
	}
}

func (m *Monitor) getAddresses() ([]common.Hash, error) {
	rows, err := m.db.Query(`SELECT eth_addr FROM accounts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []common.Hash
	for rows.Next() {
		var addrHex string
		if err := rows.Scan(&addrHex); err != nil {
			return nil, fmt.Errorf("failed to scan rows: %v", err)
		}
		addresses = append(addresses, common.HexToHash(addrHex))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return addresses, nil
}
