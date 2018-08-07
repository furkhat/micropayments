package monitor

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	reform "gopkg.in/reform.v1"

	"github.com/furkhat/micropayments/data"
)

type fakeEthConn struct {
	filterQueries []ethereum.FilterQuery
	fakeLogs      []ethtypes.Log
	lastBlock     uint64

	filterCall chan struct{}
}

func (c *fakeEthConn) FilterLogs(ctx context.Context,
	q ethereum.FilterQuery) ([]ethtypes.Log, error) {
	c.filterQueries = append(c.filterQueries, q)
	ret := c.fakeLogs
	c.fakeLogs = []ethtypes.Log{}
	return ret, nil
}

func (c *fakeEthConn) HeaderByNumber(ctx context.Context,
	number *big.Int) (*ethtypes.Header, error) {
	return &ethtypes.Header{
		Number: new(big.Int).SetUint64(c.lastBlock),
	}, nil
}

func (c *fakeEthConn) Close() {}

type mockTicker struct {
	C chan time.Time
}

func newMockTicker() *mockTicker {
	return &mockTicker{C: make(chan time.Time, 1)}
}

func (t *mockTicker) tick() {
	select {
	case t.C <- time.Now():
	default:
	}
}

func insertToTestDB(t *testing.T, str reform.Struct) {
	err := db.Insert(str)
	if err != nil {
		t.Fatal(err)
	}
}

func deleteFromTestDB(t *testing.T, rec reform.Record) {
	err := db.Delete(rec)
	if err != nil {
		t.Fatal(err)
	}
}

type testdata struct {
	mon      *Monitor
	fakeConn *fakeEthConn
}

func newTestData(t *testing.T, lastBlock, lastSeen uint64) *testdata {
	fakeConn := &fakeEthConn{
		lastBlock:  lastBlock,
		filterCall: make(chan struct{}),
	}
	if lastSeen > 0 {
		insertToTestDB(t, &data.Setting{
			Key:   data.SettingsLastSeenBlock,
			Value: "10",
		})
	}
	mon, err := NewMonitor(db, fakeConn, 1, pscAddr, ptcAddr)
	if err != nil {
		t.Fatal(err)
	}
	return &testdata{mon, fakeConn}
}

func (d *testdata) cleanUp(t *testing.T) {
	db.DeleteFrom(data.SettingTable, "WHERE key=$1", data.SettingsLastSeenBlock)
}

var (
	ptcAddr = common.HexToAddress("0x123")
	pscAddr = common.HexToAddress("0xabc")
	db      *reform.DB
)

func TestMain(m *testing.M) {
	connStr := os.Getenv("DB_CONN")
	db, _ = data.NewDBFromConnStr(connStr)
	defer data.CloseDB(db)

	os.Exit(m.Run())
}
