package usecase1

import (
	"errors"
	"example.com/go-mod-test/flag/pkg"
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

	lotteryOrder, err := pkg.NewLotteryOrder(id, uid, "", "")
	if err != nil {
		return "", errors.New("failed to create new lottery order")
	}

	err = u.lotteryOrderRepo.Save(lotteryOrder)
	if err != nil {
		return "", errors.New("failed to save lottery order")
	}

	notif, err := pkg.NewNotificationOfRegisteringLotteryOrder(lotteryOrder)
	if err != nil {
		return "", errors.New("failed to create notification")
	}

	err = u.notificationRepo.Save(notif)
	if err != nil {
		return "", errors.New("failed to save notification")
	}

	return lotteryOrder.LotteryOrderID, nil

}
