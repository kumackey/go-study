package usecase

const リリースフラグ = false

func 実装() {
	if リリースフラグ {
		新実装()
	}

	旧実装()
}
