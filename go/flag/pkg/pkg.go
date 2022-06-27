package pkg

import "time"

func GenerateID() (string, error) {
	// ランダムにID生成する感じの処理
	return "12345678901234567890", nil
}

type LotteryOrder struct {
	LotteryOrderID string
	UID            string
	Provider       string
	ProviderID     string
	OrderedAt      time.Time
}

func NewLotteryOrder(iD string, uID string, provider string, providerID string) (*LotteryOrder, error) {
	return &LotteryOrder{
		LotteryOrderID: iD,
		UID:            uID,
		Provider:       provider,
		ProviderID:     providerID,
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

type UserRepository struct{}

type User struct {
	UID   string
	Email string
}

func (r *UserRepository) FindByUID(uid string) *User {
	// 実際には永続層から該当UIDのユーザを抽出する
	return &User{
		UID:   uid,
		Email: "hoge-test-2022@example.com",
	}
}
