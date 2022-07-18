package books

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"

	// "latihan/pkg/v1/utils/constants"

	"latihan/pkg/v1/postgres"
	"latihan/pkg/v1/utils/constants"
	hlpb "latihan/proto/v1/books"
)

func (s *Server) Get(ctx context.Context, req *empty.Empty) (*hlpb.Response, error) {
	rows, err := s.config.Pg.PublicMainBook(&postgres.BooksMain{
		No: "4",
		// NameBook: "BRI",
		// Author:   "BRI",
	})
	if err != nil {

		s.logger.Errorf("[HEALTH][GET] ERROR %v", err)

		return nil, err

	}
	if len(rows) == 0 {

		s.logger.Errorf("[HEALTH][GET] ERROR no data found")

		return nil, errors.New("no data found")

	}

	var data []*hlpb.Books

	for _, row := range rows {
		data = append(data, &hlpb.Books{

			Namebook: row.NameBook,
			Author:   row.Author,
			// No:       row.No,
		})
	}

	s.logger.Infof("[HEALTH][GET] SUCCESS")

	return &hlpb.Response{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
		Data:    data,
	}, nil
}

func (s *Server) Create(ctx context.Context, req *hlpb.Books) (*hlpb.ResponseCreate, error) {
	row, err := s.config.Pg.PublicCreateBook(&postgres.InputBooksMain{
		NameBook: req.Namebook,
		Author:   req.Author,
	})

	if err != nil {
		return nil, err
	}

	return &hlpb.ResponseCreate{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
		Books: &hlpb.Books{
			Namebook: row.NameBook,
			Author:   row.Author,
		},
	}, nil
}

func (s *Server) GetAll(ctx context.Context, req *empty.Empty) (*hlpb.Response, error) {
	rows, err := s.config.Pg.GetAllData()
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("no data not found")
	}

	var data []*hlpb.Books

	for _, row := range rows {
		data = append(data, &hlpb.Books{
			Namebook: row.NameBook,
			Author:   row.Author,
		})
	}

	return &hlpb.Response{
		Success: false,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
		Data:    data,
	}, nil
}

func (s *Server) UpdateData(ctx context.Context, req *hlpb.UpdateBookData) (*hlpb.ResponseUpdateData, error) {
	rows, err := s.config.Pg.UpdateBooks(&postgres.UpdateBooks{
		No:       req.No,
		NameBook: req.Namebook,
		Author:   req.Author,
	})
	if err != nil {
		return nil, err
	}

	return &hlpb.ResponseUpdateData{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
		Books: &hlpb.Books{
			Namebook: rows.NameBook,
			Author:   rows.Author,
		},
	}, nil

}
