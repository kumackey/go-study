package status

import (
	"errors"
	"time"
)

type Status string

const (
	Created   Status = "created"
	Submitted Status = "submitted"
	Approved  Status = "approved"
	Rejected  Status = "rejected"
)

var ErrInvalidStatusTransition = errors.New("invalid status transition") // 不正な申請ステータス遷移

type Application struct {
	applicant         string     // 申請者
	createdAt         *time.Time // 作成日時
	content           string     // 申請内容
	submittedAt       *time.Time // 提出日時
	reviewer          string     // 申請レビューア
	approvedAt        *time.Time // 承認日時
	rejectedReason    string     // 否認理由
	rejectedAt        *time.Time // 否認日時
	countOfSubmission int        // 提出回数
	status            Status     // 申請ステータス
}

func (a *Application) IsCreated() bool {
	return a.status == Created
}

func (a *Application) IsApplying() bool {
	return a.status == Submitted
}

func (a *Application) IsApproved() bool {
	return a.status == Approved
}

func (a *Application) IsRejected() bool {
	return a.status == Rejected
}

func (a *Application) Submit(content string, submittedAt time.Time) error {
	if !a.IsCreated() && !a.IsRejected() {
		return ErrInvalidStatusTransition
	}

	a.content = content
	a.submittedAt = &submittedAt
	a.countOfSubmission += 1
	a.status = Submitted

	return nil
}

func (a *Application) Approve(reviewer string, approvedAt time.Time) error {
	if !a.IsApplying() {
		return ErrInvalidStatusTransition
	}

	a.reviewer = reviewer
	a.approvedAt = &approvedAt
	a.status = Approved

	return nil
}

func (a *Application) Reject(reviewer string, rejectedReason string, rejectedAt time.Time) error {
	if !a.IsApplying() {
		return ErrInvalidStatusTransition
	}

	a.reviewer = reviewer
	a.rejectedReason = rejectedReason
	a.rejectedAt = &rejectedAt
	a.status = Rejected

	return nil
}

func NewApplication(applicant string, createdAt time.Time) *Application {
	return &Application{
		applicant: applicant,
		createdAt: &createdAt,
		status:    Created,
	}
}
