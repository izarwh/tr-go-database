package repository

import (
	"context"
	"database/sql"
	"transactions_mysql/model"
)

type voucherRepository struct {
	db *sql.DB
}

func NewVoucherRepository(db *sql.DB) *voucherRepository {
	return &voucherRepository{db}
}

func (vRepo *voucherRepository) FindAllVoucher(ctx context.Context) ([]model.Voucher, error) {
	query := "select id, code, value from vouchers"
	var sliceVouchers []model.Voucher

	res, err := vRepo.db.QueryContext(ctx, query)

	if err != nil {
		return sliceVouchers, err
	}

	for res.Next() {
		var voucher_model model.Voucher
		var id, code, value = voucher_model.GetVoucher()

		res.Scan(id, code, value) //akan mengganti nilai product model yang diambil karena berupa pointer

		sliceVouchers = append(sliceVouchers, voucher_model)
	}

	return sliceVouchers, nil
}

func (vRepo *voucherRepository) FindVoucherByCode(ctx context.Context, code string) (model.Voucher, error) {
	var vouch model.Voucher

	query := "select id, code, value from vouchers where code like ?"

	err := vRepo.db.QueryRowContext(ctx, query, code).Scan(vouch.GetVoucherId(), vouch.GetVoucherCode(), vouch.GetVoucherValue())

	if err == sql.ErrNoRows {
		return vouch, err
	}

	return vouch, nil
}
