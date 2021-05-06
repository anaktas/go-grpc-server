package main

import (
	"context"
	"log"
	"net"
	//"fmt"

	db "7linternational.com/grpc-server/db"
	pb "7linternational.com/grpc-server/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedRecipesServer
}

func (s *server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Registration attempt for: %v", in.GetEmail())

	code, err := db.Register(in.GetFirstName(), in.GetLastName(), in.GetEmail(), in.GetPassword())
	if err != nil {
		log.Println(err.Error())
		return &pb.RegisterResponse{
			Code: int64(code), 
			Message: err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Code: int64(code), 
		Message: "OK",
	}, nil
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Login attempt for: %v", in.GetEmail())

	user, code, err := db.Login(in.GetEmail(), in.GetPassword())
	if err != nil {
		log.Println(err.Error())
		return &pb.LoginResponse{
			Code: int64(code), 
			Message: err.Error(), 
			UserId: 0, 
			FirstName: "", 
			LastName: "", 
			Email: "",
		}, nil
	}

	return &pb.LoginResponse{
		Code: int64(code), 
		Message: "OK", 
		UserId: user.Id, 
		FirstName: user.FirstName, 
		LastName: user.LastName, 
		Email: user.Email,
	}, nil
}

func (s *server) GetUserRecipes(ctx context.Context, in *pb.GetUserRecipesRequest) (*pb.GetUserRecipesResponse, error) {
	log.Printf("Get recipes of user: %v", in.GetUserId())

	code, err, recipes := db.GetUserRecipes(in.GetUserId())
	if err != nil {
		return &pb.GetUserRecipesResponse{
			Code: int64(code), 
			Message: err.Error(), 
			Recipes: nil, 
		}, nil
	}

	response        := new(pb.GetUserRecipesResponse)
	response.Code    = int64(code)
	response.Message = "OK"

	responseRecipes := make([]*pb.Recipe, 0)

	for _, recipe := range recipes {
		responseRecipe := new(pb.Recipe)

		responseRecipe.RecipeId   = recipe.Id
		responseRecipe.Title      = recipe.Title
		responseRecipe.Description = recipe.Description
		responseRecipe.ImageUrl   = recipe.ImageUrl
		responseRecipe.UserId     = recipe.UserId

		responseProducts := make([]*pb.Product, 0)

		for _, product := range recipe.Products {
			responseProduct := new(pb.Product)

			responseProduct.ProductId   = product.Id
			responseProduct.Title       = product.Title
			responseProduct.Description = product.Description
			responseProduct.ImageUrl    = product.ImageUrl

			responseProducts = append(responseProducts, responseProduct)
		}

		responseRecipe.Products = responseProducts

		responseRecipes = append(responseRecipes, responseRecipe)

	}

	response.Recipes = responseRecipes

	//log.Println("Server: Response: " + fmt.Sprintf("%v", response))

	return response, nil
}

func main() {
	log.Println("Listening")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRecipesServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
