package module

import (
	"backend/app"
)

func AddSubDomainItems(data []*app.SubdomainS) error {
	// var affected int64
	res, err := engine.IsTableExist(app.SubdomainS{})
	if err != nil {
		panic(err)
		return err
	}
	if !res {
		err = engine.CreateTables(app.SubdomainS{})
		if err != nil {
			panic(err)
			return err
		}

	}

	_, err = engine.Insert(&data)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
