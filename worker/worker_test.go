package worker

import (
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/reform.v1"

	"github.com/furkhat/micropayments/contract"
	"github.com/furkhat/micropayments/data"
)

var (
	db          *reform.DB
	testEthBack *testEthBackend
	worker      *Worker

	pscAddr = common.HexToAddress("0x12345")
	ptcAddr = common.HexToAddress("0x54321")
)

func TestMain(m *testing.M) {
	connStr := os.Getenv("DB_CONN")
	db, _ = data.NewDBFromConnStr(connStr)
	defer data.CloseDB(db)

	testEthBack = newTestEthBackend()

	worker = NewWorker(db, testEthBack)

	os.Exit(m.Run())
}

func TestAfterChannelCreated(t *testing.T) {
	acc := data.NewAccount()
	db.Insert(acc)
	defer db.Delete(acc)

	deposit := int64(123)
	logData, _ := channelCreatedNonIndexArgs.Pack(big.NewInt(deposit),
		common.HexToHash("0x0"))
	channelCreatedLog := &ethtypes.Log{
		Address: pscAddr,
		Topics: []common.Hash{
			contract.EthChannelCreated,
			common.HexToHash(acc.EthAddr),
			common.HexToHash("0x001"),
			common.BytesToHash([]byte{}),
		},
		Data: logData,
	}

	err := worker.AfterChannelCreated(channelCreatedLog)
	if err != nil {
		t.Fatal(err)
	}

	channel := &data.Channel{}
	if err := db.FindOneTo(channel, "agent", acc.EthAddr); err != nil {
		t.Fatal(err)
	}

	select {
	case <-time.After(time.Second):
		t.Fatal("session start was not send")
	case id := <-worker.sessStart:
		if channel.ID != id {
			t.Fatal("wrong channel id was send to session start")
		}
	}
}
