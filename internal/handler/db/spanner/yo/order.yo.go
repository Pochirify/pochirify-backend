// Code generated by yo. DO NOT EDIT.
// Package yo contains the types.
package yo

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// Order represents a row from 'Orders'.
type Order struct {
	ID            string    `spanner:"ID" json:"ID"`                       // ID
	UserID        string    `spanner:"UserID" json:"UserID"`               // UserID
	UserAddressID string    `spanner:"UserAddressID" json:"UserAddressID"` // UserAddressID
	Status        string    `spanner:"Status" json:"Status"`               // Status
	PaymentMethod string    `spanner:"PaymentMethod" json:"PaymentMethod"` // PaymentMethod
	ProductID     string    `spanner:"ProductID" json:"ProductID"`         // ProductID
	Price         int64     `spanner:"Price" json:"Price"`                 // Price
	CreateTime    time.Time `spanner:"CreateTime" json:"CreateTime"`       // CreateTime
	UpdateTime    time.Time `spanner:"UpdateTime" json:"UpdateTime"`       // UpdateTime
}

func OrderPrimaryKeys() []string {
	return []string{
		"ID",
	}
}

func OrderColumns() []string {
	return []string{
		"ID",
		"UserID",
		"UserAddressID",
		"Status",
		"PaymentMethod",
		"ProductID",
		"Price",
		"CreateTime",
		"UpdateTime",
	}
}

func OrderWritableColumns() []string {
	return []string{
		"ID",
		"UserID",
		"UserAddressID",
		"Status",
		"PaymentMethod",
		"ProductID",
		"Price",
		"CreateTime",
		"UpdateTime",
	}
}

func (o *Order) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "ID":
			ret = append(ret, &o.ID)
		case "UserID":
			ret = append(ret, &o.UserID)
		case "UserAddressID":
			ret = append(ret, &o.UserAddressID)
		case "Status":
			ret = append(ret, &o.Status)
		case "PaymentMethod":
			ret = append(ret, &o.PaymentMethod)
		case "ProductID":
			ret = append(ret, &o.ProductID)
		case "Price":
			ret = append(ret, &o.Price)
		case "CreateTime":
			ret = append(ret, &o.CreateTime)
		case "UpdateTime":
			ret = append(ret, &o.UpdateTime)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (o *Order) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "ID":
			ret = append(ret, o.ID)
		case "UserID":
			ret = append(ret, o.UserID)
		case "UserAddressID":
			ret = append(ret, o.UserAddressID)
		case "Status":
			ret = append(ret, o.Status)
		case "PaymentMethod":
			ret = append(ret, o.PaymentMethod)
		case "ProductID":
			ret = append(ret, o.ProductID)
		case "Price":
			ret = append(ret, o.Price)
		case "CreateTime":
			ret = append(ret, o.CreateTime)
		case "UpdateTime":
			ret = append(ret, o.UpdateTime)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newOrder_Decoder returns a decoder which reads a row from *spanner.Row
// into Order. The decoder is not goroutine-safe. Don't use it concurrently.
func newOrder_Decoder(cols []string) func(*spanner.Row) (*Order, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*Order, error) {
		var o Order
		ptrs, err := o.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &o, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (o *Order) Insert(ctx context.Context) *spanner.Mutation {
	values, _ := o.columnsToValues(OrderWritableColumns())
	return spanner.Insert("Orders", OrderWritableColumns(), values)
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (o *Order) Update(ctx context.Context) *spanner.Mutation {
	values, _ := o.columnsToValues(OrderWritableColumns())
	return spanner.Update("Orders", OrderWritableColumns(), values)
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (o *Order) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	values, _ := o.columnsToValues(OrderWritableColumns())
	return spanner.InsertOrUpdate("Orders", OrderWritableColumns(), values)
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (o *Order) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, OrderPrimaryKeys()...)

	values, err := o.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "Order.UpdateColumns", "Orders", err)
	}

	return spanner.Update("Orders", colsWithPKeys, values), nil
}

// FindOrder gets a Order by primary key
func FindOrder(ctx context.Context, db YORODB, id string) (*Order, error) {
	key := spanner.Key{id}
	row, err := db.ReadRow(ctx, "Orders", key, OrderColumns())
	if err != nil {
		return nil, newError("FindOrder", "Orders", err)
	}

	decoder := newOrder_Decoder(OrderColumns())
	o, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindOrder", "Orders", err)
	}

	return o, nil
}

// ReadOrder retrieves multiples rows from Order by KeySet as a slice.
func ReadOrder(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*Order, error) {
	var res []*Order

	decoder := newOrder_Decoder(OrderColumns())

	rows := db.Read(ctx, "Orders", keys, OrderColumns())
	err := rows.Do(func(row *spanner.Row) error {
		o, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, o)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadOrder", "Orders", err)
	}

	return res, nil
}

// Delete deletes the Order from the database.
func (o *Order) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := o.columnsToValues(OrderPrimaryKeys())
	return spanner.Delete("Orders", spanner.Key(values))
}
