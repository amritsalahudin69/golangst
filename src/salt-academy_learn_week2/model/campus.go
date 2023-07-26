package model

type Campus struct {
	ID         uint   `dbq:"id"`
	Name       string `dbq:"nama"`
	Email      string `dbq:"email"`
	NoTphn     string `dbq:"no_tlfn"`
	Address    string `dbq:"address"`
	Rektor     string `dbq:"rektor"`
	Akreditasi string `dbq:"akreditasi"`
}

func (m Campus) GetTableName() string {
	return `campus`
}
func ListRowsInsertCampus() []string {
	return []string{
		"nama",
		"email",
		"no_tlfn",
		"address",
		"rektor",
		"akreditasi",
	}
}
