package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyResponse struct {
	Message string                            `json:"message"`
	Items   []map[string]types.AttributeValue `json:"items"`
}

func HandleRequest(ctx context.Context, event MyEvent) (MyResponse, error) {
	// AWS設定の読み込み
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return MyResponse{}, fmt.Errorf("AWS設定の読み込みに失敗しました: %v", err)
	}

	// DynamoDBクライアントの作成
	client := dynamodb.NewFromConfig(cfg)

	// 全アイテムを取得
	result, err := client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("farao"),
	})
	if err != nil {
		return MyResponse{}, fmt.Errorf("アイテムの取得に失敗しました: %v", err)
	}

	return MyResponse{
		Message: "全アイテムの取得に成功しました",
		Items:   result.Items,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
