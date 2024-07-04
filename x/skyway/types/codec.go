package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.NewLegacyAmino()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterInterfaces registers the interfaces for the proto stuff
// nolint: exhaustruct
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendToEth{},
		&MsgConfirmBatch{},
		&MsgSendToPalomaClaim{},
		&MsgBatchSendToEthClaim{},
		&MsgCancelSendToEth{},
		&MsgSubmitBadSignatureEvidence{},
	)

	registry.RegisterInterface(
		"palomachain.paloma.skyway.EthereumClaim",
		(*EthereumClaim)(nil),
		&MsgSendToPalomaClaim{},
		&MsgBatchSendToEthClaim{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&SetERC20ToDenomProposal{},
		&SetBridgeTaxProposal{},
		&SetBridgeTransferLimitProposal{},
	)

	registry.RegisterInterface("palomachain.paloma.skyway.EthereumSigned", (*EthereumSigned)(nil),
		&OutgoingTxBatch{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterCodec registers concrete types on the Amino codec
// nolint: exhaustruct
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*EthereumClaim)(nil), nil)
	cdc.RegisterConcrete(&MsgSendToEth{}, "skyway/MsgSendToEth", nil)
	cdc.RegisterConcrete(&MsgConfirmBatch{}, "skyway/MsgConfirmBatch", nil)
	cdc.RegisterConcrete(&MsgSendToPalomaClaim{}, "skyway/MsgSendToPalomaClaim", nil)
	cdc.RegisterConcrete(&MsgBatchSendToEthClaim{}, "skyway/MsgBatchSendToEthClaim", nil)
	cdc.RegisterConcrete(&OutgoingTxBatch{}, "skyway/OutgoingTxBatch", nil)
	cdc.RegisterConcrete(&MsgCancelSendToEth{}, "skyway/MsgCancelSendToEth", nil)
	cdc.RegisterConcrete(&OutgoingTransferTx{}, "skyway/OutgoingTransferTx", nil)
	cdc.RegisterConcrete(&ERC20Token{}, "skyway/ERC20Token", nil)
	cdc.RegisterConcrete(&IDSet{}, "skyway/IDSet", nil)
	cdc.RegisterConcrete(&Attestation{}, "skyway/Attestation", nil)
	cdc.RegisterConcrete(&MsgSubmitBadSignatureEvidence{}, "skyway/MsgSubmitBadSignatureEvidence", nil)
	cdc.RegisterConcrete(&SetERC20ToDenomProposal{}, "skyway/SetERC20ToDenomProposal", nil)
	cdc.RegisterConcrete(&SetBridgeTaxProposal{}, "skyway/SetBridgeTaxProposal", nil)
	cdc.RegisterConcrete(&SetBridgeTransferLimitProposal{}, "skyway/SetBridgeTransferLimitProposal", nil)
}