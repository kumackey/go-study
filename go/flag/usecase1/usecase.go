package usecase

import (
	"errors"
	"example.com/go-mod-test/go/flag/pkg"
)

type RegisterLotteryOrderUsecase struct {
	lotteryOrderRepo *pkg.LotteryOrderRepository
	notificationRepo *pkg.NotificationRepository
}

func (u *RegisterLotteryOrderUsecase) exec(uid string) (lotteryOrderID string, err error) {
	id, err := pkg.GenerateID()
	if err != nil {
		return "", err
	}

	// 宝くじの注文を作成
	lotteryOrder, err := pkg.NewLotteryOrder(id, uid, "", "")
	if err != nil {
		return "", errors.New("failed to create new lottery order")
	}

	// その注文を永続化
	err = u.lotteryOrderRepo.Save(lotteryOrder)
	if err != nil {
		return "", errors.New("failed to save lottery order")
	}

	// 宝くじの注文の通知を作成
	notif, err := pkg.NewNotificationOfRegisteringLotteryOrder(lotteryOrder)
	if err != nil {
		return "", errors.New("failed to create notification")
	}

	// その通知を永続化。通知も飛ばしてくれるとする
	err = u.notificationRepo.Save(notif)
	if err != nil {
		return "", errors.New("failed to save notification")
	}

	return lotteryOrder.LotteryOrderID, nil
}
