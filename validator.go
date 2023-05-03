package main

import (
	"context"
	"log"

	"cloud.google.com/go/asset/apiv1/assetpb"
	"cloud.google.com/go/pubsub"
	cvassetpb "github.com/GoogleCloudPlatform/config-validator/pkg/api/validator"
	libCvAsset "github.com/GoogleCloudPlatform/config-validator/pkg/asset"
	"github.com/GoogleCloudPlatform/config-validator/pkg/gcv"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type Validator struct {
	gcv *gcv.Validator
}

type pubsubMsg struct {
	Message pubsub.Message
}

func NewValidator() (*Validator, error) {
	cv, err := gcv.NewValidator([]string{"./policy-library/policies"}, "./policy-library/lib")
	if err != nil {
		return nil, err
	}
	return &Validator{
		gcv: cv,
	}, nil
}

func (v *Validator) Handler(c *gin.Context) {
	var msg pubsubMsg
	if err := c.ShouldBindJSON(&msg); err != nil {
		log.Println(err)
		return
	}

	var tprAsset assetpb.TemporalAsset
	protoUm := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	if err := protoUm.Unmarshal([]byte(msg.Message.Data), &tprAsset); err != nil {
		log.Println(err)
		return
	}

	cvAsset := &cvassetpb.Asset{
		Name:         tprAsset.Asset.Name,
		AssetType:    tprAsset.Asset.AssetType,
		AncestryPath: libCvAsset.AncestryPath(tprAsset.Asset.Ancestors),
		Resource:     tprAsset.Asset.Resource,
		IamPolicy:    tprAsset.Asset.IamPolicy,
		OrgPolicy:    tprAsset.Asset.OrgPolicy,
	}

	log.Printf("Validating asset: %s\n", cvAsset.Name)
	violations, err := v.gcv.ReviewAsset(context.Background(), cvAsset)
	if err != nil {
		log.Println(err)
	}

	if len(violations) > 0 {
		for i, av := range violations {
			log.Printf("Violation %d (%s-%s): %s\n", i+1, av.Severity, av.Constraint, av.Message)
		}
	} else {
		log.Println("No violation detected")
	}
}
