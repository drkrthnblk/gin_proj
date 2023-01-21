package postgresrepo

import (
	"gin_proj/app/models"
	"gin_proj/pkg/common/request"
	dbconns "gin_proj/pkg/dbConns"
	"gin_proj/pkg/dbConns/postgresConn"
	"strings"
)

type OrderRepo struct{}

func (o *OrderRepo) Migrate(rctx *request.Context, order models.Order) error {
	log := rctx.Log
	client, _ := dbconns.NewDbConnSelectorFactory().GetDbConn(dbconns.POSTGRESCONN)
	postgresClient := client.(*postgresConn.PostgresClient)
	db, err := postgresClient.GetClient()
	if err != nil {
		log.Errorln("error in getting postgres client: ", err)
		return err
	}
	if err := db.AutoMigrate(&order); err != nil {
		log.Errorln("error in order table: ", err)
		return err
	}
	log.Infoln("order table successfuy migrated")
	return nil
}

func (o *OrderRepo) Create(rctx *request.Context, order models.Order) error {
	log := rctx.Log
	client, _ := dbconns.NewDbConnSelectorFactory().GetDbConn(dbconns.POSTGRESCONN)
	postgresClient := client.(*postgresConn.PostgresClient)
	db, err := postgresClient.GetClient()
	if err != nil {
		log.Errorln("error in getting postgres client: ", err)
		return err
	}
	if tx := db.Save(&order); tx.Error != nil {
		log.Errorln("error in order table: ", tx.Error)
		return tx.Error
	}
	log.Infoln("order table successfuy migrated")
	return nil
}

func (o *OrderRepo) Upsert(rctx *request.Context, order models.Order) error {
	log := rctx.Log
	client, _ := dbconns.NewDbConnSelectorFactory().GetDbConn(dbconns.POSTGRESCONN)
	postgresClient := client.(*postgresConn.PostgresClient)
	db, err := postgresClient.GetClient()
	if err != nil {
		log.Errorln("error in getting postgres client: ", err)
		return err
	}
	if tx := db.Save(&order); tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "duplicate key value violates unique constraint") {
			if tx = db.Model(&order).Where("id=?", order.ID).Updates(order); tx.Error != nil {
				log.Errorln("error in order table updation: ", tx.Error)
				return tx.Error
			}
		} else {
			log.Errorln("error in order table creation: ", tx.Error)
			return tx.Error
		}
	}
	log.Infoln("order table successfuy migrated")
	return nil
}
