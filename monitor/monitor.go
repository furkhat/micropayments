package monitor

import (
	"context"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	reform "gopkg.in/reform.v1"
)

// EthConn defines typed wrappers for the used Ethereum RPC API methods.
type EthConn interface {
	FilterLogs(ctx context.Context,
		q ethereum.FilterQuery) ([]ethtypes.Log, error)
	HeaderByNumber(ctx context.Context,
		number *big.Int) (*ethtypes.Header, error)
	Close()
}

type workerFunc func(*ethtypes.Log) error

// Monitor monitors blockchain for interested ethereum logs and handles them.
type Monitor struct {
	contracts  []common.Address
	db         *reform.DB
	ethConn    EthConn
	ethlogChan chan *ethtypes.Log
	errChan    chan error
	exit       chan struct{}
	period     int64 // In milliseconds.

	lastSeenBlock uint64

	workers map[common.Hash]workerFunc
}

// NewMonitor creates a Monitor for given url.
func NewMonitor(db *reform.DB, ethConn EthConn, period int64,
	pscAddr, ptcAddr common.Address) *Monitor {
	return &Monitor{
		contracts:  []common.Address{pscAddr, ptcAddr},
		db:         db,
		ethConn:    ethConn,
		errChan:    make(chan error),
		ethlogChan: make(chan *ethtypes.Log),
		exit:       make(chan struct{}),
		period:     period,
		workers:    map[common.Hash]workerFunc{},
	}
}

// RegisterWorker registers a worker for an ethereum log.
func (m *Monitor) RegisterWorker(logDigest common.Hash, f workerFunc) {
	m.workers[logDigest] = f
}

// Start opens a connection and starts monitoring process.
func (m *Monitor) Start() {
	ticker := time.Tick(time.Duration(m.period) * time.Millisecond)
	m.start(ticker)
}

func (m *Monitor) start(ticker <-chan time.Time) {
	go func() {
		for {
			select {
			case <-ticker:
				m.collect()
			case <-m.exit:
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case log := <-m.ethlogChan:
				if len(log.Topics) > 0 {
					h, ok := m.workers[log.Topics[0]]
					if ok {
						if err := h(log); err != nil {
							m.errChan <- err
						}
					}
				}
			case <-m.exit:
				return
			}
		}
	}()
}

// Stop stops monitoring process and closes the connection.
func (m *Monitor) Stop() {
	m.ethConn.Close()
	m.exit <- struct{}{}
}

// Err returns the first unprocessed error that was encountered during monitoring.
func (m *Monitor) Err() error {
	return <-m.errChan
}
