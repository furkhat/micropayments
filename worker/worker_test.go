package worker

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/furkhat/micropayments/contract"
	"github.com/furkhat/micropayments/data"
	"gopkg.in/reform.v1"
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

func TestAfterApprove(t *testing.T) {
	owner := data.NewAccount()
	db.Insert(owner)
	defer db.Delete(owner)
	spender := common.HexToAddress("0xaaaaaa")
	approvedAmount := int64(123)
	logData, err := approvalNonIndexArgs.Pack(big.NewInt(approvedAmount))
	approveLog := &ethtypes.Log{
		Address: pscAddr,
		Topics: []common.Hash{
			contract.EthTokenApproval,
			common.HexToHash(owner.EthAddr),
			spender.Hash(),
		},
		Data: logData,
	}

	err = worker.AfterApprove(approveLog)
	if err != nil {
		t.Fatal(err)
	}

	auth, _ := accTransactOpts(owner)
	testEthBack.testCalled(t, "PSCAddBalanceERC20", auth.From,
		gasLimitAddBalanceERC20, big.NewInt(approvedAmount))
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
