package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyTaken
	}
	s.bookings[b.SeatID] = b

	return nil
}

func (s *MemoryStore) ListBookings(movieId string) []Booking {
	var result []Booking
	for _, b := range s.bookings {
		if b.ID == movieId {
			result = append(result, b)
		}
	}
	return result
}
