package table_manager

type Table interface {
	Close()
	Creator() Player
}
