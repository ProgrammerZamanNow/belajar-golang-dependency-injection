package simple

type SimpleRepository struct {
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}
