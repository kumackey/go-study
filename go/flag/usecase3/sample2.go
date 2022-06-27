package usecase

const リリースフラグ = true

func 実装() {
	if リリースフラグ {
		新実装()
	}

	旧実装()
}
