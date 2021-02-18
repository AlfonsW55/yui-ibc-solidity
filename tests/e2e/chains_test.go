package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/datachainlab/ibc-solidity/pkg/contract"
	"github.com/datachainlab/ibc-solidity/pkg/ibc/channel"
	channeltypes "github.com/datachainlab/ibc-solidity/pkg/ibc/channel"
	ibctesting "github.com/datachainlab/ibc-solidity/pkg/testing"
	testchain0 "github.com/datachainlab/ibc-solidity/tests/e2e/config/chain0"
	testchain1 "github.com/datachainlab/ibc-solidity/tests/e2e/config/chain1"
	"github.com/stretchr/testify/suite"
)

const mnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type ChainTestSuite struct {
	suite.Suite

	coordinator ibctesting.Coordinator
	chainA      *ibctesting.Chain
	chainB      *ibctesting.Chain
}

func (suite *ChainTestSuite) SetupTest() {
	chainClientA, err := contract.CreateClient("http://127.0.0.1:8645")
	suite.Require().NoError(err)

	chainClientB, err := contract.CreateClient("http://127.0.0.1:8745")
	suite.Require().NoError(err)

	suite.chainA = ibctesting.NewChain(suite.T(), 2018, *chainClientA, testchain0.Contract, mnemonicPhrase, uint64(time.Now().UnixNano()))
	suite.chainB = ibctesting.NewChain(suite.T(), 3018, *chainClientB, testchain1.Contract, mnemonicPhrase, uint64(time.Now().UnixNano()))
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), suite.chainA, suite.chainB)
}

func (suite ChainTestSuite) TestChannel() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, ibctesting.BesuIBFT2Client)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.TransferPort, ibctesting.TransferPort, channel.UNORDERED)

	// TODO refactoring follow steps:

	// // TODO give a dynamic height
	// packet := channeltypes.NewPacket([]byte("data"), 1, chanB.PortID, chanB.ID, chanA.PortID, chanA.ID, channeltypes.Height{0, 1000000}, 0)
	// suite.Require().NoError(suite.coordinator.SendPacket(ctx, chainA, chainB, packet, chanA.CounterpartyClientID))
	// suite.Require().NoError(suite.coordinator.RecvPacket(ctx, chainB, chainA, chanB, chanA, packet, chanB.CounterpartyClientID))

	balanceA, err := chainA.SimpletokenModule.BalanceOf(chainA.CallOpts(ctx), chainA.CallOpts(ctx).From)
	suite.Require().NoError(err)

	balanceB, err := chainB.SimpletokenModule.BalanceOf(chainB.CallOpts(ctx), chainB.CallOpts(ctx).From)
	suite.Require().NoError(err)

	// TODO modify an address of recipient
	suite.Require().NoError(
		suite.chainA.WaitIfNoError(ctx)(suite.chainA.SimpletokenModule.CrossTransfer(suite.chainA.TxOpts(ctx), chanA.PortID, chanA.ID, suite.chainA.CallOpts(ctx).From, 100, 1000000)),
	)
	chainA.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainB, chainA, clientB, ibctesting.BesuIBFT2Client))

	packet, err := chainA.IBCStore.GetPacket(chainA.CallOpts(ctx), chanA.PortID, chanA.ID, 1)
	suite.Require().NoError(err)
	transferPacket := channel.NewPacket(packet.Data, packet.Sequence, packet.SourcePort, packet.SourceChannel, packet.DestinationPort, packet.DestinationChannel, channeltypes.Height(packet.TimeoutHeight), packet.TimeoutTimestamp)
	suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainB, chainA, chanB, chanA, transferPacket))

	postBalanceA, err := chainA.SimpletokenModule.BalanceOf(chainA.CallOpts(ctx), chainA.CallOpts(ctx).From)
	suite.Require().NoError(err)
	suite.Require().EqualValues(balanceA.Uint64()-100, postBalanceA.Uint64())

	postBalanceB, err := chainB.SimpletokenModule.BalanceOf(chainB.CallOpts(ctx), chainB.CallOpts(ctx).From)
	suite.Require().NoError(err)
	suite.Require().EqualValues(balanceB.Uint64()+100, postBalanceB.Uint64())

	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainA, chainB, chanA, chanB, transferPacket, []byte{1}))
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ChainTestSuite))
}
