package mysql

func ConvertPagination(perpage uint64, page uint64) (limit uint64, offset uint64) {
	offset = perpage * (page - 1)
	limit = perpage

	return limit, offset
}
