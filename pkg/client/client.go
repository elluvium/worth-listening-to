package client

import (
	"github.com/Azure/azure-sdk-for-go/storage"
	"time"
)

type Azure struct {
	storage storage.Client
}

func (az *Azure) Init(accountName, accountKey string) error {
	sg, err := storage.NewBasicClient(accountName, accountKey)
	if err != nil {
		return err
	}

	az.storage = sg
	return nil
}

func (az *Azure) AddData(pk string, rk string, data map[string]interface{}) error {
	tableService := az.storage.GetTableService()
	table := tableService.GetTableReference("posts")
	tableBatch := table.NewBatch()

	entity := storage.Entity{}
	entity.Table = table
	entity.PartitionKey = pk
	entity.RowKey = rk
	entity.TimeStamp = time.Now()
	entity.Properties = data

	tableBatch.InsertOrMergeEntity(&entity, false)

	err := tableBatch.ExecuteBatch()
	if err != nil {
		return err
	}

	return nil
}

func (az Azure) GetData() error {
	return nil
}
