package facade

import (
	"context"

	"github.com/labstack/gommon/log"
	pb "septanto.github.com/omdb/proto"
	"septanto.github.com/omdb/service"
)

func (s *Server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SeachReply, error) {
	resp, err := service.RequestSearchToOmdb(in.GetSearchword(), in.GetPagination())
	if err != nil {
		log.Info("Error request to omdb: " + err.Error())
		return &pb.SeachReply{Message: "Bad Gateway"}, nil
	}
	return &pb.SeachReply{Message: string(resp.Body())}, nil
}

func (s *Server) GetDetail(ctx context.Context, in *pb.DetailRequest) (*pb.DetailReply, error) {
	resp, err := service.RequestDetailToOmdb(in.GetImdbid())
	if err != nil {
		log.Info("Error request to omdb: " + err.Error())
		return &pb.DetailReply{Message: "Bad Gateway"}, nil
	}
	return &pb.DetailReply{Message: string(resp.Body())}, nil
}

type Server struct {
	pb.UnimplementedOmdbServiceServer
}
