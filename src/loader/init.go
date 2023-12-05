//go:generate mockgen -source=$GOFILE -destination=${GOFILE}_mock.go -package=$GOPACKAGE

package loader

type ILoader[T any] interface {
	Load() ([]T, error)
}
