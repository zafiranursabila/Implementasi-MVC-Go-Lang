package controller

import "gorm.io/gorm"

type TransactionController struct {
	DB *gorm.DB
}

func (ctrl TransactionController) Transfer (ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Transfer(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "unknown error", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}