//go:generate mockgen -source=$GOFILE -destination=${GOFILE}_mock.go -package=$GOPACKAGE

package src

type ILoader[T any] interface {
	Load() ([]T, error)
}

type IGetter[Key any, Val any] interface {
	Get(key Key) (Val, bool, error)
}
