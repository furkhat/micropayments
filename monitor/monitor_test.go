package monitor

import (
	"context"
	"crypto/ecdsa"
	cryptorand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"os"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	reform "gopkg.in/reform.v1"

	"github.com/furkhat/micropayments/data"
)

type fakeEthConn struct {
	filterQueries []ethereum.FilterQuery
	fakeLogs      []ethtypes.Log
	number        uint64
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
		Number: new(big.Int).SetUint64(c.number),
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

func newTestAccount() *data.Account {
	priv, _ := ecdsa.GenerateKey(crypto.S256(), cryptorand.Reader)
	pub := base64.URLEncoding.EncodeToString(
		crypto.FromECDSAPub(&priv.PublicKey))
	addr := hex.EncodeToString(crypto.PubkeyToAddress(priv.PublicKey).Bytes())
	return &data.Account{
		ID:         data.NewUUID(),
		EthAddr:    addr,
		PublicKey:  pub,
		PrivateKey: base64.URLEncoding.EncodeToString(crypto.FromECDSA(priv)),
	}
}

var (
	ptcAddr = common.HexToAddress("0x123")
	pscAddr = common.HexToAddress("0xabc")

	someAddr = common.HexToAddress("0x111")
	someHash = common.HexToHash("0xaaa")

	db *reform.DB
)

func TestMain(m *testing.M) {
	connStr := os.Getenv("DB_CONN")
	db, _ = data.NewDBFromConnStr(connStr)
	defer data.CloseDB(db)

	os.Exit(m.Run())
}
