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

func reject() *Application {
	app := submit()
	_ = app.Reject("たなか　れびゅあ", "申請内容に誤字があったため", time.Now())

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
	app := submit() // 申請を作成し提出

	err := app.Reject("たなか　れびゅあ", "申請内容に誤字があったため", time.Now()) // 否認
	assert.Nil(t, err)
	assert.Equal(t, false, app.IsApplying()) // 否認されたので申請中でない
	assert.Equal(t, true, app.IsRejected())  // 否認されている
}

func TestApplication_Reject_InvalidStatus(t *testing.T) {
	app := NewApplication("さとう　しんせいしゃ", time.Now()) // 申請を作成

	err := app.Reject("たなか　れびゅあ", "申請内容に誤字があったため", time.Now()) // 否認       // 申請を提出しないで承認しようとしている
	assert.Equal(t, err, ErrInvalidStatusTransition)
}

func TestApplication_Resubmit(t *testing.T) {
	app := reject()

	err := app.Resubmit("再提出した何か内容")
	assert.Nil(t, err)
	assert.Equal(t, true, app.IsApplying())  // 再提出したので申請中である
	assert.Equal(t, false, app.IsRejected()) // 再提出したので否認とならない
}

func TestApplication_Resubmit_InvalidStatus(t *testing.T) {
	app := submit()

	err := app.Resubmit("再提出した何か内容")
	assert.Equal(t, err, ErrInvalidStatusTransition)
}
