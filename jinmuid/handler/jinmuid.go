package handler

import (
	db "github.com/jinmukeji/gf-api2/jinmuid/mysqldb"
	bizcorepb "github.com/jinmukeji/proto/gen/micro/idl/jm/core/v1"
	jinmuidpb "github.com/jinmukeji/proto/gen/micro/idl/jinmuid/v1"
	sempb "github.com/jinmukeji/proto/gen/micro/idl/jm/sem/v1"
	smspb "github.com/jinmukeji/proto/gen/micro/idl/jm/sms/v1"
	subscriptionpb "github.com/jinmukeji/proto/gen/micro/idl/jm/subscription/v1"
)

// JinmuIDService 金姆ID Service
type JinmuIDService struct {
	datastore       db.Datastore
	rpcSvc          jinmuidpb.UserManagerAPIService
	semSvc          sempb.SemAPIService
	encryptKey      string
	smsSvc          smspb.SmsAPIService
	bizSvc          bizcorepb.JinmuhealthAPIService
	subscriptionSvc subscriptionpb.SubscriptionManagerAPIService
}

// NewJinmuIDService 构建JinmuIDService
func NewJinmuIDService(datastore db.Datastore, smsSvc smspb.SmsAPIService, semSvc sempb.SemAPIService, rpcUserManagerSvc jinmuidpb.UserManagerAPIService, bizSvc bizcorepb.JinmuhealthAPIService, subscriptionSvc subscriptionpb.SubscriptionManagerAPIService, encryptKey string) *JinmuIDService {
	j := &JinmuIDService{
		datastore:       datastore,
		smsSvc:          smsSvc,
		semSvc:          semSvc,
		rpcSvc:          rpcUserManagerSvc,
		encryptKey:      encryptKey,
		bizSvc:          bizSvc,
		subscriptionSvc: subscriptionSvc,
	}
	return j
}