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

func (az *Azure) AddData() error {
	tableService := az.storage.GetTableService()
	table := tableService.GetTableReference("posts2")

	tableBatch := table.NewBatch()

	entity := storage.Entity{}
	entity.Table = table
	entity.PartitionKey = "5"
	entity.RowKey = "0"
	entity.TimeStamp = time.Now()
	entity.Properties = map[string]interface{}{"fina": 17}

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
