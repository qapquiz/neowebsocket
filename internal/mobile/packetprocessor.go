package mobile

import (
	"github.com/qapquiz/neowebsocket/pkg/websocket"
	"github.com/qapquiz/packet"
	"go.uber.org/zap"
)

const (
	csLogin                            uint16 = 10001
	csChat                             uint16 = 10002
	csDummyJoin                        uint16 = 10003
	csDummySit                         uint16 = 10005
	csDummyStandUp                     uint16 = 10006
	csDummyReady                       uint16 = 10007
	csDummyUnready                     uint16 = 10008
	csDummyStart                       uint16 = 10009
	csDummyGeb                         uint16 = 10010
	csDummyJua                         uint16 = 10011
	csDummyGerd                        uint16 = 10012
	csDummyFaak                        uint16 = 10013
	csDummyTing                        uint16 = 10014
	csDummyKnock                       uint16 = 10015
	csDummyDarkKnock                   uint16 = 10016
	csDummyChangeSeat                  uint16 = 10017
	csEmoticon                         uint16 = 10018
	csLobbyCreateRoom                  uint16 = 10019 // ไม่มีใน mobile client
	csLobbyRefreshNumUserPlayingInRoom uint16 = 10020
	csDummyUserProfile                 uint16 = 10021
	csShowMeJuaCard                    uint16 = 10022
	csDummyInbox                       uint16 = 10024
	csDummyAcceptInbox                 uint16 = 10025
	csDummyChooseDailyReward           uint16 = 10026
	csDummyEnd                         uint16 = 10028
	csAniDone                          uint16 = 10031
	csMakeOrder                        uint16 = 10032
	csSticker                          uint16 = 10033
	csDummyQuickJoin                   uint16 = 10034 // ปุ่ม เล่นเลย ห้องไหนก็ได้
	csDummyQuickJoinWithFilter         uint16 = 10035 // ปุ่ม เริ่มเล่น
	csDummyRequestGameTurnData         uint16 = 10037
	csOKDialogReconnectMoka            uint16 = 10038
	csOKDialogReconnectEndGame         uint16 = 10039
	csRequestDailyEmergencyChip        uint16 = 10041
	csSoundBackground                  uint16 = 10042
	csSoundEffect                      uint16 = 10043
	csVibrate                          uint16 = 10047
	csBuyLotto                         uint16 = 10044
	csGetLottoChip                     uint16 = 10045
	csClosePopupLottoLose              uint16 = 10046
	csChangeServer                     uint16 = 10048
	csUserCCU                          uint16 = 10049
	csReceiveMobileApplicationState    uint16 = 10051
	csCheckReachableServer             uint16 = 10054
	csAction                           uint16 = 10062
	csRequestBuyTicket                 uint16 = 10063
	csLobbyCreateRoomMobile            uint16 = 10064 // สร้างห้องส่วนตัว
	csRequestShareDeepLink             uint16 = 10067
	csRequestClaimDeepLink             uint16 = 10068
	csRequestChangeNameAndProfilePic   uint16 = 10058
	csDummyObserverJoin                uint16 = 10070
	csDummyJoinAutoChangeServer        uint16 = 10071
	csChangeServerToJoin               uint16 = 10072
	csObserverRequestReconnectRoom     uint16 = 10073
	csObserverRequestCancelRoom        uint16 = 10074
	csDummyObserverLeave               uint16 = 10075
	csUserRequestChangeName            uint16 = 10076
	csRequestCheckChangeName           uint16 = 10077
	csUserRequestChangeFreeName        uint16 = 10079
	csUserSaveProfileItem              uint16 = 10080
	csRequestJoinTournamentServer      uint16 = 10082
	csRequestCheckUserBlacklist        uint16 = 10083
	csRequestJoinMatchingRoom          uint16 = 10084
	csRequestCancelMatching            uint16 = 10085
	csRequestWaitingUserReady          uint16 = 10086
	csRequestLeaveTournament           uint16 = 10087
	csRequestBuyCredit                 uint16 = 10088
	csRequestRecvTournamentReward      uint16 = 10089
	csRequestTournaamentRoomList       uint16 = 10090
	csRequestSlotMachineStart          uint16 = 10096
	csRequestShareChest                uint16 = 10097
	csRequestSlotMachineHistory        uint16 = 10098
	scError                            uint16 = 20000 // Send uint16 to Client start here!
	scLoggedIn                         uint16 = 20001
	scChat                             uint16 = 20002
	scDummyJoin                        uint16 = 20003
	scDummySit                         uint16 = 20005
	scDummyStandUp                     uint16 = 20006
	scDummyReady                       uint16 = 20007
	scDummyUnready                     uint16 = 20008
	scDummyStartRound                  uint16 = 20009
	scDummyGeb                         uint16 = 20010
	scDummyJua                         uint16 = 20011
	scDummyGerd                        uint16 = 20012
	scDummyFaak                        uint16 = 20013
	scDummyTing                        uint16 = 20014
	scDummyKnock                       uint16 = 20015
	scDummyRoundscore                  uint16 = 20018
	scDummySpecialGebHua               uint16 = 20019
	scDummySpecialSpeto2Club           uint16 = 20020
	scDummySpecialSpetoQSpades         uint16 = 20021
	scDummyChangeSeat                  uint16 = 20022
	scGiveXp                           uint16 = 20023
	scDummyAllscore                    uint16 = 20024
	scEmoticon                         uint16 = 20025
	scDummyAniTingPeeHua               uint16 = 20027
	scDummyAniTingPeeSpeto             uint16 = 20028
	scDummyAniWasFaakedSpeto           uint16 = 20039
	scLobbyRoomListRefresh             uint16 = 20040
	scLevelUp                          uint16 = 20045
	scDummyAITakePlace                 uint16 = 20046
	scDummyUserProfile                 uint16 = 20047
	scDummyInbox                       uint16 = 20050
	scDummyNewInbox                    uint16 = 20051
	scDummyDailyReward                 uint16 = 20052
	scDummyChooseDailyReward           uint16 = 20053
	scSetActivePlayer                  uint16 = 20055
	scDummyWarningPeeHua               uint16 = 20056
	scDummyWarningPeeSpeto             uint16 = 20057
	scDummyKick                        uint16 = 20060
	scGMAnnounce                       uint16 = 20061
	scAni                              uint16 = 20062
	scRemoveUserFromRoom               uint16 = 20063
	scAlert                            uint16 = 20064
	scBatchItem                        uint16 = 20067 // Get items from batch.
	scOrderResult                      uint16 = 20068
	scItemShop                         uint16 = 20069
	scSticker                          uint16 = 20072
	scReconnectRoomState               uint16 = 20073
	scAnotherUserReconnected           uint16 = 20074
	scDummyResponseGameTurnData        uint16 = 20075
	scShowDialogReconnectMoka          uint16 = 20076
	scShowDialogReconnectEndgame       uint16 = 20077
	scResponseDailyEmergencyChip       uint16 = 20080
	scLottoResult                      uint16 = 20081
	scLottoRewardHistory               uint16 = 20082
	scLottoServerStatus                uint16 = 20083
	scLottoReward                      uint16 = 20084
	scLottoRewardChip                  uint16 = 20085
	scChangeServer                     uint16 = 20088
	scUserCCU                          uint16 = 20089
	scStartWaitReadyTimeout            uint16 = 20092
	scStopWaitReadyTimeout             uint16 = 20093
	scUpdateFirstPurchaseGold          uint16 = 20094
	scResponseCheckReachableServer     uint16 = 20096
	scRemoveBanner                     uint16 = 20101
	scRemoveInterstitial               uint16 = 20102
	scResponseDailyChips               uint16 = 20104
	scFlashDeal                        uint16 = 20107
	scAction                           uint16 = 20108
	scShowBuyTicketPopup               uint16 = 20109
	scResponseReceivedTicket           uint16 = 20110
	scResponsePersuadeChipToRoom       uint16 = 20111
	scResponseUserTicketLeft           uint16 = 20113
	scResponseClaim                    uint16 = 20116
	scDummyObserverJoin                uint16 = 20118
	scObserverJoinAfterStarted         uint16 = 20119
	scObserverStartRound               uint16 = 20120
	scObserverSeeJua                   uint16 = 20121
	scResponseUserJoinAutoChangeServer uint16 = 20122
	scChangeServerToJoin               uint16 = 20123
	scNumberUserObserver               uint16 = 20124
	scDialogReconnectEndgameObserve    uint16 = 20125
	scDialogObserverReconnect          uint16 = 20126
	scDialogForceObserverLeave         uint16 = 20127
	scResponseDialogChangeName         uint16 = 20129
	scResultChangeName                 uint16 = 20130
	scResponseUserSummaryPay           uint16 = 20131
	scSendProfileItem                  uint16 = 20132
	scSendProfileItemExpired           uint16 = 20133
	scSendDialogBuyGold                uint16 = 20134
	scResponseJoinTournamentServer     uint16 = 20135
	scResponseUserCheckBlackList       uint16 = 20136
	scResponseJoinMatchingRoom         uint16 = 20137
	scResponseJoinWaitingRoom          uint16 = 20138
	scResponseCancelMatching           uint16 = 20139
	scResponseWaitingUserReady         uint16 = 20140
	scResponseAllUserJoinTable         uint16 = 20141
	scResponseAlertHint                uint16 = 20142
	scMakeUserBlacklist                uint16 = 20143
	scResponseTournamentRound          uint16 = 20144
	scResponseLeaveTournament          uint16 = 20145
	scResponseBuyCredit                uint16 = 20146
	scMakeTournamentDialogResult       uint16 = 20147
	scResponseTournamentRecvReward     uint16 = 20148
	scResponseTournamentRoomList       uint16 = 20149
	scResponsePersuadeChip             uint16 = 20150
	scDialogReconnectEndGameTournament uint16 = 20151
	scResponseSlotMachineResult        uint16 = 20156
	scResponseSlotMachinePayoutRate    uint16 = 20159
	scResponseSlotMachineHistory       uint16 = 20160
	scMakeCancelSlotMachine            uint16 = 20163
)

// PacketProcessor will hold every packet and packet function for mobile client
type PacketProcessor struct {
	mapper map[uint16]func(*websocket.Remote, *packet.Reader)
}

// NewPacketProcessor will hold packet id and
func NewPacketProcessor() PacketProcessor {
	return PacketProcessor{
		mapper: map[uint16]func(*websocket.Remote, *packet.Reader){
			csLogin: receiveLogin,
		},
	}
}

// GetPacketFunc will return packet function that associate with packet id
func (pp PacketProcessor) GetPacketFunc(packetID uint16) func(*websocket.Remote, *packet.Reader) {
	packetFunc, ok := pp.mapper[packetID]
	if !ok {
		zap.S().Error("there is no packetID: %d", packetID)
	}

	return packetFunc
}

func receiveLogin(remote *websocket.Remote, pr *packet.Reader) {

}
