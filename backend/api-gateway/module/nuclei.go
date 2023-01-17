package module

import "common"

func AddNucleiItems(data []*common.Nucleivulns) error {

	if !mysql_db.HasTable(&common.Nucleivulns{}) {
		mysql_db.CreateTable(&common.Nucleivulns{})
	}
	result := mysql_db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
