// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldSpouseID holds the string denoting the spouse_id field in the database.
	FieldSpouseID = "spouse_id"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeSpouse holds the string denoting the spouse edge name in mutations.
	EdgeSpouse = "spouse"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeInfo holds the string denoting the info edge name in mutations.
	EdgeInfo = "info"
	// EdgeRentals holds the string denoting the rentals edge name in mutations.
	EdgeRentals = "rentals"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PetsTable is the table that holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "owner_id"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "users"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "users"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// SpouseTable is the table that holds the spouse relation/edge.
	SpouseTable = "users"
	// SpouseColumn is the table column denoting the spouse relation/edge.
	SpouseColumn = "spouse_id"
	// CardTable is the table that holds the card relation/edge.
	CardTable = "cards"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "owner_id"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "metadata"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "id"
	// InfoTable is the table that holds the info relation/edge.
	InfoTable = "infos"
	// InfoInverseTable is the table name for the Info entity.
	// It exists in this package in order to avoid circular dependency with the "info" package.
	InfoInverseTable = "infos"
	// InfoColumn is the table column denoting the info relation/edge.
	InfoColumn = "id"
	// RentalsTable is the table that holds the rentals relation/edge.
	RentalsTable = "rentals"
	// RentalsInverseTable is the table name for the Rental entity.
	// It exists in this package in order to avoid circular dependency with the "rental" package.
	RentalsInverseTable = "rentals"
	// RentalsColumn is the table column denoting the rentals relation/edge.
	RentalsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldParentID,
	FieldSpouseID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
