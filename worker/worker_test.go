package worker

import (
	"encoding/base64"
	"math/big"
	"os"
	"strings"
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
	logData, _ := approvalNonIndexArgs.Pack(big.NewInt(approvedAmount))
	approveLog := &ethtypes.Log{
		Address: pscAddr,
		Topics: []common.Hash{
			contract.EthTokenApproval,
			common.HexToHash(owner.EthAddr),
			spender.Hash(),
		},
		Data: logData,
	}

	err := worker.AfterApprove(approveLog)
	if err != nil {
		t.Fatal(err)
	}

	auth, _ := accTransactOpts(owner)
	testEthBack.testCalled(t, "PSCAddBalanceERC20", auth.From,
		gasLimitAddBalanceERC20, big.NewInt(approvedAmount))
}

func TestAfterTransfer(t *testing.T) {
	acc1 := data.NewAccount()
	db.Insert(acc1)
	defer db.Delete(acc1)

	acc2 := data.NewAccount()
	db.Insert(acc2)
	defer db.Delete(acc2)

	transferLog := &ethtypes.Log{
		Address: pscAddr,
		Topics: []common.Hash{
			contract.EthTokenTransfer,
			common.HexToHash(acc1.EthAddr),
			common.HexToHash(acc2.EthAddr),
		},
	}

	testEthBack.balanceEth = big.NewInt(10)
	testEthBack.balancePSC = big.NewInt(20)
	testEthBack.balancePTC = big.NewInt(30)
	expectedEth := base64.URLEncoding.EncodeToString(big.NewInt(10).Bytes())
	expectedPSC := uint64(20)
	expectedPTC := uint64(30)

	err := worker.AfterTransfer(transferLog)
	if err != nil {
		t.Fatal(err)
	}

	db.Reload(acc1)
	db.Reload(acc2)

	if strings.TrimSpace(string(acc1.EthBalance)) != expectedEth ||
		acc1.PSCBalance != expectedPSC || acc1.PTCBalance != expectedPTC {
		t.Fatalf("balance of sender not updated")
	}

	if strings.TrimSpace(string(acc2.EthBalance)) != expectedEth ||
		acc2.PSCBalance != expectedPSC || acc2.PTCBalance != expectedPTC {
		t.Fatal("balance of receiver not updated")
	}
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
			contract.EthTokenApproval,
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
}

func TestAfterChannelClose(t *testing.T) {
	t.Skip("TODO")
}
