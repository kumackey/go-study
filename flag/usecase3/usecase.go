package usecase1

import (
	"errors"
	"example.com/go-mod-test/flag/pkg"
)

const registersLotteryWithHoge = false

type HogeClient struct{}

func (c *HogeClient) postLotteries(_ string) (string, error) {
	// emailを使って、Hoge社の宝くじの注文するAPIを叩いて、Hoge社管理の宝くじIDが返ってくるとしましょう
	return "12345", nil
}

type RegisterLotteryOrderUsecase struct {
	hoge             *HogeClient
	userRepo         *pkg.UserRepository
	lotteryOrderRepo *pkg.LotteryOrderRepository
	notificationRepo *pkg.NotificationRepository
}

func (u *RegisterLotteryOrderUsecase) exec(uid string) (lotteryOrderID string, err error) {
	id, err := pkg.GenerateID()
	if err != nil {
		return "", err
	}

	if !registersLotteryWithHoge {
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

	user := u.userRepo.FindByUID(uid)

	hogeLotteryOrderID, err := u.hoge.postLotteries(user.Email)

	lotteryOrder, err := pkg.NewLotteryOrder(id, uid, "hoge", hogeLotteryOrderID)
	if err != nil {
		return "", errors.New("failed to create new lottery order")
	}

	err = u.lotteryOrderRepo.Save(lotteryOrder)
	if err != nil {
		return "", errors.New("failed to save lottery order")
	}

	return lotteryOrder.LotteryOrderID, nil
}
