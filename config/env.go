package env

type Config struct {
	WriterDBURL string `env:"WRITER_DB_URL"`
	ReaderDBURL string `env:"READER_DB_URL"`
}
