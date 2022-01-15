package flag

import (
	"errors"
	"time"
)

type LotteryOrder struct {
	LotteryOrderID string
	UID            string
	OrderedAt      time.Time
}

func NewLotteryOrder(uID string) (*LotteryOrder, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}

	return &LotteryOrder{
		LotteryOrderID: id,
		UID:            uID,
		OrderedAt:      time.Now(),
	}, nil
}

type LotteryOrderRepository struct{}

func (r *LotteryOrderRepository) Save(_ *LotteryOrder) error {
	// クリーニングのオーダの永続化を行う
	return nil
}

type Notification struct{}

func NewNotificationOfRegisteringLotteryOrder(_ *LotteryOrder) (*Notification, error) {
	// 宝くじの注文をした旨の通知を作る
	return &Notification{}, nil
}

type NotificationRepository struct{}

func (r *NotificationRepository) Save(_ *Notification) error {
	// 通知の永続化を行う
	return nil
}

type AClient struct{}

func GenerateID() (string, error) {
	// ランダムにID生成する感じの処理
	return "12345678901234567890", nil
}

type RegisterLotteryOrderUsecase struct {
	aClient          *AClient
	lotteryOrderRepo *LotteryOrderRepository
	notificationRepo *NotificationRepository
}

func (u *RegisterLotteryOrderUsecase) exec(uid string) (lotteryOrderID string, err error) {
	lotteryOrder, err := NewLotteryOrder(uid)
	if err != nil {
		return "", errors.New("failed to create new lottery order")
	}

	err = u.lotteryOrderRepo.Save(lotteryOrder)
	if err != nil {
		return "", errors.New("failed to save lottery order")
	}

	notif, err := NewNotificationOfRegisteringLotteryOrder(lotteryOrder)
	if err != nil {
		return "", errors.New("failed to create notification")
	}

	err = u.notificationRepo.Save(notif)
	if err != nil {
		return "", errors.New("failed to save notification")
	}

	return lotteryOrder.LotteryOrderID, nil
}
