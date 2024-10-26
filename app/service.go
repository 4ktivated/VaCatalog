package app

import (
	"context"
	"math/rand"
	"some_code/apperror"
	"some_code/dbase"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// response := parser.HHparser("golang", "1", "0", "12")

type Service struct {
	rnd     *rand.Rand
	VacsDAO *dbase.VacsDAO
}

func NewService(VacsDAO *dbase.VacsDAO) *Service {
	return &Service{
		VacsDAO: VacsDAO,
		rnd:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *Service) Main_Page(ctx context.Context, url string, ttlDays int) (*dbase.Vac, error) {
	Vac := &dbase.Vac{
		Title:   `bson:"Title"`,
		Company: `bson:"Company"`,
		URL:     `bson:"URL"`,
		Salary:  `bson:"Salary"`,
		Info:    `bson:"Info"`,
	}

	for it := 0; it < 10; it++ {
		Vac.ID = s.generateRandomID()
		err := s.VacsDAO.InsertVacs(ctx, Vac)
		if err == nil {
			return Vac, nil
		}
		if !mongo.IsDuplicateKeyError(err) {
			return nil, err
		}
	}
	return nil, apperror.ErrCollision
}


func (s *Service) GetbyLang(ctx context.Context, Vac string) (string, error) {
	sURL, err := s.VacsDAO.FindbyLang(ctx, Vac)
	if err != nil {
		return "", err
	}
	return sURL.URL, nil
}



var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (s *Service) generateRandomID() string {
	const idLength = 6
	id := make([]rune, idLength)
	for i := range id {
		id[i] = symbols[s.rnd.Intn(len(symbols))]
	}
	return string(id)
}

func getExpirationTime(ttlDays int) *time.Time {
	if ttlDays <= 0 {
		return nil
	}
	t := time.Now().Add(time.Hour * 24 * time.Duration(ttlDays))
	return &t
}
