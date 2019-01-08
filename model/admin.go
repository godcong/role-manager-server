package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Node 节点
type Node struct {
	Name           string
	Type           string
	Property       string
	OrganizationID primitive.ObjectID
	NodeIP         string
	NodePort       string
}
