package exception

type CollectionIsNotPresentError struct {
}

func (e CollectionIsNotPresentError) Error() string {
	return "COLLECTION_IS_NOT_PRESENT_ERROR"
}
