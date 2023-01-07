package db

import (
	"context"
	"encoding/json"
	"fmt"

	"pilot_server/apps"
	"pilot_server/apps/funcs"
	"pilot_server/ent"
	"pilot_server/ent/node"

	_ "github.com/lib/pq"
)

func GetClient() *ent.Client {
	postgresinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		apps.Conf.DB.Addr, apps.Conf.DB.Port, apps.Conf.DB.User, apps.Conf.DB.PW, apps.Conf.DB.DBName)

	client, err := ent.Open("postgres", postgresinfo)
	if err != nil {
		apps.Logs.Error(fmt.Printf("failed connecting to postgresql: %v", err))
		return nil
	}
	//defer client.Close()

	// Run the auto migration tool.
	ctx := context.Background()
	if err = client.Schema.Create(ctx); err != nil {
		apps.Logs.Error(fmt.Printf("failed creating schema resources: %v", err))
	}

	return client
}

func UpdateNode(procList funcs.PodProcList) error {
	apps.Logs.Info(procList.NodeName)
	apps.Logs.Info(procList.TimeStamp)

	body, err := json.Marshal(procList.PodList)
	if err != nil {
		return err
	}
	// apps.Logs.Info(string(body))

	// DB
	ctx := context.Background()
	client := GetClient()

	n, err := client.Node.
		UpdateOneID(procList.NodeName).
		SetTimestamp(procList.TimeStamp).
		SetPodInfo(string(body)).
		Save(ctx)

	if err != nil {
		apps.Logs.Error("failed updating node: ", err)
		return err
	}

	apps.Logs.Info("updated: ", n)
	return nil
}

func CreateNode(procList funcs.PodProcList) error {
	apps.Logs.Info(procList.NodeName)
	apps.Logs.Info(procList.TimeStamp)

	body, err := json.Marshal(procList.PodList)
	if err != nil {
		return err
	}
	apps.Logs.Info(string(body))

	// DB
	ctx := context.Background()
	client := GetClient()

	apps.Logs.Info("Node Name: ", procList.NodeName)
	names, err := client.Node.
		Query().
		Unique(true).
		Select(node.FieldID).
		Strings(ctx)
	for n := range names {
		apps.Logs.Info("DB Node Name: ", n)
	}

	// Check if exists
	cnt, err := client.Node.
		Query().
		Unique(true).
		Where(node.ID(procList.NodeName)).
		Count(ctx)
	apps.Logs.Info("count: ", cnt)
	if cnt > 0 {
		// Update
		if err = UpdateNode(procList); err != nil {
			return err
		}
	} else {
		n, err := client.Node.
			Create().
			SetID(procList.NodeName).
			SetTimestamp(procList.TimeStamp).
			SetPodInfo(string(body)).
			Save(ctx)

		if err != nil {
			apps.Logs.Error("failed creating node: ", err)
			return err
		}
		apps.Logs.Info("created: ", n)
	}

	return nil
}
