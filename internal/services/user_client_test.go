package services

// import (
// 	"context"
// 	"net"
// 	"testing"

// 	"github.com/AhmedEnnaime/SnapEvent/pb"
// 	"github.com/AhmedEnnaime/SnapEvent/utils"
// 	"github.com/stretchr/testify/require"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func TestClientSignUp(t *testing.T) {
// 	t.Parallel()

// 	userServer, serverAddress := startTestUserServer(t)
// 	userClient := newTestUserClient(t, serverAddress)

// 	user := utils.NewUser()
// 	expectedID := user.Id
// 	req := &pb.CreateUserRequest{
// 		Name:     user.Name,
// 		Birthday: user.Birthday,
// 		Email:    user.Email,
// 		Password: user.Password,
// 		Gender:   user.Gender,
// 	}
// 	res, err := userClient.CreateUser(context.Background(), req)
// 	require.NoError(t, err)
// 	require.NotNil(t, res)
// 	require.Equal(t, expectedID, res.User.Id)

// 	other, err := userServer.h.us.GetByID(uint(res.User.Id))
// 	require.NoError(t, err)
// 	require.NotNil(t, other)

// 	// requireSameUser(t, user, other)

// }

// func startTestUserServer(t *testing.T) (*UserServer, string) {
// 	h, cleaner := setUp(t)
// 	defer cleaner(t)
// 	userServer := NewUserServer(h)
// 	grpcServer := grpc.NewServer()
// 	pb.RegisterUserServiceServer(grpcServer, userServer)

// 	listener, err := net.Listen("tcp", ":0")
// 	require.NoError(t, err)

// 	go grpcServer.Serve(listener)

// 	return userServer, listener.Addr().String()
// }

// func newTestUserClient(t *testing.T, serverAddress string) pb.UserServiceClient {
// 	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	require.NoError(t, err)
// 	return pb.NewUserServiceClient(conn)
// }

// func requireSameUser(t *testing.T, user1 *pb.User, user2 *pb.User) {
// 	require.Equal(t, user1, user2)
// }
