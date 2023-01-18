package module

import (
	"common"
	"fmt"
)

func AddNucleiItems(data []*common.Nucleivulns) error {

	res, err := engine.IsTableExist(common.Nucleivulns{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if !res {
		err = engine.CreateTables(common.Nucleivulns{})
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

	}

	_, err = engine.Insert(&data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
