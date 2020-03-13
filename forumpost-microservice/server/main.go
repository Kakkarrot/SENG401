package main

import (
	"fmt"
	"log"
	"context"
	"net"
	"os"
	"os/signal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	forumpb "../proto"
)

type ForumServiceServer struct {}

type ForumItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (s *ForumServiceServer) CreateForum(ctx context.Context, req *forumpb.CreateForumReq) (*forumpb.CreateForumRes, error) {
	// read forum to be added
	forum := req.GetForum()
	// convert to local struct
	data := ForumItem{
		AuthorID: forum.GetAuthorId(),
		Title:    forum.GetTitle(),
		Content:  forum.GetContent(),
	}
	// insert to mongodb
	result, err := forumdb.InsertOne(mongoCtx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err),
		)
	}
	// get generated ID
	oid := result.InsertedID.(primitive.ObjectID)
  	forum.Id = oid.Hex()
	// return forum back
	return &forumpb.CreateForumRes{Forum: forum}, nil
}

func (s *ForumServiceServer) ReadForum(ctx context.Context, req *forumpb.ReadForumReq) (*forumpb.ReadForumRes, error) {
	// read requested forum ID
	oid, err := primitive.ObjectIDFromHex(req.GetId());
	if(err != nil){
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	// find forum on mongoDB
	result := forumdb.FindOne(ctx, bson.M{"_id": oid})
	data := ForumItem{}
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Forum does not exist. id requested %s: %v", req.GetId(), err))
	}
	// setup response of forum found
	response := &forumpb.ReadForumRes{
		Forum: &forumpb.Forum{
			Id: oid.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}
	// return forum found
	return response, nil
}

func (s *ForumServiceServer) ListForums(req *forumpb.ListForumReq, stream forumpb.ForumService_ListForumsServer) error {
	// get all forums 
	data := &ForumItem{}
	cursor, err := forumdb.Find(context.Background(), bson.M{})
	if err != nil{
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())
	// send forums as a gRPC stream
	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		stream.Send(&forumpb.ListForumRes{
			Forum: &forumpb.Forum{
				Id: data.ID.Hex(),
				AuthorId: data.AuthorID,
				Title:    data.Title,
				Content:  data.Content,
			},
		})
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	// stop sending
	return nil
}

func (s *ForumServiceServer)  UpdateForum(ctx context.Context, req *forumpb.UpdateForumReq) (*forumpb.UpdateForumRes, error){
	// read forum to be update
	forum := req.GetForum()
	oid, err := primitive.ObjectIDFromHex(forum.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Error getting Forum id %v", err),)
	}
	// convert values to be updated into local struct
	update := bson.M{
		"authord_id": forum.GetAuthorId(),
		"title":      forum.GetTitle(),
		"content":    forum.GetContent(),
	}
	forumEditID := bson.M{"_id": oid}
	// find and update based on forum id
	result := forumdb.FindOneAndUpdate(ctx, forumEditID, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))
	decoded := ForumItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Forum was not updated. forum id requested: %v", err))
	}
	// return updated forum
	return &forumpb.UpdateForumRes{
		Forum: &forumpb.Forum{
			Id:       decoded.ID.Hex(),
			AuthorId: decoded.AuthorID,
			Title:    decoded.Title,
			Content:  decoded.Content,
		},
	}, nil
}

func (s *ForumServiceServer) DeleteForum(ctx context.Context, req *forumpb.DeleteForumReq) (*forumpb.DeleteForumRes, error) {
	// read id of forum to be deleted
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	// delete forum
	_, err = forumdb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Forum was not deleted. forum id requested = %s: %v", req.GetId(), err))
	}
	// return true if successful
	return &forumpb.DeleteForumRes{
		Success: true,
	}, nil
}

var db *mongo.Client
var forumdb *mongo.Collection
var mongoCtx context.Context

func main() {
	// SETUP VARIABLES
	portNo := ":50051"
	mongoURI := "mongodb://localhost:27017"

	// SETUP LISTENER
	fmt.Printf("Starting server on port %s\n", portNo);
	listener, err := net.Listen("tcp", portNo)
	if err != nil {
		log.Fatalf("Setup on port %s failed: %v", portNo, err);
	}

	// SETUP GRPC SERVER
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	forumsrv := &ForumServiceServer{}

	// REGISTER GRPC SERVICES
	forumpb.RegisterForumServiceServer(s, forumsrv)

	// SETUP MONGODB
	fmt.Printf("Connecting to MongoDB at: %s\n", mongoURI)
	mongoCtx = context.Background()
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	} else {
		fmt.Printf("Connected to MongoDB\n")
	}

	// specify database collections
	forumdb = db.Database("mydb").Collection("forum")

	// START SERVER
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("server failed: %v", err)
		}
	}()
	fmt.Printf("Server running on port: %s\n", portNo)

	// KILL SERVER
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c

	// TEARDOWN
	fmt.Println("\nkill server")
	s.Stop()
	listener.Close()
	db.Disconnect(mongoCtx)
	fmt.Println("bye.")
}