package status

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func submit() *Application {
	app := NewApplication("さとう　しんせいしゃ", time.Now())
	_ = app.Submit("審査される何か内容", time.Now())

	return app
}

func TestApplication_Submit(t *testing.T) {
	app := NewApplication("さとう　しんせいしゃ", time.Now()) // 申請を作成
	assert.Equal(t, false, app.IsApplying())        // 申請を作成したばかりなので申請中でない

	err := app.Submit("審査される何か内容", time.Now()) // 申請を提出
	assert.Nil(t, err)
	assert.Equal(t, true, app.IsApplying())  // 申請中である
	assert.Equal(t, false, app.IsApproved()) // 申請中であるので承認されていない
}

func TestApplication_Submit_AfterReject(t *testing.T) {
	app := submit()
	_ = app.Reject("たなか　れびゅあ", "申請内容に誤字があったため", time.Now())

	err := app.Submit("再提出した何か内容", time.Now())
	assert.Nil(t, err)
	assert.Equal(t, true, app.IsApplying())  // 再提出したので申請中である
	assert.Equal(t, false, app.IsRejected()) // 再提出したので否認とならない
}

func TestApplication_Submit_InvalidStatus(t *testing.T) {
	app := submit()
	_ = app.Approve("たなか　れびゅあ", time.Now()) // 承認

	err := app.Submit("審査される何か内容", time.Now()) // 承認されているのに提出しようとしている
	assert.Equal(t, err, ErrInvalidStatusTransition)
}

func TestApplication_Submit_LimitExceeded(t *testing.T) {
	app := NewApplication("さとう　しんせいしゃ", time.Now())
	for i := 0; i < maxSubmissions; i++ {
		// 提出上限まで提出->否認を繰り返す
		_ = app.Submit("申請内容", time.Now())
		_ = app.Reject("たなか　れびゅあ", "申請内容に誤字があったため", time.Now())
	}

	err := app.Submit("何度も提出", time.Now()) // 提出上限を超えているので再提出できない
	assert.Equal(t, err, ErrSubmissionLimitExceeded)
}

func TestApplication_Approve(t *testing.T) {
	app := submit() // 申請を作成し提出

	err := app.Approve("たなか　れびゅあ", time.Now()) // 承認
	assert.Nil(t, err)
	assert.Equal(t, false, app.IsApplying()) // 承認されたので申請中でない
	assert.Equal(t, true, app.IsApproved())  // 承認されている
}

func TestApplication_Approve_InvalidStatus(t *testing.T) {
	app := NewApplication("さとう　しんせいしゃ", time.Now()) // 申請を作成

	err := app.Approve("たなか　れびゅあ", time.Now()) // 申請を提出しないで承認しようとしている
	assert.Equal(t, err, ErrInvalidStatusTransition)
}

func TestApplication_Reject(t *testing.T) {
	app := submit() // 提出

	err := app.Reject("たなか　れびゅあ", "申請内容に誤字があったため", time.Now()) // 否認
	assert.Nil(t, err)
	assert.Equal(t, false, app.IsApplying()) // 否認されたので申請中でない
	assert.Equal(t, true, app.IsRejected())  // 否認されている
}

func TestApplication_Reject_InvalidStatus(t *testing.T) {
	app := NewApplication("さとう　しんせいしゃ", time.Now()) // 申請を作成

	err := app.Reject("たなか　れびゅあ", "申請内容に誤字があったため", time.Now()) // 申請を提出しないで否認しようとしている
	assert.Equal(t, err, ErrInvalidStatusTransition)
}
