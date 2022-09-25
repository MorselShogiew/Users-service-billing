package usecases

// i

func (r *ResizeService) GetUserBalance(reqID string, id int) (float64, error) {
	return r.resizeAPIRepo.GetUserBalance(id)
}
func (r *ResizeService) СreditingFunds(reqID string, id int, value float64) error {
	return r.resizeAPIRepo.СreditingFunds(id, value)
}
func (r *ResizeService) DebitingFunds(reqID string, id int, value float64) error {
	return r.resizeAPIRepo.DebitingFunds(id, value)
}
func (r *ResizeService) TransferFunds(reqID string, idFrom int, idTo int, value float64) error {
	return r.resizeAPIRepo.TransferFunds(idFrom, idTo, value)
}
