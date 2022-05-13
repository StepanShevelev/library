package protobuff

import (
	"context"
	cfg "github.com/StepanShevelev/library/config"
	mydb "github.com/StepanShevelev/library/db"
	proto "github.com/StepanShevelev/library/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type LibInfoServer struct {
	proto.UnimplementedLibInfoServer

	Config *cfg.Config
}


// NewServer ...
func NewServer(c *cfg.Config) *grpc.Server {
	s := grpc.NewServer()
	srv := &LibInfoServer{Config: c}
	proto.RegisterLibInfoServer(s, srv)
	return s
}

// RunServer ...
func RunServer(s *grpc.Server, c *cfg.Config) {
	l, err := net.Listen("tcp", ":"+c.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
	log.Printf("server listening at %v", l.Addr())
}

func (s *LibInfoServer) FindBook(ctx context.Context, r *proto.LibRequest, c *cfg.Config) (*proto.LibResponse, error) {
	db, err := mydb.New(c)
	if err != nil {
		log.Fatal(err)
	}
	db.SetDB()
	name, err := mydb.Client.GiveBookByAuthor(r.GetId())
	if err != nil {
		return nil, err
	}
	return &proto.LibResponse{Name: name}, nil
}

func (s *LibInfoServer) FindAuthor(ctx context.Context, r *proto.LibRequest, c *cfg.Config) (*proto.LibResponse, error) {
	db, err := mydb.New(c)
	if err != nil {
		log.Fatal(err)
	}
	db.SetDB()
	name, err := mydb.Client.GiveAuthorByBook(r.GetId())
	if err != nil {
		return nil, err
	}
	return &proto.LibResponse{Name: name}, nil
}
