package monitor

import (
	"context"
	"database/sql"
	"math/big"
	"strconv"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	reform "gopkg.in/reform.v1"

	"github.com/furkhat/micropayments/data"
)

// EthConn defines typed wrappers for the used Ethereum RPC API methods.
type EthConn interface {
	FilterLogs(ctx context.Context,
		q ethereum.FilterQuery) ([]ethtypes.Log, error)
	HeaderByNumber(ctx context.Context,
		number *big.Int) (*ethtypes.Header, error)
	Close()
}

type WorkerFunc func(*ethtypes.Log) error

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

	workers map[common.Hash]WorkerFunc
}

// NewMonitor creates a Monitor for given url.
func NewMonitor(db *reform.DB, ethConn EthConn, period int64,
	pscAddr, ptcAddr common.Address) (*Monitor, error) {
	number, err := lastSeenBlockNumber(db)
	if err != nil {
		return nil, err
	}
	return &Monitor{
		contracts:     []common.Address{pscAddr, ptcAddr},
		db:            db,
		ethConn:       ethConn,
		errChan:       make(chan error),
		ethlogChan:    make(chan *ethtypes.Log),
		exit:          make(chan struct{}),
		lastSeenBlock: number,
		period:        period,
		workers:       map[common.Hash]WorkerFunc{},
	}, nil
}

func lastSeenBlockNumber(db *reform.DB) (uint64, error) {
	setting := &data.Setting{}
	err := db.FindByPrimaryKeyTo(setting, data.SettingsLastSeenBlock)
	if err == sql.ErrNoRows {
		err = db.Insert(&data.Setting{
			Key:   data.SettingsLastSeenBlock,
			Value: "0",
		})
		return 0, err
	}
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(setting.Value, 10, 64)
}

// RegisterWorker registers a worker for an ethereum log.
func (m *Monitor) RegisterWorker(logDigest common.Hash, f WorkerFunc) {
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
			lastUpdate := m.lastSeenBlock
			select {
			case log := <-m.ethlogChan:
				if len(log.Topics) > 0 {
					worker, ok := m.workers[log.Topics[0]]
					if ok {
						if err := worker(log); err != nil {
							m.errChan <- err
						}
					}
				}
				if log.BlockNumber > lastUpdate {
					m.updateLastSeenBlock(int(log.BlockNumber))
				}
			case <-m.exit:
				return
			}
		}
	}()

	for {
		select {
		case <-ticker:
			m.collect()
		case <-m.exit:
			return
		}
	}
}

func (m *Monitor) updateLastSeenBlock(num int) {
	settings := &data.Setting{
		Key:   data.SettingsLastSeenBlock,
		Value: strconv.Itoa(num),
	}
	if err := m.db.Update(settings); err != nil {
		m.errChan <- err
	}
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
