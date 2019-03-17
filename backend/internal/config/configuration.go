package config

type Configuration interface {
	Get(option string) string
	SetFromDB() error
	GetFromMatrix(smap string, key1 string, key2 string) (string, bool)
}

var impl Configuration

func Set(conf Configuration) {
	impl = conf
}

func Get(option string) string {
	return impl.Get(option)
}

func SetFromDB() error {
	return impl.SetFromDB()
}

func GetFromMatrix(smap string, key1 string, key2 string) (string, bool) {
	return impl.GetFromMatrix(smap, key1, key2)
}
