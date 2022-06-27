package usecase

import (
	"errors"
	"example.com/go-mod-test/go/flag/pkg"
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

	if registersLotteryWithHoge {
		// Emailを出すために該当ユーザを抽出
		user := u.userRepo.FindByUID(uid)

		// Hoge社にAPIで連携。このときに通知も飛ぶ
		hogeLotteryOrderID, err := u.hoge.postLotteries(user.Email)

		// 内部で管理するため、宝くじの注文を作成
		lotteryOrder, err := pkg.NewLotteryOrder(id, uid, "hoge", hogeLotteryOrderID)
		if err != nil {
			return "", errors.New("failed to create new lottery order")
		}

		// その注文を永続化
		err = u.lotteryOrderRepo.Save(lotteryOrder)
		if err != nil {
			return "", errors.New("failed to save lottery order")
		}

		return lotteryOrder.LotteryOrderID, nil
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
