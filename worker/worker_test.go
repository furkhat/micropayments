package worker

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/furkhat/micropayments/data"
	"gopkg.in/reform.v1"
)

var (
	db          *reform.DB
	testEthBack *testEthBackend
	worker      *Worker
)

func TestMain(m *testing.M) {
	connStr := os.Getenv("DB_CONN")
	db, _ = data.NewDBFromConnStr(connStr)
	defer data.CloseDB(db)

	testEthBack = newTestEthBackend(common.HexToAddress("0x12345"))

	worker = NewWorker(db, testEthBack)

	os.Exit(m.Run())
}

func TestAfterApprove(t *testing.T) {
	t.Skip("TODO")
}

func TestAfterTransfer(t *testing.T) {
	t.Skip("TODO")
}

func AfterChannelCreated(t *testing.T) {
	t.Skip("TODO")
}

func AfterChannelClose(t *testing.T) {
	t.Skip("TODO")
}
