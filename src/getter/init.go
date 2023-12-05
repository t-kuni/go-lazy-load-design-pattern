//go:generate mockgen -source=$GOFILE -destination=${GOFILE}_mock.go -package=$GOPACKAGE

package getter

type IGetter[Key any, Val any] interface {
	Get(key Key) (Val, bool, error)
}
