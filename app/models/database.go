package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type Asn struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ASN    string             `json:"asn,omitempty" bson:"asn,omitempty"`
	STATUS string             `json:"status,omitempty" bson:"status,omitempty"`
}

type Country struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ISO    string             `json:"iso,omitempty" bson:"iso,omitempty"`
	STATUS string             `json:"status,omitempty" bson:"status,omitempty"`
}

type GenericFirewallItem struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VALUE  string             `json:"value,omitempty" bson:"value,omitempty"`
	STATUS string             `json:"status,omitempty" bson:"status,omitempty"`
}
