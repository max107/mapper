package mapper

func New[K, V comparable](valueMap map[K]V) Mapper[K, V] {
	keyMap := make(map[V]K, len(valueMap))

	for k, v := range valueMap {
		keyMap[v] = k
	}

	return Mapper[K, V]{
		valueMap: valueMap,
		keyMap:   keyMap,
	}
}

type Mapper[K, V comparable] struct {
	valueMap map[K]V
	keyMap   map[V]K
}

func (m *Mapper[K, V]) Keys() []K   { return Keys[K, V](m.valueMap) }
func (m *Mapper[K, V]) Values() []V { return Values[K, V](m.valueMap) }

func (m *Mapper[K, V]) FromKey(key K) V   { return m.valueMap[key] }
func (m *Mapper[K, V]) FromValue(key V) K { return m.keyMap[key] }
