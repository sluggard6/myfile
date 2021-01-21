package model

func Create(model interface{}) (int64, error) {
	result := db.Create(model)
	return result.RowsAffected, result.Error
}

func GetById(model interface{}, id uint) (interface{}, error) {
	result := db.First(model, id)
	return model, result.Error
}
