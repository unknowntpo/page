package page

import (
	"context"
	"io"
	"log"

	"github.com/unknowntpo/page/domain"
	"github.com/unknowntpo/page/pkg/errors"

	connect "github.com/bufbuild/connect-go"
	pb "github.com/unknowntpo/page/gen/proto/page"
)

func NewPageServer(uCase domain.PageUsecase) *pageServer {
	return &pageServer{useCase: uCase}
}

type pageServer struct {
	useCase domain.PageUsecase
}

// NewList creates a new list for a given user and list key.
// Return value:
// - Status OK if operation succeed.
// - error if exists.
//
// An error object indicating the error if the operation was unsuccessful. The possible errors are:
// - connect.CodeInvalidArgument: if the ListKey or UserID are not valid.
// - connect.CodeAlreadyExists: if a list with the same list key already exists for the given user.
// - connect.CodeInternal: if an internal error occurs.
func (s *pageServer) NewList(ctx context.Context, req *connect.Request[pb.NewListRequest]) (*connect.Response[pb.NewListResponse], error) {
	// verify inpput
	// TODO: Put verification inside domain
	if req.Msg.ListKey == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, domain.ErrInvalidListKey)
	}
	if req.Msg.UserID <= 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, domain.ErrInvalidUserID)
	}
	if err := s.useCase.NewList(ctx, req.Msg.UserID, domain.ListKey(req.Msg.ListKey)); err != nil {
		switch {
		case errors.Is(err, domain.ErrListAlreadyExists):
			return nil, connect.NewError(connect.CodeAlreadyExists, domain.ErrListAlreadyExists)
		default:
			log.Println("failed on s.useCase.NewList", err)
			return nil, connect.NewError(connect.CodeInternal, domain.ErrInternal)
		}
	}
	res := connect.NewResponse(&pb.NewListResponse{
		Status: "OK",
	})
	res.Header().Set("Page-Version", "v1")
	return res, nil
}

// GetHead retrieves the head page of the specified list for a given user.
// - An error object indicating the error if the operation was unsuccessful. The possible errors are:
// - connect.CodeNotFound: if the specified list was not found for the given user.
// - connect.CodeInternal: if an internal error occurs.
func (s *pageServer) GetHead(ctx context.Context, req *connect.Request[pb.GetHeadRequest]) (*connect.Response[pb.GetHeadResponse], error) {
	pageKey, err := s.useCase.GetHead(ctx, req.Msg.UserID, domain.ListKey(req.Msg.ListKey))
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrListNotFound):
			return nil, connect.NewError(connect.CodeNotFound, domain.ErrListNotFound)
		default:
			log.Println("failed on s.useCase.GetHead", err)
			return nil, connect.NewError(connect.CodeInternal, domain.ErrInternal)
		}
	}
	res := connect.NewResponse(&pb.GetHeadResponse{
		PageKey: string(pageKey),
	})
	res.Header().Set("Page-Version", "v1")
	return res, nil
}

// GetPage streams the content of a page in a list for a given user.
// Return value:
// - page content: the conent of this page.
// - next: next pageKey.
// - error if exists.
// An error object indicating the error if the operation was unsuccessful. The possible errors are:
// - connect.CodeNotFound: if the specified page was not found for the given user.
// - connect.CodeInternal: if an internal error occurs.
func (s *pageServer) GetPage(ctx context.Context, stream *connect.BidiStream[pb.GetPageRequest, pb.GetPageResponse]) error {
	for {
		req, err := stream.Receive()
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				return nil
			default:
				log.Println("failed on s.useCase.GetPage", err)
				return connect.NewError(connect.CodeAborted, nil)
			}
		}
		page, err := s.useCase.GetPage(ctx, domain.PageKey(req.PageKey))
		if err != nil {
			switch {
			case errors.KindIs(err, errors.ResourceNotFound):
				return connect.NewError(connect.CodeNotFound, errors.New(errors.ResourceNotFound, "resource not found"))
			default:
				// TODO: log error
				log.Println("failed on s.useCase.GetPage", err)
				return connect.NewError(connect.CodeInternal, errors.New(errors.Internal, "internal server error"))
			}
		}
		res := connect.NewResponse(&pb.GetPageResponse{
			PageContent: page.Content,
			Next:        string(page.Next),
		})
		res.Header().Set("Page-Version", "v1")
		if err := stream.Send(res.Msg); err != nil {
			// TODO: hwo to handle this error ?
			log.Println("failed on stream.Send", err)
			return connect.NewError(connect.CodeAborted, nil)
		}
	}
}

// SetPage sets the content of a page in a list for a given user.
// Return value:
// - pageKey: the pageKey assigned to this page.
// - error if exists.
// An error object indicating the error will be returned if the operation was unsuccessful. The possible errors are:
// - connect.CodeInternal: if an internal error occurs.
// - connect.CodeAborted: if the operation was aborted due to an error.
// - domain.ErrInternal: if an internal error occurs in the domain layer.
func (s *pageServer) SetPage(ctx context.Context, stream *connect.BidiStream[pb.SetPageRequest, pb.SetPageResponse]) error {
	for {
		req, err := stream.Receive()
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				return nil
			default:
				// TODO: Which error should we handle ?
				log.Println("failed on s.useCase.SetPage", err)
				return connect.NewError(connect.CodeInternal, domain.ErrInternal)
			}
		}
		p := domain.Page{}
		p.SetContent(req.PageContent)
		pageKey, err := s.useCase.SetPage(ctx, req.UserID, domain.ListKey(req.ListKey), p)
		if err != nil {
			// TODO: Which error should we handle ?
			log.Println("failed on s.useCase.SetPage", err)
			return connect.NewError(connect.CodeAborted, domain.ErrInternal)
		}
		res := connect.NewResponse(&pb.SetPageResponse{
			PageKey: string(pageKey),
		})
		res.Header().Set("Page-Version", "v1")
		if err := stream.Send(res.Msg); err != nil {
			// TODO: hwo to handle this error ?
			log.Println("failed on stream.Send", err)
			return connect.NewError(connect.CodeAborted, nil)
		}
	}
}
