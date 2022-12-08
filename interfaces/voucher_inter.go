package interfaces

import (
	"context"
	"transactions_mysql/model"
)

type VoucherInter interface {
	FindAllVoucher(ctx context.Context) ([]model.Voucher, error)
	FindVoucherByCode(ctx context.Context, code string) (model.Voucher, error)
}
