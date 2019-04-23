package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type ASN struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ASN    string             `json:"asn,omitempty" bson:"asn,omitempty"`
	STATUS string             `json:"status,omitempty" bson:"status,omitempty"`
}
